FROM golang:alpine3.18 AS builder

COPY app/ /go/app

RUN go build ./app/hello-world.go

FROM scratch

COPY --from=builder /go/app/ /go/app

ENTRYPOINT [ "./go/app/hello-world"]