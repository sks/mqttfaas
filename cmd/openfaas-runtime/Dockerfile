FROM golang AS builder

RUN apt-get update && \
    apt-get install -y gcc-aarch64-linux-gnu

ENV GO111MODULE on

WORKDIR /go/src/github.com/sks/mqttfaas

COPY go.mod go.sum ./

ENV GO111MODULE on

RUN go mod vendor

COPY . .

RUN go build -ldflags "-w -s -linkmode external -extldflags -static" -o /go/bin/openfaas-runtime cmd/openfaas-runtime/main.go

FROM docker:dind

COPY --from=builder /go/bin/openfaas-runtime /bin/openfaas-runtime

COPY cmd/openfaas-runtime/var/data/docker-compose.yml /var/data/docker-compose.yml

COPY faasswarm.tar.gz /var/data/faasswarm.tar.gz

ENTRYPOINT [ "/bin/openfaas-runtime" ]