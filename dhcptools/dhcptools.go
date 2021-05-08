package dhcptools
import (
	"github.com/c-robinson/iplib"
)


func Hello() string {
	return "Hello World."
}

func GetFirstIp(ip iplib.Net) string {
	return ip.FirstAddress().String()
}

func GetLastIp(ip iplib.Net) string {
	return ip.LastAddress().String()
}

func GetAccessPool(ip iplib.Net) string {
	net3 := ip.Enumerate(0,0)
	return net3[len(net3)-13].String()
}