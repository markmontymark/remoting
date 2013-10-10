package ui

type Foods struct {
	collection map[int]Food
	PublicAPI
}

func NewFoods () Foods {
	obj := Foods{}
	obj.collection = make(map[int]Food,0)
	return obj
}

func(this Foods) List()[]Food {
	retval := make([]Food,0)
	for _,v := range this.collection {
		retval = append(retval,v)
	}
	return retval
}

func(this Foods) View(id int)Food{
	for ratingId,rating := range this.collection {
		if id == ratingId {
			return rating
		}
	}
	return Food{}
}

func(this Foods) Add(i Food) Food{
		
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

func(this Foods) Delete(id int) {
	for Id, _ := range this.collection {
		if Id == id {
			delete(this.collection,Id) // this borks things that are linked with this rating
			return  //Default http code for DELETE is 200
		}
	}
}


