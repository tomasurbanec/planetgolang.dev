build:
	go build -o bin/crawler .

scrape:
	bin/crawler scrape

generate:
	rm -rf dist/*
	mkdir -p dist
	bin/crawler generate
	cp -r static/* dist/

upload:
	aws s3 rm --recursive --profile planetgolang-data-updater "s3://planetgolang/"
	aws s3 cp --recursive --profile planetgolang-data-updater dist "s3://planetgolang" --acl bucket-owner-full-control

deploy: scrape generate upload

invalidate:
	aws cloudfront create-invalidation --distribution-id ELHTE4P8I823B --paths "/*" --profile planetgolang
