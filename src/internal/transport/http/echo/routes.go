package echo

func (r *rest) routing() {
	g := r.echo.Group("/api/v1")

	g.Group("/user")

	g.POST("/user/login", r.userController.login)
	g.POST("/user", r.userController.register)
	g.GET("user/:id", nil)

	g.GET("user/:user_id/friends", nil)
	g.POST("user/:user_id/friends/:target_user_id", nil)
	g.PATCH("user/:user_id/friends/:target_user_id", nil)
	g.DELETE("user/:user_id/friends/:target_user_id", nil)

	g.GET("user/:user_id/game", nil)
	g.POST("user/:user_id/game", nil)

	g.GET("user/:user_id/game/champions", nil)
	g.POST("user/:user_id/game/champions/:champion_name", nil)
	g.GET("user/:user_id/game/champions/:champion_name/skins", nil)
	g.POST("user/:user_id/game/champions/:champion_name/skins/{skin_name}", nil)
}
