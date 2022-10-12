# ip-crawler

Service to crawl across the internet and find all available IP addresses on port 80.

[Frontend](https://github.com/yanislav-igonin/ip-crawler-frontend) repo.

## How to run
1. Add `.env` file from `.env.example` and fill it with your data.
2. `go run main.go`

**In development** I prefer to run it with [nodemon](https://github.com/remy/nodemon) for live reload with Makefile:
```sh
make dev
```