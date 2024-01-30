package usecases

type IUsersUsecase interface {
	RegisterUsers()
}

type usersUsecase struct {
}

// NewUsersUsecase ...
func NewUsersUsecase() IUsersUsecase {
	return &usersUsecase{}
}

func (u *usersUsecase) RegisterUsers() {}
