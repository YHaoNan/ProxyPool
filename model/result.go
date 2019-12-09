package model


type ProxyType string

const (
	HTTP ProxyType = "HTTP"
	HTTPS = "HTTPS"
	BOTH = "BOTH"
)

type Proxy struct {
	Address string `json:"address"`
	Port int `json:"port"`
	Pos string `json:"position"`
	ISP string `json:"isp"`
	Type ProxyType `json:"type"`
}

