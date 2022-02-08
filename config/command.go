package config

type Command struct {
	Name   string
	Cmd    string
	SubCmd *Command
	Desc   string
}
