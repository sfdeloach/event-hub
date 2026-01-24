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

### setup

#### Tailwind CSS

 - Download [v.1.18](https://github.com/tailwindlabs/tailwindcss/releases/download/v4.1.18/tailwindcss-linux-x64) to `/server`
   - **OPTIONAL:** Download [watchman-bin](https://aur.archlinux.org/packages/watchman-bin) but it works without it.
 - Rename to `tailwindcss-4.1.18` and make executable with `chmod +x`
 - **TODO:** Develop `up.sh` and `dn.sh` scripts to start Tailwind and Docker Compose
   - `./tailwindcss-4.1.18 -i ./static/css/input.css -o ./static/css/output.css --watch`

### resources

- [The files & folders of Go projects](https://changelog.com/gotime/278)
- [How I write HTTP services in Go after 13 years](https://grafana.com/blog/how-i-write-http-services-in-go-after-13-years/)

## database

- 🚧 Postgres
- 🚧 Models
  - Users
  - Events
    - gorm.Model (ID, CreatedAt, UpdatedAt, DeletedAt)
    - OnAirAt
    - OffAirAt
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
