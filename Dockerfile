FROM golang:alpine  AS builder

RUN apk update && \
    apk add curl git build-base && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/sks/mqttfaas

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

COPY . .

RUN CGO_ENABLED=0 go build \
    -ldflags "-w -s -linkmode external -extldflags -static" \
    -o /tmp/mqttfaas ./cmd/mqttfaas

FROM scratch

ARG VERSION

ENV MQTTFAAS_VERSION ${VERSION:-v-DEV}

ENV DOCKER_API_VERSION "1.37"

COPY --from=builder /tmp/mqttfaas /bin/mqttfaas

ENTRYPOINT [ "/bin/mqttfaas" ]
