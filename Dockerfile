# Use a Go 1.21 or newer image
FROM golang:1.21

# Install Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application into the container
COPY . .

# Expose the port Delve will use (default is 2345)
EXPOSE 2345

# Start Delve in headless mode to listen for remote connections
CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient", "--log"]
