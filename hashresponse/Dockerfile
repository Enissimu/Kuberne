from golang:1.24.3-alpine as builder

workdir /app
copy go.mod go.sum ./
run go mod download

copy . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 3000
CMD ["./main"]