include .env

postgres:
	docker run --name ${DOCKER_IMAGE_NAME} -p 5432:${POSTGRESS_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:${POSTGRES_VERSION}
createdb:
	docker exec -it ${DOCKER_IMAGE_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}
dropdb:
	docker exec -it ${DOCKER_IMAGE_NAME} dropdb ${POSTGRES_DB}
migrateup:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRESS_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRESS_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test