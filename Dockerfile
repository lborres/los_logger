FROM rust:1.83-alpine AS build

# Install necessary build dependencies for Rust
RUN apk add --no-cache build-base musl-dev openssl-dev pkgconfig

WORKDIR /app

COPY . .

RUN cargo build --release

FROM alpine:latest AS runner

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=build /app/target/release/los_logger .

CMD ["/los_logger"]
