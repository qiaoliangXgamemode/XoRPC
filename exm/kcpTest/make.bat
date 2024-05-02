set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

go build ./server.go


SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
go build clientA.go