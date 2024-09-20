## Generate proto buffer code

    protoc -I=./api --go_out=. --go-grpc_out=. movie.proto