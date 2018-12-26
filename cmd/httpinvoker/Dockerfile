FROM golang AS builder

RUN apt-get update && \
    apt-get install -y gcc-aarch64-linux-gnu

ENV GO111MODULE on

WORKDIR /go/src/github.com/sks/mqttfaas

COPY go.mod go.sum ./

ENV GO111MODULE on

RUN go mod vendor

COPY . .

RUN go build -ldflags "-w -s -linkmode external -extldflags -static" -o /go/bin/httpinvoker ./cmd/httpinvoker

FROM alpine

COPY --from=builder /go/bin/httpinvoker /bin/httpinvoker

ENTRYPOINT [ "/bin/httpinvoker" ]