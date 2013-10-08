package ui

type Lister interface {
	List() []interface{}
}

type Viewer interface {
	View() interface{}
}

type Deleter interface {
	Delete(id interface{}) 
}

type Adder interface {
	Add(obj interface{})
}

type PublicAPI interface {
	Lister
	Viewer
	Deleter
	Adder
}

