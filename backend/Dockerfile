FROM alpine:latest

# Install dependencies
RUN apk update && apk add --no-cache \
    libc6-compat

# Set working directory to /usr/src/app
WORKDIR /usr/src/app

# Copy the binary and .env file
COPY submission .env ./

# Expose the port
EXPOSE 8000

# Run the application
CMD ["./submission"]
