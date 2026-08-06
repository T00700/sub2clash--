package test

import (
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestTuic_Basic_SimpleLink(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tuic://uuid123:pass123@127.0.0.1:8080#Tuic%20Proxy"

	expected := proxy.Proxy{
		Type: "tuic",
		Name: "Tuic Proxy",
		Tuic: proxy.Tuic{
			Server:   "127.0.0.1",
			Port:     8080,
			UUID:     "uuid123",
			Password: "pass123",
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestTuic_Basic_FullConfig(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tuic://uuid123:pass123@proxy.example.com:443?congestion_control=bbr&alpn=h3,spdy/3.1&sni=proxy.example.com&allow_insecure=1&disable_sni=0&udp_relay_mode=native#Tuic%20Full"

	expected := proxy.Proxy{
		Type: "tuic",
		Name: "Tuic Full",
		Tuic: proxy.Tuic{
			Server:               "proxy.example.com",
			Port:                 443,
			UUID:                 "uuid123",
			Password:             "pass123",
			ALPN:                 []string{"h3", "spdy/3.1"},
			SNI:                  "proxy.example.com",
			UdpRelayMode:         "native",
			CongestionController: "bbr",
			SkipCertVerify:       true,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestTuic_Basic_ClashStyleParams(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tuic://uuid123:pass123@proxy.example.com:443?heartbeat-interval=10000&reduce-rtt=1&request-timeout=8000&udp-relay-mode=quic&congestion-controller=cubic&max-udp-relay-packet-size=1500&fast-open=1&max-open-streams=32&skip-cert-verify=1&token=TOKEN&ip=1.2.3.4#Tuic%20ClashStyle"

	expected := proxy.Proxy{
		Type: "tuic",
		Name: "Tuic ClashStyle",
		Tuic: proxy.Tuic{
			Server:                "proxy.example.com",
			Port:                  443,
			UUID:                  "uuid123",
			Password:              "pass123",
			Token:                 "TOKEN",
			Ip:                    "1.2.3.4",
			HeartbeatInterval:     10000,
			ReduceRtt:             true,
			RequestTimeout:        8000,
			UdpRelayMode:          "quic",
			CongestionController:  "cubic",
			MaxUdpRelayPacketSize: 1500,
			FastOpen:              true,
			MaxOpenStreams:        32,
			SkipCertVerify:        true,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestTuic_Basic_BareBoolParams(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tuic://uuid123:pass123@127.0.0.1:8080?allow_insecure&disable_sni=1#Tuic%20BareBool"

	expected := proxy.Proxy{
		Type: "tuic",
		Name: "Tuic BareBool",
		Tuic: proxy.Tuic{
			Server:         "127.0.0.1",
			Port:           8080,
			UUID:           "uuid123",
			Password:       "pass123",
			DisableSni:     true,
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

func TestTuic_Basic_IPv6Address(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tuic://uuid123:pass123@[2001:db8::1]:8080#Tuic%20IPv6"

	expected := proxy.Proxy{
		Type: "tuic",
		Name: "Tuic IPv6",
		Tuic: proxy.Tuic{
			Server:   "2001:db8::1",
			Port:     8080,
			UUID:     "uuid123",
			Password: "pass123",
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestTuic_Basic_NoNameFallback(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tuic://uuid123:pass123@127.0.0.1:8080"

	expected := proxy.Proxy{
		Type: "tuic",
		Name: "127.0.0.1:8080",
		Tuic: proxy.Tuic{
			Server:   "127.0.0.1",
			Port:     8080,
			UUID:     "uuid123",
			Password: "pass123",
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestTuic_Error_MissingServer(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tuic://uuid123:pass123@:8080"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestTuic_Error_MissingPort(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tuic://uuid123:pass123@127.0.0.1"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestTuic_Error_InvalidPort(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tuic://uuid123:pass123@127.0.0.1:99999"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestTuic_Error_InvalidProtocol(t *testing.T) {
	p := &parser.TuicParser{}
	input := "tcp://uuid123:pass123@127.0.0.1:8080"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
