build:
	go build -o 10db github.com/gtalent/tendb/exec
start-postgres:
	docker run --restart=always --name tendb-postgres -e POSTGRES_DB=tendb -e POSTGRES_PASSWORD=postgres -v /var/lib/10db/pgdata/:/var/lib/postgresql/data -p 5432:5432 -d postgres:alpine
pg-cli:
	docker exec -ti tendb-postgres psql -h localhost -U postgres -d tendb
migrate: build
	./10db migrate
generate:
	go generate github.com/gtalent/tendb/db
generate-models:
	genna -c "postgres://postgres:postgres@localhost:5432/tendb?sslmode=disable" -o db/generated_models.go -p db
	chmod -x db/generated_models.go
install-deps:
	go get github.com/gtalent/tendb/exec
	go get github.com/gtalent/tendb/assets_generate
	go get github.com/dizzyfool/genna
