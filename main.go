
package main

import "flag"
import "net"
import "fmt"
import "bufio"
// import "os"
import "strconv"
import "time"

func main() {

  // command line flags/parameters and default values
  hostPtr := flag.String("h", "127.0.0.1", "host to connect to")
  portPtr := flag.Int("p", 25, "port to connect to")
  numPtr := flag.Int("n", 10, "number of times to connect")
  delayPtr := flag.Duration("d", 2 * time.Second, "delay between connections")

  // parse command line flags/parameters
  flag.Parse()

  service := *hostPtr+":"+strconv.Itoa(*portPtr)
  fmt.Println("Connecting to "+service)

  // connect to socket -n times
  for i := 1; i < 1 + *numPtr; i++ {
    // show the loop counter
    fmt.Printf("%d ", i)

    // connect to the socket
    conn, err := net.Dial("tcp", service)

    // handle error
    if err != nil {
      fmt.Println("Connection Error: "+err.Error())
    } else {

      // read from socket
      for {
        // track the time it takes
        start := time.Now()

        // try to read the buffer
        message, err := bufio.NewReader(conn).ReadString('\n')

        elapsed := time.Since(start)
        elapsedDisplay := elapsed-(elapsed%time.Millisecond)

        // handle error
        if err != nil {
          fmt.Printf("Read Error (%v): %s\n", elapsedDisplay, err.Error())
        }

        // print the response if we got one
        if message != "" {
          fmt.Printf("response (%v): %s", elapsedDisplay, message)
        }

        // close the connection and exit the read loop
        conn.Close()
        break
      }
    }
    // wait for duration -d before the next connection attempt
    time.Sleep(*delayPtr)
  }
}
