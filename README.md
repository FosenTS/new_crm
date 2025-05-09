# CRM System Backend

A backend system for managing goods and orders in a warehouse environment.

## Features

- JWT-based authentication
- Role-based access control (Employee and Boss roles)
- Product inventory management
- Order management with automatic stock updates
- Multiple order statuses
- PostgreSQL database
- Docker deployment support

## API Documentation

### Authentication

#### Register User
```
POST /register
Content-Type: application/json

{
    "username": "string",
    "password": "string",
    "role": "employee" | "boss"
}
```

#### Login
```
POST /login
Content-Type: application/json

{
    "username": "string",
    "password": "string"
}
```
Response: `{ "token": "string" }`

### Products

#### Get All Products
```
GET /products
Authorization: Bearer <token>
```

#### Create Product (Boss only)
```
POST /products
Authorization: Bearer <token>
Content-Type: application/json

{
    "name": "string",
    "description": "string",
    "quantity": number,
    "price": number
}
```

#### Update Product (Boss only)
```
PUT /products/:id
Authorization: Bearer <token>
Content-Type: application/json

{
    "name": "string",
    "description": "string",
    "quantity": number,
    "price": number
}
```

### Orders

#### Create Order
```
POST /orders
Authorization: Bearer <token>
Content-Type: application/json

{
    "items": [
        {
            "product_id": number,
            "quantity": number
        }
    ]
}
```

#### Get All Orders
```
GET /orders
Authorization: Bearer <token>
```

#### Update Order Status (Boss only)
```
PUT /orders/:id/status
Authorization: Bearer <token>
Content-Type: application/json

{
    "status": "pending" | "processing" | "completed" | "cancelled"
}
```

## Deployment

1. Create a `.env` file with the following variables:
```
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=crm_db
JWT_SECRET=your-secret-key-here
SERVER_PORT=8080
```

2. Build and run using Docker Compose:
```bash
docker-compose up --build
```

The application will be available at `http://localhost:8080`

## Development

1. Install Go 1.21 or later
2. Install PostgreSQL
3. Set up environment variables
4. Run the application:
```bash
go mod download
go run main.go
```

## Security Considerations

- All passwords are hashed using bcrypt
- JWT tokens are used for authentication
- Role-based access control is implemented
- HTTPS is recommended for production deployment
- Database credentials should be changed in production
- JWT secret should be changed in production 