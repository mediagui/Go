package repository

type FindTask interface{
	FindByName(name string) Task

}

