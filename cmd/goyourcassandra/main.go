package main

import (
	"net/http"
	"strconv"

	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"

	"github.com/at15/goyourcassandra/pkg/server"
)

var log = dlog.NewApplicationLogger()

func main() {
	port := 8088
	cmd := cobra.Command{
		Use: "goyourcassandra",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("start http server")
			srv, err := server.New()
			if err != nil {
				log.Fatalf("error create server %s", err)
				return
			}
			addr := "localhost:" + strconv.Itoa(port)
			log.Infof("listen on %s", addr)
			http.ListenAndServe(addr, srv.HandlerWithLogger())
		},
	}
	cmd.Flags().IntVar(&port, "port", 8088, "port to listen to")
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
