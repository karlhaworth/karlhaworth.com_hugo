.PHONY: all

serve:
	hugo serve

install:
	cd ./themes/karlhaworth-com && \
	npm install

build-all: install build-tailwind build-pdf remove-extra-pdf-info build-site

remove-extra-pdf-info:
	sed -z -r -i 's#<div class="number[^<]*</div>\n##g' themes/karlhaworth-com/layouts/partials/header.html

build-site:
	hugo --minify

build-pdf:
	hugo --minify --destination pdf_public -b $(pwd)/pdf_public
	go install
	go run main.go

build-tailwind:
	cd ./themes/karlhaworth-com && \
	npm run build-tw
