package echo

func (r *rest) routing() {
	g := r.echo.Group("/api/v1")

	g.POST("/account/login", r.accountController.login)
	g.POST("/account", r.accountController.register)
	g.GET("/account/:id", nil)

	g.GET("/account/:account_id/friends", nil)
	g.POST("/account/:account_id/friends/:target_account_id", nil)
	g.PATCH("/account/:account_id/friends/:target_account_id", nil)
	g.DELETE("/account/:account_id/friends/:target_account_id", nil)

	g.GET("/account/:account_id/game", nil)
	g.POST("/account/:account_id/game", nil)

	g.GET("/account/:account_id/game/champions", nil)
	g.POST("/account/:account_id/game/champions/:champion_name", nil)
	g.GET("/account/:account_id/game/champions/:champion_name/skins", nil)
	g.POST("/account/:account_id/game/champions/:champion_name/skins/:skin_name", nil)
}
