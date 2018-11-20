# go-todo-echo-sample

## About This

echoとVue.jsを使用したTodoのサンプル。

echoについてはTodoのCRUDのAPIを実装。データはメモリ内に保持している。

Vue.jsについては一覧の表示のみ。

## Usage

### Build & Start

```
$ go build
$ ./go-todo-echo-sample
```

### CURL

* create

    ```
    $ curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d "{\"name\":\"programming\"}"  http://localhost:8080/todos

    ```

* update

    ```
    $ curl -v -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d "{\"name\":\"trainning\"}"  http://localhost:8080/todos/1
    ```

* select

    ```
    $ curl http://localhost:8080/todos
    ```

    ```
    $ curl http://localhost:8080/todos/1
    ```

* delete

    ```
    $ curl -v -H "Accept: application/json" -H "Content-type: application/json" -X DELETE http://localhost:8080/todos/1
    ```

### HTML

ブラウザで「http://localhost:8080/」を表示。

「Show Todos!」ボタンを押下。
