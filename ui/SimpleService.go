

package ui

import (
   "code.google.com/p/gorest"
   //"appengine"
   //"appengine/datastore"
   //"appengine/user"
)

type SimpleService struct{
    //Service level config
    gorest.RestService    `root:"/simple-service/" consumes:"application/json" produces:"application/json"`

	
// For the User type
listUsers   gorest.EndPoint `method:"GET"    path:"/users/"               output:"[]User"`
viewUser    gorest.EndPoint `method:"GET"    path:"/user/{UserId:string}"  output:"User"`
addUser     gorest.EndPoint `method:"POST"   path:"/user/"                postdata:"User"`
deleteUser  gorest.EndPoint `method:"DELETE" path:"/user/{UserId:string}"`



// For the Rating type
listRatings   gorest.EndPoint `method:"GET"    path:"/ratings/"               output:"[]Rating"`
viewRating    gorest.EndPoint `method:"GET"    path:"/rating/{RatingId:string}"  output:"Rating"`
addRating     gorest.EndPoint `method:"POST"   path:"/rating/"                postdata:"Rating"`
deleteRating  gorest.EndPoint `method:"DELETE" path:"/rating/{RatingId:string}"`



// For the Food type
listFoods   gorest.EndPoint `method:"GET"    path:"/foods/"               output:"[]Food"`
viewFood    gorest.EndPoint `method:"GET"    path:"/food/{FoodId:string}"  output:"Food"`
addFood     gorest.EndPoint `method:"POST"   path:"/food/"                postdata:"Food"`
deleteFood  gorest.EndPoint `method:"DELETE" path:"/food/{FoodId:string}"`



// For the BusStop type
listBusStops   gorest.EndPoint `method:"GET"    path:"/busstops/"               output:"[]BusStop"`
viewBusStop    gorest.EndPoint `method:"GET"    path:"/busstop/{BusStopId:string}"  output:"BusStop"`
addBusStop     gorest.EndPoint `method:"POST"   path:"/busstop/"                postdata:"BusStop"`
deleteBusStop  gorest.EndPoint `method:"DELETE" path:"/busstop/{BusStopId:string}"`



// For the Neighborhood type
listNeighborhoods   gorest.EndPoint `method:"GET"    path:"/neighborhoods/"               output:"[]Neighborhood"`
viewNeighborhood    gorest.EndPoint `method:"GET"    path:"/neighborhood/{NeighborhoodId:string}"  output:"Neighborhood"`
addNeighborhood     gorest.EndPoint `method:"POST"   path:"/neighborhood/"                postdata:"Neighborhood"`
deleteNeighborhood  gorest.EndPoint `method:"DELETE" path:"/neighborhood/{NeighborhoodId:string}"`



// For the Workplace type
listWorkplaces   gorest.EndPoint `method:"GET"    path:"/workplaces/"               output:"[]Workplace"`
viewWorkplace    gorest.EndPoint `method:"GET"    path:"/workplace/{WorkplaceId:string}"  output:"Workplace"`
addWorkplace     gorest.EndPoint `method:"POST"   path:"/workplace/"                postdata:"Workplace"`
deleteWorkplace  gorest.EndPoint `method:"DELETE" path:"/workplace/{WorkplaceId:string}"`



// For the Restaurant type
listRestaurants   gorest.EndPoint `method:"GET"    path:"/restaurants/"               output:"[]Restaurant"`
viewRestaurant    gorest.EndPoint `method:"GET"    path:"/restaurant/{RestaurantId:string}"  output:"Restaurant"`
addRestaurant     gorest.EndPoint `method:"POST"   path:"/restaurant/"                postdata:"Restaurant"`
deleteRestaurant  gorest.EndPoint `method:"DELETE" path:"/restaurant/{RestaurantId:string}"`



// For the MeetingPlace type
listMeetingPlaces   gorest.EndPoint `method:"GET"    path:"/meetingplaces/"               output:"[]MeetingPlace"`
viewMeetingPlace    gorest.EndPoint `method:"GET"    path:"/meetingplace/{MeetingPlaceId:string}"  output:"MeetingPlace"`
addMeetingPlace     gorest.EndPoint `method:"POST"   path:"/meetingplace/"                postdata:"MeetingPlace"`
deleteMeetingPlace  gorest.EndPoint `method:"DELETE" path:"/meetingplace/{MeetingPlaceId:string}"`



// For the Route type
listRoutes   gorest.EndPoint `method:"GET"    path:"/routes/"               output:"[]Route"`
viewRoute    gorest.EndPoint `method:"GET"    path:"/route/{RouteId:string}"  output:"Route"`
addRoute     gorest.EndPoint `method:"POST"   path:"/route/"                postdata:"Route"`
deleteRoute  gorest.EndPoint `method:"DELETE" path:"/route/{RouteId:string}"`



}

