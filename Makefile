
# populate environment with .env file
env:
	@export $(grep -v '^#' .env | xargs) > /dev/null 2>&1

.PHONY: download
download: env
	@go run ./download/main.go $(day)
