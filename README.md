# A Todo list API build with Golang and Sqlite3

## API

- curl -i -X POST http://localhost:8888/todo/create -d '{"content":"buy me a cup of coffee"}'
- curl -i -X GET  http://localhost:8888/todos
- curl -i -X GET  http://localhost:8888/todo/1
- curl -i -X DELETE http://localhost:8888/todo/1
