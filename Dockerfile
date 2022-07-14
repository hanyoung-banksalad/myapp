FROM golang:1.17.10 as golang

ARG GH_ACCESS_TOKEN
RUN git config --global url."https://${GH_ACCESS_TOKEN}@github.com/".insteadOf "https://github.com/"

WORKDIR /myapp
COPY . .

ARG test
RUN if [ "$test" = "true" ]; then make lint ; fi
RUN if [ "$test" = "true" ]; then make test ; fi

WORKDIR /myapp/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/myapp


FROM debian:stable-20220509-slim

RUN apt-get update \
    && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=golang /go/bin /app
ENTRYPOINT ["app/myapp"]
