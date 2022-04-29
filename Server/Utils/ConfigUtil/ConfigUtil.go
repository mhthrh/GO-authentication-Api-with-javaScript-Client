package ConfigUtil

import (
	"fmt"
	"github.com/mhthrh/Voiceflex/Utils/CryptoUtil"
	"github.com/mhthrh/Voiceflex/Utils/DirectUtil"
	"github.com/mhthrh/Voiceflex/Utils/FileUtil"
	"github.com/mhthrh/Voiceflex/Utils/JsonUtil"
)

type Users struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}
type Treading struct {
	MinThread int `json:"MinThread"`
	MaxThread int `json:"MaxThread"`
}
type DataBase struct {
	Name   string `json:"Name"`
	Host   string `json:"Host"`
	Port   int    `json:"Port"`
	User   Users
	Dbname string `json:"Dbname"`
	Driver string `json:"Driver"`
}
type FTP struct {
	FtpName string `json:"FtpName"`
	IP      string `json:"IP"`
	Port    int    `json:"Port"`
	User    Users  `json:"User"`
}
type Authenticate struct {
	User Users  `json:"User"`
	Role string `json:"Role"`
}
type Serv struct {
	IP   string `json:"IP"`
	Port int    `json:"Port"`
}

type Config struct {
	AppName    string       `json:"AppName"`
	IsTest     bool         `json:"IsTest"`
	Version    string       `json:"Version"`
	ExpireDate string       `json:"ExpireDate"`
	DB         []DataBase   `json:"DB"`
	Ftp        []FTP        `json:"Ftp"`
	Thread     Treading     `json:"Thread"`
	Login      Authenticate `json:"SignUp"`
	Server     Serv         `json:"Server"`
}

func ReadConfig(file string) *Config {
	ut := DirectUtil.Ut{}
	d, _ := ut.GetPath()

	var jsonMap *Config

	JsonUtil.New(nil, nil).Json2Struct([]byte(func() string {
		k := CryptoUtil.NewKey()
		k.Text, _ = FileUtil.New(d, file).Read()
		k.Decrypt()
		return k.Result
	}()), &jsonMap)

	return jsonMap
}

func WriteConfig() {
	cfg := &Config{
		AppName:    "MyServices",
		IsTest:     true,
		Version:    "1.0.0",
		ExpireDate: "01-01-2023",
		DB: []DataBase{{
			Name: "PostgresSQL",
			Host: "127.0.0.1",
			Port: 5432,
			User: Users{
				UserName: "postgres",
				Password: "123456",
			},
			Dbname: "Voiceflex",
			Driver: "postgres",
		}, {
			Name: "Oracle",
			Host: "127.0.0.1",
			Port: 1501,
			User: Users{
				UserName: "admin",
				Password: "admin",
			},
			Dbname: "",
			Driver: "OraDB",
		}},
		Ftp: []FTP{{
			FtpName: "MyFtp",
			IP:      "127.0.0.1",
			Port:    21,
			User: Users{
				UserName: "FtpUser",
				Password: "FtpPAss",
			},
		}, {
			FtpName: "YourFtp",
			IP:      "127.0.0.1",
			Port:    21,
			User: Users{
				UserName: "admin",
				Password: "admin",
			},
		}},
		Thread: Treading{
			MinThread: 1,
			MaxThread: 25,
		},
		Login: Authenticate{
			User: Users{
				UserName: "admin",
				Password: "admin",
			},
			Role: "Admin",
		}, Server: Serv{
			IP:   "localhost",
			Port: 8585,
		},
	}
	s := JsonUtil.New(nil, nil).Struct2Json(cfg)
	k := CryptoUtil.NewKey()
	k.Text = s
	k.Encrypt()
	fmt.Println(fmt.Sprintf("%s\n", s), k.Result)

}
