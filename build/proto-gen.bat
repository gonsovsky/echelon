protoc.exe -I %GOPATH%\src --go_out=%GOPATH%\src %GOPATH%\src\echelon\service\proto-files\record.proto
protoc.exe -I %GOPATH%\src --go_out=plugins=grpc:%GOPATH%\src %GOPATH%\src\echelon\service\proto-files\filter-service.proto