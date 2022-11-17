## Resume Website for Karl Haworth using Hugo

### Print styles produce PDF resume. Best printed in Chrome.

#### Technologies Utilized

- [Tailwind](http://tailwindcss.com)
- [Sass](https://sass-lang.com)
- [Hugo](http://gohugo.io)
- [Chrome Headless](https://github.com/go-rod/rod)

### Local

#### Prerequisites

- GO_VERSION >= 1.19.2
- NODE_VERSION >= 16.15.1
- CHROME_VERSION >= 103.0
- HUGO_VERSION >= 0.54.0

#### Serve

```bash
make serve
```

#### Build

```bash
make build
```

#### Tailwind Updates

```bash
make tailwind
```

### CloudFlare Publish

```bash
make CF_PAGES_URL="${CF_PAGES_URL}" publish-cloudflare
```

#### ENV Vars for CloudFlare

|Name|Value|
|---|--|
|GO_VERSION|1.19.2|
|NODE_VERSION|16.15.1|