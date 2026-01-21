# DEVELOPMENT DOCKERFILE

FROM ubuntu:noble

ARG TAILWIND_VERSION=v4.1.18

WORKDIR /server

# Install prerequisites and dependencies
RUN apt-get update && \
    apt-get install -y software-properties-common && \
    add-apt-repository universe && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y vim tree curl golang-go watchman && \
    rm -rf /var/lib/apt/lists/*

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Install and setup output directory for TailwindCSS CLI
RUN curl -fL -o /usr/local/bin/tailwind https://github.com/tailwindlabs/tailwindcss/releases/download/${TAILWIND_VERSION}/tailwindcss-linux-x64 && \
    chmod +x /usr/local/bin/tailwind && \
    mkdir -p ./static/css && touch ./static/css/style.css

# Default port for templ proxy
EXPOSE 7331

# Start tailwind and templ generator in watch mode
CMD ["./start.sh"]
