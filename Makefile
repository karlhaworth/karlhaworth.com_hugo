.PHONY: all

serve:
	foreman start -f themes/karlhaworth-com/Procfile.dev

build-all: install-tailwind build-tailwind build-pdf remove-extra-pdf-info build-site

remove-extra-pdf-info:
	sed -z -r -i 's#<div class="number[^<]*</div>\n##g' themes/karlhaworth-com/layouts/partials/header.html

build-site:
	hugo --minify

build-pdf:
	hugo --minify --destination pdf_public -b $(pwd)/pdf_public
	go install
	go run main.go

install-tailwind:
	wget https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -O /usr/local/bin/tailwindcss
	chmod +x /usr/local/bin/tailwindcss

install-tailwind-macos:
	wget https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64 -O /usr/local/bin/tailwindcss
	chmod +x /usr/local/bin/tailwindcss

build-tailwind:
	tailwindcss -i ./themes/karlhaworth-com/assets/main.css -o ./themes/karlhaworth-com/assets/style.css
