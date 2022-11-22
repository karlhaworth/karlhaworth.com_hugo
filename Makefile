.PHONY: all

serve:
	hugo serve

build:
	hugo

build-cloudflare:
	hugo --minify
	# hugo --minify -b ${CF_PAGES_URL}

publish-cloudflare: tailwind-install tailwind-build build-cloudflare make-pdf

tailwind-install:
	cd ./themes/karlhaworth-com && \
	npm install

tailwind-build:
	cd ./themes/karlhaworth-com && \
	npm run build-tw

make-pdf:
	PWD=$(pwd)
	hugo --minify --destination pdf_public -b ${PWD}/pdf_public
	go install
	go run main.go