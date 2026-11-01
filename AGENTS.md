# AGENTS.md — Universal Agent Rules (Giēnagon-Go)

Source of truth (external): `~/agent-rules-pack/AGENTS.md` and `~/Documents/shamii/AGENTS.md`.
This copy exists in-repo so contributors and tools can follow the rules locally.

---

## 1. Base rules (always)
- Match existing conventions — read neighboring code first, don't introduce new libs/patterns
- Minimal changes — edit existing files, don't refactor unrequested things
- Verify before "done" — run typecheck/build/tests after non-trivial changes, report failures honestly
- Plan before touching — for multi-file changes, state a 2–4 bullet plan first
- No destructive actions — never push, reset, drop data, or delete files unprompted
- Say what you don't know — ask when ambiguous, don't guess on hard-to-reverse things

## 2. Explainer docs (after every meaningful change)
Write a doc at `docs/explainers/YYYY-MM-DD-<short-slug>.md` with:
1. Background — how the system worked before
2. Intuition before details — essence of the change + analogy
3. (Optional) Interactive figure — standalone HTML if it genuinely helps
4. Literate code diff — logical order with prose before each file
5. Quiz — exactly 5 medium-difficulty questions, answers in a <details> block

## 3. The quiz rule (your personal rule)
No code ships to review/production until you can pass the quiz in the explainer. Agents must remind you the explainer + quiz exists before calling a task "done".

## 4. Microworlds (on demand or when stuck)
When you ask to "explain X" or are confused, agents can build a microworld — a throwaway visual UI for that one concept. Lives in `microworlds/`, plain single-file HTML/JS.

## 5. Explainer housekeeping
Explainers commit alongside their change. If a later change invalidates an older one, note it at the top — don't delete.

---

## AI-as-reviewer operating mode (how we use assistants here)

We follow the learning-first workflow while porting from TypeScript/Express to Go:

1) Port one small thing yourself first, then get a review
- Take a single route (e.g., `/health`, `login`) and implement it in Go.
- Ask: "Review this like a senior Go dev. What's non-idiomatic and why?"

2) Ask for hints, not solutions
- Ask for which stdlib/package to consult first. Escalate only if stuck: hint → outline → code.

3) Use compiler errors as lessons
- Paste error messages and ask: "What does this mean and why does Go care?"

4) Don't port line by line
- Embrace Go patterns: `if err != nil`, structs + functions (not classes), goroutines (not async/await).

5) Predict, then run
- State expected output/behavior before running. Compare after.

6) Rebuild from memory
- The next day, redo without looking. If you can't, revisit the explanation.

7) Keep a TS → Go cheat sheet
- Use `notes.md` in the repo. Add a row whenever something clicks.

8) Use your old API as the answer key
- Run Node and Go versions side by side. Same input → same JSON. Debug diffs.

## Tutor prompt to reuse
> I'm porting this Express route to Go to learn. Don't write the full solution. First, list the Go concepts I'll need (just names). Then I'll attempt it and you review.

