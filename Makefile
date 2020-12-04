tidy:
	go mod tidy

# ==============================================================================
# Build
build-linux:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/godataimport cmd/main.go

build-macos:
	export GO111MODULE=on
	env GOOS=darwin go build -o bin/godataimport cmd/main.go

build-windows:
	export GO111MODULE=on
	env GOOS=windows GOARCH=386 go build -o bin/godataimport cmd/main.go

clean:
	rm -rf ./bin Gopkg.lock
