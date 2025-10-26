# Dockerfile

# Base image (local, but still Go)
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

# Platform parameters
ARG TARGETOS
ARG TARGETARCH

# Compile code for the specified platform
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o build/app ./src

# Test phase
FROM golang:1.22 AS tester

WORKDIR /app
COPY --from=builder /app/build/app /usr/local/bin/app
COPY ./tests ./tests

CMD ["bash", "./tests/test_main.sh"]
