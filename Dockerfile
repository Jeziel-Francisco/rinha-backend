FROM golang:1.22.1 as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0  go build -o cmd ./cmd

FROM alpine:3.14.10
COPY --from=builder /app/cmd /server
ENTRYPOINT ["/server"]