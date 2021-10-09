build:
	gotip build -o example-embed github.com/pieterclaerhout/example-embed

run: build
	./example-embed
