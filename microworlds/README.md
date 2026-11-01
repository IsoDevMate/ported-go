# Microworlds

- Purpose: tiny, throwaway single-file HTML/JS pages that illustrate exactly one concept (e.g., goroutines + channels visualization).
- Location: `microworlds/`
- Naming: `YYYY-MM-DD-concept.html`
- Keep them self-contained (no build steps).

Starter HTML you can copy:

```html
<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Microworld — concept-name</title>
    <style>
      body { font-family: system-ui, sans-serif; margin: 2rem; }
      .box { padding: 1rem; border: 1px solid #ccc; border-radius: 8px; }
    </style>
  </head>
  <body>
    <h1>Concept name</h1>
    <div class="box" id="app">Interactive area</div>
    <script>
      // Minimal JS demo here.
      // Keep in one file so it's easy to share and delete later.
    </script>
  </body>
</html>
```
