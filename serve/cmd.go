package serve

import (
	"github.com/spf13/cobra"
	"time"
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Cmd struct {
	ctx      context.Context
	CobraCmd cobra.Command
	config   Config
}

func NewCmd(ctx context.Context) *Cmd {
	var ret Cmd

	ret = Cmd{
		ctx: ctx,
		CobraCmd: cobra.Command{
			Use:   "serve",
			Short: "starts the http-db server",
			RunE: func(cmd *cobra.Command, args []string) error {

				logrus.Infoln("Start server")

				server := NewServer(&ret.config)

				serverDone := make(chan bool, 1)
				defer close(serverDone)

				go func() {
					select {
					case <-serverDone:
						return
					case <-ret.ctx.Done():
					}

					server.Shutdown()
				}()

				err := server.Run()

				if err != nil && err != http.ErrServerClosed {
					return err
				}

				return nil
			},
		},
	}

	ret.initFlags()

	return &ret
}

func (c *Cmd) initFlags() {
	c.CobraCmd.PersistentFlags().StringVarP(&c.config.address, "address", "a", "0.0.0.0", "")
	c.CobraCmd.PersistentFlags().IntVarP(&c.config.port, "port", "p", 8080, "")
	c.CobraCmd.PersistentFlags().DurationVarP(&c.config.timeout, "timeout", "t", 30*time.Second, "")
}
