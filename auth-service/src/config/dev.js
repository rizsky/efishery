const host = process.env.DB_HOST || 'localhost';


const newLocal = 'kikyn0secret';
module.exports = {
  server: {
    port: 9000
  },
  database: {
    url: `mongodb://${host}/efishery-auth-service`
  },
  properties: {
    useNewUrlParser: true,
    useMongoClient: true,
    useUnifiedTopology: true 
  },
  key: {
    privateKey: newLocal,
    tokenExpireInMinutes: 1440
  },
  swaggerOption : {
    apis: ["src/routes/**/*.js"],
    swaggerDefinition: {
      info: {
        title: "Auth service",
        version: '1.0.0', 
        description: "API efishery auth",
        contact: {
          name: "Muhammad Rizsky"
        },
        servers: ["http://localhost:9000"]
      }
    },
  }
};
