package echo

func (r *rest) routing() {
	g := r.echo.Group("/api/v1")

	g.POST("/account/login", r.accountController.login)           
	g.POST("/account", r.accountController.register)              
	g.GET("/account/:account_id", r.accountController.getAccount) 

	g.GET("/account/:account_id/friends", r.accountController.getAccountFriends)                         
	g.POST("/account/:account_id/friends/:target_account_id", r.accountController.addAccountToFriends)   
	g.PATCH("/account/:account_id/friends/:target_account_id", r.accountController.blockAccountFriend)   
	g.DELETE("/account/:account_id/friends/:target_account_id", r.accountController.deleteAccountFriend) 

	g.GET("/account/:account_id/game", r.accountController.getAccountGameAccounts) 
	// g.POST("/account/:account_id/game", r.accountController.createAccountGameAccount)

	g.GET("/account/:account_id/game/champions", r.accountController.getAccountGameAccountChampions)                          
	g.POST("/account/:account_id/game/champions/:champion_name", r.accountController.createAccountGameAccountChampion)        
	g.GET("/account/:account_id/game/champions/:champion_name/skins", r.accountController.getAccountGameAccountChampionSkins)
	g.POST("/account/:account_id/game/champions/:champion_name/skins/:skin_name", r.accountController.createAccountGameAccountChampionSkin)

	g.GET("/champions", nil)
	g.GET("/champions/:champion_name", nil)
	g.GET("/champions/:champion_name/skins", nil)
	g.GET("/champions/:champion_name/skins/:skin_name", nil)

}
