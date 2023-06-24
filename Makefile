main:
	go run main.go

compile-wasm:
	env GOOS=js GOARCH=wasm go build -o wasm/revdriller.wasm revdriller