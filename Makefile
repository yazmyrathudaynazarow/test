docker-run:
	sudo docker run --name testProject -p 9042:9042 -d scylladb/scylla

dev:
	swag init --dir ./ -g $(SRC_DIR)/cmd/main.go
	go run cmd/main.go