package request

type (
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Register struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Tag      string `json:"tag"`
		Email    string `json:"email"`
	}
)
