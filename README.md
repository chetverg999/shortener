# URL Shortener

## Project Overview
The URL Shortener is a simple and efficient web service written in Go that converts long URLs into short, easily shareable links. The service is backed by MongoDB for persistent storage and supports dynamic generation of short links.

### Key Features
- **Shorten URLs:** Convert long URLs into short, unique, and shareable links.
- **Redirects:** Automatically redirect users to the original URL when they access the short link.
- **MongoDB Integration:** Stores URLs and their corresponding shortened versions in a database.
- **UTF-8 Validation Support:** Handles URL validation to ensure proper encoding.
- **Customizable Environment:** Easily configure database connection and server settings using `.env`.

---

## Installation

### Prerequisites
1. **Go:** Ensure you have Go installed on your machine ([Download Go](https://golang.org/dl/)).
2. **MongoDB:** Set up a local or cloud instance of MongoDB.
3. **Git:** Clone the repository to your local machine.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```
2. Set up your environment variables in a `.env` file:
   ```env
   DB=mongodb://localhost:27017
   DB_name=shortener_db
   DB_collection=urls
   PORT=:8080
   HOST=http://localhost:8080/
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

---

## Usage

### Shorten a URL
Send a `POST` request to the root endpoint with the original URL in the request body:
```bash
curl -X POST -d "https://example.com" http://localhost:8080/
```
Response:
```
http://localhost:8080/abc123
```

### Redirect to Original URL
Access the shortened URL in your browser or via `curl`:
```bash
curl http://localhost:8080/abc123
```
Redirects to `https://example.com`.

---

## Directory Structure
```
url-shortener/
├── cmd/
│   ├── server/     
├── internal/
│   ├── env/          # Environment variable handling
│   ├── shortener/    # URL shortening logic
│   └── storage/      # MongoDB integration and database logic
├── handlers/         # HTTP handlers for POST and GET requests
├── .env              # Environment configuration (ignored in Git)
├── main.go           # Application entry point
└── go.mod            # Module dependencies
```

---

## Development

### Running Tests
Run tests with the following command:
```bash
go test ./...
```

### Debugging MongoDB Issues
- Ensure MongoDB is running on the specified host and port.
- Check if UTF-8 validation is disabled if encountering encoding issues:
  ```bash
  mongodb://localhost:27017/?enableUtf8Validation=false
  ```

---

## Contributions
Contributions are welcome! Please fork the repository and create a pull request with your improvements.

---


## Author
Created by [chetverg999](https://github.com/chetverg999).
