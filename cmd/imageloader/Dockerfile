FROM golang AS builder

RUN apt-get update && \
    apt-get install -y gcc-aarch64-linux-gnu

ARG GOARCH=amd64
ARG CC

ENV GO111MODULE on

WORKDIR /go/src/github.com/sks/mqttfaas

COPY go.mod go.sum ./

ENV GO111MODULE on

RUN go mod vendor

COPY . .

RUN CC=${CC} CGO_ENABLED=0 GOARCH=${GOARCH} \
    go build -ldflags "-w -s -linkmode external -extldflags -static" -o /go/bin/imageloader ./cmd/imageloader

ARG BASE_IMAGE
FROM alpine

ARG VERSION

ENV MQTT_FAAS_VERSION ${VERSION:-v-DEV}

EXPOSE 8000

COPY --from=builder /go/bin/imageloader /bin/imageloader

ENTRYPOINT [ "/bin/imageloader" ]