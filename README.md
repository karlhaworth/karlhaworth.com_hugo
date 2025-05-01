# Project Overview

This project allows a user to generate/build a resume (CV) PDF from website content and host a resume website put together from json files using Hugo. This reduces the duplication needed to keep a separate resume from website content updated. Sections can be hidden using certain key-value pairs. 

The resume PDF rebuilds each time there is a change and can used on a wide array of free static hosting providers such as GitHub Pages, CloudFlare Pages, Netlify, and more!

This project shifted from using provider runners for build to github actions as the software stack is a little more reliable and predictable. Using headless Chromium for resume pdf generation seemed to cause issues with CloudFlare runners. 

There was a decision to keep the website as light as possible, therefore frameworks such as React were not utilized. Lightbox and jQuery only load on the pages they are utilized. 

## Resume Website for Karl Haworth using Hugo

Resume website for Karl Haworth using [Hugo](http://gohugo.io) as a static site generator (SSG) with [Tailwind](http://tailwindcss.com) and [Sass](https://sass-lang.com) for styling. [Go-Rod](http://go-rod.github.io) is used to create a PDF from the generated content which serves as a resume using print media styling. A makefile is used to simplify the commands needed to generate the content.

Using GitHub as a code repository and [CloudFlare Pages](https://pages.cloudflare.com) to build and host this website. ($0)
Using [CloudFlare DNS](https://www.cloudflare.com/dns/). ($0)
Using [CloudFlare Domains](https://www.cloudflare.com/products/registrar/) to purchase domain at cost. (~$8/yr)

### Intent

For myself, this creates a marketable and favorable image during search results. The generation allows resume content to be intermingled with other content that might not be found elsewhere, while offering an easy print layout that serves as a resume.

The intent is to keep this simple to update and host as well as cheap to run.

### Print styles produce PDF resume. Best printed in Chrome.

### Local

- Go - `brew install go`
- Hugo - `brew install hugo`
- Foreman - `brew install foreman`

#### Serving Locally

The below will run foreman to serve Hugo, build CSS, build PDF, and convert SVGs to fonts while watching files for changes.

```bash
make serve
```

### Building for Publishing

This is performed in the GitHub Workflow and builds CSS, a PDF, converts SVGs, and builds using Hugo

```bash
make build-all
```
