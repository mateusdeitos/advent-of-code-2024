set dotenv-load

download day:
	export SESSION=$SESSION && go run ./download/main.go {{day}}
