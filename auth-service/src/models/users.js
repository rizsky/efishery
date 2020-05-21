import mongoose from 'mongoose';
import helper from '../helpers/helper'

const Schema = mongoose.Schema;
const UsersSchema = new Schema({
  phone: {
    type: String,
    required: true,
    index: { unique: true }
  },
  password: {
    type: String
  },
  role: {
    type: String,
    enum : ['user', 'admin'],
    default: 'user'
  },
  name: {
    type: String,
    required: true,
    index: { unique: true }
  },
  timestamp: {type: Date, default: Date.now}
});

UsersSchema.set('toJSON', {
  transform: function(doc, ret, options) {
    delete ret.__v
    return ret;
  }
});

UsersSchema.pre('save', function(next) {
  this.password = helper.generatePass(4)
  next();
});

UsersSchema.methods.getTokenData = function() {
  return {
    name: this.name,
    phone: this.phone,
    role: this.role,
    timestamp: this.timestamp
  }
};

UsersSchema.methods.verifyPassword = function(candidatePassword=false){
    if(candidatePassword === this.password) {
      return true
  }
}
module.exports = mongoose.model('Users', UsersSchema);
