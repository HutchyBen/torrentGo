package main

import (
	"github.com/hutchybean/torrentGo/sources"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.App{
		Commands: []*cli.Command{
			{
				Name:        "1337x",
				Description: "Search Torrents from 1337x.to",
				Usage:       "Use search term",
				Action:      sources.LEET,
			},
		},
	}

	app.Run(os.Args)
}
