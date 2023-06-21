include .env
export $(shell sed 's/=.*//' .env)

make run:
	clear
	go run main.go