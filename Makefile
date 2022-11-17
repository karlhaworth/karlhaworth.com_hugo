.PHONY: all

serve:
	hugo serve

build:
	hugo build

build-cloudflare: tailwind-install tailwind-build
	hugo --minify -b ${CF_PAGES_URL}
	hugo --minify --destination pdf_public
	pwd
	cat ./pdf_public/index.html
	go install
	go run main.go

tailwind-install:
	cd ./themes/karlhaworth-com && \
	npm install

tailwind-build:
	cd ./themes/karlhaworth-com && \
	npm run build-tw

make-pdf:
	hugo --minify --destination pdf_public
	go install
	go run main.go