# Golang Redis
- Below instructions are for mac/ linux based system

## Prerequisite:
- Install [golang](https://golang.org/doc/install) 1.12 or above
- Install [docker](https://docs.docker.com/get-docker/) and get the docker running
- Start the redis container using `docker-compose.yaml`
```bash
docker-compose up -d redis
```
- Execute main.go
```bash
go run main.go
```