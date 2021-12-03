package usecase

import (
	"SuperPanel/domain"
	"context"
	"time"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func (this *userUsecase) GetByID(ctx context.Context, id uint) (domain.User, error) {
	// جزییات پیاده سازی اینجا قراار میگیره
	user, err := this.userRepo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func NewUserUsecase(a domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       a,
		contextTimeout: timeout,
	}
}
