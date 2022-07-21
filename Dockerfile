# Download latest golang image
FROM golang:latest as base

FROM base as dev

# Create a directory for the app
RUN mkdir /app

# Set working directory
WORKDIR /app

# Copy all files from current directory to working directory
COPY . /app

# Install depencies
RUN go get ./...

RUN cd cmd && go build -o EmbedDiscordVideo

CMD ["/app/cmd/EmbedDiscordVideo"]

