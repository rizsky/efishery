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
│   └── dev.js
├── controllers
│   ├── auth.js
│   └── user.js
├── db
│   └── db.js
├── helpers
│   ├── helper.js
│   └── response.js
├── models
│   └── users.js
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
