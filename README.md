# Swaggo Demo

Check [[gin-swagger][https://github.com/swaggo/gin-swagger]] for details.

1. Install swag
```bash
$ go install github.com/swaggo/swag/cmd/swag
```

2. Generate swagger docs
```bash
$ swag init
$ tree docs
docs
├── docs.go
├── swagger.json
└── swagger.yaml
```

3. Run app
```bash
$ go run .
```

4. Check [[http://localhost:8080/swagger/index.html]].
