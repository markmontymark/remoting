package ui

type Ratings struct {
	collection map[string]Rating
	PublicAPI
}

func NewRatings () Ratings {
	obj := Ratings{}
	obj.collection = make(map[string]Rating,0)
	return obj
}

func(this Ratings) List()[]Rating {
	retval := make([]Rating,0)
	for _,v := range this.collection {
		retval = append(retval,v)
	}
	return retval
}

func(this Ratings) View(id string)Rating{
	for ratingId,rating := range this.collection {
		if id == ratingId {
			return rating
		}
	}
	return Rating{}
}

func(this Ratings) Add(i Rating) Rating{
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

func(this Ratings) Delete(id string) {
	for Id, _ := range this.collection {
		if Id == id {
			delete(this.collection,Id) // this borks things that are linked with this rating
			return  //Default http code for DELETE is 200
		}
	}
}


