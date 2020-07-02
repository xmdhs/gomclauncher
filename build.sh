CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gml-windows.exe -trimpath -ldflags "-w -s" main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gml-linux -trimpath -ldflags "-w -s" main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gml-darwin -trimpath -ldflags "-w -s" main.go
