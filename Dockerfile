FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./

COPY web/  ./bin/web
COPY sqlite/scheduler_creator.sql ./bin/sqlite

RUN go mod tidy

ENV PATH="$PATH:/app/bin"

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 TODO_DBFILE="" go build -o /app/bin/todorun ./cmd/todo/main.go

CMD ["/app/bin/todorun"]

#ENTRYPOINT ["/app/bin/todorun"]