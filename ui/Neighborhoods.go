package ui

type Neighborhoods struct {
	collection map[string]Neighborhood
	PublicAPI
}

func NewNeighborhoods () Neighborhoods {
	obj := Neighborhoods{}
	obj.collection = make(map[string]Neighborhood,0)
	return obj
}

func(this Neighborhoods) List()[]Neighborhood {
	retval := make([]Neighborhood,0)
	for _,v := range this.collection {
		retval = append(retval,v)
	}
	return retval
}

func(this Neighborhoods) View(id string)Neighborhood{
	for ratingId,rating := range this.collection {
		if id == ratingId {
			return rating
		}
	}
	return Neighborhood{}
}

func(this Neighborhoods) Add(i Neighborhood) Neighborhood{
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

func(this Neighborhoods) Delete(id string) {
	for Id, _ := range this.collection {
		if Id == id {
			delete(this.collection,Id) // this borks things that are linked with this rating
			return  //Default http code for DELETE is 200
		}
	}
}


