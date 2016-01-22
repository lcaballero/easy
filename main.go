package main

import (
	"fmt"
	"os"

	"github.schq.secious.com/Logrhythm/easy/cli"
	"github.schq.secious.com/Logrhythm/easy/compress"
)

func main() {
	cli.NewCli().Run(os.Args)
}

func Run() {
	name := "service.logs.zip"
	f1 := "/var/log/persistent/anubis.log"
	f2 := "/var/log/persistent/link.log"

	compress.Compress(name, f1, f2)
	fmt.Printf("writing zip: %s\n", name)
}
