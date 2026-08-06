package proxy

// https://github.com/MetaCubeX/mihomo/blob/Meta/adapter/outbound/tuic.go
type Tuic struct {
	Server                string   `yaml:"server" proxy:"server"`
	Port                  int      `yaml:"port" proxy:"port"`
	Token                 string   `yaml:"token,omitempty" proxy:"token,omitempty"`
	UUID                  string   `yaml:"uuid,omitempty" proxy:"uuid,omitempty"`
	Password              string   `yaml:"password,omitempty" proxy:"password,omitempty"`
	Ip                    string   `yaml:"ip,omitempty" proxy:"ip,omitempty"`
	HeartbeatInterval     int      `yaml:"heartbeat-interval,omitempty" proxy:"heartbeat-interval,omitempty"`
	ALPN                  []string `yaml:"alpn,omitempty" proxy:"alpn,omitempty"`
	ReduceRtt             bool     `yaml:"reduce-rtt,omitempty" proxy:"reduce-rtt,omitempty"`
	RequestTimeout        int      `yaml:"request-timeout,omitempty" proxy:"request-timeout,omitempty"`
	UdpRelayMode          string   `yaml:"udp-relay-mode,omitempty" proxy:"udp-relay-mode,omitempty"`
	CongestionController  string   `yaml:"congestion-controller,omitempty" proxy:"congestion-controller,omitempty"`
	DisableSni            bool     `yaml:"disable-sni,omitempty" proxy:"disable-sni,omitempty"`
	MaxUdpRelayPacketSize int      `yaml:"max-udp-relay-packet-size,omitempty" proxy:"max-udp-relay-packet-size,omitempty"`

	FastOpen             bool       `yaml:"fast-open,omitempty" proxy:"fast-open,omitempty"`
	MaxOpenStreams       int        `yaml:"max-open-streams,omitempty" proxy:"max-open-streams,omitempty"`
	CWND                 int        `yaml:"cwnd,omitempty" proxy:"cwnd,omitempty"`
	SkipCertVerify       bool       `yaml:"skip-cert-verify,omitempty" proxy:"skip-cert-verify,omitempty"`
	Fingerprint          string     `yaml:"fingerprint,omitempty" proxy:"fingerprint,omitempty"`
	Certificate          string     `yaml:"certificate,omitempty" proxy:"certificate,omitempty"`
	PrivateKey           string     `yaml:"private-key,omitempty" proxy:"private-key,omitempty"`
	ReceiveWindowConn    int        `yaml:"recv-window-conn,omitempty" proxy:"recv-window-conn,omitempty"`
	ReceiveWindow        int        `yaml:"recv-window,omitempty" proxy:"recv-window,omitempty"`
	DisableMTUDiscovery  bool       `yaml:"disable-mtu-discovery,omitempty" proxy:"disable-mtu-discovery,omitempty"`
	MaxDatagramFrameSize int        `yaml:"max-datagram-frame-size,omitempty" proxy:"max-datagram-frame-size,omitempty"`
	SNI                  string     `yaml:"sni,omitempty" proxy:"sni,omitempty"`
	ECHOpts              ECHOptions `yaml:"ech-opts,omitempty" proxy:"ech-opts,omitempty"`

	UDPOverStream        bool `yaml:"udp-over-stream,omitempty" proxy:"udp-over-stream,omitempty"`
	UDPOverStreamVersion int  `yaml:"udp-over-stream-version,omitempty" proxy:"udp-over-stream-version,omitempty"`
}
