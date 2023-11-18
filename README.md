todo rest-app project using gin framework for routing

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
