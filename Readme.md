# 📚 GORM-MySQL API with Go Fiber, Protobuf & .env Configuration

This project is a RESTful API built in Go using:

- [Fiber](https://gofiber.io/) — fast web framework
- [GORM](https://gorm.io/) — ORM for MySQL
- [Protocol Buffers (Protobuf)](https://developers.google.com/protocol-buffers) — for consistent API data contracts
- `.env` file loading via [`godotenv`](https://github.com/joho/godotenv)
- Clean modular architecture using controllers, repositories, and models

---

## 🚀 Features

- 📦 Modular structure (controllers, models, routes, repository)
- 🔁 Auto-migrate DB schema on startup
- ✅ .env-based configuration
- 🧪 Easy to extend and test

---

## 📂 Project Structure

The project is structured as follows:

- `cmd/`: Entrypoint for the application
	+ `main.go`: Initializes app, DB, and routes
- `config/`: Configuration and environment loading
	+ `db.config.go`: Connects to MySQL using GORM
- `controllers/`: Fiber route handlers (HTTP layer)
	+ `book.controller.go`: Handles HTTP requests/responses
- `models/`: GORM models and Protobuf mapping logic
	+ `book.go`: GORM model
	+ `book.pb.go`: Generated from `.proto` file
- `protocol/`: Source Protobuf definition files
	+ `book.proto`: Protobuf schema definition
- `repository/`: Data access layer (repository pattern)
	+ `book.repository.go`: Abstracts database operations
- `routes/`: HTTP route registration
	+ `book.routes.go`: Registers routes with Fiber
- `.env`: Environment variable definitions
- `go.mod`: Go module definition
- `go.sum`: Go module checksums


## 🧼 License

This project is distributed under the [MIT License](https://opensource.org/licenses/MIT).

---

## 👨‍💻 Maintainer

Designed by [Heinirich](https://github.com/Heinirich)