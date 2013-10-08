package ui

import (
	"fmt"
	"net/http"
	"time"

	"code.google.com/p/gorest"

	"appengine"
	//"appengine/datastore"
	"appengine/user"

	"encoding/json"
)

var ratings Ratings
var users Users //map[int]User 
var foods Foods
var busstops BusStops
var routes Routes
var neighborhoods Neighborhoods
var workplaces Workplaces
var restaurants Restaurants
var meetingplaces MeetingPlaces

type Mappable interface {
	GetMapLink()
}

type Ided struct {
	Id string
}

type Locator interface {
	Locate()
}

type Worker interface {
	Work()
}

type Address struct {
	street string
	neighborhood string
}
/*
type Food struct {
	name string
	rating Rating // 0-9 scale, 1-5 stars?
	Rater
}
*/

type Location struct {
	Name string
	has_wifi bool
	num_wifi int
	num_outlets int
	Locator
	Rater
	Seating
}

type BusStopLocation struct {
	Locator
	Rater
	Mappable
}
func (this BusStopLocation) Locate () {
}
func (this BusStopLocation) Rate () {
}
type NeighborhoodLocation struct {
	Location
}

type Seating struct {
	num_chairs, 
		num_outlets, 
		available_chairs, 
		available_outlets int
}

/*
type WorkplaceLocation struct {
	Location
	Locator
	Rating
	Rater
	Seating
	WifiUpness Rating
}

type FoodLocation struct {
	Location
	Locator
	Rating
	Rater	
	Seating
}

type MeetingLocation struct {
	Location
	Locator
	Rating
	Rater	
}

*/

//type Route struct {
	//busStops []BusStopLocation
//}

type WorkLog struct {
	timeIn time.Time
	timeOut time.Time
	Route
	Location
}

// or is this more of a query (ie give me []worklog entries for date y-m-d
type DaysWork struct {
	date time.Time
	worklog []WorkLog
}

func init() {

	ratings = NewRatings() //make(map[int]Rating)
	users = NewUsers() //	users = make(map[int]User)
	foods = NewFoods()
	busstops = NewBusStops()
	routes = NewRoutes()
	neighborhoods = NewNeighborhoods()
	workplaces = NewWorkplaces()
	restaurants = NewRestaurants()
	meetingplaces = NewMeetingPlaces()

	gorest.RegisterService(new(SimpleService))
	http.Handle("/",gorest.Handle())
	http.HandleFunc("/user", userHandler)

	http.HandleFunc("/dump",dumpHandler)
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

func dumpHandler(w http.ResponseWriter, r *http.Request) {
	b,_ := json.Marshal(ratings)
	fmt.Fprintf(w, "Ratings\n----\n%s\n\n", b)
	b,_  = json.Marshal(users)
	fmt.Fprintf(w, "Users\n----\n%s\n\n", b)
	b,_  = json.Marshal(foods)
	fmt.Fprintf(w, "Foods\n----\n%s\n\n",  b)
	b,_  = json.Marshal(busstops)
	fmt.Fprintf(w, "BusStops\n----\n%s\n\n",  b)
	b,_  = json.Marshal(routes)
	fmt.Fprintf(w, "Routes\n----\n%s\n\n",  b)
	b,_  = json.Marshal(neighborhoods)
	fmt.Fprintf(w, "Neighborhoods\n----\n%s\n\n",  b)
	b,_  = json.Marshal(workplaces)
	fmt.Fprintf(w, "Workplaces\n----\n%s\n\n",  b)
	b,_  = json.Marshal(restaurants)
	fmt.Fprintf(w, "Restaurants\n----\n%s\n\n",  b)
	b,_  = json.Marshal(meetingplaces)
	fmt.Fprintf(w, "MeetingPlaces \n----\n%s\n\n", b)
}
