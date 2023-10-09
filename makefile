postgres-container:
	docker run --name todo_container -e POSTGRES_PASSWORD=todo \
	-e POSTGRES_USER=todo -e PGPORT=7899 -p 7899:7899 -d postgres:14.8-alpine

postgres-run:
	docker start todo_container


postgres-create:
	docker exec -it todo_container createdb --username=todo --owner=todo todo_db
postgres-drop:
	docker exec -it todo_container dropdb  --username=todo todo_db
migration-up:
	docker exec -i todo_container psql -U todo -d todo_db < migration/up.sql
migration-down:
	docker exec -i todo_container psql -U todo -d todo_db < migration/down.sql

test:
	go test -v -cover ./...

test-coverage-view:
	touch coverage.out
	go test -coverprofile coverage.out ./...
	go tool cover -html=coverage.out