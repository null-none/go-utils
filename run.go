package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net"
    "net/http"
    "fmt"
    "crypto/md5"
    "encoding/hex"
)


func getClientIPByRequest(req *http.Request) (ip string, err error) {
    ip, port, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        log.Printf("debug: Getting req.RemoteAddr %v", err)
        return "", err
    } else {
        log.Printf("debug: With req.RemoteAddr found IP:%v; Port: %v", ip, port)
    }

    userIP := net.ParseIP(ip)
    if userIP == nil {
        message := fmt.Sprintf("debug: Parsing IP from Request.RemoteAddr got nothing.")
        return "", fmt.Errorf(message)
    }
    log.Printf("debug: Found IP: %v", userIP)
    return userIP.String(), nil
}

func getMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func main() {

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        ip, err := getClientIPByRequest(c.Request)

        if err != nil {
          ip = "localhost"
        } 

        c.JSON(200, gin.H{
          "ip": ip,
          "md5": getMD5Hash("test"),
        }) 

    })
    r.Run(":8000")
}
