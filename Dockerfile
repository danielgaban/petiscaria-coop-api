# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required dependencies
RUN go mod download

# Command to run the executable
CMD ["go", "run", "main.go"]

# Expose port 8080 to the outside world
EXPOSE 80

