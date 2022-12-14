.PHONY: all

serve:
	hugo serve

publish-cloudflare: npm-modules-install tailwind-build build-cloudflare build-pdf

build-cloudflare:
	hugo --minify -b ${CF_PAGES_URL}

build-pdf:
	PWD=$(pwd)
	hugo --minify --destination pdf_public -b ${PWD}/pdf_public
	go install
	go run main.go

# Tailwind and Lightbox2
npm-modules-install:
	cd ./themes/karlhaworth-com && \
	npm install

tailwind-build:
	cd ./themes/karlhaworth-com && \
	npm run build-tw
