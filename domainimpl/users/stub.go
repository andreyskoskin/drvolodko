package users

import (
	"github.com/google/uuid"

	"github.com/andreyskoskin/drvolodko/domain"
)

type StubUsers struct{}

func (StubUsers) GetUser(_ domain.UserID) (domain.GetUserResponse, error) {
	return domain.GetUserResponse{
		ID:   "123",
		Name: "John Smith",
	}, nil
}

func (StubUsers) ListUsers() (domain.ListUsersResponse, error) {
	return domain.ListUsersResponse{
		Items: []domain.ListUsersResponseItem{
			{
				ID:   "123",
				Name: "John Smith",
			},
		},
	}, nil
}

func (StubUsers) AddUser(_ domain.AddUserRequest) (domain.AddUserResponse, error) {
	var id, err = uuid.NewUUID()
	if err != nil {
		return domain.AddUserResponse{}, err
	}

	return domain.AddUserResponse{
		ID: domain.UserID(id.String()),
	}, nil
}

func (StubUsers) KillUser(_ domain.UserID) error {
	return nil
}
