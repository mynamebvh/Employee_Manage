dev:	
	air
migrate:
	go run ./db/migrate/migrate.go
seeder: 
	go run ./db/seeder/seeder.go
doc:
	swag init -g routes/index.go