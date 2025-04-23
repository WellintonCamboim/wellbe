FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
RUN cd cmd/api && go build -o /wellbe-app && cd ../..

FROM alpine:latest
WORKDIR /app
COPY --from=builder /wellbe-app .
COPY configs/app.env .

EXPOSE 8000
CMD ["./wellbe-app"]