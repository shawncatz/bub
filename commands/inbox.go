package commands

var cmdInbox = &Command{
	Run:          gitInbox,
	GitExtension: false,
	Usage:        "inbox -c",
	Long:         `Get list of PRs awaiting your review`,
}

func init() {
	CmdRunner.Use(cmdInbox)
}

func gitInbox(command *Command, args *Args) {

}
