package ui

import (
	"fmt"
	"net/http"

	"code.google.com/p/gorest"
	"appengine"
	"appengine/user"
)

var locationStore map[int]Location

type streetType struct {
	name, abbrev string
}

type Address struct {
	number int
	name string
	full string
	streetType
}

type Location struct {
	Address
	Name string
	has_wifi bool
	num_wifi int
	num_outlets int
}


type SimpleService struct{
    //Service level config
    gorest.RestService    `root:"/simple-service/" consumes:"application/json" produces:"application/json"`

    //End-Point level configs: Field names must be the same as the corresponding method names,
    // but not-exported (starts with lowercase)
    //discover    gorest.EndPoint `method:"GET"  path:"/discover/"      output:"Discover"`
    //userDetails gorest.EndPoint `method:"GET"  path:"/users/{Id:int}" output:"User"`
    listLocations   gorest.EndPoint `method:"GET"  path:"/locations/"         output:"[]Location"`
    //listItems   gorest.EndPoint `method:"GET"  path:"/items/"         output:"[]Item"`
    //addItem     gorest.EndPoint `method:"POST" path:"/items/"         postdata:"Item"`

    //On a real app for placeOrder below, the POST URL would probably be just /orders/, this is just to
    // demo the ability of mixing post-data parameters with URL mapped parameters.
    //placeOrder  gorest.EndPoint `method:"POST"   path:"/orders/new/{UserId:int}/{RequestDiscount:bool}" postdata:"Order"`
    //viewOrder   gorest.EndPoint `method:"GET"    path:"/orders/{OrderId:int}"                           output:"Order"`
    //deleteOrder gorest.EndPoint `method:"DELETE" path:"/orders/{OrderId:int}"`
    //listOrders  gorest.EndPoint `method:"GET"  path:"/orders/"         output:"[]Order"`
}

func(serv SimpleService) ListLocations()[]Location{
	serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
	retval := make([]Location,0)
	for _,v := range locationStore {
		retval = append(retval,v)
	}
	return retval
}



func init() {
	gorest.RegisterService(new(SimpleService))
	http.Handle("/",gorest.Handle())
	http.HandleFunc("/squirrel", appHandler)
	http.HandleFunc("/user", userHandler)
}

func appHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, quirreld!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Hello, %v!", u)
}
