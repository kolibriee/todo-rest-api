todo rest-app project using gin framework for routing

docker:
```
docker build -f .\build\Dockerfile --progress=plain -t todo-rest-api .
docker run -it --name my-todo-app --network my_network -p 8000:8000 todo-rest-api
postgres:
docker run --name todo-rest-api-postgres --network my_network -e POSTGRES_PASSWORD=your_password -d -p 5433:5432 postgres
```

DB diagram: https://dbdiagram.io/d/64e32ad702bd1c4a5e1d82e6

DB migration: https://github.com/golang-migrate/migrate  
command: migrate -database "postgres://your_user:your_password@your_host:your_port/your_database?sslmode=disable" -path /path_to_migrations up

.env:
```
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_DBNAME=
DB_SSLMODE=
PASSWORD_HASH_SALT=
TOKEN_SECRET_KEY=
```
