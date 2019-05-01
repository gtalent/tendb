build:
	go build -o 10db github.com/gtalent/tendb/exec
start-postgres:
	docker run --restart=always --name tendb-postgres -e POSTGRES_DB=tendb -e POSTGRES_PASSWORD=postgres -v /var/lib/10db/pgdata/:/var/lib/postgresql/data -p 5432:5432 -d postgres:alpine
pg-cli:
	docker exec -ti tendb-postgres psql -h localhost -U postgres -d tendb
