SWAG=github.com/swaggo/swag/cmd/swag@v1.16.6

swag:
	go install $(SWAG)
	swag init -g cmd/main.go -o ./docs