package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var bindAddr string
	var bindPort int
	var interfaceName string

	app := &cli.App{
		Name:  "dnsrecap",
		Usage: "Recap domain name by DNS packets",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "addr",
				Aliases:     []string{"a"},
				Usage:       "Bind address of HTTP server",
				Value:       "127.0.0.1",
				Destination: &bindAddr,
			},
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "Bind port of HTTP server",
				Value:       5080,
				Destination: &bindPort,
			},
			&cli.StringFlag{
				Name:        "interface",
				Aliases:     []string{"i"},
				Usage:       "Interface name to capture packet",
				Destination: &interfaceName,
			},
		},

		Action: func(c *cli.Context) error {
			db := newDatabase()
			go runServer(bindAddr, bindPort, db)
			if err := capture(interfaceName, db); err != nil {
				return err
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	logger.Info("Shutting down...")
	if err != nil {
		logger.Errorw("Exit caused by error", "error", err)
	}
}
