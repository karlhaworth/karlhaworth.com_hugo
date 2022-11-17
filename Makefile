.PHONY: all

serve:
	hugo serve

build:
	hugo build

build-cloudflare: tailwind-install tailwind-build
	hugo --minify -b ${CF_PAGES_URL}

tailwind-install:
	cd ./themes/karlhaworth-com && \
	npm install

tailwind-build:
	cd ./themes/karlhaworth-com && \
	npm run build-tw

make-pdf:
	go install
	go run main.go