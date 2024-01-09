clean:
	rm -rf ./public/*
	echo "" > ./public/.gitkeep
.PHONY: clean

ssg:
	go run ./main.go
	cp ./static/* ./public
.PHONY: ssg

dev: ssg
	rsync -vr ./public/ hey@pgs.sh:/docs-local
.PHONY: dev
