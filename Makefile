.PHONY: all

serve:
	hugo serve

build:
	hugo build

build-cloudflare: tailwind-install tailwind-build
	hugo

tailwind-install:
	cd ./themes/karlhaworth-com && \
	npm install

tailwind-build:
	cd ./themes/karlhaworth-com && \
	npm run build-tw