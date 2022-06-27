package response

type (
	LoginData struct {
		Token string `json:"token"`
	}
	Login struct {
		Data LoginData `json:"data"`
	}

	RegisterData struct {
		Token string `json:"token"`
	}
	Register struct {
		Data RegisterData `json:"data"`
	}
)
