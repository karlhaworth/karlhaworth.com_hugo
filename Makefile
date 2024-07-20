.PHONY: all

serve:
	hugo serve

install:
	cd ./themes/karlhaworth-com && \
	npm install

build-all: install build-tailwind build-site build-pdf

build-site:
	hugo --minify

build-pdf:
	PWD=$(pwd)
	hugo --minify --destination pdf_public -b ${PWD}/pdf_public
	go install
	go run main.go

build-tailwind:
	cd ./themes/karlhaworth-com && \
	npm run build-tw
