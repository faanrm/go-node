import { DataTypes } from 'sequelize';
import db from '../config/sql.js';

const User = db.define('users', {
  user_id: {
    type: DataTypes.INTEGER(11),
    allowNull: false,
    primaryKey: true,
    autoIncrement: true
  },
  name: {
    type: DataTypes.STRING(100),
    allowNull: false
  },
  email : {
    type : DataTypes.STRING(100),
    allowNull : false , 
    unique : true ,
    validate : {
      isEmail : true
    }
  },
  password : {
   type : DataTypes.STRING(100),
   allowNull: false 
  }
}, {
  timestamps: false,
  freezeTableName: true,
  tableName: 'users'
});

export default User;
