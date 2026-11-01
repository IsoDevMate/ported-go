# 2026-11-01 — Go project structure and wiring (why the pieces feel scattered)

## Background

The repo is a TypeScript/Express app being ported to Go. Right now it contains:

```
gienagon-go/
├── go.mod                  ← module root: "gienagon-backend"
├── go.sum
├── cmd/api/
│   ├── main.go             ← package main, current entry point
│   ├── api.go              ← package main, currently just `package main`
│   └── notes.md
├── docs/explainers/
├── microworlds/
├── AGENTS.md
├── notes.md
└── REVIEW_PROMPTS.md
```

`main.go` builds a chi router, registers `/`, then calls `http.ListenAndServe(":3333", r)` — with no error handling, and with the `PORT` lookup placed *after* that call. The earlier version of `api.go` held a tutorial-shaped skeleton (`application` struct, `config`, `dbConfig`, `mount()`, `run()`) but did not compile.

The feeling of "everything at once" comes from three different problems being mixed into one task:
1. Go **syntax** (capitalisation, `err != nil`, unused imports)
2. Go **project layout** (what goes in which folder/package)
3. **Application architecture** (app struct → config → db → router → server)

Solve them in that order, in small slices.

## Intuition before details

Think of a Go module as a **building**, and packages as **rooms**.

- `go.mod` is the street address. Everything under that folder is one module, imported by paths that start with the module name (`gienagon-backend/...`).
- A **directory** is a room (a package). Every `.go` file in the same room is the *same* package and can use every other file's functions directly — no import needed. That is why `main.go` and `api.go` sit beside each other under `package main`.
- **Capitalisation is the door.** `HealthHandler` has a door to the street (exported). `healthHandler` does not (package-private).
- `cmd/api` is the room that gets turned into a **binary** because it holds `func main()`.

That's the whole tree. Go has no default deep nesting convention: you add a folder only when you need a real boundary (testable unit, shared code, or privacy via `internal/`).

## Literate walkthrough

### 1. Reading the tree

| Path | What it is | How you "import" it |
|---|---|---|
| `go.mod` | module root, name = `gienagon-backend` | — |
| `cmd/api/main.go` | `package main` → binary | built with `go run ./cmd/api` |
| `cmd/api/api.go` | same package as `main.go` | *no import needed* |
| `internal/config/` (future) | private package | `import "gienagon-backend/internal/config"` |
| `internal/handler/` (future) | private package | `import "gienagon-backend/internal/handler"` |

Import path = **module path + folder path**, nothing to do with the GitHub URL.

### 2. Why `api.go` broke (the compiler as teacher)

The version pasted earlier was:

```go
import ( "contezt"; "database/sql"; ... )   // no `package` line

type config struct { apiUrl stiing }        // `stiing` doesn't exist

func (app *application) mount() http.handler {   // lowercase h
	r.use(midddleware.RealIP)                    // `use`, `midddleware`
	http.ListenAndServe(":3333", r)              // mount() serving?
}                                                // missing return

func (app *application) run(n http.Handler) error {
}                                                // missing return
```

Each error teaches a Go rule:

| Error | Why Go cares |
|---|---|
| `expected 'package', found 'import'` | Every file, without exception, begins with a package clause. |
| `"contezt" is not in std` | Import paths are literal folder names; there is no fuzzy matching. |
| `undefined: stiing` | Types must exist; Go won't guess you meant `string`. |
| `undefined: http.handler` | Only uppercase names are exported. `http.Handler` is the interface. |
| `r.use undefined` | Same rule, on chi's router type: `r.Use`. |
| `missing return` | A declared return type is a promise; every path must return. |
| `declared and not used` | Unused imports/variables are compile errors, not warnings — Go refuses to ship dead weight. |

Two **design** errors hide in there as well:
- `mount()` should only *build and return* the router. Serving belongs in `run()`.
- `mount()` returning a router and also calling `ListenAndServe` makes the function untestable and ignores the error.

### 3. The shape the tutorial is aiming at

```
main()
  ├─ cfg  := load env into config{}          → internal/config
  ├─ db   := sql.Open(...) + Ping            → internal/db
  ├─ app  := &application{config: cfg, db: db}
  ├─ mux  := app.mount()                     → builds chi router, registers routes
  └─ err  := app.run(mux)                    → http.Server + graceful shutdown
```

`mount()` returns `http.Handler`; `run()` owns `http.Server`, timeouts, and shutdown. Handlers live in their own package and receive `*application` (or a purpose-built dependency struct) so they can reach config/db.

### 4. The bug in `main.go` today

```go
http.ListenAndServe(":3333", r)   // blocks forever

port := os.Getenv("PORT")         // never runs while the server is up
```

`ListenAndServe` blocks until the server stops, so everything after it is effectively dead code. Read config **before** starting the server, pass the address into it, and handle the returned error. Predict what `go run ./cmd/api` does today before changing it.

### 5. Order of work (one slice per session)

1. `main.go`: env → addr → `http.Server` → error handling → log "listening on ..."
2. `GET /health` returning JSON (still in `cmd/api` is fine)
3. Move config into `internal/config` and call it from `main`
4. Move `/health` into `internal/handler` and mount it
5. `GET /api/gyms` with `database/sql` + `lib/pq`
6. `go mod tidy` (all deps are currently `// indirect`)

Each step is verified by running the binary, not by reading code.

## Quiz

1. `main.go` and `api.go` are in the same folder. Do you need to import one from the other? Why or why not?
2. What is the import path for a future package at `internal/config/config.go`, given `go.mod` says `module gienagon-backend`?
3. Why does `http.handler` fail when `http.Handler` works?
4. In `main()`, why does the `PORT` lookup never execute?
5. Which folder would produce the compiled binary, and what command builds it?

<details>
<summary>Answers</summary>

1. No. All `.go` files in one directory form a single package, so they already share the same namespace. Importing is only for crossing *directory* boundaries.
2. `gienagon-backend/internal/config`. The import path is the module path plus the folder path.
3. Go's visibility rule: identifiers beginning with an uppercase letter are exported; lowercase ones are package-private. `Handler` is the exported interface type; `handler` does not exist.
4. `http.ListenAndServe` blocks for the lifetime of the server, so lines after it only run once the server has stopped or failed.
5. `cmd/api/`, because it is `package main` with `func main()`. `go run ./cmd/api` (or `go build ./cmd/api`) targets it.

</details>
