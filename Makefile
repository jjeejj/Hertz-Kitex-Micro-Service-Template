# start the environment of FreeCar
.PHONY: start
start:
	docker-compose up -d

# stop the environment of FreeCar
.PHONY: stop
stop:
	docker-compose down

# migrate mysql database

# run the api
.PHONY: api
api:
	go run ./server/cmd/api

# gen kitex rpc client
#gen_client:
#	kitex

.PHONY: oss
oss:
	go run ./server/cmd/oss
