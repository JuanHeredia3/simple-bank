# Build stage
FROM golang:1.25-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY wait-for.sh /app/wait-for.sh
COPY start.sh /app/start.sh
COPY db/migration ./db/migration

RUN chmod +x /app/start.sh /app/wait-for.sh

EXPOSE 8080 9090
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]