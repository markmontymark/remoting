package ui

type Workplaces struct {
	collection map[string]Workplace
	PublicAPI
}

func NewWorkplaces () Workplaces {
	obj := Workplaces{}
	obj.collection = make(map[string]Workplace,0)
	return obj
}

func(this Workplaces) List()[]Workplace {
	retval := make([]Workplace,0)
	for _,v := range this.collection {
		retval = append(retval,v)
	}
	return retval
}

func(this Workplaces) View(id string)Workplace{
	for ratingId,rating := range this.collection {
		if id == ratingId {
			return rating
		}
	}
	return Workplace{}
}

func(this Workplaces) Add(i Workplace) Workplace{
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

func(this Workplaces) Delete(id string) {
	for Id, _ := range this.collection {
		if Id == id {
			delete(this.collection,Id) // this borks things that are linked with this rating
			return  //Default http code for DELETE is 200
		}
	}
}


