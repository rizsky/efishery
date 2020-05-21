import mongoose from 'mongoose';
import response from '../helpers/response';

const Users = mongoose.model('Users');

exports.create = function(req, res) {
  const newUser = new Users(req.body);  
  newUser.save(function(err, user) {
    if (err) return response.sendBadRequest(res, err);
    response.sendCreated(res, user);
  });};

