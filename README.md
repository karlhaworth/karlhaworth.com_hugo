# Project Overview

A Hugo-based resume website that generates a printable PDF from site content using Go + go-rod. The site keeps resume data in JSON and reuses it for both the website and a PDF resume to avoid duplication.

Quick links:
- Local serve: `make serve` (uses Procfile.dev)
- Build for publish: `make build-all` (used by GitHub Actions)

Prerequisites
- Go (>= 1.21): `brew install go` or from https://golang.org
- Hugo: `brew install hugo` (ensure the version supports SCSS if using it)
- Node.js (>=18) + npm or pnpm: for tailwind, svgo, svgtofont
- Chrome/Chromium: required by go-rod for PDF generation
- Foreman or goreman: used to run Procfile.dev (gem: `gem install foreman` or Go port: `go install github.com/mattn/goreman@latest`)

Recommended local setup
1. Install system deps (example macOS):
   brew install go hugo node
   go install github.com/mattn/goreman@latest
2. Install Node tooling (project-optional):
   npm install -g svgo svgtofont  # or use npx as in Makefile

Serving locally
- Use the Makefile (recommended):
  make serve
- Or run Procfile.dev directly with goreman/foreman:
  goreman start -f Procfile.dev
- To run Hugo-only server:
  make hugo-serve

Building for publishing
- CI (GitHub Actions) runs the same steps as `make build-all` which:
  - builds Tailwind CSS
  - converts SVGs to fonts
  - generates the site via Hugo
  - produces a PDF via go-rod

Notes & tips
- PDF generation requires Chrome/Chromium available in CI or local machine — GitHub Actions is used to avoid flaky third-party runners.
- Tailwind is built by the Makefile; you can also manage it via npm if you prefer (add package.json and devDependencies).

Hosting (cheap)
- Primary: Cloudflare Pages (free) — connect the GitHub repo and use the existing GitHub Actions or Pages' native build. Cloudflare Pages + Cloudflare DNS keeps hosting costs at $0.
- Alternatives: GitHub Pages and Netlify both offer free static hosting and are compatible with this Hugo site.
- Domain: Use Cloudflare Registrar (~$8/yr) or any registrar; point DNS to Cloudflare for free DNS management.

Cost summary: Hosting and DNS = $0 (Pages + Cloudflare DNS), Domain ≈ $8/yr (optional).

If any tooling preference should be added (npm/pnpm, goreman vs foreman), say which and the README will be adjusted.