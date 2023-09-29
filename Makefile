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
	aws s3 rm --recursive --profile planetgolang-data-updater "s3://www.planetgolang.dev/"
	aws s3 cp --recursive --profile planetgolang-data-updater dist "s3://www.planetgolang.dev" --acl bucket-owner-full-control

deploy: scrape generate upload

invalidate:
	aws cloudfront create-invalidation --distribution-id E2DG4KJ4HYJIJ4 --paths "/*" --profile planetgolang-data-updater
