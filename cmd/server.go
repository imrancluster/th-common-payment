package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	raven "github.com/getsentry/raven-go"
	"github.com/imrancluster/th-common-payment/router"

	"github.com/imrancluster/th-common-payment/config"
	"github.com/imrancluster/th-common-payment/conn"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the application server",
	Long:  "Start the application server",
	Run:   serve,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.Init()
		if err := conn.ConnectDB(); err != nil {
			return fmt.Errorf("cannot connect to the db: %v", err)
		}
		appCfg := config.App()
		if dsn := appCfg.Sentry; dsn != "" {
			if err := raven.SetDSN(dsn); err != nil {
				return err
			}
			raven.SetTagsContext(map[string]string{"service": "reader"})
			raven.SetEnvironment(appCfg.Env)
			raven.SetRelease(appCfg.Version)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	router := router.Router()
	serveApp := config.App()

	server := http.Server{
		Addr:    ":" + serveApp.Port,
		Handler: router,
	}

	go func() {
		log.Println("server started and listening on port:", serveApp.Port)
		server.ListenAndServe()
	}()
	fmt.Println("I'm here")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	log.Println("Signal ", <-ch, " received")

	ctx, errCtx := context.WithTimeout(context.Background(), time.Second*5)
	if errCtx != nil {
		log.Fatalf("failed to serve: %v", errCtx)
	}
	server.Shutdown(ctx)
}
