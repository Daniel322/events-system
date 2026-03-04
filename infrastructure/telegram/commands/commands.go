package tg_commands

import "events-system/interfaces"

var COMMANDS = interfaces.Commands{
	"start":   StartCmd,
	"info":    InfoCmd,
	"event":   EventCmd,
	"help":    HelpCmd,
	"default": DefaultCmd,
}
