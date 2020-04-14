# A Todo list API build with Golang and Sqlite3

## Build

`go build`

## Run

`./todolist-api`

## API

- curl -i -X POST http://localhost:8888/todo/create -H "Content-Type: application/json" -d '{"content":"buy me a cup of coffee", "title": "hello", "image_url": "http://..." }'
- curl -i -X GET  http://localhost:8888/todos
- curl -i -X GET  http://localhost:8888/todo/:id
- curl -i -X DELETE http://localhost:8888/todo/:id
- curl -i -X PUT http://localhost:8888/todo/111-11111-1111111 -H "Content-Type: application/json" -d '{"completed": true, "contents": "update"}'

## Struct

```go
type Todo struct {
  ID      uuid.UUID   `json:"id"`
  Contents string `json:"contents"`
  Title   string `json:"title"`
  Completed bool `json:"completed"`
  ImageURL string `json:"image_url"`
}
```

## Access with Axios and jQuery

```js
// Axios
axios.post('http://localhost:8888/todo/create', {content:"buy a book"}).then(resp => console.log(resp.data));
axios.get('http://localhost:8888/todos').then(resp => console.log(resp.data));

// jQuery
$.ajax({
  url: 'http://localhost:8888/todo/create',
  method: 'POST',
  contentType: 'application/json',
  data: JSON.stringify({content:"buy a book"}),
  success: function (todo) {
    console.log(todo)
  },
  error: function (xhr, status) {
    console.error("error");
  }
});

$.ajax({
  url: "http://localhost:8888/todos",
  type: "GET",
  crossDomain: true,
  dataType: "json",
  success: function (todos) {
    console.log(todos)
  },
  error: function (xhr, status) {
    console.error("error");
  }
});
```

## License

[MIT](https://github.com/aztack/todolist-api/blob/master/LICENSE)
