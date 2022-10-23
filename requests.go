package main

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const baseURL = "https://hostloc.com"

type Account struct {
	UserName string
	Password string
	Cookies  Cookies
	Client   resty.Client
}

type Space struct {
}

func New(account *Account) (*Account, error) {
	if account.UserName == "" {
		return account, errors.New("请输入用户名")
	}
	if account.Password == "" {
		return account, errors.New("请输入密码")
	}

	account.Client = *resty.New()

	account.Client.
		SetRetryCount(3).
		SetRetryWaitTime(10 * time.Second).
		SetBaseURL(baseURL).
		SetHeaders(map[string]string{
			"user-agent":   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
			"content-type": "application/x-www-form-urlencoded",
		})

	return account, nil
}

// Login 登录
func (account *Account) Login() []*http.Cookie {
	formhash := account.FormHash()
	if formhash == "" {
		logger.Error("请检查网络是否连通.")
		os.Exit(0)
	} else {
		logger.Info(fmt.Sprintf("获取到 [formhash]: %v", formhash))
	}

	cookies := account.GetCookies()
	client := account.Client

	params := map[string]string{
		"mod":         "logging",
		"action":      "login",
		"loginsubmit": "yes",
		"infloat":     "yes",
		"lssubmit":    "yes",
		"inajax":      "1",
	}

	data := map[string]string{
		"fastloginfield": "username",
		"username":       account.UserName,
		"cookietime":     "2592000",
		"password":       account.Password,
		"formhash":       formhash,
		"quickforward":   "yes",
		"handlekey":      "ls",
	}

	r, err := client.R().
		SetQueryParams(params).
		SetFormData(data).
		SetCookies(cookies).
		Post("member.php")
	checkErr(err)

	return r.Cookies()
}

// GetCookies 从主页获取 cookies
func (account *Account) GetCookies() []*http.Cookie {
	client := account.Client

	r, err := client.R().Get("")
	checkErr(err)

	logger.Info("初始化 [Cookies] 成功")
	return r.Cookies()
}

// FormHash 获取一个奇怪的哈希
func (account *Account) FormHash() string {
	formhash := ""
	errorTime := 0
	client := account.Client

	for {
		task := time.NewTimer(5 * time.Second)

		params := map[string]string{
			"mod":        "logging",
			"action":     "login",
			"infloat":    "yes",
			"handlekey":  "login",
			"inajax":     "1",
			"ajaxtarget": "fwin_content_login",
		}

		r, err := client.R().
			SetQueryParams(params).
			Get("member.php")
		checkErr(err)

		if strings.Contains(r.String(), "formhash") == true {
			rule := regexp.MustCompile(` name="formhash" value="([\da-zA-z]+)"`)
			result := rule.FindStringSubmatch(r.String())
			if len(result) <= 1 {
				logger.Error("正则匹配 [formhash] 失败! ")
				errorTime += 1
				if errorTime >= 3 {
					logger.Error("重试次数大于 3， 退出")
					return formhash
				}
				checkErr(err)
			}
			formhash = result[len(result)-1]
			return formhash
		} else {
			logger.Error("获取 [formhash] 失败! ")
			errorTime += 1
			if errorTime >= 3 {
				logger.Error("重试次数大于 3， 退出")
				return formhash
			}
			checkErr(err)
		}

		<-task.C
	}

}

// CheckLoginStatus 通过主页检查是否登录
func (account *Account) CheckLoginStatus(cookies []*http.Cookie) bool {
	client := account.Client

	r, err := client.R().
		SetCookies(cookies).
		Get("")
	checkErr(err)

	if strings.Contains(r.String(), account.UserName) == true {
		logger.Info(fmt.Sprintf("用户 [%v] 登录成功", account.UserName))
		account.Client.SetCookies(cookies)
		return true
	} else {
		logger.Info(fmt.Sprintf("用户 [%v] 登录失败", account.UserName))
		return false
	}
}

// CheckCoin 查看金币
func (account *Account) CheckCoin() string {
	client := account.Client

	params := map[string]string{
		"mod":        "spacecp",
		"ac":         "credit",
		"showcredit": "1",
		"inajax":     "1",
	}

	r, err := client.R().
		SetQueryParams(params).
		Get("home.php")
	checkErr(err)

	rule := regexp.MustCompile(`金钱: <span id="hcredit_2">([0-9]+)`)
	result := rule.FindStringSubmatch(r.String())
	coins := result[len(result)-1]

	logger.Info(fmt.Sprintf("用户 [%v] 当前金钱 [%v]", account.UserName, coins))
	return coins
}

// AccessSpace 访问他人空间
func (account *Account) AccessSpace(spaceID int) {
	client := account.Client

	r, err := client.R().
		Get(fmt.Sprintf("space-uid-%v.html", spaceID))
	checkErr(err)

	if strings.Contains(r.String(), account.UserName) == true {
		logger.Info(fmt.Sprintf("用户 [%v] 空间访问成功 [Space UID: %v]", account.UserName, spaceID))
	} else {
		logger.Error(fmt.Sprintf("用户 [%v] 空间访问失败 [Space UID: %v]", account.UserName, spaceID))
	}

}

func telegramNotice(message string) {
	tg := resty.New()

	_, err := tg.R().
		SetHeaders(map[string]string{
			"user-agent":   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
			"content-type": "application/x-www-form-urlencoded",
		}).SetQueryParams(map[string]string{
		"chat_id": config.Telegram.ChatId,
		"text":    message,
	}).Post(fmt.Sprintf("%v/bot%v/sendMessage", config.Telegram.Url, config.Telegram.Api))
	if err != nil {
		checkErr(err)
	} else {
		logger.Info("Telegram 推送成功")
	}
}
