import mongoose from 'mongoose';
import response from '../helpers/response';
import jwt from 'jsonwebtoken';
import config from 'config';

const privateKey = config.key.privateKey;
const Users = mongoose.model('Users');

  exports.authenticate = function(req, res) {
    Users.findOne({ phone: req.body.phone })
    .exec(function(err, user) {
      if (err) throw err;  
      if (!user) {
        response.sendUnauthorized(res, 'Authentication failed.');
      } else if (user) {
        if (user.verifyPassword(req.body.password)){
            const token = jwt.sign(user.getTokenData(), privateKey);
            res.json({
              success: true,
              message: 'Token created.',
              token: token            
            });
          } else {
            response.sendUnauthorized(res, 'Authentication failed.');
          }
      }
    });
  }

  exports.verifyToken = function(req, res) {
    let token = req.body.token || req.query.token || req.headers.authorization;
    if (token) {
      token = token.replace('Bearer ', '')
      jwt.verify(token, privateKey, function(err, decoded) {
        if (err) {
          response.sendUnauthorized(res, 'Failed to authenticate token.');
        } else {
          res.json({
            success: true,
            message: 'jwt claimed',
            data: decoded            
          });

        }
      });
    } else {
      response.sendUnauthorized(res, 'No token provided.');
    }
  };