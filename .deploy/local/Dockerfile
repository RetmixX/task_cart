FROM golang:1.24 AS builder


WORKDIR /app

COPY . .


RUN go get && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go


FROM alpine:latest


WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/.deploy/local/.env .deploy/local/.env
COPY --from=builder /app/migrations/ migrations/


RUN apk add --no-cache libc6-compat

#RUN ./server migrations up

#CMD ["./server"]