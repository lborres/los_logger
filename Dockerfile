FROM rust:1.83-alpine AS build

# Install necessary build dependencies for Rust
RUN apk add --no-cache build-base musl-dev openssl-dev pkgconfig tzdata

# Set the working directory
WORKDIR /app

# Copy source code and build
COPY . .
RUN cargo build --release

# Final image
FROM alpine:latest AS runner

# Install runtime dependencies and timezone data
ENV TZ=Asia/Manila
RUN apk add --no-cache ca-certificates tzdata

# Create a non-root user
ARG USERNAME=appuser
ARG USER_ID=1000
ARG GROUP_ID=1000

RUN addgroup -g ${GROUP_ID} ${USERNAME} \
    && adduser -D -u ${USER_ID} -G ${USERNAME} ${USERNAME}

# Set working directory and switch to non-root user
WORKDIR /app

# Create necessary directories
RUN mkdir -p ./logs/los \
  && chown -R ${USER_ID}:${GROUP_ID} /app

# Copy the built binary
COPY --from=build /app/target/release/los_logger .

USER ${USERNAME}

# Run the application
CMD ["./los_logger"]
