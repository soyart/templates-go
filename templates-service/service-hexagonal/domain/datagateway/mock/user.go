package mock

import (
	"context"
	"fmt"

	"example.com/servicehex/domain/datagateway"
	"example.com/servicehex/domain/entity"
)

type mockRepoUser struct {
	users map[string]entity.User
}

func NewMockRepoUser() datagateway.DataGatewayUser {
	return &mockRepoUser{users: make(map[string]entity.User)}
}

func (m *mockRepoUser) CreateUser(ctx context.Context, user entity.User) error {
	_, found := m.users[user.ID]
	if found {
		return fmt.Errorf("duplicate user id %s", user.ID)
	}

	m.users[user.ID] = user
	return nil
}

func (m *mockRepoUser) GetUser(ctx context.Context, userID string) (entity.User, error) {
	u, found := m.users[userID]
	if !found {
		return u, fmt.Errorf("userID %s not found", userID)
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

func (m *mockRepoUser) UpdateUser(ctx context.Context, userID string, user entity.User) error {
	savedUser, found := m.users[userID]
	if !found {
		return fmt.Errorf("no such userID %s", user.ID)
	}
	if savedUser.ID != userID {
		panic("mock bug - unexpected userID")
	}

	m.users[userID] = user
	return nil
}

func (m *mockRepoUser) DeleteUser(ctx context.Context, userID string) error {
	savedUser, found := m.users[userID]
	if !found {
		return fmt.Errorf("")
	}

	if savedUser.ID != userID {
		panic("mock bug - unexpected userID")
	}

	delete(m.users, userID)
	return nil
}
