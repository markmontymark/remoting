package ui

type Restaurants struct {
	collection map[int]Restaurant
	PublicAPI
}

func NewRestaurants () Restaurants {
	obj := Restaurants{}
	obj.collection = make(map[int]Restaurant,0)
	return obj
}

func(this Restaurants) List()[]Restaurant {
	retval := make([]Restaurant,0)
	for _,v := range this.collection {
		retval = append(retval,v)
	}
	return retval
}

func(this Restaurants) View(id int)Restaurant{
	for ratingId,rating := range this.collection {
		if id == ratingId {
			return rating
		}
	}
	return Restaurant{}
}

func(this Restaurants) Add(i Restaurant) Restaurant{
    for id,item := range this.collection {
        if id == i.Id {
            item = i
            return item
        }
    }

    //Item Id not in database, so create new
    i.Id = len(this.collection)
    this.collection[i.Id] = i
    return i
}

func(this Restaurants) Delete(id int) {
	for Id, _ := range this.collection {
		if Id == id {
			delete(this.collection,Id) // this borks things that are linked with this rating
			return  //Default http code for DELETE is 200
		}
	}
}