func(serv SimpleService) ListUsers()[]User{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return users.List()
}

func(serv SimpleService) ViewUser(id string)User{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return users.View(id)
}

func(serv SimpleService) AddUser(i User){
   itemAdded := users.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/users/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteUser(id string) {
   users.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListRatings()[]Rating{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return ratings.List()
}

func(serv SimpleService) ViewRating(id string)Rating{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return ratings.View(id)
}

func(serv SimpleService) AddRating(i Rating){
   itemAdded := ratings.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/ratings/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteRating(id string) {
   ratings.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListFoods()[]Food{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return foods.List()
}

func(serv SimpleService) ViewFood(id string)Food{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return foods.View(id)
}

func(serv SimpleService) AddFood(i Food){
   itemAdded := foods.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/foods/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteFood(id string) {
   foods.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListBusStops()[]BusStop{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return busstops.List()
}

func(serv SimpleService) ViewBusStop(id string)BusStop{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return busstops.View(id)
}

func(serv SimpleService) AddBusStop(i BusStop){
   itemAdded := busstops.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/busstops/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteBusStop(id string) {
   busstops.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListNeighborhoods()[]Neighborhood{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return neighborhoods.List()
}

func(serv SimpleService) ViewNeighborhood(id string)Neighborhood{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return neighborhoods.View(id)
}

func(serv SimpleService) AddNeighborhood(i Neighborhood){
   itemAdded := neighborhoods.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/neighborhoods/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteNeighborhood(id string) {
   neighborhoods.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListWorkplaces()[]Workplace{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return workplaces.List()
}

func(serv SimpleService) ViewWorkplace(id string)Workplace{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return workplaces.View(id)
}

func(serv SimpleService) AddWorkplace(i Workplace){
   itemAdded := workplaces.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/workplaces/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteWorkplace(id string) {
   workplaces.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListRestaurants()[]Restaurant{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return restaurants.List()
}

func(serv SimpleService) ViewRestaurant(id string)Restaurant{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return restaurants.View(id)
}

func(serv SimpleService) AddRestaurant(i Restaurant){
   itemAdded := restaurants.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/restaurants/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteRestaurant(id string) {
   restaurants.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListMeetingPlaces()[]MeetingPlace{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return meetingplaces.List()
}

func(serv SimpleService) ViewMeetingPlace(id string)MeetingPlace{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return meetingplaces.View(id)
}

func(serv SimpleService) AddMeetingPlace(i MeetingPlace){
   itemAdded := meetingplaces.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/meetingplaces/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteMeetingPlace(id string) {
   meetingplaces.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}



func(serv SimpleService) ListRoutes()[]Route{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return routes.List()
}

func(serv SimpleService) ViewRoute(id string)Route{
   serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
   return routes.View(id)
}

func(serv SimpleService) AddRoute(i Route){
   itemAdded := routes.Add(i)
   serv.ResponseBuilder().Created(
      "http://localhost:8080/simple-service/routes/"+string(itemAdded.Id))
}

func(serv SimpleService) DeleteRoute(id string) {
   routes.Delete(id)
   serv.ResponseBuilder().SetResponseCode(404).Overide(true)
   return
}




