protoc.exe -I $env:GOPATH\src --go_out=$env:GOPATH\src $env:GOPATH\src\echelon\internal\proto-files\domain\repository.proto
protoc.exe -I %GOPATH%\src --go_out=plugins=grpc:%GOPATH%\src %GOPATH%\src\echelon\internal\proto-files\service\repository-service.proto