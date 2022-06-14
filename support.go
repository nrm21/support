package support

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// SetupCloseHandler detects ctrl-c pressed by user to break loop and close program
func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("--- Ctrl + C pressed in Terminal ---")
		os.Exit(0)
	}()
}

// ReadConfigFileContents reads and return the contents of a file text or binary.  Don't
// throw log.fatal messages here, instead send them back to calling program to handle them
func ReadConfigFileContents(filename string) ([]byte, error) {
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// GetOutboundIP gets the preferred outbound ip of the current node at runtime
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP
}
