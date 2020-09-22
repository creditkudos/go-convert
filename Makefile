.PHONY: dist

build:
	docker-compose build go_convert

get-protos:
	docker-compose run go_convert go get github.com/creditkudos/protos