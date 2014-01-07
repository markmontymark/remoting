

package ui

import (
   "code.google.com/p/gorest"
   "appengine"
   "appengine/datastore"
	//"fmt"
	//"os"
   //"appengine/user"
)

type SimpleService struct{
    //Service level config
    gorest.RestService    `root:"/simple-service/" consumes:"application/json" produces:"application/json"`

options     		gorest.EndPoint `method:"OPTIONS" path:"/"`
optionsFoods     	gorest.EndPoint `method:"OPTIONS" path:"/foods/{FoodId:int}"`
	
// For the User type
listUsers   gorest.EndPoint `method:"GET"    path:"/users/"               output:"[]User"`
viewUser    gorest.EndPoint `method:"GET"    path:"/user/{UserId:int}"  output:"User"`
addUser     gorest.EndPoint `method:"POST"   path:"/user/"                postdata:"User"`
deleteUser  gorest.EndPoint `method:"DELETE" path:"/user/{UserId:int}"`



// For the Rating type
listRatings   gorest.EndPoint `method:"GET"    path:"/ratings/"               output:"[]Rating"`
viewRating    gorest.EndPoint `method:"GET"    path:"/rating/{RatingId:int}"  output:"Rating"`
addRating     gorest.EndPoint `method:"POST"   path:"/rating/"                postdata:"Rating"`
deleteRating  gorest.EndPoint `method:"DELETE" path:"/rating/{RatingId:int}"`



// For the Food type
listFoods   gorest.EndPoint `method:"GET"    path:"/foods/"               output:"[]Food"`
viewFood    gorest.EndPoint `method:"GET"    path:"/food/{FoodId:int}"  output:"Food"`
addFood     gorest.EndPoint `method:"POST"   path:"/food/"                postdata:"Food"`
deleteFood  gorest.EndPoint `method:"DELETE" path:"/food/{FoodId:int}"`
deleteFoods  gorest.EndPoint `method:"DELETE" path:"/foods/{FoodId:int}"`



// For the BusStop type
listBusStops   gorest.EndPoint `method:"GET"    path:"/busstops/"               output:"[]BusStop"`
viewBusStop    gorest.EndPoint `method:"GET"    path:"/busstop/{BusStopId:int}"  output:"BusStop"`
addBusStop     gorest.EndPoint `method:"POST"   path:"/busstop/"                postdata:"BusStop"`
deleteBusStop  gorest.EndPoint `method:"DELETE" path:"/busstop/{BusStopId:int}"`



// For the Neighborhood type
listNeighborhoods   gorest.EndPoint `method:"GET"    path:"/neighborhoods/"               output:"[]Neighborhood"`
viewNeighborhood    gorest.EndPoint `method:"GET"    path:"/neighborhood/{NeighborhoodId:int}"  output:"Neighborhood"`
addNeighborhood     gorest.EndPoint `method:"POST"   path:"/neighborhood/"                postdata:"Neighborhood"`
deleteNeighborhood  gorest.EndPoint `method:"DELETE" path:"/neighborhood/{NeighborhoodId:int}"`



// For the Workplace type
listWorkplaces   gorest.EndPoint `method:"GET"    path:"/workplaces/"               output:"[]Workplace"`
viewWorkplace    gorest.EndPoint `method:"GET"    path:"/workplace/{WorkplaceId:int}"  output:"Workplace"`
addWorkplace     gorest.EndPoint `method:"POST"   path:"/workplace/"                postdata:"Workplace"`
deleteWorkplace  gorest.EndPoint `method:"DELETE" path:"/workplace/{WorkplaceId:int}"`



// For the Restaurant type
listRestaurants   gorest.EndPoint `method:"GET"    path:"/restaurants/"               output:"[]Restaurant"`
viewRestaurant    gorest.EndPoint `method:"GET"    path:"/restaurant/{RestaurantId:int}"  output:"Restaurant"`
addRestaurant     gorest.EndPoint `method:"POST"   path:"/restaurant/"                postdata:"Restaurant"`
deleteRestaurant  gorest.EndPoint `method:"DELETE" path:"/restaurant/{RestaurantId:int}"`



// For the MeetingPlace type
listMeetingPlaces   gorest.EndPoint `method:"GET"    path:"/meetingplaces/"               output:"[]MeetingPlace"`
viewMeetingPlace    gorest.EndPoint `method:"GET"    path:"/meetingplace/{MeetingPlaceId:int}"  output:"MeetingPlace"`
addMeetingPlace     gorest.EndPoint `method:"POST"   path:"/meetingplace/"                postdata:"MeetingPlace"`
deleteMeetingPlace  gorest.EndPoint `method:"DELETE" path:"/meetingplace/{MeetingPlaceId:int}"`



// For the Route type
listRoutes   gorest.EndPoint `method:"GET"    path:"/routes/"               output:"[]Route"`
viewRoute    gorest.EndPoint `method:"GET"    path:"/route/{RouteId:int}"  output:"Route"`
addRoute     gorest.EndPoint `method:"POST"   path:"/route/"                postdata:"Route"`
deleteRoute  gorest.EndPoint `method:"DELETE" path:"/route/{RouteId:int}"`



}

