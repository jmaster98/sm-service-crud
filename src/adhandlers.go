package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handlepostadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		advertisement := PostAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&advertisement)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to add advertisement ")
			return
		}
		//set response variables
		var advertisementposted bool
		var id, message string

		//communcate with the database
		querystring := "SELECT * FROM public.addadvertisement('" + advertisement.UserID + "','" + advertisement.AdvertisementType + "','" + advertisement.EntityID + "','" + advertisement.Price + "','" + advertisement.Description + "')"

		//retrieve message from database tt set to JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&advertisementposted, &id, &message)

		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add advertisement")
			return
		}

		//set JSON object variables for respinse
		postAdvertisementResult := PostAdvertisementResult{}
		postAdvertisementResult.AdvertisementPosted = advertisementposted
		postAdvertisementResult.ID = id
		postAdvertisementResult.Message = message

		//convert struct back to JSON
		js, jserr := json.Marshal(postAdvertisementResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to post advert")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve URL from ad service
		getadvertisementid := r.URL.Query().Get("id")
		advertisementid := AdvertisementID{}
		advertisementid.AdvertisementID = getadvertisementid

		//set response variables
		var id, userid, advertisementtype, entityid, price, description string

		//communcate with the database
		querystring := "SELECT * FROM public.getadvertisement('" + advertisementid.AdvertisementID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&id, &userid, &advertisementtype, &entityid, &price, &description)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to get advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to get advertisement")
			return
		}
		//fmt.Println("This is Advertisement!: " + id)
		advertisement := getAdvertisement{}
		advertisement.AdvertisementID = id
		advertisement.UserID = userid
		advertisement.AdvertisementType = advertisementtype
		advertisement.EntityID = entityid
		advertisement.Price = price
		advertisement.Description = description

		//convert struct back to JSON
		js, jserr := json.Marshal(advertisement)
		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to get advertisement")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdateadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		advertisement := UpdateAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&advertisement)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to update advertisement")
			return
		}

		//set response variables
		var advertisementupdated bool
		var msg string

		//communcate with the database
		querystring := "SELECT * FROM public.updateadvertisement('" + advertisement.AdvertisementID + "','" + advertisement.UserID + "','" + advertisement.AdvertisementType + "','" + advertisement.EntityID + "','" + advertisement.Price + "','" + advertisement.Description + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&advertisementupdated, &msg)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to update advertisement")
			fmt.Println("Error in communicating with database to update advertisement")
			return
		}

		updateAdvertisementResult := UpdateAdvertisementResult{}
		updateAdvertisementResult.AdvertisementUpdated = advertisementupdated
		updateAdvertisementResult.Message = msg

		//convert struct back to JSON
		js, jserr := json.Marshal(updateAdvertisementResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to update advertisement")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleremoveadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve ID from advert service
		getadvertisementid := r.URL.Query().Get("id")

		advertisementid := AdvertisementID{}
		advertisementid.AdvertisementID = getadvertisementid

		var advertisementDeleted bool
		querystring := "SELECT * FROM public.deleteadvertisement('" + advertisementid.AdvertisementID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&advertisementDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete advertisement")
			fmt.Println("Error in communicating with database to delete advertisement")
			return
		}

		//set response variables

		deleteAdvertisementResult := DeleteAdvertisementResult{}
		deleteAdvertisementResult.AdvertisementDeleted = advertisementDeleted
		deleteAdvertisementResult.AdvertisementID = getadvertisementid

		if advertisementDeleted {
			deleteAdvertisementResult.Message = "Advert Successfully Deleted!"
		} else {
			deleteAdvertisementResult.Message = "Unable to Delete Advert!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(deleteAdvertisementResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete advert")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}
