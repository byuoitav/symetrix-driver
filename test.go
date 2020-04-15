package main

import (
        "bufio"
        "fmt"
        "net"
	"os"
	"strconv"
	"strings"
	)

const (
	maxVolumeLevel = 56000
)

func main() {
        address := "10.66.76.8"
        block := "3\n"

        s, err := net.ResolveUDPAddr("udp4", address + ":48631")
        c, err := net.DialUDP("udp4", nil, s)
        if err != nil {
			fmt.Println("unable to establish TCP client")
			return
        }

        defer c.Close()
        fmt.Printf("result= %sType: %T\n", block, block)
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">> ")
        block, _ = reader.ReadString('\n')
        fmt.Print(block)
        fmt.Printf("result= %sType: %T\n", block, block)
        text := fmt.Sprintf("GS %v", block)
        fmt.Print(text)

        data := []byte(text)

        _, err = c.Write(data)

        if err != nil {
                fmt.Println(err)
                return
        }
        buffer := make([]byte, 1024)
        fmt.Print("test\n")
        n, _, err2 := c.ReadFromUDP(buffer)
        fmt.Print("test\n")
        if err2 != nil {
                fmt.Println(err)
                return
		}
	val :=string(buffer[0:n])
	result, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64)
	//fmt.Printf("result= %d\nType: %T\n", result, result)
	if result > maxVolumeLevel {
		fmt.Println("fader is above max volume level")
		result = 56000
	}
	volume := int(float64(result * 100.0) / maxVolumeLevel)
        fmt.Printf("Volume = %v\n", volume)
        return
}
