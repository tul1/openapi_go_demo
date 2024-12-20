# Go OpenAPI Demo

## Steps to run the API

1. Install `oapi-codegen`:

   ```bash
   go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
   ```

2. Generate the server code:

   ```bash
   oapi-codegen -generate models,gin-server -package openapi -o ./openapi/api.gen.go api.yaml
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

4. The API will be accessible at `http://localhost:8080`.

## Testing the API

### Using `curl`

- Get all TODOs:

  ```bash
  curl http://localhost:8080/todos
  ```

- Create a new TODO:

  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"task":"Learn Go"}' http://localhost:8080/todos
  ```

- Get a TODO by ID:

  ```bash
  curl http://localhost:8080/todos/1
  ```

- Delete a TODO by ID:

  ```bash
  curl -X DELETE http://localhost:8080/todos/1
  ```

### Using Postman

1. Import the OpenAPI spec (`api.yaml`) into Postman.
2. Test the endpoints interactively.
