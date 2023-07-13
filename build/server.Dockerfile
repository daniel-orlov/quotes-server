# Use the latest GoLang base image
FROM golang:1.19-buster AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# Use a distroless image as the final base image
FROM gcr.io/distroless/static:nonroot

# Copy the binary from the build stage to the final image
COPY --from=build /app/server /app/server

# Expose any necessary ports
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["/app/server"]
