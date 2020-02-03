package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handledeleteuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hande Delete User Has Been Called..")
		userid := UserID{}
		err := json.NewDecoder(r.Body).Decode(&userid)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided...")
			return
		}

		var userDeleted bool
		querystring := "SELECT * FROM public.deleteuser('" + userid.UserID + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&userDeleted)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			fmt.Println(err.Error())
			return
		}
		deleteUserResult := DeleteUserResult{}
		deleteUserResult.UserDeleted = userDeleted

		js, jserr := json.Marshal(deleteUserResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}

func (s *Server) handleupdateuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Update User Has Been Called...")
		user := updateUser{}
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided...")
			return
		}

		var userUpdated bool
		querystring := "SELECT * FROM public.updateuser('" + user.UserID + "','" + user.Username + "','" + user.Password + "','" + user.Name + "','" + user.Surname + "','" + user.Email + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&userUpdated)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			fmt.Println(err.Error())
			return
		}

		updateUserResult := UpdateUserResult{}
		updateUserResult.UserUpdated = userUpdated

		js, jserr := json.Marshal(updateUserResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleregisteruser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Register User Has Been Called!")
		user := User{}
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided...")
			return
		}
		var userCreated bool
		var username, userid string

		querystring := "SELECT * FROM public.registeruser('" + user.Username + "','" + user.Password + "','" + user.Name + "','" + user.Surname + "','" + user.Email + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&userCreated, &username, &userid)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			fmt.Println(err.Error())
			return
		}

		regUserResult := RegisterUserResult{}
		regUserResult.UserCreated = userCreated
		regUserResult.Username = username
		regUserResult.UserID = userid

		js, jserr := json.Marshal(regUserResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get User Has Been Called...")
		userid := UserID{}
		err := json.NewDecoder(r.Body).Decode(&userid)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided...")
			return
		}

		var id, username, name, surname, email string

		querystring := "SELECT * FROM public.getuser('" + userid.UserID + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&id, &username, &name, &surname, &email)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			fmt.Println(err.Error())
			return
		}

		user := getUser{}
		user.UserID = id
		user.Username = username
		user.Name = name
		user.Surname = surname
		user.Email = email

		js, jserr := json.Marshal(user)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlerespond() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Println("Hit!")
		return
	}
}
