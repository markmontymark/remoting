package ui

type Routes struct {
	collection map[int]Route
	PublicAPI
}

func NewRoutes () Routes {
	obj := Routes{}
	obj.collection = make(map[int]Route,0)
	return obj
}

func(this Routes) List()[]Route {
	retval := make([]Route,0)
	for _,v := range this.collection {
		retval = append(retval,v)
	}
	return retval
}

func(this Routes) View(id int)Route{
	for ratingId,rating := range this.collection {
		if id == ratingId {
			return rating
		}
	}
	return Route{}
}

func(this Routes) Add(i Route) Route{
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

func(this Routes) Delete(id int) {
	for Id, _ := range this.collection {
		if Id == id {
			delete(this.collection,Id) // this borks things that are linked with this rating
			return  //Default http code for DELETE is 200
		}
	}
}


