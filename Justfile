build:
	go mod tidy
	go build main.go -o hexstring

install:
	just build
	cp hexstring ~/.local/bin
