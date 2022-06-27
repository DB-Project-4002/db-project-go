package response

type (
	Login struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	RegisterData struct {
		Token string `json:"token"`
	}
	Register struct {
		Data RegisterData `json:"data"`
	}
)
