import express from 'express';

import user from './user';
import auth from './auth';
import response from '../helpers/response';
const swaggerUi = require('swagger-ui-express');
import config from '../config/dev';
import swaggerJSDoc from 'swagger-jsdoc'

const specs = swaggerJSDoc(config.swaggerOption)


const routes  = express.Router();
routes.use('/api-docs', swaggerUi.serve, swaggerUi.setup(specs));

routes.use('/auth', auth)
routes.use('/user', user);

routes.get('/', (req, res) => {
  res.status(200).json({ message: 'Ok' });
});

routes.use(function(req, res) {
  response.sendNotFound(res);
});

module.exports = routes;
