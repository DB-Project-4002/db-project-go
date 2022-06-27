package request

type (
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Register struct {
		Name     string `json:"name" form:"name" validate:"required"`
		Password string `json:"password" form:"password" validate:"required"`
		Tag      string `json:"tag" form:"tag" validate:"required"`
		Email    string `json:"email" form:"email" validate:"required,email"`
	}
)
