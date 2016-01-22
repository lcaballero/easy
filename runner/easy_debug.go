package runner

import (
	cmd "github.schq.secious.com/Logrhythm/Godeps/_workspace/src/github.com/codegangsta/cli"
	"fmt"
	"github.schq.secious.com/Logrhythm/easy/tasks/copy_dx_logs"
	"github.schq.secious.com/Logrhythm/easy/tasks"
	"github.schq.secious.com/Logrhythm/easy/conf"
)

func Run(cmd *cmd.Context) {
	conf := conf.NewConf(cmd)
	f := NewFinder()
	f.Run()

	imps := f.Imports

	if conf.Debug() {
		fmt.Println(imps)
	}

	for k,v := range Gatherers {
		fmt.Println(k, v.Usage())
	}
}

var Gatherers = map[string]tasks.Gatherer{
	"copy_dx_logs":&copy_dx_logs.Copier{},
}
