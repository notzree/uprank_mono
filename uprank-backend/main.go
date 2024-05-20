package main 

import "flag"

func main () {
	listen_addr := flag.String("listen_addr", ":8080", "server listen address")
	flag.Parse()

}