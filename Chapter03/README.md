## Start consul registry service

    docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0

## Adding additional instances of each service

    go run main.go --port <PORT>

## Testing API

    curl -v localhost:8083/movie?id=1