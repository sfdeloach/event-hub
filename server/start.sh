#!/bin/bash
set -e
echo "Starting TailwindCSS and Templ generator in watch mode..."

# Trap to clean up background process on exit
trap "kill 0" EXIT

tailwind -i ./tailwind.css -o ./static/css/style.css --watch=always &
go tool templ generate --watch --proxy="http://localhost:8080" --proxybind="0.0.0.0" --open-browser=false --cmd="go run ."
