# CRM Backend

This is the backend system for a CRM application developed using Go and the Gin framework.

## Features

- User management
- Customer management
- Interaction tracking
- Analytics and reporting
- Email integration
- Activity notifications

## Getting Started

### Prerequisites

- Go (version 1.17 or later)
- MongoDB
- Docker (optional, for containerization)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/crm-backend.git
   cd crm-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up MongoDB:
   - Install MongoDB
   - Start the MongoDB server
   - Create a new database for the CRM application

4. Configure environment variables:
   - Create a `.env` file in the project root
   - Add necessary environment variables (e.g., database connection string, email service API keys)

5. Run the application:
   ```bash
   go run main.go
   ```

## Project Structure

```
crm-backend/
├── controllers/
├── models/
├── routes/
├── services/
├── utils/
├── middlewares/
├── config/
├── tests/
├── Dockerfile
├── go.mod
├── go.sum
└── main.go
```

## API Documentation

API documentation is available using Swagger. After running the application, visit:
```
http://localhost:8080/swagger/index.html
```

## Testing

Run the tests using the following command:
```
go test ./...
```

## Deployment

1. Build the Docker image:
   ```bash
   docker build -t crm-backend .
   ```

2. Run the Docker container:
   ```bash
   docker run -p 8080:8080 crm-backend
   ```

## Contributing

1. Fork the repository
2. Create a new branch
3. Make your changes
4. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.