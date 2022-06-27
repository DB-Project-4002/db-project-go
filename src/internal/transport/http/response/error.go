package response

type (
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	ErrorResp struct {
		Error Error `json:"error"`
	}
)
