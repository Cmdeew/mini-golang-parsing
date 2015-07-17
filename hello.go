/*	
	Author: Thomas Rieux-Laucat
*/

package hello

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"appengine"
	"appengine/urlfetch"

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

func guestbookKey(c appengine.Context) *datastore.Key {
        // The string "default_guestbook" here could be varied to have multiple guestbooks.
        return datastore.NewKey(c, "Video", "Video", 0, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World ! Let's start\n\n")

	//Get
	url := "https://api.dailymotion.com/channel/music/videos"
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	res, err := client.Get(url)
	if err != nil{
		panic(err.Error())
	}
	//Read
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err.Error())
	}
	//Decode
	var data cont
	err = json.Unmarshal([]byte(body), &data)
	if err != nil{
		panic(err.Error())
	}
	//write
	fmt.Fprint(w, data.List)
	for track := range data.List{
	    v := Video{
	            Id: track.Id,
	            Title: track.Title,
	            Channel: track.Channel,
	            Owner: track.Owner,
	    }	
	    key := datastore.NewIncompleteKey(c, "Video", guestbookKey(c))
	    _, err := datastore.Put(c, key, &v)
	    if err != nil {
	            http.Error(w, err.Error(), http.StatusInternalServerError)
	            return
	    }
	    http.Redirect(w, r, "/", http.StatusFound)	
	}
}