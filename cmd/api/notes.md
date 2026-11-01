#I think I'm falling in love with neovim again and also the direction this repo is steering in that 
#this repo the commments the queestions I'm asking myself show who I am and how i think when sandboxed to new environments


# TS → Go cheat sheet

Add entries as you port. Examples to start:

```
process.env.PORT         → os.Getenv("PORT")
app.use(cors())          → r.Use(cors.Handler(...))
try/catch                → if err != nil { ... }
setInterval              → time.NewTicker + goroutine
import 'dotenv/config'   → godotenv.Load()
fetch(...)               → net/http (Client, Request) or resty
Express middleware       → http.Handler / middleware funcs
class + methods          → struct + funcs (methods on pointer receivers when mutating)
await db.query(...)      → database/sql (QueryContext, Scan) or GORM/pgx
Promise.all              → multiple goroutines + sync.WaitGroup or errgroup

// Language
null/undefined           → nil
=== / !==                → Go has typed equality; beware nil interfaces
JSON.parse/stringify     → encoding/json (Unmarshal, Marshal)
Type guards              → type assertions, type switches
```

Tips:
- Prefer stdlib first; add deps only with clear payoff.
- Keep handlers small; push logic into functions you can unit test.


Here are some of the problems i have encountered as am tyring to port this api to golang

#i dont know the what routing layer to use betweeen chi and mux am also wondering  why we are not using the context package upfront 
#either way i was not using the ctx variable i declared on the  `context.Withcancel` but i need to know why we dont have the *response writer
of the http request handler being utilized by mux

noticed that the `mux` is now a native part of the go std library ... as of go 1.22. Previously it was part of the gorilla mux
