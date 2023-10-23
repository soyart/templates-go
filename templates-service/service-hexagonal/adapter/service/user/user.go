package user

import (
	"context"
	"fmt"
	"reflect"

	"github.com/pkg/errors"

	"example.com/servicehex/domain/datagateway"
	"example.com/servicehex/domain/entity"
	"example.com/servicehex/internal/pwhash"
)

func register(
	ctx context.Context,
	repo datagateway.DataGatewayUser,
	user entity.User,
) error {
	if len(user.Password) == 0 {
		return fmt.Errorf("missing password")
	}

	if len(user.Username) == 0 {
		return fmt.Errorf("missing username")
	}

	hashed, err := pwhash.Hash(user.Password)
	if err != nil {
		return errors.WithStack(err)
	}

	user.Password = hashed

	if err := repo.CreateUser(ctx, user); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func login(
	ctx context.Context,
	repo datagateway.DataGatewayUser,
	username string,
	password []byte,
) (
	entity.User,
	error,
) {
	user, err := repo.GetUserByUsername(ctx, username)
	if err != nil {
		return entity.User{}, nil
	}

	if err := pwhash.Compare(user.Password, []byte(password)); err != nil {
		return entity.User{}, errors.WithStack(err)
	}

	// Prevent leak
	user.Password = nil

	return user, nil
}

func changePassword(
	ctx context.Context,
	repo datagateway.DataGatewayUser,
	userId string,
	password []byte,
	newPassword []byte,
) error {
	user, err := repo.GetUser(ctx, userId)
	if err != nil {
		return errors.WithStack(err)
	}

	if reflect.DeepEqual(password, newPassword) {
		return fmt.Errorf("new password is identical to old password")
	}

	hashed, err := pwhash.Hash(newPassword)
	if err != nil {
		return errors.WithStack(err)
	}

	user.Password = hashed

	if err := repo.UpdateUser(ctx, userId, user); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func deleteUser(
	ctx context.Context,
	repo datagateway.DataGatewayUser,
	userId string,
) error {
	_, err := repo.GetUser(ctx, userId)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := repo.DeleteUser(ctx, userId); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
