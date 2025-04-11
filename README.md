## Getting Started

### Clone the Repository

```bash
git clone https://github.com/mdayat/fullstack2024-test.git
cd fullstack2024-test
```

### Install Deps and Run

```bash
go mod tidy && make run
```

> The application will be available at http://localhost:8080.

## Tech Stack

1. Go
2. sqlc
3. chi

## Directory Structure

- `configs`
- `internal`
  - `handlers`
  - `dtos`
  - `httputil`
  - `dbutil`
  - `retryutil`
- `repository`

> NOTE: Don't forget to provide `.env` file for `DATABASE_URL` and `REDIS_URL` environment variables.
