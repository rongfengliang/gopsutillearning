build: clean compile
clean: 
	rm -rf pprof-demo-*
compile:
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o pprof-demo-mac main.go
	GOOS=linux GOARCH=amd64 go build -o pprof-demo-linux main.go
	GOOS=windows GOARCH=amd64 go build -o pprof-demo-windows.exe main.go