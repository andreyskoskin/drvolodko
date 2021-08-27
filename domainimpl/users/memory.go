package users

import (
	"sync"

	"github.com/google/uuid"

	"github.com/andreyskoskin/drvolodko/domain"
)

type MemoryUsers struct {
	users sync.Map
}

type memoryUser struct {
	id   uuid.UUID
	name string
}

func (m *MemoryUsers) GetUser(id domain.UserID) (domain.GetUserResponse, error) {
	var item, ok = m.users.Load(string(id))
	if !ok {
		return domain.GetUserResponse{}, domain.ErrNotFound
	}

	var user, _ = item.(memoryUser)

	return domain.GetUserResponse{
		ID:   domain.UserID(user.id.String()),
		Name: user.name,
	}, nil
}

func (m *MemoryUsers) ListUsers() (domain.ListUsersResponse, error) {
	var users []domain.ListUsersResponseItem

	m.users.Range(func(_, value interface{}) bool {
		var user, _ = value.(memoryUser)
		users = append(users, domain.ListUsersResponseItem{
			ID:   domain.UserID(user.id.String()),
			Name: user.name,
		})
		return true
	})

	return domain.ListUsersResponse{
		Items: users,
	}, nil
}

func (m *MemoryUsers) AddUser(request domain.AddUserRequest) (domain.AddUserResponse, error) {
	var id, err = uuid.NewUUID()
	if err != nil {
		return domain.AddUserResponse{}, err
	}

	m.users.Store(id.String(), memoryUser{
		id:   id,
		name: request.Name,
	})

	return domain.AddUserResponse{
		ID: domain.UserID(id.String()),
	}, nil
}

func (m *MemoryUsers) KillUser(id domain.UserID) error {
	m.users.Delete(id)
	return nil
}
