package main

import (
	"flag"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"strconv"
	"time"
)

var (
	fileName string
	message  string
)

func init() {
	log()

	flag.StringVar(&fileName, "c", "config.json", "配置文件路径, 默认当前路径下的 config.json")
	flag.Parse()

	readConfig()
}

func main() {
	logger.Info("签到任务开始!")
	accounts := config.Accounts
	for k, v := range accounts {
		logger.Info(fmt.Sprintf("正在进行第 %v 个账号签到", k+1))

		account := &Account{
			UserName: v.Username,
			Password: v.Password,
		}

		LOC, err := New(account)
		if err != nil {
			checkErr(err)
			continue
		}

		rCookies := LOC.Login()
		if LOC.CheckLoginStatus(rCookies) == false {
			message += fmt.Sprintf("用户 [%v] 登陆失败\v", account.UserName)
			continue
		}

		logger.Info(fmt.Sprintf("用户 [%v] 准备执行空间访问任务", account.UserName))
		// 嗯, OC!
		oc := LOC.CheckCoin()
		if config.Time < 5 {
			config.Time = 5
			logger.Warn(fmt.Sprintf("由于设定的请求频率过快, 已将任务间隔设置为 %v s", config.Time))
		}
		uids := randomUID()
		for i := 0; i < 20; i++ {
			time.Sleep(time.Duration(config.Time) * time.Second)
			logger.Info(fmt.Sprintf("用户 [%v] 正在进行第 %v 次空间访问.", account.UserName, i+1))
			LOC.AccessSpace(uids[i])
		}
		cc := LOC.CheckCoin()
		message += fmt.Sprintf("用户 [%v] 金钱: %v -> %v\n", account.UserName, oc, cc)
	}

	if config.Telegram.Enable == true {

		if config.Telegram.Api == "" || config.Telegram.ChatId == "" {
			logger.Info("Telegram 配置不全, 取消推送")
		} else {
			if message == "" {
				message = "没有账号进行了签到"
			}
			message = "[LOC 签到小助手]\n\n" + message
			fmt.Println(message)
			intChatId, err := strconv.ParseInt(config.Telegram.ChatId, 10, 64)
			checkErr(err)
			bot, err := tgbotapi.NewBotAPI(config.Telegram.Api)
			text := tgbotapi.NewMessage(intChatId, message)
			_, err = bot.Send(text)
			checkErr(err)
		}
	}
}

func randomUID() []int {
	var nums []int
	for i := 0; i < 20; i++ {
		exist := false
		uid := rand.Intn(58425) + 1
		for _, n := range nums {
			if uid == n {
				exist = true
			}
		}
		if exist == false {
			nums = append(nums, uid)
		} else {
			i -= 1
		}
	}
	return nums
}
