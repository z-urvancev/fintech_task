ARG store_type

FROM golang:alpine as builder
WORKDIR /server
COPY . .
RUN go mod tidy
RUN go build -o main cmd/main.go
FROM alpine
WORKDIR /server
COPY --from=builder /server/main /server/main
COPY --from=builder /server/.env /server/.env
COPY --from=builder /server/config/config.yml /server/config/config.yml
CMD ["./main", "--store_type=${store_type}"]
