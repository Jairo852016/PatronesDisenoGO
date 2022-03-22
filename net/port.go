package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {

		// 1,2,3....99
		// sitio:1, sitio:2, sitio>99
		// 1-> Open, 2->Closed, ..
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			//conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", port))
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open \n", port)
		}(i)
	}
	wg.Wait()
}
