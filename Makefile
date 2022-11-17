.PHONY: all

serve:
	hugo serve

build:
	hugo build

build-cloudflare: tailwind
	hugo -b $CF_PAGES_URL

tailwind:
	cd ./themes/karlhaworth-com && \
	npm install && \
	npm run build-tw