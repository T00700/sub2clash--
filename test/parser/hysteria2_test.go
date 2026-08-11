package test

import (
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestHysteria2_Basic_SimpleLink(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria2://password123@127.0.0.1:8080#Hysteria2%20Proxy"

	expected := proxy.Proxy{
		Type: "hysteria2",
		Name: "Hysteria2 Proxy",
		Hysteria2: proxy.Hysteria2{
			Server:         "127.0.0.1",
			Port:           8080,
			Password:       "password123",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria2_Basic_AltPrefix(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hy2://password123@proxy.example.com:443?insecure=1&sni=proxy.example.com#Hysteria2%20Alt"

	expected := proxy.Proxy{
		Type: "hysteria2",
		Name: "Hysteria2 Alt",
		Hysteria2: proxy.Hysteria2{
			Server:         "proxy.example.com",
			Port:           443,
			Password:       "password123",
			SNI:            "proxy.example.com",
			SkipCertVerify: true,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria2_Basic_WithObfs(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria2://password123@127.0.0.1:8080?obfs=salamander&obfs-password=obfs123#Hysteria2%20Obfs"

	expected := proxy.Proxy{
		Type: "hysteria2",
		Name: "Hysteria2 Obfs",
		Hysteria2: proxy.Hysteria2{
			Server:         "127.0.0.1",
			Port:           8080,
			Password:       "password123",
			Obfs:           "salamander",
			ObfsPassword:   "obfs123",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria2_Basic_IPv6Address(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria2://password123@[2001:db8::1]:8080#Hysteria2%20IPv6"

	expected := proxy.Proxy{
		Type: "hysteria2",
		Name: "Hysteria2 IPv6",
		Hysteria2: proxy.Hysteria2{
			Server:         "2001:db8::1",
			Port:           8080,
			Password:       "password123",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria2_Basic_FullConfig(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria2://password123@proxy.example.com:443?insecure=1&sni=proxy.example.com&obfs=salamander&obfs-password=obfs123#Hysteria2%20Full"

	expected := proxy.Proxy{
		Type: "hysteria2",
		Name: "Hysteria2 Full",
		Hysteria2: proxy.Hysteria2{
			Server:         "proxy.example.com",
			Port:           443,
			Password:       "password123",
			SNI:            "proxy.example.com",
			Obfs:           "salamander",
			ObfsPassword:   "obfs123",
			SkipCertVerify: true,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria2_Basic_NoPassword(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria2://@127.0.0.1:8080#No%20Password"

	expected := proxy.Proxy{
		Type: "hysteria2",
		Name: "No Password",
		Hysteria2: proxy.Hysteria2{
			Server:         "127.0.0.1",
			Port:           8080,
			Password:       "",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria2_Error_MissingServer(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria2://password123@:8080"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestHysteria2_Error_MissingPort(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria2://password123@127.0.0.1"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestHysteria2_Error_InvalidPort(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria2://password123@127.0.0.1:99999"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestHysteria2_Error_InvalidProtocol(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria://example.com:8080"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestHysteria2_Basic_PortHopping(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hysteria2://password123@proxy.example.com:443?mport=443,30000-30100&hop-interval=30#Hysteria2%20PortHop"

	expected := proxy.Proxy{
		Type: "hysteria2",
		Name: "Hysteria2 PortHop",
		Hysteria2: proxy.Hysteria2{
			Server:         "proxy.example.com",
			Port:           443,
			Password:       "password123",
			Ports:          "443,30000-30100",
			HopInterval:    30,
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria2_Basic_PortHoppingAliases(t *testing.T) {
	p := &parser.Hysteria2Parser{}
	input := "hy2://password123@proxy.example.com:443?ports=20000-30000&hop_interval=abc#Alias"

	expected := proxy.Proxy{
		Type: "hysteria2",
		Name: "Alias",
		Hysteria2: proxy.Hysteria2{
			Server:         "proxy.example.com",
			Port:           443,
			Password:       "password123",
			Ports:          "20000-30000",
			HopInterval:    0, // 非法值宽松回退
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}
