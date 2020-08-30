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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "sort",
						Usage: "Set sorting of torrents. Can be: size[A/D], time[A/D], seeders[A/D], leechers[A/D]",
					},

					&cli.StringFlag{
						Name:  "category",
						Usage: "Filter by category. Can be: tv, movies, games, music, applications, anime, xxx, other, documentaries",
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
