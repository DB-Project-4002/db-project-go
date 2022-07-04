package echo

func (r *rest) routing() {

	r.echo.POST("/account/login", r.accountController.login)
	r.echo.POST("/account", r.accountController.register)
	r.echo.GET("/account/:account_id", r.accountController.getAccount)

	r.echo.GET("/account/:account_id/friends", r.accountController.getAccountFriends)
	// r.echo.POST("/account/:account_id/friends/:target_account_id", r.accountController.addAccountToFriends)
	r.echo.POST("/account/:account_id/friends", r.accountController.addAccountToFriendsByUsername)
	r.echo.PATCH("/account/:account_id/friends/:target_account_id", r.accountController.blockAccountFriend)
	r.echo.DELETE("/account/:account_id/friends/:target_account_id", r.accountController.deleteAccountFriend)

	r.echo.GET("/account/:account_id/game", r.accountController.getAccountGameAccounts)
	// r.echo.POST("/account/:account_id/game", r.accountController.createAccountGameAccount)

	r.echo.GET("/account/:account_id/game/champions", r.accountController.getAccountGameAccountChampions)
	r.echo.POST("/account/:account_id/game/champions/:champion_name", r.accountController.createAccountGameAccountChampion)
	r.echo.GET("/account/:account_id/game/champions/:champion_name/skins", r.accountController.getAccountGameAccountChampionSkins)
	r.echo.POST("/account/:account_id/game/champions/:champion_name/skins/:skin_name", r.accountController.createAccountGameAccountChampionSkin)

	r.echo.GET("/game/champions", r.championController.getChampions)
	r.echo.GET("/game/champions/:champion_name", r.championController.getChampion)
	r.echo.GET("/game/champions/:champion_name/skins", r.championController.getChampionSkins)
	r.echo.GET("/game/champions/:champion_name/skins/:skin_name", r.championController.getChampionSkin)

}
