# Build stage
FROM golang:1.21-alpine3.18 AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration
RUN chmod +x *.sh && chmod +x main

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
