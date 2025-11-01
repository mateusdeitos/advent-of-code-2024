set dotenv-load

download year day:
	export SESSION=$SESSION && go run ./download/main.go {{year}} {{day}}
