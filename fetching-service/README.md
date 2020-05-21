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


## Cache explanation
Why i use big-cache over redis for this project, because i saw its too overkill using redis because just 1 key that needed store and the currency api said its update over 60 minutes to take the new one, so i set cache running in memory for 10 minutes as long server alive, and the limit from free are 100 hits. With this cache strategy it would be keep the service run like for unlimited (paid membership of currency converter lol).

## Explanation for endpoint {GET} /v1/admin/list
1. i added `query_param` based on `area_provinsi` and `weekly` this `weekly` i adjust to from time to time with standart of `RFC3339 golang package time` (see the swagger i already attach the default value). both of them is optional, wether one of them is exist or both. i already cover the logic (and if both of them null too) 

2. i do some adjustment and scope on my own vision and assumption because its not clear wether what data that needed for (avg, median, min, and max) i assumption this is `price idr` field


