package main

import (
	"dhcpmaker/dhcptools"
	"flag"
	"fmt"
	"github.com/c-robinson/iplib"
	"os"
	"strings"
)

func main (){
	ip := flag.String("ip", "10.4.20.0/26", "IP to subnet")
	name := flag.String("name", "DEVICE", "Name of device")
	flag.Parse()
	fmt.Println(*ip)
	fmt.Println("test")
	dhcptools.Hello()
	_, ipnet, err := iplib.ParseCIDR(*ip)
	if err != nil {
		fmt.Println("Invalid ip ")
	}
	cide, _ := ipnet.Mask.Size()
	fmt.Println("mask: " , cide)
	fmt.Println("name: " , name)
	fmt.Println(dhcptools.GetFirstIp(ipnet))
	data := strings.Split(*ip, "/")
	address, mask := data[0], data[1]
	switch mask {
	case "32":
		fmt.Println(address, ": Not a supported network!")
		os.Exit(1)
	case "26":
		fmt.Printf("Name: %11v_%v_%v\n", *name, address,mask)
		fmt.Printf("Default Gateway: %11v\n",dhcptools.GetFirstIp(ipnet))
		fmt.Printf("Start IP: %11v\n", dhcptools.GetAccessPool(ipnet))
		fmt.Printf("End IP: %11v\n", dhcptools.GetLastIp(ipnet))
	case "31":
		fmt.Printf("Name: %v_%v_%v\n", *name, address,mask)
		fmt.Printf("Default Gateway: %v\n",dhcptools.GetFirstIp(ipnet))
		fmt.Printf("Start IP: %v\n", dhcptools.GetLastIp(ipnet))
		fmt.Printf("End IP: %v\n", dhcptools.GetLastIp(ipnet))
	default:
		fmt.Printf("Name: %v_%v_%v\n", *name, address,mask)
		fmt.Printf("Default Gateway: %v\n",dhcptools.GetFirstIp(ipnet))
		fmt.Printf("Start IP: %-15v\n", dhcptools.GetAccessPool(ipnet))
		fmt.Printf("End IP: %v\n", dhcptools.GetLastIp(ipnet))
	}
}