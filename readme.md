"Эшелон"

Написать простой сервер gRPC реализующий одну функцию A.Задача функции–проверка значений полей...

Установка:

dep init

(dep ensure -add google.golang.org/gRPC github.com/golang/protobuf/protoc-gen-go)

./build/proto-gen.bat (.sh)


Запуск:

go run .\server\main.go

go run .\client\main.go