## installs templ command-line tool and gets templ package
deps:
	go install github.com/a-h/templ/cmd/templ@latest
	go get github.com/a-h/templ

## generate templ Go file and binary
build:
	templ generate
	go build -o bin/todo *.go

## run the app
run: build
	./bin/todo
