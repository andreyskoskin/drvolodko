package domain

type (
	UserID string

	ListUsersResponse struct {
		Items []ListUsersResponseItem `json:"items"`
	}

	ListUsersResponseItem struct {
		ID   UserID `json:"id"`
		Name string `json:"name"`
	}

	ListUsers interface {
		ListUsers() (ListUsersResponse, error)
	}

	GetUserResponse struct {
		ID   UserID `json:"id"`
		Name string `json:"name"`
	}

	GetUser interface {
		GetUser(id UserID) (GetUserResponse, error)
	}

	AddUserRequest struct {
		Name string `json:"name"`
	}

	AddUserResponse struct {
		ID UserID `json:"id"`
	}

	AddUser interface {
		AddUser(request AddUserRequest) (AddUserResponse, error)
	}

	KillUser interface {
		KillUser(id UserID) error
	}
)
