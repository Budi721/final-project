package main

import (
	"github.com/itp-backend/backend-a-co-create/app"
	"github.com/itp-backend/backend-a-co-create/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
