package main  

import (  
    "fmt"  
    "net"  
    "os"  
    "bufio"
)  

func sender(conn net.Conn) {
    r := bufio.NewReader(os.Stdin) 
    for{
        rawLine, _, _ := r.ReadLine()
        line := string(rawLine)
        conn.Write([]byte(line))  
        fmt.Println("send over") 
    }
}  

func main() {    
    tcpAddr, err := net.ResolveTCPAddr("tcp4", "192.168.112.98:4999")  
    if err != nil {  
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())  
        os.Exit(1)  
    }  

    conn, err := net.DialTCP("tcp", nil, tcpAddr)  
    if err != nil {  
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())  
        os.Exit(1)  
    }  

    fmt.Println("connect success")  
    sender(conn)  
}  