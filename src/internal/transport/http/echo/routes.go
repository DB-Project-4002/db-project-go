package echo

func (r *rest) routing() {

	r.echo.POST("/account/login", r.accountController.login)
	r.echo.POST("/account", r.accountController.register)
	r.echo.GET("/account/:account_id", r.accountController.getAccount)

	ag := r.echo.Group("/account/:account_id", r.accountMiddleware.OnlyAccountOwner)

	ag.GET("/friends", r.accountController.getAccountFriends)
	// ag.POST("/friends/:target_account_id", r.accountController.addAccountToFriends)
	ag.POST("/friends", r.accountController.addAccountToFriendsByUsername)
	ag.PATCH("/friends/:target_account_id", r.accountController.blockAccountFriend)
	ag.DELETE("/friends/:target_account_id", r.accountController.deleteAccountFriend)

	ag.GET("/game", r.accountController.getAccountGameAccounts)
	// ag.POST("/game", r.accountController.createAccountGameAccount)

	ag.GET("/game/champions", r.accountController.getAccountGameAccountChampions)
	ag.POST("/game/champions/:champion_name", r.accountController.createAccountGameAccountChampion)
	ag.GET("/game/champions/:champion_name/skins", r.accountController.getAccountGameAccountChampionSkins)
	ag.POST("/game/champions/:champion_name/skins/:skin_name", r.accountController.createAccountGameAccountChampionSkin)

	r.echo.GET("/game/champions", r.championController.getChampions)
	r.echo.GET("/game/champions/:champion_name", r.championController.getChampion)
	r.echo.GET("/game/champions/:champion_name/skins", r.championController.getChampionSkins)
	r.echo.GET("/game/champions/:champion_name/skins/:skin_name", r.championController.getChampionSkin)

}
