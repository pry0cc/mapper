package main

import (
    "strings"
    "os/exec"
    "runtime"
    "fmt"
    "bufio"
	"sync"
    "os"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(30)
    os.Mkdir("output", 0755)

    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        ipport := strings.ToLower(sc.Text())
	    wg.Add(1)
		go ScanLine(ipport, &wg)
    }

    wg.Wait()
}


// ScanLine takes input as ip:port and scans using nmap
func ScanLine(ipport string, wg *sync.WaitGroup) (err error) {

    defer wg.Done()

    s := strings.Split(ipport, ":")

    ip := s[0]
    port := s[1]
    filename := fmt.Sprintf("output/%s_%s.xml", ip, port)

    fmt.Println("nmap -T4 -Pn -sV  -oX ", filename, " -p", port, " ", ip)
    cmd := exec.Command("nmap", "-T4", "-sV",  "-Pn", "-oX",filename, "-p", port, ip)
    
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
    err = cmd.Run()
    
    if err != nil {
        return
    }

    return nil
}

