# Start with the base golang image
FROM golang:1.20-alpine AS build

# Set the working directory to /app
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the binary
RUN go build -o simple-go-app

# Use a small image for the final build
FROM alpine:3.14

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the build image
COPY --from=build /app/simple-go-app .

# Expose port 8080 for the app
EXPOSE 8080

# Set the command to run the app
CMD ["./simple-go-app"]
