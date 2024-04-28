build:
	go mod tidy
	go build -o hexstring main.go

install:
	just build
	cp hexstring ~/.local/bin
