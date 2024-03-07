FROM golang:1.21.2 as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0  go build -o cmd ./cmd

FROM alpine:3.14.10
COPY --from=builder /app/cmd /server
ENTRYPOINT ["/server"]