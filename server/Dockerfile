FROM ubuntu:noble

WORKDIR /server

# Install prerequisites
RUN apt-get update
RUN apt-get install -y software-properties-common

# Update packages and install dependencies
RUN add-apt-repository universe
RUN apt-get update && apt-get upgrade

# Install helpful tools for use inside the container
RUN apt-get install -y vim tree

# Install Go, Curl (for downloading TailwindCSS), and Watchman (TailwindCSS dependency)
RUN apt-get install -y curl golang-go watchman

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Install and configure TailwindCSS CLI v4.1.18
RUN curl -fL -o tailwind https://github.com/tailwindlabs/tailwindcss/releases/download/v4.1.18/tailwindcss-linux-x64
RUN chmod +x tailwind
RUN mv tailwind /usr/local/bin/
RUN mkdir -p ./static/css
RUN touch ./static/css/style.css

# Start tailwind and templ generator in watch mode
CMD ["./start.sh"]

# Default port for templ proxy
EXPOSE 7331
