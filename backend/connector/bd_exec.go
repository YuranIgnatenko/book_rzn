package connector

var insert_user = `INSERT Users(Id, Login, Password, Name, Phone, Email, Type, Token) 
VALUES ( '%s', '%s', '%s','%s', '%s', '%s', '%s');`
