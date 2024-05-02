set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

go build -o ./debug/test_server ./test_server.go
go build  -o ./debug/test_node ./test_node.go


SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
go build -o ./debug/test_node.exe ./test_node.g
go build  -o ./debug/test_server.exe test_server.go