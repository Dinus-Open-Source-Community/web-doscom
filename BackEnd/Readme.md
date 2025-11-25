# web-doscom Backend

Backend API for Dinus Open Source Community (DOSCOM)

## Features

- User registration and login (JWT-based)
- Role-based access: User, Admin, Super_Admin
- CRUD for users, works, activities, blogs, gallery, pengurus
- Secure password hashing (bcrypt)
- RESTful API with Gin
- Swagger/OpenAPI documentation (auto-generated with swag)
- Live reload for development (air)
- PostgreSQL database

## Requirements

- Go 1.18+
- PostgreSQL
- [air](https://github.com/air-verse/air) (for live reload, optional)
- [swag](https://github.com/swaggo/swag) (for Swagger docs)

## Setup

1. **Clone the repo**

```sh
git clone https://github.com/Dinus-Open-Source-Community/web-doscom.git
cd web-doscom/BackEnd
```

2. **Configure environment**

- Copy `.env.example` to `.env` and edit DB credentials as needed.

3. **Run database migrations**

- Use the SQL files in `migrations/` to set up your database tables.

4. **Install dependencies and tools**

```sh
go mod tidy
go install github.com/air-verse/air@latest
go install github.com/swaggo/swag/cmd/swag@latest
```

5. **Generate Swagger docs**

```sh
swag init -g cmd/api/main.go -o docs
```

6. **Run the server (with live reload)**

```sh
air
# or
GO_ENV=development go run ./cmd/api
```

## API Endpoints

- `POST /api/v1/register` — Register user (fields: username, email, password, fullname, role)
- `POST /api/v1/login` — Login (returns JWT)
- `GET /api/v1/users/:id` — Get user by ID
- `GET /api/v1/users` — List users
- `PUT /api/v1/users/:id` — Update user
- `DELETE /api/v1/users/:id` — Delete user
- `POST /api/v1/users/superadmin` — Create superadmin (admin only)
- ...and more for works, activities, blogs, gallery, pengurus

## Auth

- Use the JWT token from `/api/v1/login` in the `Authorization: Bearer <token>` header for protected endpoints.
- Only Admin/Super_Admin can create superadmin users.

## Swagger UI

- After running `swag init`, access docs at: `http://localhost:3001/swagger/index.html`

## Development

- Use `air` for hot reload during development.
- Use `swag` to update API docs after changing handler comments.

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](LICENSE)
