## Resume Website for Karl Haworth using Hugo

Resume website for Karl Haworth using [Hugo](http://gohugo.io) as a static site generator (SSG) with [Tailwind](http://tailwindcss.com) and [Sass](https://sass-lang.com) for styling. [Go-Rod](http://go-rod.github.io) is used to create a PDF from the generated content which serves as a resume using print media styling. A makefile is used to simplify the commands needed to generate the content.

Using GitHub as a code repository and [CloudFlare Pages](https://pages.cloudflare.com) to build and host this website. ($0)
Using [CloudFlare DNS](https://www.cloudflare.com/dns/). ($0)
Using [CloudFlare Domains](https://www.cloudflare.com/products/registrar/) to purchase domain at cost. (~$8/yr)

### Intent

For myself, this creates a marketable and favorable image during search results. The generation allows resume content to be intermingled with other content that might not be found elsewhere, while offering an easy print layout that serves as a resume.

The intent is to keep this simple to update and host as well as cheap to run.

#### Move from 11ty

Previously a static stie generator named [11ty](https://www.11ty.dev) was used. With 1tty, the image processing was extra, the scripts grew and were not friendly, and the freedom of choice with 11ty made things more complicated than they needed to be. Hugo seemed to provide a more opinionated route making for less scripts which resulted in what feels like a better project structure and less pieces to manage.

Bootstrap was used however many hacks were needed to make it appear the way I wanted, while tailwind provided less opinionated defaults.

### Print styles produce PDF resume. Best printed in Chrome.

### Local

#### Prerequisites

- GO_VERSION >= 1.19.2
- NODE_VERSION >= 16.15.1
- CHROME_VERSION >= 103.0
- HUGO_VERSION >= 0.109.0

#### Serve

```bash
make serve
```

#### Build Local

```bash
hugo
```

#### Build PDF

```bash
make build-pdf
```

#### Tailwind Updates

```bash
make tailwind
```

### CloudFlare Publish added to CloudFlare pages config

```bash
make CF_PAGES_URL="${CF_PAGES_URL}" publish-cloudflare
```

#### ENV Vars for CloudFlare

|Name|Value|Env|
|---|---|---|
|GO_VERSION|1.19.2|ALL|
|NODE_VERSION|16.15.1|ALL|
|CF_PAGES_URL|https://karlhaworth.com|PROD|
|HUGO_VERSION|0.109.0|ALL|
