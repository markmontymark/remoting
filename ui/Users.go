package ui

type Users struct {
	collection map[string]User
	PublicAPI
}

func NewUsers () Users {
	obj := Users{}
	obj.collection = make(map[string]User,0)
	return obj
}

func(this Users) List()[]User {
	retval := make([]User,0)
	for _,v := range this.collection {
		retval = append(retval,v)
	}
	return retval
}

func(this Users) View(id string)User{
	for ratingId,rating := range this.collection {
		if id == ratingId {
			return rating
		}
	}
	return User{}
}

func(this Users) Add(i User) User{
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

func(this Users) Delete(id string) {
	for Id, _ := range this.collection {
		if Id == id {
			delete(this.collection,Id) // this borks things that are linked with this rating
			return  //Default http code for DELETE is 200
		}
	}
}


