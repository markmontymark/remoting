package ui

type BusStops struct {
	collection map[string]BusStop
	PublicAPI
}

func NewBusStops () BusStops {
	obj := BusStops{}
	obj.collection = make(map[string]BusStop,0)
	return obj
}

func(this BusStops) List()[]BusStop {
	retval := make([]BusStop,0)
	for _,v := range this.collection {
		retval = append(retval,v)
	}
	return retval
}

func(this BusStops) View(id string)BusStop{
	for ratingId,rating := range this.collection {
		if id == ratingId {
			return rating
		}
	}
	return BusStop{}
}

func(this BusStops) Add(i BusStop) BusStop{
    for id,item := range this.collection {
        if id == i.Id {
            item = i
            return item
        }
    }

    //Item Id not in database, so create new
    i.Id = string(len(this.collection))
    this.collection[i.Id] = i
    return i
}

func(this BusStops) Delete(id string) {
	for Id, _ := range this.collection {
		if Id == id {
			delete(this.collection,Id) // this borks things that are linked with this rating
			return  //Default http code for DELETE is 200
		}
	}
}


