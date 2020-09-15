package config

type Logger_ struct {
	Level string `yaml:level`
	Dir   string `yaml:dir`
}

type Service_ struct {
	Name     string `yaml:name`
	TTL      int64  `yaml:ttl`
	Interval int64  `yaml:interval`
	Address  string `yaml:address`
}

type Client_ struct {
	Retry   int32 `yaml:retry`
	Timeout int32 `yaml:timeout`
}

type MSA_ struct {
	Account string `yaml:account`
	Group   string `yaml:group`
}

type ConfigSchema_ struct {
	Service Service_ `yaml:service`
	Logger  Logger_  `yaml:logger`
	Client  Client_  `yaml:client`
	MSA     MSA_     `yaml:msa`
}
