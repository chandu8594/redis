# Golang Redis
- Below instructions are for mac/ linux based system

## Prerequisite:
- Install [golang](https://golang.org/doc/install) 1.12 or above
- Install [docker](https://docs.docker.com/get-docker/) and get the docker running
- Pull latest redis image and start the container using below commands in the terminal
```bash
docker pull redis
docker run --name redis-test-instance -p 6379:6379 -d redis
```
- Execute main.go
```bash
go run main.go
```