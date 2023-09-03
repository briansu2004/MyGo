# Go WebSockets

## Build

```dos
go mod tidy

go run main.go
```

http://localhost:9080/

![local.png](doc/local.png)

start client_local.html

![client.png](doc/client.png)

![connected.png](doc/connected.png)

## Containerization

```dos
call docker_build.bat
call docker_run.bat
```

![docker_run.png](doc/docker_run.png)

## Docker Hub

```dos
docker tag websockets-go-1st:latest briansu2004/websockets-go-1st:latest

docker push briansu2004/websockets-go-1st:latest

docker pull briansu2004/websockets-go-1st:latest

(later)
docker run -it --rm -p 9080:9080 briansu2004/websockets-go-1st:latest
```

## Chrome DevTools

```dos
let ws = new WebSocket("ws://localhost:9080/ws")

ws.send("Nice")

ws.send(JSON.stringify({data: "test mssage", type: "test"}))
```

![devtools.png](doc/devtools.png)
