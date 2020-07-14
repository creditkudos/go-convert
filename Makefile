.PHONY: dist

build:
	docker-compose build go_convert

sh:
	docker-compose run go_convert sh

get-protos:
	docker-compose run go_convert go get github.com/creditkudos/protos