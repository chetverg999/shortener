# URL Shortener

A simple and efficient web service written in Go that converts long URLs into short, easily shareable links. The service utilizes MongoDB for persistent storage and supports dynamic generation of short links.

## Features

- **Shorten URLs**: Convert lengthy URLs into concise, unique, and shareable links.
- **Redirects**: Automatically redirect users to the original URL when they access the shortened link.
- **MongoDB Integration**: Store original and shortened URLs in a MongoDB database.
- **Environment Configuration**: Manage database connections and server settings using a `.env` file.

## Project Structure

```plaintext
shortener/
├── cmd/
│   └── server/       # Main application entry point
├── internal/
│   ├── env/          # Environment variable handling
│   ├── shortener/    # URL shortening logic
│   ├── storage/      # MongoDB integration and database logic
│   └── handlers/     # HTTP handlers for POST and GET requests
├── .env              # Environment configuration (ignored in Git)
├── go.mod            # Module dependencies
└── Makefile          # Task automation
```

## Getting Started

### Prerequisites

- **Go**: Version 1.23.1 or higher. [Download Go](https://golang.org/dl/)
- **MongoDB**: Ensure you have access to a MongoDB instance. [Install MongoDB](https://docs.mongodb.com/manual/installation/)
- **Git**: To clone the repository. [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/chetverg999/shortener.git
   cd shortener
   ```

2. **Set up environment variables**:

   - Create a `.env` file in the root directory.
   - Refer to `.env.example` for the required variables.

   ```bash
   cp .env.example .env
   ```

   - Update the `.env` file with your MongoDB connection string and any other necessary configurations.

3. **Install dependencies**:

   ```bash
   make deps
   ```

4. **Build the application**:

   ```bash
   make build
   ```

5. **Run the application**:

   ```bash
   make run
   ```

   The server should now be running, and you can access it at `http://localhost:8080` (or the port specified in your `.env` file).

## Usage

- **Shorten a URL**:
  - Send a POST request to `/` with a Text body containing the `url` field.
  - Example using `curl`:

    ```bash
    curl -X POST http://localhost:8080/ -H "Content-Type: text/plain; charset=utf-8" -d "https://example.com"
    ```

- **Access the shortened URL**:
  - Navigate to `http://localhost:8080/{shortID}`, where `{shortID}` is the unique identifier returned from the `/shorten` endpoint.
  - This will redirect you to the original URL.

## Testing

- To run tests:

  ```bash
  make test
  ```

## Formatting

- To format the codebase:

  ```bash
  make fmt
  ```

## Cleaning Up

- To remove the built binary and clean up:

  ```bash
  make clean
  ```

## Configuration

- **Environment Variables**:
  - `PORT`: Port on which the server runs (default: `8080`).
  - `MONGODB_URI`: Connection string for MongoDB.
  - `DATABASE_NAME`: Name of the MongoDB database.
  - `COLLECTION_NAME`: Name of the collection to store URLs.

  Ensure these variables are set in your `.env` file.

## Dependencies

- [Go Modules](https://github.com/golang/go/wiki/Modules)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [godotenv](https://github.com/joho/godotenv)
- [MongoDB Go Driver](https://go.mongodb.org/mongo-driver)

---
