package main

import (
	"encoding/json"
	"os"
)

var (
	config = &Config{}
)

type Config struct {
	Time     int `json:"time"`
	Telegram struct {
		Enable bool   `json:"enable"`
		Url    string `json:"url"`
		Api    string `json:"api"`
		ChatId string `json:"chat_id"`
	} `json:"telegram"`
	Accounts []struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"accounts"`
}

type Cookies struct {
	HkCM2132ConnectIsBind  string `json:"hkCM_2132_connect_is_bind"`
	HkCM2132Nofavfid       string `json:"hkCM_2132_nofavfid"`
	HkCM2132Smile          string `json:"hkCM_2132_smile"`
	HkCM2132Saltkey        string `json:"hkCM_2132_saltkey"`
	HkCM2132Lastvisit      string `json:"hkCM_2132_lastvisit"`
	HkCM2132Auth           string `json:"hkCM_2132_auth"`
	HkCM2132Lastcheckfeed  string `json:"hkCM_2132_lastcheckfeed"`
	HkCM2132Visitedfid     string `json:"hkCM_2132_visitedfid"`
	HkCM2132ForumLastvisit string `json:"hkCM_2132_forum_lastvisit"`
	HkCM2132Ulastactivity  string `json:"hkCM_2132_ulastactivity"`
	HkCM2132HomeReadfeed   string `json:"hkCM_2132_home_readfeed"`
	HkCM2132Sid            string `json:"hkCM_2132_sid"`
	HkCM2132Lip            string `json:"hkCM_2132_lip"`
	HkCM2132Lastact        string `json:"hkCM_2132_lastact"`
}

func readConfig() {
	jsonFile, err := os.ReadFile(fileName)
	checkErr(err)
	err = json.Unmarshal(jsonFile, config)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