func(serv SimpleService) ListUsers()[]User{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day

   return users.List()
}

func(serv SimpleService) ViewUser(id int)User{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return users.View(id)
}

func(serv SimpleService) AddUser(i User){
   itemAdded := users.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/users/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteUser(id int) {
   users.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListRatings()[]Rating{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return ratings.List()
}

func(serv SimpleService) ViewRating(id int)Rating{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return ratings.View(id)
}

func(serv SimpleService) AddRating(i Rating){
   itemAdded := ratings.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/ratings/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteRating(id int) {
   ratings.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListFoods()[]Food{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin","http://localhost")
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Headers","origin, x-requested-with, content-type")
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Methods","POST,GET,PUT,OPTIONS,DELETE")
   return foods.List()
}

func(serv SimpleService) ViewFood(id int)Food{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return foods.View(id)
}

func(serv SimpleService) AddFood(i Food){
   itemAdded := foods.Add(i)
	ctx := appengine.NewContext( serv.Context.Request())
	/*key, err := */datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "food", nil), &itemAdded)
	/*
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	*/

   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/foods/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteFood(id int) {
   foods.Delete(id)
   serv.ResponseBuilder().SetResponseCode(204).Overide(true)
   return
}
func(serv SimpleService) DeleteFoods(id int) {
	serv.DeleteFood(id)
}



func(serv SimpleService) ListBusStops()[]BusStop{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return busstops.List()
}

func(serv SimpleService) ViewBusStop(id int)BusStop{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return busstops.View(id)
}

func(serv SimpleService) AddBusStop(i BusStop){
   itemAdded := busstops.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/busstops/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteBusStop(id int) {
   busstops.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}


func(serv SimpleService) ListNeighborhoods()[]Neighborhood{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return neighborhoods.List()
}

func(serv SimpleService) ViewNeighborhood(id int)Neighborhood{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return neighborhoods.View(id)
}

func(serv SimpleService) AddNeighborhood(i Neighborhood){
   itemAdded := neighborhoods.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/neighborhoods/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteNeighborhood(id int) {
   neighborhoods.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListWorkplaces()[]Workplace{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return workplaces.List()
}

func(serv SimpleService) ViewWorkplace(id int)Workplace{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return workplaces.View(id)
}

func(serv SimpleService) AddWorkplace(i Workplace){
   itemAdded := workplaces.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/workplaces/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteWorkplace(id int) {
   workplaces.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListRestaurants()[]Restaurant{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return restaurants.List()
}

func(serv SimpleService) ViewRestaurant(id int)Restaurant{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return restaurants.View(id)
}

func(serv SimpleService) AddRestaurant(i Restaurant){
   itemAdded := restaurants.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/restaurants/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteRestaurant(id int) {
   restaurants.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListMeetingPlaces()[]MeetingPlace{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return meetingplaces.List()
}

func(serv SimpleService) ViewMeetingPlace(id int)MeetingPlace{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return meetingplaces.View(id)
}

func(serv SimpleService) AddMeetingPlace(i MeetingPlace){
   itemAdded := meetingplaces.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/meetingplaces/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteMeetingPlace(id int) {
   meetingplaces.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListRoutes()[]Route{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return routes.List()
}

func(serv SimpleService) ViewRoute(id int)Route{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day
   return routes.View(id)
}

func(serv SimpleService) AddRoute(i Route){
   itemAdded := routes.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/routes/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteRoute(id int) {
   routes.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}




func(serv SimpleService) Options() {
	//fmt.Fprintf( os.Stderr, "in OptionsFood with id %v\n",id)
   serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin","http://localhost")
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Headers","origin, x-requested-with, content-type")
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Methods","POST,GET,PUT,OPTIONS,DELETE")
   serv.ResponseBuilder().SetResponseCode(200).Overide(true)
   return
}
func(serv SimpleService) OptionsFoods(id int) {
	serv.Options()
	return
}
