# Mess Backend
Backend for the mess app

## Setup
- Install Golang 1.22
- Add your `ATLAS_URI` environment variable to `.env`
- Optionally mention your `PORT` number
## Running natively
Make sure to export your environment variables in your `.env`. You can use tools like `direnv` to automate this process for yourself. Then run the server for testing using `find -name "*.go" | entr -r go run .`
(entr is generic nodemon... :)
## Run locally
```bash
docker build --tag mess-backend . && docker run --env-file .env -p 8000:8000 -dit mess-backend
```