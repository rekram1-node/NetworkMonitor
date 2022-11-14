package monitor

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var sites []string = []string{
	"health.aws.amazon.com",
	"status.azure.com",
	"status.cloud.mongodb.com",
}

func ConnectedToInternet(hostname string) bool {
	fmt.Println("Checking Internet Connection...")
	address := hostname + ":" + strconv.Itoa(443)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)

	if err != nil {
		fmt.Printf("Connected: %t\n", false)
		return false
	}
	defer conn.Close()
	fmt.Printf("Connected: %t\n", true)
	return true
}
