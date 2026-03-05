package command

func RegisterCommandHandlers(cmds *Commands) {
	cmds.Register("login", LoginHandler)
	cmds.Register("register", RegisterHandler)
	cmds.Register("reset", ResetHandler)
}
