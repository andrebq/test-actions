FROM golang:1.13 as build

WORKDIR /go/src/app
COPY go.mod go.sum /go/src/app/

RUN go mod download

COPY . /go/src/app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app/call-wh ./call-wh && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app/dummy-webhook .

FROM alpine

COPY --from=build /go/bin/app/call-wh /usr/local/bin/
COPY --from=build /go/bin/app/dummy-webhook /usr/local/bin/

CMD ["/usr/local/bin/call-wh"]
