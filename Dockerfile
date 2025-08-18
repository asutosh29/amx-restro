FROM golang:1.24 
RUN mkdir -p /app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

# migrations
#  RUN migrate -path database/migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(localhost:3306)/${MYSQL_DATABASE}" up 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./server_go ./cmd/main.go

RUN chmod +x "/app/entrypoint.sh"
ENTRYPOINT ["/app/entrypoint.sh"]
# run the server
# CMD ["./entrypoiny.sh"]
CMD ["./server_go"]


