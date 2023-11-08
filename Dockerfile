# PREFACE: I have ONLY included what is needed to run the application, nothing more.

# Use Alpine base image with GO 1.21
FROM golang:1.21-alpine

# Set workdir
WORKDIR /app

# Copy source to working dir
COPY . .

# Build image
RUN go build -o main .

# Run image
CMD ["./main"]