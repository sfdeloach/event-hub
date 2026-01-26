# Event Hub

## server

- ✅ Go standard library net/http
- ✅ Templ templates
- ✅ Live reloads in development
- ✅ Docker development containers
- ✅ Alpine.js
- ✅ Htmx - installed, not used yet
- ✅ Tailwind CSS
- 🚧 Create New Category Form
- 🚧 Add database to network
- 🚧 GORM
- 🚧 CRUD Ops for Category Table
- 🚧 CRUD Ops for Event Table
- 🚧 Session-based Auth stored in Redis

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
    - Category (foreign key)
    - Title
    - Date
    - Time (purposely kept separate?)
    - Description (accepts limited html tags <a> <em> <strong>)
  - Event Categories
    - gorm.Model
    - Category

## cache

- 🚧 Redis

## proxy

- 🚧 Nginx reverse proxy
