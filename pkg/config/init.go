package config

func init() {
	Server = (&server{}).Load("conf/app.ini").Init()
	Mysql = (&mysql{}).Load("conf/database.ini").Init()
}