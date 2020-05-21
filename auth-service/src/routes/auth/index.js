import express from 'express';

import auth from '../../controllers/auth';

const routes  = express.Router();
/**
 * @swagger
 * 
 * /auth:
 *    post:
 *      tags:
 *          - auth
 *      summary: this endpoint for sign in use only on port 9000
 *      description: Input phone and password to retrive the jwt token that crafted from mongodb data
 *      consumes:
 *        - application/json
 *      parameters:
 *        - name: body
 *          in: body
 *          schema:
 *            type: object
 *            properties:
 *              phone:
 *                type: string
 *              password:
 *                type: string
 *      responses:
 *        200:
 *          description: jwt is returned in response.
 */

routes.route('/')
  .post(auth.authenticate);

/**
 * @swagger
 * 
 * /auth/private-claim:
 *    get:
 *      tags:
 *          - auth
 *      summary: this endpoint for sign in use only on port 9000
 *      description: Input phone and password to retrive the jwt token that crafted from mongodb data
 *      consumes:
 *        - application/json
 *      parameters:
 *        - default: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiRXhjZWxlbmN5IiwicGhvbmUiOiIrNjI4MjI3MjQyMjMyMCIsInJvbGUiOiJ1c2VyIiwidGltZXN0YW1wIjoiMjAyMC0wNS0xOVQwNDo1MzowNS40NDNaIiwiaWF0IjoxNTg5ODY0MDE0fQ.pnQXvaa9vFlIM4AKnptrbzbfztEGPXWU-ISDH-K1Vsg
 *          name: authorization
 *          in: header
 *      responses:
 *        200:
 *          description: jwt is returned in response.
 */

routes.route('/private-claim')
  .get(auth.verifyToken);

module.exports = routes;
