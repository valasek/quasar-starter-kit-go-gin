// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valasek/quasar-starter-kit-go-gin/server/api"
	"github.com/valasek/quasar-starter-kit-go-gin/server/logger"
	"github.com/valasek/quasar-starter-kit-go-gin/server/routes"

	"github.com/gin-gonic/autotls"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the server on URL and port defined in config.yaml",
	Long: `Starts the server on URL and port defined in config.yaml.
	
Command first tests connection to DB, checks if DB contains at least one record in tables
projects, rates, consultants and holidays. If succeeds it will start server.`,
	Run: func(cmd *cobra.Command, args []string) {

		// prepare the server
		ssl := viper.GetBool("ssl")
		url := viper.GetString("url")
		port := viper.GetString("port")
		apiInst := api.NewAPI()

		r := routes.SetupRouter(apiInst)

		// run the server
		var address = url
		if len(port) > 0 {
			address = address + ":" + port
		}
		srv := &http.Server{
			Addr:    address,
			Handler: r,
		}

		// run the server with gracefull shutdown
		go func() {
			if ssl == true {
				logger.Log.Info("starting server on: https://", address)
				if err := autotls.Run(r, url); err != nil && err != http.ErrServerClosed {
					logger.Log.Error(fmt.Sprintf("listen: %s", err))
				}
			} else {
				logger.Log.Info("starting server on: http://", address)
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Log.Error(fmt.Sprintf("listen: %s", err))
				}
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		// kill (no param) default send syscanll.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Log.Info("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Log.Error("server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 1 seconds.
		select {
		case <-ctx.Done():
			logger.Log.Info("timeout of 1 second")
		}
		logger.Log.Info("Server exiting")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
