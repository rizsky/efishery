# Efishery Fetch Service

this is fetch service with based golang using gin as web framework, bigcache as cache lib in memory of server

## Project Layout
```
fetching-service
.
├── README.md
├── cache
│   └── cache.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
├── module
│   └── fetch
│       ├── handlers
│       ├── helpers
│       ├── middlewares
│       ├── models
│       ├── repositories
│       ├── routes
│       └── services
└── refresh.yml
```

## Start this service
i provide docker-compose in this root project. But if want install standalone, you need :
- go ver 1.1.4

run this command on **terminal** :
- go get -v

if you have `refresh` you can go with `refresh run` yet alternatives you can start service with 
- `go run*.go/`
- `go run main.go serve`

I Provide swagger-ui for these endpoints too, you can access it on 
`localhost:9001/api-docs/index.html` 
