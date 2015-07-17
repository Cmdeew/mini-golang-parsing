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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World ! Let's start\n\n")

	url := "https://api.dailymotion.com/channel/music/videos"
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	res, err := client.Get(url)
	if err != nil{
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	fmt.Fprint(w, string(body))
	if err != nil{
		panic(err.Error())
	}
	//Decode
	var data cont
	err = json.Unmarshal([]byte(body), &data)
	if err != nil{
		panic(err.Error())
	}

	fmt.Fprint(w, "\n!!!!!!!!!! DEBUG !!!!!!!!!!!!!!!!\n")
	fmt.Fprint(w, data)
	/*for i, track := range data.List{
		fmt.Fprint(w, track.Id, i)
	}*/
}