# AI-as-reviewer prompts (use these during the port)

## Senior review (after you code)
Paste your Go code and ask:

> Review this like a senior Go dev. What's non-idiomatic and why? What would you change and why? No full rewrite unless requested.

## Hints, not solutions
When stuck:

> Don't give code yet. Which standard library packages or concepts should I look at for this problem?

If still stuck:

> Give me a higher-level outline or pseudocode.

Then, only if needed:

> Now show a minimal code example.

## Error explanations
When the compiler yells:

> Explain what this error means and why Go cares. What's the minimal change to fix it, and what are the tradeoffs?

## Predict, then run
Before running a new handler/test:

> Here's what I expect to happen: ... What are 1–2 things that could surprise me?

## Diff-check against TS
When validating parity with the Node API:

> Here is the Node response and the Go response for the same input. Help me spot the semantic differences and likely root causes.
