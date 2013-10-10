package ui

type MeetingPlaces struct {
	collection map[int]MeetingPlace
	PublicAPI
}

func NewMeetingPlaces () MeetingPlaces {
	obj := MeetingPlaces{}
	obj.collection = make(map[int]MeetingPlace,0)
	return obj
}

func(this MeetingPlaces) List()[]MeetingPlace {
	retval := make([]MeetingPlace,0)
	for _,v := range this.collection {
		retval = append(retval,v)
	}
	return retval
}

func(this MeetingPlaces) View(id int)MeetingPlace{
	for ratingId,rating := range this.collection {
		if id == ratingId {
			return rating
		}
	}
	return MeetingPlace{}
}

func(this MeetingPlaces) Add(i MeetingPlace) MeetingPlace{
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

func(this MeetingPlaces) Delete(id int) {
	for Id, _ := range this.collection {
		if Id == id {
			delete(this.collection,Id) // this borks things that are linked with this rating
			return  //Default http code for DELETE is 200
		}
	}
}


