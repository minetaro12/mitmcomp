FROM golang:1.20.0 AS builder

WORKDIR /work
COPY . ./
RUN  go build

FROM gcr.io/distroless/base:latest
WORKDIR /app
COPY --from=builder /work/mitmcomp /app

EXPOSE 8080

ENTRYPOINT ["/app/mitmcomp"]