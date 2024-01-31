package usecases

import (
	"context"
	dto "dating-services/src/handler/requests/data_transfer_object"
	"dating-services/src/infra/persistence/model"
	"dating-services/src/infra/persistence/postgres"
	"errors"
	"fmt"
)

type IUsersUsecase interface {
	RegisterUsers(ctx context.Context, do dto.IUsersRequest) (err error)
	LoginUsers(ctx context.Context, do dto.IUsersRequest) (resp *model.Users, err error)
	UpgradeUser(ctx context.Context, do dto.IUsersRequest) (resp *model.Carts, err error)
	PayUpgradeUser(ctx context.Context, do dto.IUsersRequest) (resp *model.PackageOrders, err error)
}

type usersUsecase struct {
	usersRepo        postgres.IUsersRepository
	packagesRepo     postgres.IPackagesRepository
	cartsRepo        postgres.ICartsRepository
	packageOrderRepo postgres.IPackageOrdersRepository
}

// NewUsersUsecase ...
func NewUsersUsecase(u postgres.IUsersRepository, p postgres.IPackagesRepository, c postgres.ICartsRepository, po postgres.IPackageOrdersRepository) IUsersUsecase {
	return &usersUsecase{
		usersRepo:        u,
		packagesRepo:     p,
		cartsRepo:        c,
		packageOrderRepo: po,
	}
}

func (u *usersUsecase) RegisterUsers(ctx context.Context, do dto.IUsersRequest) (err error) {
	registerUser, ok := do.(*dto.UsersRegisterRequestDto)
	if !ok {
		err := errors.New("type assertion failed to dto.UsersRegisterRequestDto")
		return err
	}
	return u.usersRepo.RegisterUsers(ctx, postgres.UserRegisterRequest{
		Email:    registerUser.Email,
		Password: registerUser.Password,
		Username: registerUser.Username,
		Gender:   registerUser.Gender,
	})
}

func (u *usersUsecase) LoginUsers(ctx context.Context, do dto.IUsersRequest) (resp *model.Users, err error) {
	loginUser, ok := do.(*dto.UsersLoginRequestDto)
	if !ok {
		err := errors.New("type assertion failed to dto.UsersLoginRequestDto")
		return resp, err
	}
	return u.usersRepo.LoginUsers(ctx, postgres.UserLoginRequest{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	})
}

func (u *usersUsecase) UpgradeUser(ctx context.Context, do dto.IUsersRequest) (resp *model.Carts, err error) {
	upgradeUser, ok := do.(*dto.UpgradeUserPackageRequestDto)
	if !ok {
		err := errors.New("type assertion failed to dto.UpgradeUserPackageRequestDto")
		return resp, err
	}
	u.createDummyPackage(ctx)

	packageData, err := u.packagesRepo.PackagesDetail(ctx, upgradeUser.PackageId)
	if err != nil {
		return resp, err
	}

	if packageData.ID == 0 {
		return resp, fmt.Errorf("Data Package not found")
	}

	return u.cartsRepo.CreateCart(ctx, postgres.CreateCartRequest{
		UserId:     upgradeUser.UserId,
		PackageId:  upgradeUser.PackageId,
		TotalPrice: packageData.Price + (packageData.Price * 10 / 100),
	})
}

func (u *usersUsecase) PayUpgradeUser(ctx context.Context, do dto.IUsersRequest) (resp *model.PackageOrders, err error) {
	upgradeUser, ok := do.(*dto.PayUpgradeUserRequestDto)
	if !ok {
		err := errors.New("type assertion failed to dto.PayUpgradeUserRequestDto")
		return resp, err
	}

	cartData, err := u.cartsRepo.CartsDetail(ctx, upgradeUser.CartId)
	if err != nil {
		return resp, err
	}

	if cartData.ID == 0 {
		return resp, fmt.Errorf("Data Cart not found")
	}

	err = u.cartsRepo.UpdateCartPaid(ctx, upgradeUser.CartId)
	if err != nil {
		return resp, err
	}

	return u.packageOrderRepo.CreatePackageOrder(ctx, postgres.CreatePackageOrderRequest{
		CartId:     int(cartData.ID),
		PaymentId:  upgradeUser.PaymentMethodId,
		TotalPrice: cartData.TotalPrice,
	})
}

func (u *usersUsecase) createDummyPackage(ctx context.Context) {
	if p, err := u.packagesRepo.PackagesDetailByName(ctx); err == nil && p.ID == 0 {
		u.packagesRepo.CreatePackage(ctx)
	}
}
