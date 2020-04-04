FROM golang:1.14-alpine3.11 as builder

WORKDIR /go/src/vshuki
COPY . .

ENV CGO_ENABLED=0

RUN go build


FROM alpine:3.11

COPY --from=builder /go/src/vshuki/vshuki .

ENTRYPOINT [ "./vshuki" ]
