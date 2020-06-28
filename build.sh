CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gml-windows.exe -ldflags "-w -s" main.go
gzip -f gml-windows.exe
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gml-linux -ldflags "-w -s" main.go
gzip -f gml-linux.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gml-darwin -ldflags "-w -s" main.go
gzip -f gml-darwin.exe