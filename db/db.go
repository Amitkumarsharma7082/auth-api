package db

import "auth-api/model"

/*
key (Email) -> Value(complete struct)
Users -> amit@.com -> Name, Email, Password
*/
var Users = map[string]model.User{}
