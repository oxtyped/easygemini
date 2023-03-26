package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"git.sr.ht/~adnano/go-gemini"
	"git.sr.ht/~adnano/go-gemini/certificate"
	"github.com/spf13/cobra"
)

var addr string
var port string
var site string

func init() {
	serveCmd.Flags().StringVarP(&addr, "addr", "b", "0.0.0.0", "ip to listen to")
	serveCmd.Flags().StringVarP(&port, "port", "p", "1965", "port to listen to")
	serveCmd.Flags().StringVarP(&site, "site", "s", "", "fqdn of your domain name (default to *)")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start ezgem server",
	Run: func(cmd *cobra.Command, args []string) {

		path := "sites/"

		server := &gemini.Server{}

		certificates := &certificate.Store{}
		if site != "" {
			certificates.Register(fmt.Sprintf("*.%s", site))
		} else {
			certificates.Register("*")
		}

		err := certificates.Load("certs")
		if err != nil {
			log.Fatalf("error loading certificate: %#v", err)
		}
		server.GetCertificate = certificates.Get

		mux := &gemini.Mux{}

		mux.HandleFunc("/", func(ctx context.Context, w gemini.ResponseWriter, r *gemini.Request) {

			pathAccess := fmt.Sprintf("%s%s", path, r.URL.Path)
			file, err := os.Stat(pathAccess)
			if err != nil {
				log.Fatalf("error accessing static path: %#v", err)
				w.WriteHeader(51, "File not Found")
				return

			}

			if file.IsDir() {
				pathAccess = pathAccess + "index.gmi"
			}

			reader, err := ioutil.ReadFile(pathAccess)
			if err != nil {
				w.WriteHeader(50, "")
				return
			}
			w.Write(reader)
		})
		server.Handler = mux

		ctx := context.Background()
		log.Printf("Starting server at %s:%s", addr, port)
		server.Addr = fmt.Sprintf("%s:%s", addr, port)
		err = server.ListenAndServe(ctx)
		if err != nil {
			log.Fatalf("error serving and listening: %#v", err)
		}

	},
}
