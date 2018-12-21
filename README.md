# A Todo list API build with Golang and Sqlite3

## Build

`go build`

## Run

`./todolist-api`

## API

- curl -i -X POST http://localhost:8888/todo/create -d '{"content":"buy me a cup of coffee"}'
- curl -i -X GET  http://localhost:8888/todos
- curl -i -X GET  http://localhost:8888/todo/:id
- curl -i -X DELETE http://localhost:8888/todo/:id

## Struct

```go
type Todo struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}
```

## License

[MIT](https://github.com/aztack/todolist-api/blob/master/LICENSE)
