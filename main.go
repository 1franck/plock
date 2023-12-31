package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var err error

	address := flag.String("addr", "", "address to listen on")
	directory := flag.String("dir", "./", "directory to serve")
	ssl := flag.Bool("ssl", false, "enable self-signed SSL (need openssl)")
	flag.Parse()

	if *address == "" {
		*address = "localhost:8080"
	}

	if *directory == "" || *directory == "./" {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		*directory = cwd
	}

	sslGenerated := false
	if *ssl && commandExists("openssl") {
		var cmd *exec.Cmd

		if !fileExists("server.key") {
			log.Println("generating private key...")
			cmd = exec.Command("openssl", "genrsa", "-out", "server.key", "2048")
			err = cmd.Run()
			if err != nil {
				panic(err)
			}
		}

		if !fileExists("server.crt") {
			log.Println("generating self-signed certificate...")
			cmd = exec.Command("openssl", "req", "-new", "-x509", "-sha256", "-key", "server.key", "-out", "server.crt", "-days", "3650", "-subj", "/C=US/ST=YourState/L=YourCity/O=YourOrganization/CN=YourCommonName")
			output, err := cmd.Output()
			if err != nil {
				log.Println(string(output))
				panic(err)
			}
		}

		sslGenerated = true
	} else if *ssl && !commandExists("openssl") {
		log.Println("openssl not found, cannot use ssl")
		*ssl = false
	}

	path, err := filepath.Abs(*directory)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(path); err != nil {
		panic(fmt.Sprintf("directory %s does not exist", path))
	}

	fs := http.FileServer(http.Dir(*directory))

	http.Handle("/", fs)

	addressProtocol := "http"
	if *ssl && sslGenerated {
		addressProtocol = "https"
	}
	log.Printf("Plock serving %s on %s://%s...\n", *directory, addressProtocol, *address)

	if *ssl && sslGenerated {
		err = http.ListenAndServeTLS(*address, "server.crt", "server.key", nil)
	} else {
		err = http.ListenAndServe(*address, nil)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func commandExists(command string) bool {
	cmd := exec.Command(command)
	err := cmd.Run()
	return err == nil
}
