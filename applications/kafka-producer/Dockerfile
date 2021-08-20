# buider stage
FROM golang:1.16-stretch as builder

# Set working directory
WORKDIR /src

# Copy go-modules definition for the application
COPY go.mod .

# Download required go-modules
RUN go mod download

# Copy application code
COPY . .

# Compile application
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags netgo -o /bin/app -ldflags "-w -s"

# final stage
FROM debian:11-slim

# Set working directory
WORKDIR /bin

# Copy application executable
COPY --from=builder --chown=1001 /bin/app .

## Metrics
EXPOSE 9090
## KubeProbes
# EXPOSE 9091

# Set user
USER 1001

# Run application
ENTRYPOINT ["/bin/app"]
