package models

type Config struct {
	Host       string `json:"Hostname"`
	Login      string `json:"Username"`
	Password   string `json:"Password"`
	SkipVerify bool   `json:"SkipVerify"`
	ConfigPath string `json:"-"`
}
