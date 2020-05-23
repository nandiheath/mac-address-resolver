package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

func main() {

	fmt.Println("Listening to :8080")

	http.HandleFunc("/", handleConnection)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnection(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		fmt.Println("cannot parse ip address")
	}
	writer.Write([]byte(fmt.Sprintf("Your IP: %s\n", ip)))

	r, w := io.Pipe()

	c1 := exec.Command("arp", "-a")
	c2 := exec.Command("awk", "{print $2 $4}")

	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()

	output := b2.String()

	reg, _ := regexp.Compile(`\((.+)\)(.+)`)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		arrs := reg.FindStringSubmatch(line)
        if len(arrs) < 3 {
            continue
        }
        deviceIp := arrs[1]
        deviceMac := arrs[2]

        fmt.Println(deviceIp)
        if deviceIp == ip {
           writer.Write([]byte(fmt.Sprintf("Your Mac: %s\n", deviceMac)))
        }
	}

	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}
