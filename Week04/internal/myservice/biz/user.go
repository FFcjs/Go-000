package biz

//User
type User struct{
	Name string
}

//UserRepo
type UserRepo interface{
	CreateUser(*User)
}

//NewUserUsecase
func NewUserUsecase(repo UserRepo)*UserUsecase{
	return &UserUsecase{repo:repo}
}

//UserUsecase
type UserUsecase struct{
	repo UserRepo
}


//Get
func (uc *UserUsecase) CreateUser(u *User){
	uc.repo.CreateUser(u)
}
