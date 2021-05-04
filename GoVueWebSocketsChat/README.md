

# Go Vue WebSockets Chat App

A simple chat web app written in Go, Vue, and WebSockets.


## Build

```
cd ./src
go mod tidy
go run .
```

## Containerization

```
call docker_build.bat
call docker_run.bat
```

## Docker Hub

```
docker tag websockets-go-vue-chat:latest briansu2004/websockets-go-vue-chat:latest

docker push briansu2004/websockets-go-vue-chat:latest

docker pull briansu2004/websockets-go-vue-chat:latest

(later)
docker run -it --rm -p 8000:8000 briansu2004/websockets-go-vue-chat:latest
```

## Chat

http://localhost:8000

![chat.png](doc/chat.png)


## Feedback

Please drop a note to BrianSu2004@hotmail.com for any guides, suggestions, recommendations, ideas, and comments etc. Thanks!

