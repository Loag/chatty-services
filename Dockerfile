# Start from the official Golang base image
FROM golang:1.21 AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy go module files first to leverage Docker caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o dist ./cmd/app/main.go

# Use a minimal alpine image for the runtime container
FROM scratch

# Copy the compiled Go binary into our runtime container
COPY --from=build /app/dist .

# Expose port (if your app serves on a particular port)
# EXPOSE 8080

# Command to run the application
CMD ["./dist"]
