/*
	Author: Thomas Rieux-Laucat
*/

package hello

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"appengine"
	"appengine/datastore"
	//"appengine/urlfetch"

)

type cont 		struct {
	Page 		int `json:"page"`
	Limit 		int `json:"limit"`
	Explicit 	bool `json:"explicit"`
	Total 		int `json:"total"`
	Has_more 	bool `json:"has_more"`
	List 		[]Video `json:"list"`
}

type Video 	struct {
	Id 		string `json:"id"`
	Title 	string `json:"title"`
	Channel string `json:"channel"`
	Owner 	string `json:"owner"`
}

func init() {
	http.HandleFunc("/home", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World ! Let's start\n\n")

	//Getter
	//url := "https://api.dailymotion.com/channel/music/videos"
	c := appengine.NewContext(r)
	//client := urlfetch.Client(c)
	//res, err := client.Get(url)
	/*if err != nil{
		panic(err.Error())
	}*/
	//local
	file, err := os.Open("data.json")
	//Read
	body, err := ioutil.ReadAll(file) //res.Body / file -> travail local
	if err != nil{
		panic(err.Error())
	}
	//Decode
	var data cont
	err = json.Unmarshal([]byte(body), &data)
	if err != nil{
		panic(err.Error())
	}

	for i , track := range data.List {
	k := datastore.NewKey(c, "Video", track.Id, 0, nil)

		e := Video{
			Id: 		track.Id,
			Title:		track.Title,
			Channel:	track.Channel,
			Owner:		track.Owner,
		}
		fmt.Fprint(w, "##--->KEY : ", k)
		fmt.Fprint(w, "\n")
		fmt.Fprintf(w, "Envoi de: ID: [%q], Title: [%q]", e.Id, e.Title)
		fmt.Fprint(w, "\n")
		if _, err := datastore.Put(c, k, &e); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//DEBUG
		/*e2 := new(Video)
		if err = datastore.Get(c, k, e2); err != nil {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
        	return
    	}
    	*/
    	i++
	}
}