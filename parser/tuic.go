package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	P "github.com/bestnite/sub2clash/model/proxy"
)

type TuicParser struct{}

func (p *TuicParser) SupportClash() bool {
	return false
}

func (p *TuicParser) SupportMeta() bool {
	return true
}

func (p *TuicParser) GetPrefixes() []string {
	return []string{"tuic://"}
}

func (p *TuicParser) GetType() string {
	return "tuic"
}

// queryGetFirst 按顺序返回第一个非空参数值, 兼容下划线(dae 标准)与连字符(clash 风格)两种命名
func queryGetFirst(query url.Values, keys ...string) string {
	for _, key := range keys {
		if v := query.Get(key); v != "" {
			return v
		}
	}
	return ""
}

// parseBoolParam 解析布尔参数: 仅当参数存在时生效, 1/true/yes/on 或裸参数名(空值)视为 true
func parseBoolParam(query url.Values, keys ...string) bool {
	for _, key := range keys {
		values, ok := query[key]
		if !ok {
			continue
		}
		switch strings.ToLower(values[0]) {
		case "", "1", "true", "yes", "on":
			return true
		default:
			return false
		}
	}
	return false
}

// parseIntParam 解析整数参数, 缺失或非法时返回 0
func parseIntParam(value string) int {
	if value == "" {
		return 0
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return n
}

func (p *TuicParser) Parse(config ParseConfig, proxy string) (P.Proxy, error) {
	if !hasPrefix(proxy, p.GetPrefixes()) {
		return P.Proxy{}, fmt.Errorf("%w: %s", ErrInvalidPrefix, proxy)
	}

	link, err := url.Parse(proxy)
	if err != nil {
		return P.Proxy{}, fmt.Errorf("%w: %s", ErrInvalidStruct, err.Error())
	}

	server := link.Hostname()
	if server == "" {
		return P.Proxy{}, fmt.Errorf("%w: %s", ErrInvalidStruct, "missing server host")
	}
	portStr := link.Port()
	if portStr == "" {
		return P.Proxy{}, fmt.Errorf("%w: %s", ErrInvalidStruct, "missing server port")
	}
	port, err := ParsePort(portStr)
	if err != nil {
		return P.Proxy{}, fmt.Errorf("%w: %s", ErrInvalidPort, err.Error())
	}

	uuid := link.User.Username()
	password, _ := link.User.Password()

	query := link.Query()

	var alpn []string
	if alpnStr := queryGetFirst(query, "alpn"); alpnStr != "" {
		alpn = strings.Split(alpnStr, ",")
	}

	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	result := P.Tuic{
		Server:                server,
		Port:                  port,
		UUID:                  uuid,
		Password:              password,
		Token:                 queryGetFirst(query, "token"),
		Ip:                    queryGetFirst(query, "ip"),
		HeartbeatInterval:     parseIntParam(queryGetFirst(query, "heartbeat_interval", "heartbeat-interval")),
		ALPN:                  alpn,
		ReduceRtt:             parseBoolParam(query, "reduce_rtt", "reduce-rtt"),
		RequestTimeout:        parseIntParam(queryGetFirst(query, "request_timeout", "request-timeout")),
		UdpRelayMode:          queryGetFirst(query, "udp_relay_mode", "udp-relay-mode"),
		CongestionController:  queryGetFirst(query, "congestion_control", "congestion-controller"),
		DisableSni:            parseBoolParam(query, "disable_sni", "disable-sni"),
		MaxUdpRelayPacketSize: parseIntParam(queryGetFirst(query, "max_udp_relay_packet_size", "max-udp-relay-packet-size")),
		FastOpen:              parseBoolParam(query, "fast_open", "fast-open"),
		MaxOpenStreams:        parseIntParam(queryGetFirst(query, "max_open_streams", "max-open-streams")),
		SkipCertVerify:        parseBoolParam(query, "allow_insecure", "allow-insecure", "skip-cert-verify"),
		SNI:                   queryGetFirst(query, "sni"),
	}

	return P.Proxy{
		Type: p.GetType(),
		Name: remarks,
		Tuic: result,
	}, nil
}

func init() {
	RegisterParser(&TuicParser{})
}
