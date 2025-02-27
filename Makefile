clean:
	rm -rf ./public/*
.PHONY: clean

ssg:
	go run ./main.go
	cp ./static/* ./public
.PHONY: ssg

dev: ssg
	rsync -vr --delete ./public/ pgs.sh:/docs-local
.PHONY: dev

fmt:
	deno fmt
	go fmt ./...
.PHONY: fmt

lint:
	golangci-lint run -E goimports -E godot --timeout 10m
.PHONY: lint
