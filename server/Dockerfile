FROM ubuntu:noble

WORKDIR /app

# Update packages and install dependencies
RUN apt-get update && apt-get upgrade

# Install Go and dependencies
RUN apt-get install -y curl golang-go
COPY go.mod go.sum ./
RUN go mod download

# Install TailwindCSS CLI
RUN curl -fL -o tailwind https://github.com/tailwindlabs/tailwindcss/releases/download/v4.1.18/tailwindcss-linux-x64
RUN chmod +x tailwind
RUN mv tailwind /usr/local/bin/
RUN mkdir -p ./static/css
RUN touch ./static/css/style.css

# TODO: create 'start.sh' script to run both tailwind and go templ in parallel

CMD ["go", "tool", "templ", "generate", "--watch", \
     "--proxy=http://localhost:8080", \
     "--proxybind=0.0.0.0", \
     "--open-browser=false", \
     "--cmd", "go run ."]

EXPOSE 7331
