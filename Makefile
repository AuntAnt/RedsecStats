#Build the app

.PHONY: all build clean

all: build

build:
	@echo "Building..."
	make build_macos
	make build_linux
	make build_win

build_linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/redsec_stats_linux ./src

build_macos:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./bin/redsec_stats_macos ./src
	
build_win:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/redsec_stats_win.exe ./src
	
clean:
	rm ./bin/*

