# Blossom Server Admin UI

SvelteKit SPA served by the Go server at `/admin`. Built to static files and embedded via `//go:embed`.

## Develop

```sh
pnpm install
pnpm dev
```

Dev server proxies `/api`, `/list`, and `/stats` to `http://localhost:8000`. Run the Go server alongside.

## Build

```sh
pnpm build
```

Outputs to `build/`. The Go server embeds this at build time.

## Auth

Two options:
- **NIP-07** — browser extension (nos2x, Alby, etc.).
- **nsec** — paste an nsec. Kept in memory only; never persisted. Reload = relogin.
