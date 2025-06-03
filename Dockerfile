FROM golang:1.22.4-alpine AS builder

WORKDIR /app
COPY . .
RUN cd cmd/api && go build -o /wellbe-app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /wellbe-app .
COPY configs/app.env .

EXPOSE 8080
CMD ["./wellbe-app"]