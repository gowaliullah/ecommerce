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


---

## 🔗 API Endpoints

### 👤 **Auth Routes**
| Method | Endpoint | Description |
|--------|-----------|-------------|
| `POST` | `/api/auth/register` | Register a new user |
| `POST` | `/api/auth/login` | Login user and receive JWT token |

---

### 🛒 **Product Routes**
| Method | Endpoint | Description | Access |
|--------|-----------|-------------|--------|
| `GET` | `/api/products` | Get all products | Public |
| `GET` | `/api/products/:id` | Get product by ID | Public |
| `POST` | `/api/products` | Create product | Admin |
| `PUT` | `/api/products/:id` | Update product | Admin |
| `DELETE` | `/api/products/:id` | Delete product | Admin |

---

### 📦 **Order Routes**
| Method | Endpoint | Description | Access |
|--------|-----------|-------------|--------|
| `GET` | `/api/orders` | Get all orders (Admin) | Admin |
| `POST` | `/api/orders` | Create a new order | Authenticated User |
| `GET` | `/api/orders/:id` | Get order details | Authenticated User |
| `PUT` | `/api/orders/:id` | Update order status | Admin |

---

### 🧑‍💼 **Admin Routes**
| Method | Endpoint | Description | Access |
|--------|-----------|-------------|--------|
| `GET` | `/api/admin/dashboard` | Get summary stats (users, sales, orders) | Admin |
| `GET` | `/api/admin/users` | Get all registered users | Admin |
| `PUT` | `/api/admin/users/:id/role` | Update user role | Admin |

---

## ⚙️ Environment Variables

Create a `.env` file in your root directory:

```bash
# PostgreSQL
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=gocommerce

# JWT
JWT_SECRET=your_jwt_secret_key
JWT_EXPIRE_HOURS=24

# Server
PORT=8080
