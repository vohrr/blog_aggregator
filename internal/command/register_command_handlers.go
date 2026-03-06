package command

func RegisterCommandHandlers(cmds *Commands) {
	cmds.Register("login", LoginHandler)
	cmds.Register("register", RegisterHandler)
	cmds.Register("reset", ResetHandler)
	cmds.Register("users", UsersHandler)
	cmds.Register("agg", AggHandler)
	cmds.Register("addfeed", AddFeedHandler)
	cmds.Register("feeds", FeedsHandler)
}
