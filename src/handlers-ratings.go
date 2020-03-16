package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handleratebuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handleaddchat Has Been Called!")
		//get JSON payload
		rating := StartRating{}
		err := json.NewDecoder(r.Body).Decode(&rating)
		fmt.Println("Handle rate buyer Has Been Called..")
		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to rate buyer ")
			return
		}
		//set response variables
		var buyerrated bool
		var advertisementid string

		//communcate with the database
		querystring := "SELECT * FROM public.ratebuyer('" + rating.AdvertisementID + "','" + rating.SellerID + "','" + rating.BuyerID + "','" + rating.BuyerRating + "','" + rating.BuyerComments + "')"

		//retrieve message from database tt set to JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&buyerrated, &advertisementid)

		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to rate buyer")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to rate buyer")
			return
		}

		//set JSON object variables for response
		startratingResult := StartRatingResult{}
		startratingResult.BuyerRated = buyerrated
		startratingResult.AdvertisementID = advertisementid

		if buyerrated {
			startratingResult.Message = "Buyer sucessfully rated!"
		} else {
			startratingResult.Message = "Buyer has not been rated!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(startratingResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to rate buyer")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
