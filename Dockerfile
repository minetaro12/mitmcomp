FROM golang:1.20.0 AS builder

WORKDIR /work
COPY . ./
RUN apt update && \
    apt install --no-install-recommends -y libvips-dev && \
    go build

FROM debian:bullseye-slim
RUN apt update && \
    apt install --no-install-recommends -y libvips ca-certificates && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /work/mitmcomp /app

EXPOSE 8080

ENTRYPOINT ["/app/mitmcomp"]