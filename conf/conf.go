package conf

import (
	cmd "github.schq.secious.com/Logrhythm/Godeps/_workspace/src/github.com/codegangsta/cli"
)

const (
	output  = "output"
	pattern = "pattern"
	list    = "list"
	debug   = "debug"
)

type Conf struct {
	cli *cmd.Context
}

func NewConf(cmd *cmd.Context) *Conf {
	return &Conf{
		cli: cmd,
	}
}
func (c *Conf) Debug() bool {
	return c.cli.Bool(pattern)
}
func (c *Conf) Pattern() string {
	return c.cli.String(pattern)
}
func (c *Conf) List() []string {
	return c.cli.StringSlice(list)
}
func (c *Conf) Output() string {
	return c.cli.String(output)
}

func (c *Conf) HasList() bool {
	return c.cli.IsSet(list)
}
func (c *Conf) HasOutput() bool {
	return c.cli.IsSet(output)
}
func (c *Conf) HasPattern() bool {
	return c.cli.IsSet(pattern)
}
func (c *Conf) HasDebug() bool {
	return c.cli.IsSet(debug)
}
