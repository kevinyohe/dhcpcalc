package dhcptools

import (
	iplib "github.com/c-robinson/iplib"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello World."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestDFG(t *testing.T){
	want := "10.1.1.1"
	_, ip3, _ := iplib.ParseCIDR("10.1.1.0/24")
	got := GetFirstIp(ip3)
	if got != want {
		t.Errorf("GetFirstIp() = %q, want %q", got, want)
	} else {
		t.Log("TEST OK ", got, " = ", want)
	}
}

func TestDFG26(t *testing.T){
	want := "10.1.1.65"
	_, ip3, _ := iplib.ParseCIDR("10.1.1.64/26")
	got := GetFirstIp(ip3)
	if got != want {
		t.Errorf("GetFirstIp() = %q, want %q", got, want)
	} else {
		t.Log("TEST OK ", got, " = ", want)
	}
}


func TestDFG31(t *testing.T){
	want := "10.1.1.4"
	_, ip3, _ := iplib.ParseCIDR("10.1.1.4/31")
	got := GetFirstIp(ip3)
	if got != want {
		t.Errorf("GetFirstIp() = %q, want %q", got, want)
	} else {
		t.Log("TEST OK ", got, " = ", want)
	}
}



func TestLastIp31(t *testing.T){
	want := "10.1.1.5"
	_, ip3, _ := iplib.ParseCIDR("10.1.1.4/31")
	got := GetLastIp(ip3)
	if got != want {
		t.Errorf("GetLastIp() = %q, want %q", got, want)
	} else {
		t.Log("TEST OK ", got, " = ", want)
	}
}

func TestLastIp24(t *testing.T){
	want := "10.1.1.254"
	_, ip3, _ := iplib.ParseCIDR("10.1.1.0/24")
	got := GetLastIp(ip3)
	if got != want {
		t.Errorf("GetLastIp() = %q, want %q", got, want)
	} else {
		t.Log("TEST OK ", got, " = ", want)
	}
}

func TestLastIp26(t *testing.T){
	want := "10.1.1.62"
	_, ip3, _ := iplib.ParseCIDR("10.1.1.0/26")
	got := GetLastIp(ip3)
	if got != want {
		t.Errorf("GetLastIp() = %q, want %q", got, want)
	} else {
		t.Log("TEST OK ", got, " = ", want)
	}
}


func TestGetAccessPool(t *testing.T) {
	want := "10.1.1.50"
	_, ip3, _ := iplib.ParseCIDR("10.1.1.0/26")
	got := GetAccessPool(ip3)
	if got != want{
		t.Errorf("GetAccessPool() = %q, want %q", got, want)
	}
}