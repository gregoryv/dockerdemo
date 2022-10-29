package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gregoryv/cmdline"
)

func main() {
	var (
		cli      = cmdline.NewBasicParser()
		server   = cli.Option("-s, --server, $SERVER").Url("") // required
		duration = cli.Option("-d, --duration, $DURATION").Duration("5s")
		bind     = cli.Option("-b, --bind, $BIND").String(":8088")
	)
	cli.Parse()

	if server.String() == "" {
		fmt.Println("must specify --server, try --help")
		os.Exit(1)
	}

	for _, e := range os.Environ() {
		log.Print(e)
	}

	http.HandleFunc("/", noop)
	log.Print("listen ", bind)
	go func() {
		log.Fatal(http.ListenAndServe(bind, nil))
	}()

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	log.Print("start")
loop:
	for {
		<-time.After(time.Second)
		select {
		case <-ctx.Done():
			break loop
		default:
		}
		// todo
		resp, err := http.Get(server.String())
		if err != nil {
			log.Print(err)
		} else {
			var buf bytes.Buffer
			io.Copy(&buf, resp.Body)
			resp.Body.Close()
			log.Println(resp.Status, buf.String())
		}
	}
	defer log.Print("stop")
}

func noop(w http.ResponseWriter, r *http.Request) {
	data := os.Getenv("HOSTNAME")
	fmt.Fprint(w, "from ", data)
}
