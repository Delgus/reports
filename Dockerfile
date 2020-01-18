### pre-built
FROM golang:1.13 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux make build

### report1
FROM alpine:latest as report1
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bin/report1 .
CMD ["./report1"]

### report2
FROM alpine:latest as report2
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bin/report2 .
CMD ["./report2"]