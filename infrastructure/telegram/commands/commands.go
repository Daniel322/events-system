package tg_commands

import "events-system/interfaces"

var COMMANDS = interfaces.Commands{
	"start":   StartCmd,
	"event":   EventCmd,
	"help":    HelpCmd,
	"default": DefaultCmd,
}
