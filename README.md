# Swaggo Demo

Check [gin-swagger](https://github.com/swaggo/gin-swagger) for details.

1. Install swag
```bash
$ go get -u github.com/swaggo/swag/cmd/swag
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

4. Check swagger [index](http://localhost:8080/swagger/index.html) page.
