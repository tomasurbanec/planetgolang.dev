build:
	go build -o bin/crawler .

scrape: build
	bin/crawler scrape

generate: build
	rm -rf dist/*
	mkdir -p dist
	bin/crawler generate
	cp -r static/* dist/

upload:
	aws s3 rm --recursive --profile planetgolang "s3://planetgolang.dev/"
	aws s3 cp --recursive --profile planetgolang dist "s3://planetgolang.dev" --acl bucket-owner-full-control

deploy: scrape generate upload

invalidate:
	aws cloudfront create-invalidation --distribution-id ELHTE4P8I823B --paths "/*" --profile planetgolang
