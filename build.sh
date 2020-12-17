LDFLAGS="-X 'main.buildDate=$(date)' -X 'main.gitHash=$(git rev-parse HEAD) -X 'main.buildOn=$(go version)' -w -s "

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gml-windows.exe -trimpath -ldflags "${LDFLAGS}" main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gml-linux -trimpath -ldflags "${LDFLAGS}" main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gml-darwin -trimpath -ldflags "${LDFLAGS}" main.go

# sha256
sha256sum gml* > gml-sha256
cat gml-sha256

# gzip
gzip --best gml*