# Event Hub

## server

- ✅ Go standard library net/http
- ✅ Templ templates
- ✅ Live reloads in development
- ✅ Docker development containers
- ✅ Alpine.js
- 🚧 Htmx - installed, not used yet
- 🚧 Tailwind CSS
- 🚧 Session-based Auth stored in Redis
- 🚧 GORM

## database

- 🚧 Postgres
- 🚧 Models
  - Users
  - Events
    - gorm.Model
    - ValidAt & ExpiresAt
    - Title
    - Date
    - Time (purposely kept separate?)
    - Description
    - Foreign Key for Link(s)
  - Links
    - gorm.Model
    - Href
    - Display

## cache

- 🚧 Redis

## proxy

- 🚧 Nginx reverse proxy
