# Go Chi RESTful API
RESTful API with the Go programming language.
## Tools
- go 1.16, You can install Go using gvm
- IDE for go, I suggest vscode(free) or GoLand(paid)

## How to run

Install dependencies

```
make install_deps
```

run the service

```
make run_service
```

## What I learn on this
- how to design REST API endpoints
- how to create your API using the chi library and Go
- how to run your Go project at the command line

## Endpoints
- GET / - Returns a "Hello World!" message.
- GET /posts - Get a list of posts.
- POST /posts - Create a post.
- GET /posts/{id} - Get a single post.
- PUT /posts/{id} - Update a single post.
- DELETE /posts/{id} - Delete a single post.