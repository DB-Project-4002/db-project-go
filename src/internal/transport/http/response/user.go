package response

type (
	Login struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	Register struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}
)
