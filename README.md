## Resume Website for Karl Haworth using Hugo

### Local

#### Serve

```bash
hugo serve
```

#### Build

```bash
hugo build
```

#### Tailwind Updates

```bash
cd themes/karlhaworth-com && npm run build-tw
```

### CloudFlare Build

```bash
cd themes/karlhaworth-com && npm install && npm run build-tw && cd ../../ && hugo -b $CF_PAGES_URL
```