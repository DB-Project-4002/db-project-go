package response

type (
	Login struct {
		Data struct {
			Token string `json:"token"`
		}
	}

	Register struct {
		Data struct {
			Token string `json:"token"`
		}
	}
)
