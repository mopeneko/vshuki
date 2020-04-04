FROM golang:1.14-alpine3.11 as builder

ENV PROJECT_WORKDIR /go/src/vshuki/

WORKDIR ${PROJECT_WORKDIR}

COPY go.mod go.sum ${PROJECT_WORKDIR}
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -o vshuki


FROM alpine:3.11

COPY --from=builder /go/src/vshuki/vshuki .

ENTRYPOINT [ "./vshuki" ]
