  # Efishery  Services
  ## Brief Description
  This project have 2 services 
  1. auth service `running on port 9000`
  2. fetching service `running on port 9001`
  
  ## Installation
  The most easiest way is use docker-compose, i already provided dockerfile for all needed instance and docker-compose, do in current path :
  `docker-compose up`
  
or giving arg `-d` for running on background
`docker-compose up -d`

## How to start
After compose up, i already configured the ***swagger-ui*** already running, but its running on each service (not centralized )
so if you wanna using swagger i recommend go to these url:
1. `localhost:9000/api-docs` for auth service
2. `localhost:9001/api-docs/index.html` for fetching service


# Efishery Auth Service

 
this is auth service with based node.js and mongodb using express.js as the framework.

## Project Layout

```

auth-service

├── README.md

├── package.json

└── src

  

src

├── app.js

├── config

│ └── dev.js

├── controllers

│ ├── auth.js

│ └── user.js

├── db

│ └── db.js

├── helpers

│ ├── helper.js

│ └── response.js

├── models

│ └── users.js

└── routes

├── auth

├── index.js

└── user

```

  

## Start this service

i provide docker-compose in this root project. But if want install standalone, you need :

- node version ^8

- mongo db version ^3

  

Make sure the mongodb already running on local. Clone all the folder and do this command on **terminal** :

- npm install

- npm run start

  

see the terminal log for successfuly running the project.

  

I Provide swagger-ui for these endpoints too, you can access it on :

`localhost:9000/api-docs`

  

# Efishery Fetch Service

  

this is fetch service with based golang using gin as web framework, bigcache as cache lib in memory of server

  

## Project Layout

```

fetching-service

.

├── README.md

├── cache

│ └── cache.go

├── docs

│ ├── docs.go

│ ├── swagger.json

│ └── swagger.yaml

├── go.mod

├── go.sum

├── main.go

├── module

│ └── fetch

│ ├── handlers

│ ├── helpers

│ ├── middlewares

│ ├── models

│ ├── repositories

│ ├── routes

│ └── services

└── refresh.yml

```

  

## Start this service

i provide docker-compose in this root project. But if want install standalone, you need :

- go ver 1.1.4

  

run this command on **terminal** :

- go get -v

  

if you have `refresh` you can go with `refresh run` yet alternatives you can start service with

-  `go run*.go/`

-  `go run main.go serve`

  

I Provide swagger-ui for these endpoints too, you can access it on

`localhost:9001/api-docs/index.html`

## Cache explanation
Why i use big-cache over redis for this project, because i saw its too overkill using redis because just 1 key that needed store and the currency api said its update over 60 minutes to take the new one, so i set cache running in memory for 10 minutes as long server alive, and the limit from free are 100 hits. With this cache strategy it would be keep the service run like for unlimited (paid membership of currency converter lol).
