# Go OpenAPI Demo

## Steps to Run the API

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

### Using `restish`

- Install [`restish`](https://rest.sh/#/):

  ```bash
  brew install restish
  ```

- Get all TODOs:

  ```bash
  restish get /todos
  ```

- Create a new TODO:

  ```bash
  restish post /todos --body '{"task":"Learn Go"}'
  ```

- Get a TODO by ID:

  ```bash
  restish get /todos/1
  ```

- Delete a TODO by ID:

  ```bash
  restish delete /todos/1
  ```

### Using Postman

1. Import the OpenAPI spec (`api.yaml`) into Postman.
2. Test the endpoints interactively.
