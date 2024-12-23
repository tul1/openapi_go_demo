# Go OpenAPI Demo

## Steps to Run the API

### 1. Install Dependencies

1. Install `oapi-codegen`:

   ```bash
   go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
   ```

2. Install Go modules:

   ```bash
   go mod tidy
   ```

### 2. Generate the Server Code

Generate the server code from the OpenAPI spec:

```bash
oapi-codegen -generate models,gin-server -package openapi -o ./openapi/api.gen.go api.yaml
```

### 3. Set Up SQLite with Docker

Run a SQLite database using Docker:

1. Pull the official SQLite Docker image:

   ```bash
   docker pull nouchka/sqlite3
   ```

2. Create a folder to store the SQLite database file:

   ```bash
   mkdir -p $(pwd)/data
   ```

3. Run the SQLite Docker container:

   ```bash
   docker run --rm -d \
     -v $(pwd)/data:/data \
     --name sqlite-todos nouchka/sqlite3
   ```

   This starts a SQLite container and maps your local `./data` folder to the container's `/data` directory.

4. Verify the SQLite container is running:

   ```bash
   docker ps
   ```

5. (Optional) Connect to the SQLite shell:

   ```bash
   docker exec -it sqlite-todos sqlite3 /data/todos.db
   ```

   Inside the SQLite shell:
   - List tables with `.tables`
   - Check the schema of the `todos` table with `.schema todos`
   - Exit the SQLite shell with `.exit`

### 4. Automatic Table Creation

You do not need to manually create the `todos` table. 

Both database handlers (`SQLHandler` and `GORMHandler`) automatically create the `todos` table if it doesn’t already exist:

- **SQLHandler**: Runs a `CREATE TABLE IF NOT EXISTS` statement during initialization.
- **GORMHandler**: Uses GORM’s `AutoMigrate` to automatically create or update the schema based on the `model.Todo` struct.

### 5. Run the Server

Run the API server:

```bash
go run main.go
```

The API will be accessible at `http://localhost:8080`.

### 6. Testing the API

#### Using `restish`

1. Install [`restish`](https://rest.sh/#/):

   ```bash
   brew install restish
   ```

2. Test the endpoints:

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

#### Using Postman

1. Import the OpenAPI spec (`api.yaml`) into Postman.
2. Test the endpoints interactively.

### 7. Switching Between `database/sql` and GORM

Set the `USE_GORM` environment variable to `true` to use the GORM implementation:

```bash
USE_GORM=true go run main.go
```

If `USE_GORM` is not set, the server will default to using `database/sql`.
