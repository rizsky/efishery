import express from 'express';

import user from '../../controllers/user';

const routes  = express.Router();

/**
 * @swagger
 * 
 * /user:
 *    post:
 *      tags:
 *          - user
 *      summary: this endpoint for sign up use only on port 9000
 *      description: Input phone, name and role and get return the password after signup
 *      consumes:
 *        - application/json
 *      produces:
 *        - application/json
 *      parameters:
 *        - name: body
 *          in: body
 *          schema:
 *            type: object
 *            properties:
 *              phone:
 *                type: string
 *              role:
 *                type: string
 *              name:
 *                type: string
 *              
 *      responses:
 *        201:
 *          description: password return.
 */

routes.route('/')
  .post(user.create);

module.exports = routes;
