all:
	docker-compose up -d
	
build:
	docker build -t my-golang -f Dockerfile .

migrateup:
	migrate -path db/migration -database "mysql://me:hithere@tcp(localhost:3306)/managementEquipment?x-no-lock=true" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://me:hithere@tcp(localhost:3306)/managementEquipment?x-no-lock=true" -verbose down

.PHONY: all build migrateup migratedown