FROM golang:1.22

WORKDIR /app

COPY . .
RUN go mod download

WORKDIR /app/cmd/app
RUN GOOS=linux go build -o /app/main

EXPOSE 8000
WORKDIR /app
CMD ["/app/main"]

# docker build -f .\build\Dockerfile --progress=plain -t todo-rest-api .
# docker run -it --name my-todo-app --network my_network -p 8000:8000 todo-rest-api
# postgres:
# docker run --name todo-rest-api-postgres --network my_network -e POSTGRES_PASSWORD=your_password -d -p 5433:5432 postgres