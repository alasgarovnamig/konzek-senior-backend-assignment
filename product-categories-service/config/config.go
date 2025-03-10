package config

type Config struct {
	Server      Server
	Database    Database
	KeyCloak    KeyCloak
	DEBUG       bool
	RunK8s      bool
	ServiceName string
}

type Database struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
}

type Server struct {
	HTTPAddr string
	RPCAddr  string
	WinSrv   bool
	Debug    bool
}

type KeyCloak struct {
	Realm                     string
	BaseUrl                   string
	KeyCloakClient            KeyCloakClient
	RealmKonzekRS256PublicKey string
}

type KeyCloakClient struct {
	ClientId     string
	ClientSecret string
}
