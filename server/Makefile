DB_URL=postgresql://root:It123456@@localhost:5432/erpoo?sslmode=disable


postgres:
	docker run -d --name postgres-erp -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=It123456@ postgres

createdb:
	docker exec -it postgres-erp createdb --username=root --owner=root erpoo

dropdb:
	docker exec -it postgres-erp dropdb erpoo

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1