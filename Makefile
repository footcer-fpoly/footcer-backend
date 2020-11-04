pro:
    sudo docker system prune
	docker rmi -f web-service:1.0
	docker-compose build
	docker-compose up
dev:
	cd cmd/dev; go run main.go