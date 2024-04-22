clean:
	rm -rf ./public/*
	echo "" > ./public/.gitkeep
.PHONY: clean

ssg:
	go run ./main.go
	cp ./static/* ./public
.PHONY: ssg

dev: ssg
	rsync -vr ./public/ pgs.sh:/docs-local
.PHONY: dev

fmt:
	deno fmt
	go fmt ./...
.PHONY: fmt

lint:
	golangci-lint run -E goimports -E godot --timeout 10m
.PHONY: lint
