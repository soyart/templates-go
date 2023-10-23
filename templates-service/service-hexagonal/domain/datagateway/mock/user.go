package mock

import (
	"context"
	"fmt"

	"example.com/servicehex/domain/entity"
)

type mockRepoUser struct {
	users map[string]entity.User
}

func (m *mockRepoUser) CreateUser(ctx context.Context, user entity.User) error {
	_, found := m.users[user.Id]
	if found {
		return fmt.Errorf("duplicate user id %s", user.Id)
	}

	m.users[user.Id] = user
	return nil
}

func (m *mockRepoUser) GetUser(ctx context.Context, userId string) (entity.User, error) {
	u, found := m.users[userId]
	if !found {
		return u, fmt.Errorf("userId %s not found", userId)
	}

	return u, nil
}

func (m *mockRepoUser) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	for _, u := range m.users {
		if u.Username == username {
			return u, nil
		}
	}

	return entity.User{}, fmt.Errorf("no such username %s", username)
}

func (m *mockRepoUser) UpdateUser(ctx context.Context, userId string, user entity.User) error {
	savedUser, found := m.users[userId]
	if !found {
		return fmt.Errorf("no such userId %s", user.Id)
	}
	if savedUser.Id != userId {
		panic("mock bug - unexpected userId")
	}

	m.users[userId] = user
	return nil
}

func (m *mockRepoUser) DeleteUser(ctx context.Context, userId string) error {
	savedUser, found := m.users[userId]
	if !found {
		return fmt.Errorf("")
	}

	if savedUser.Id != userId {
		panic("mock bug - unexpected userId")
	}

	delete(m.users, userId)
	return nil
}
