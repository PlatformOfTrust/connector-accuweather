FROM golang:1.15.2
WORKDIR /go/src/github.com/PlatformOfTrust/connector-accuweather/
COPY ./ .
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GO111MODULE on
RUN go get 
RUN go test ./...
RUN go build -v -o ./main

FROM alpine:latest
WORKDIR /connector/
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=0 /go/src/github.com/PlatformOfTrust/connector-accuweather/main ./main
EXPOSE 8080

CMD [ "./main"]
