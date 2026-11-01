# Explainers: how to write them

- Create files as `docs/explainers/YYYY-MM-DD-<short-slug>.md`.
- Prefer small, focused explainers per meaningful change.
- Keep prose first, code second. Treat it like a literate walkthrough.

Template sections (see `TEMPLATE.md`):
1. Background — what existed before
2. Intuition — the core idea in plain language
3. Optional interactive figure — link to a `microworlds/*.html` if helpful
4. Literate code diff — snippets + explanations in logical order
5. Quiz — exactly 5 questions; answers hidden in a `<details>` block

Validation rule: You should be able to pass the quiz before shipping the change.
