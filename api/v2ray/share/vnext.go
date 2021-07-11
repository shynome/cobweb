package share

type VNEXT struct {
	Version string `json:"v"`          // version
	Remark  string `json:"ps"`         //remark
	Address string `json:"add"`        // server addr
	Port    string `json:"port"`       //
	ID      string `json:"id"`         //vmess id
	AlertID uint32 `json:"aid,string"` //alertid
	Network string `json:"net"`        // 'ws'|'kcp'|'tcp'|'h2'
	Type    string `json:"Type"`       // 伪装类型
	Host    string `json:"host"`       // 伪装的域名 中间逗号(,)隔开 (http,ws,h2)
	Path    string `json:"path"`       // ws/h2 stream path
	TLS     string `json:"tls"`        // tls
}
