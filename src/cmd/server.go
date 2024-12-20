package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
