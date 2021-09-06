package config

import (
	"EasyGo-Server/pkg/helper"

	"gopkg.in/ini.v1"
)

var Mysql *mysql 

type mysql struct {
	Host string `ini:"host"`
	Port int `ini:"port"`
	Database string `ini:"database"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	source *ini.File 
}

func (s *mysql) Load(path string) *mysql {
	var err error 
	// check if config file exists 
	exists, err := helper.PathExists(path) 
	if !exists || err != nil {
		return s 
	}
	s.source, err = ini.Load(path) 
	if err != nil {
		panic(err)
	}
	return s 
}

func (s *mysql)Init() *mysql {
	// check if config loaded successfully 
	if s.source == nil {
		return s 
	}
	// map config to struct
	err := s.source.Section("mysql").MapTo(s)
	if err != nil {
		panic(err)
	}
	return s 
}