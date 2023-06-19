FROM golang:1.19 AS builder
WORKDIR /dws/
COPY main.go go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o dummy-web-service

FROM busybox:1.36
COPY --from=builder /dws/dummy-web-service /
ENTRYPOINT ["/dummy-web-service"]
CMD [""]
