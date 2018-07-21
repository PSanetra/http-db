package main

import (
	"github.com/psanetra/http-db/serve"
	"github.com/spf13/cobra"
	"context"
	"os"
	"os/signal"
	"syscall"
)

type Cmd struct {
	ctx context.Context
	cancelFunc context.CancelFunc
	rootCmd cobra.Command
}

func newCmd() *Cmd {

	ctx, cancelFunc := context.WithCancel(context.Background())

	var cmd *Cmd

	cmd = &Cmd{
		ctx: ctx,
		cancelFunc: cancelFunc,
		rootCmd: cobra.Command{
			Use:   "http-db",
			SilenceErrors: true,
			SilenceUsage: true,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Help()
			},
		},
	}

	cmd.rootCmd.AddCommand(&serve.NewCmd(ctx).CobraCmd)

	return cmd
}

func (c *Cmd) Execute() error {
	return c.rootCmd.Execute()
}

func (c *Cmd) CancelOnSignals() {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-signalChan
		c.Cancel()
	}()
}

func (c *Cmd) Cancel() {
	c.cancelFunc()
}
