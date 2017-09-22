package latlong

import (
	"testing"
)

const validiP string = "8.8.8.8"
const invalidiP string = "qwerty"

func TestLatlongInvalidIp(t *testing.T) {

	ip := invalidiP
	_, err := Latlong(ip)

	if (err == nil) {
		t.Errorf("Expected error on wrong ip %s" , ip)
	}
}

func TestLatlongValidIp(t *testing.T) {

	ip := validiP
	_, err := Latlong(ip)

	if (err != nil) {
		t.Errorf("Expected response on valid ip %s\nError: %s" , ip, err)
	}
}

func TestLatlongInvalidUrl(t *testing.T) {

	ip := validiP
	invalidUrls := []string{"aaa", "bbb"}
	SetUrls(invalidUrls)
	_, err := Latlong(ip)

	if (err == nil) {
		t.Errorf("Expected error on invalid urls %v." , invalidUrls)
	}
}

func TestLatlongValidIp(t *testing.T) {

	ip := validiP
	_, err := Latlong(ip)

	if (err != nil) {
		t.Errorf("Expected response on valid ip %s\nError: %s" , ip, err)
	}
}