dev:	
	air
migrate:
	go run ./db/migrate/migrate.go
seeder: 
	go run ./db/seeder/seeder.go
doc:
	swag init -g routes/index.go
consumer:
	go run ./services/consumer/consumer.go
docker-dev:
	sudo docker-compose -f docker-compose.dev.yml up -d