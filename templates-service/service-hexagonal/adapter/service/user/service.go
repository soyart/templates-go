package user

import (
	"context"

	"github.com/pkg/errors"

	"example.com/servicehex/domain/core"
	"example.com/servicehex/domain/datagateway"
	"example.com/servicehex/domain/entity"
)

type service struct {
	repo datagateway.DataGatewayUser
}

func New(repo datagateway.DataGatewayUser) core.PortUser {
	return &service{repo: repo}
}

func (s *service) Register(
	ctx context.Context,
	user entity.User,
) error {
	if err := register(ctx, s.repo, user); err != nil {
		return errors.Wrapf(err, "failed to register user %s", user.Username)
	}

	return nil
}

func (s *service) Login(
	ctx context.Context,
	username string,
	password []byte,
) (
	entity.User,
	error,
) {
	user, err := login(ctx, s.repo, username, password)
	if err != nil {
		return entity.User{}, errors.Wrapf(err, "failed to login user %s", username)
	}

	return user, nil
}

func (s *service) ChangePassword(
	ctx context.Context,
	userID string,
	password []byte,
	newPassword []byte,
) error {
	if err := changePassword(ctx, s.repo, userID, password, newPassword); err != nil {
		return errors.Wrapf(err, "failed to change password for userID %s", userID)
	}

	return nil
}

func (s *service) DeleteUser(ctx context.Context, userID string) error {
	if err := deleteUser(ctx, s.repo, userID); err != nil {
		return errors.Wrapf(err, "failed to delete userID %s", userID)
	}

	return nil
}
