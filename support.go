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

// ReadConfigFileContents reads and return the contents of a file text or binary
func ReadConfigFileContents(filename string) []byte {
	configFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	contents, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}
	configFile.Close()

	return contents
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

// Unmarshals the config contents from a YAML file into memory
func GetConfigContentsFromYaml(filename string) (Config, error) {
	var conf Config
	file, err := support.ReadConfigFileContents(filename)
	if err != nil {
		fmt.Printf("The file was not found. err = %s\n", err)
		return conf, err
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		fmt.Printf("There was an error decoding the yaml file. err = %s\n", err)
		return conf, err
	}

	return conf, nil
}
