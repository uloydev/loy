package main

import (
	"github.com/uloydev/loy/loy"
	"github.com/uloydev/loy/view"
)

func main() {
	app := loy.New()

	app.Get("/", func(c *loy.Context) error {
		return c.Render(view.Index())
	})

	if err := app.Start(); err != nil {
		app.Logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
