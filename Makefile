run:
	go run ./cmd/ -c ./configs/config-local.yaml

migrate:
	go run ./cmd -c ./configs/config-local.yaml -migrate

vendor:
	go mod vendor

build:
	docker-compose up -d --build

down:
	docker-compose down

seed:
	go run ./cmd -c ./configs/config-local.yaml -seed