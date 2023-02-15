# start the environment of FreeCar
.PHONY: start
start:
	docker-compose up -d

# stop the environment of FreeCar
.PHONY: stop
stop:
	docker-compose down

# migrate mysql database
.PHONY: migrate
migrate:
	go run server/cmd/auth/model/migrate/main.go

# run the auth
.PHONY: auth
auth:
	go run ./server/cmd/auth

# run the api
.PHONY: api
api:
	go run ./server/cmd/api