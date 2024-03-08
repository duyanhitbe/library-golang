package config

import "os"

type Env struct {
	DriverName    string
	DataSource    string
	ServerAddress string
}

func NewEnv() *Env {
	driverName := os.Getenv("DRIVER_NAME")
	dataSource := os.Getenv("DATA_SOURCE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	return &Env{
		DriverName:    driverName,
		DataSource:    dataSource,
		ServerAddress: serverAddress,
	}
}
