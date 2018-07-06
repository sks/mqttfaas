FROM golang:alpine AS builder

WORKDIR /go/src/main

COPY main.go /go/src/main/gocat.go

RUN CGO_ENABLED=false go build -o /bin/gocatter .

FROM scratch 

LABEL mqtt_faas="go_catter" mqtt_faas_topic="/cat/input|/wildcat/input/#"

COPY --from=builder /bin/gocatter /bin/gocatter

ENTRYPOINT [ "/bin/gocatter" ]