.PHONY: all

serve:
	hugo serve

build:
	hugo build

build-cloudflare: tailwind
	hugo

tailwind:
	cd ./themes/karlhaworth-com && \
	npm install && \
	npm run build-tw