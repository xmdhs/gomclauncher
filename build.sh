CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gml-windows.exe -ldflags "-w -s" main.go
gzip gml-windows.exe
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gml-linux.exe -ldflags "-w -s" main.go
gzip gml-linux.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gml-darwin.exe -ldflags "-w -s" main.go
gzip gml-darwin.exe