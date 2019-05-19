package config

import "fmt"

type ServerCfg struct {
	RpcPort   string
	HttpPort  string
	HttpsPort string
	TlsCert   string
	TlsKey    string
}

type DbCfg struct {
	Addr     string
	Port     string
	DbName   string
	User     string
	Password string
}

type Composition struct {
	Server *ServerCfg
	Db     *DbCfg
}

func (db *DbCfg) ConnString() string {
	uri := "mongodb://"
	if len(db.User) > 0 && len(db.Password) > 0 {
		uri = fmt.Sprintf("%s%s:%s@", uri, db.User, db.Password)
	}

	if addr := db.Addr; len(addr) > 0 {
		uri = fmt.Sprintf("%s%s", uri, addr)
	}

	if port := db.Port; len(port) > 0 {
		uri = fmt.Sprintf("%s:%s", uri, port)
	}

	return uri
}
