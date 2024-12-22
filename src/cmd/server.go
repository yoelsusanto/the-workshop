package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/yoelsusanto/the-workshop/src/app"
	"github.com/yoelsusanto/the-workshop/src/factory"
)

var port string

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	Long:  `Start a simple HTTP server that responds with Hello World`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	ServerCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
}

func startServer() {
	ctx := context.Background()

	rootRouter := mux.NewRouter()

	rootRouter.Path("/").Methods(http.MethodGet).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
		})

	app.Bind(rootRouter, &app.Handler{
		MainQueue: factory.NewMainQueue(ctx),
	})

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, rootRouter); err != nil {
		log.Fatal(err)
	}
}
