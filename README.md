# myapp

[![Actions Status](https://github.com/banksalad/myapp/workflows/ci/badge.svg)](https://github.com/banksalad/myapp/actions) ![Golang Badge](https://badgen.net/badge/Language/Go/cyan) ![GRPC Badge](https://badgen.net/badge/Use/gRPC/blue)

> first go-rpc app

## Docs
<!-- TODO: Update to the actual document url -->
- [TechSpec]()
- [Go in Banksalad](https://docs.google.com/document/d/1mPUGKlfA6pFLMUuUCHv54ejnUDrrldJ5z06AbvinRQA)

## Getting Started

### Start a Server
```sh
$ git config --global url."https://${GITHUB_ACCESS_TOKEN}@github.com/".insteadOf "https://github.com/"  # insert your github access token
$ make init
$ make run

# Use Docker
$ docker build --build-arg GH_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN} --tag myapp .  # insert your github access token
$ docker run --rm -p 18081:18081 -p 18082:18082 myapp

# Use Onebox
$ make deploy-to-local-k8s
```

### Test & Lint
```sh
$ make test

$ make lint
```

## APIs
<!-- TODO: Update to actual urls -->
- [myapp.proto](https://github.com/banksalad/idl/blob/master/protos/apis/v1/myapp/myapp.proto)
- [myapp.swagger.json](https://github.com/banksalad/idl/blob/master/gen/swagger/apis/v1/myapp/myapp.swagger.json)

## Directory Structure
```
.
├── client.go # dependency service 들에 대한 client
├── cmd       # server를 실행시키기 위한 main.go
│   └── ...
├── config    # 설정 파일
│   └── ...
└── server
    ├── grpc_server.go  # main gRPC server
    ├── http_server.go  # HTTP <-> gRPC 변환해주는 grpc-gateway layer
    └── handler         # gRPC server handlers
```
