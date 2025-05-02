# Makefile equivalent to Taskfile.yml
envfile := .env

include $(envfile)
export $(shell sed 's/=.*//' $(envfile))

.PHONY: hello setup build test topcmd playingcmd popularcmd upcomingcmd

hello:
	@echo "Makefile works ok!!!"

setup:
	@test -f .env && echo "<.env> file already exists" || @cp .env.example .env

build:
	@go build -o tmdbcli .

test:
	@go test ./... -v

topcmd:
	@go run main.go movies --type top

playingcmd:
	@go run main.go movies --type playing

popularcmd:
	@go run main.go movies --type popular

upcomingcmd:
	@go run main.go movies --type upcoming
