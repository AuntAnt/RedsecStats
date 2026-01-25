#Build the app

.PHONY: all build clean

all: build

build:
	@echo "Building..."
	make build_macos
	make build_linux
	make build_win

build_linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/redsec_revives_linux ./src

build_macos:
	GOOS=darwin GOARCH=arm64 go build -o ./bin/redsec_revives_macos ./src
	
build_win:
	GOOS=windows GOARCH=amd64 go build -o ./bin/redsec_revives_win.exe ./src
	
clean:
	rm ./bin/*

