package usecase

import (
	"context"

	r "mongodbrebe/internal/repository/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

type UserUsecase interface {
	InsertUsers(ctx context.Context, count int) error
	AggregateUsers(ctx context.Context) ([]bson.D, error)
}

type userUsecase struct {
	userRepository r.UserRepository
}

func NewUserUsecase(userRepository r.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) InsertUsers(ctx context.Context, count int) error {
	if err := u.userRepository.InsertUsers(ctx, count); err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) AggregateUsers(ctx context.Context) ([]bson.D, error) {
	users, err := u.userRepository.AggregateUsersToSex(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
