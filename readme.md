# 🛍️ GoCommerce — E-commerce Backend API (Go + PostgreSQL)

A fully functional **e-commerce backend** built with **Golang**, **PostgreSQL**, and **JWT authentication**.  
This project is designed to power a modern e-commerce platform with features like product management, user authentication, and order handling.

---

## 🚀 Tech Stack

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **Query:** `database/sql` + `pgx` driver 
- **Authentication:** JWT (JSON Web Token)
- **Password Hashing:** bcrypt
- **API Format:** RESTful JSON
- **Architecture:** Domain-driven design
- **Environment Management:** `.env` file
- **Dependency Management:** Go Modules

---

## 🗂️ Project Structure

gocommerce-backend/
│
├── cmd/
│ └── main.go # App entry point
│
├── internal/
│ ├── config/ # Environment & DB config
│ ├── models/ # Data models (structs)
│ ├── handlers/ # Route handlers (controllers)
│ ├── repositories/ # DB queries
│ ├── services/ # Business logic
│ ├── middlewares/ # JWT, auth, logging
│ └── routes/ # API route definitions
│
├── pkg/
│ └── utils/ # Helper functions
│
├── go.mod
├── go.sum
├── .env.example
└── README.md


🧭 API Endpoints
👤 Auth
Method	Endpoint	Description
POST	/api/auth/register	Register a new user
POST	/api/auth/login	Login and get JWT token


🛒 Products
Method	Endpoint	Description
GET	/api/products	Get all products
GET	/api/products/:id	Get product by ID
POST	/api/products	Create product (Admin only)
PUT	/api/products/:id	Update product (Admin only)
DELETE	/api/products/:id	Delete product (Admin only)


📦 Orders
Method	Endpoint	Description
GET	/api/orders	Get all orders (Admin)
POST	/api/orders	Create new order (User)
PUT	/api/orders/:id	Update order status (Admin)