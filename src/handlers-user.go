package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The function handling the request of a user trying to log in.
func (s *Server) handleloginuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Login User Has Been Called...")

		// retrieving details that the user entered to log in from URL with GET.
		getusername := r.URL.Query().Get("username")
		getpassword := r.URL.Query().Get("password")

		// Creating new struct for user login
		userLogin := UserLogin{}

		// setting struct variables to retrieved variables.
		userLogin.Username = getusername
		userLogin.Password = getpassword

		// declaring variables to catch response from database.
		var userid, username string
		var successLogin bool

		// build query string.
		querystring := "SELECT * FROM public.loginuser('" + userLogin.Username + "','" + userLogin.Password + "')"

		// querying the database and scanning database results into variables.
		err := s.dbAccess.QueryRow(querystring).Scan(&userid, &username, &successLogin)

		// checking for any errors with reading db response into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in scanning the return variables from sql function into specified variables.")
			return
		}
		loginUserResult := LoginUserResult{}

		// checking to determine whether this user was found and if his login details were correct accoding to our database.
		if userid == "00000000-0000-0000-0000-000000000000" {
			loginUserResult.UserLoggedIn = successLogin
			loginUserResult.UserID = ""
			loginUserResult.Username = username
			loginUserResult.Message = "Wrong username and password combination for user: " + username + " !"
		} else {
			loginUserResult.UserLoggedIn = successLogin
			loginUserResult.UserID = userid
			loginUserResult.Username = username
			loginUserResult.Message = "Welcome! "  + username
		}

		// converting response struct to JSON payload to send to service that called this function.
		js, jserr := json.Marshal(loginUserResult)

		// check to see if any errors occured with coverting to JSON.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to log user in")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

// The function handling the request to delete a user.
func (s *Server) handledeleteuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Delete User Has Been Called..")

		// retrieving the ID of the user that is requested to be deleted.
		getuserid := r.URL.Query().Get("id")
		userid := UserID{}
		userid.UserID = getuserid
		
		// declaring variable to catch response from database.
		var userDeleted bool

		// building query string.
		querystring := "SELECT * FROM public.deleteuser('" + userid.UserID + "')"

		// querying the database and reading response from database into variable.
		err := s.dbAccess.QueryRow(querystring).Scan(&userDeleted)

		// check for errors with reading response from database into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to delete user")
			return
		}

		// declaring result struct for delete user.
		deleteUserResult := DeleteUserResult{}
		deleteUserResult.UserDeleted = userDeleted
		deleteUserResult.UserID = getuserid

		// determine if user was infact deleted.
		if userDeleted {
			deleteUserResult.Message = "User Successfully Deleted!"
		} else {
			deleteUserResult.Message = "Unable to Delete User!"
		}

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(deleteUserResult)

		// check to see if any errors occured with coverting to JSON.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}

// The function handling the request to update a users details.
func (s *Server) handleupdateuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Update User Has Been Called...")
		// declare a updateUser struct.
		user := updateUser{}

		// convert received JSON payload into the declared struct.
		err := json.NewDecoder(r.Body).Decode(&user)

		//check for errors when converting JSON payload into struct.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to update user")
			return
		}

		// declare variables to catch response from database.
		var userUpdated bool
		var msg string

		// building query string.
		querystring := "SELECT * FROM public.updateuser('" + user.UserID + "','" + user.Username + "','" + user.Password + "','" + user.Name + "','" + user.Surname + "','" + user.Email + "')"

		// query the database and read results into variables.
		err = s.dbAccess.QueryRow(querystring).Scan(&userUpdated, &msg)

		// check for errors with reading database result into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to update user")
			return
		}

		// instansiate result struct.
		updateUserResult := UpdateUserResult{}
		updateUserResult.UserUpdated = userUpdated
		updateUserResult.Message = msg

		// convert struct into JSON payload to send to service that called this fuction.
		js, jserr := json.Marshal(updateUserResult)

		// check for errors in converting struct to JSON.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to update user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

// The function handling the request to register a user.
func (s *Server) handleregisteruser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// declare a user struct.
		user := User{}

		// convert received JSON payload into user struct.
		err := json.NewDecoder(r.Body).Decode(&user)

		// check for errors with converting received JSON payload into user struct.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to register user ")
			return
		}

		// declare variables to catch response from database.
		var userCreated string
		var username, userid, msg string

		// create query string.
		querystring := "SELECT * FROM public.registeruser('" + user.Username + "','" + user.Password + "','" + user.Name + "','" + user.Surname + "','" + user.Email + "')"

		// query database and read response from database into variables.
		err = s.dbAccess.QueryRow(querystring).Scan(&userCreated, &username, &userid, &msg)

		// check for any errors with reading database response into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to register user")
			return
		}

		// instansiate result struct.
		regUserResult := RegisterUserResult{}
		regUserResult.UserCreated = userCreated
		regUserResult.Username = username
		regUserResult.UserID = userid
		regUserResult.Message = msg

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(regUserResult)

		//check for errors with converting struct into JSON payload.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to register user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

// The function handling the request to get a users details based on their userID
func (s *Server) handlegetuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get User Has Been Called...")
		// retrieving the ID of the user that is requested.
		getuserid := r.URL.Query().Get("id")
		userid := UserID{}
		userid.UserID = getuserid

		// declare variables to catch response from database.
		var id, username, name, surname, email string

		// create query string.
		querystring := "SELECT * FROM public.getuser('" + userid.UserID + "')"

		//query database and read response from database into variables.
		err := s.dbAccess.QueryRow(querystring).Scan(&id, &username, &name, &surname, &email)
		
		// check for errors with reading reponse from database into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get user")
			return
		}

		// instansiate response struct.
		user := getUser{}
		user.UserID = id
		user.Username = username
		user.Name = name
		user.Surname = surname
		user.Email = email

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(user)
		
		// check for errors when converting struct into JSON payload.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to get user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
