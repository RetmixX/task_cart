package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"task_cart/app"
)

var rootCMD = &cobra.Command{
	Use:   "Start service",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		go func() {
			sgn := make(chan os.Signal, 1)
			signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)

			select {
			case <-ctx.Done():
			case <-sgn:
			}
			cancel()
		}()

		app.Run(ctx)

	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
