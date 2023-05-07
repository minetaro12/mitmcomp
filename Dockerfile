FROM golang:1.20.4 AS builder

WORKDIR /work
COPY . ./
RUN  go build -o main

FROM gcr.io/distroless/base:latest
WORKDIR /app
COPY --from=builder /work/main /app

EXPOSE 8080

ENTRYPOINT ["/app/main"]