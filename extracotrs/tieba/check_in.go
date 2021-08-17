package tieba

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yann0917/check-in/global"
	"github.com/yann0917/check-in/notification"
	"github.com/yann0917/check-in/utils"
)

const (
	appName = "贴吧"
	host    = "https://tieba.baidu.com"
	referer = "https://tieba.baidu.com/index.html"
)

var apis = map[string]string{
	"one_key_sign": "/tbmall/onekeySignin1",
	"sign_add":     "/sign/add",
	"tbs":          "http://tieba.baidu.com/dc/common/tbs",
}

func NewClient() *fiber.Agent {
	client := global.HttpClient

	if cookies, err := utils.ParseCookiesMap(global.Config.TieBa.Cookie); err == nil {
		for key, val := range cookies {
			client.Cookie(key, val)
		}
	}
	client.Referer(referer)
	client.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")

	return client
}

// GetTbs get tbs
func GetTbs() (resp TbsResp) {
	client := NewClient()
	client.ContentType(fiber.MIMETextHTML)
	client.Add(fiber.HeaderUpgradeInsecureRequests, "1")
	req := client.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(apis["tbs"])

	if err := client.Parse(); err != nil {
		panic(err)
	}

	_, body, errs := client.Bytes()
	// fmt.Println(string(body))
	_ = json.Unmarshal(body, &resp)
	fmt.Println(errs)
	fiber.ReleaseAgent(client)
	return
}

// GetForumList get 贴吧列表
func GetForumList() (tbs string, list []Forum) {

	client := NewClient()
	req := client.Request()
	client.ContentType(fiber.MIMETextHTML)
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(host)

	if err := client.Parse(); err != nil {
		panic(err)
	}

	_, body, errs := client.Bytes()
	fmt.Println(errs)

	// 获取 tbs
	re := regexp.MustCompile(`<script[\S\s]+?</script>`)
	scriptList := re.FindAllString(string(body), -1)

	for _, s := range scriptList {
		// 正则匹配 tbs
		re := regexp.MustCompile(`PageData.tbs = "([\S\s]+?)";`)
		tbsList := re.FindStringSubmatch(s)
		if len(tbsList) > 0 {
			// fmt.Println(tbsList)
			tbs = tbsList[1]
		}
		// 正则匹配 关注的贴吧列表
		forums := regexp.MustCompile(`{"forums":([\S\s]+?),"directory"`)
		forumList := forums.FindStringSubmatch(s)
		if len(forumList) > 0 {
			// fmt.Println(forumList[1])
			_ = json.Unmarshal([]byte(forumList[1]), &list)
		}
	}

	return
}

// SignAdd 签到
func SignAdd() {

	tbs, list := GetForumList()
	total, levelUpper7 := 0, 0

	for _, forum := range list {
		if forum.LevelId >= 7 {
			levelUpper7++
		}
	}

	// 如果等级全部 >= 7 级
	if len(list) == levelUpper7 {
		OneKeySignIn(tbs)
		return
	} else {
		for _, forum := range list {
			if forum.IsSign == 0 {
				client := NewClient()
				args := fiber.AcquireArgs()
				args.Set("ie", "utf-8")
				args.Set("kw", forum.ForumName)
				args.Set("tbs", tbs)
				client.Form(args)
				fiber.ReleaseArgs(args)

				req := client.Request()
				req.Header.SetMethod(fiber.MethodPost)
				req.SetRequestURI(host + apis["sign_add"])

				if err := client.Parse(); err != nil {
					panic(err)
				}

				var resp Response
				resp.Data = new(SignAddData)

				_, body, errs := client.Struct(&resp)
				fmt.Println(string(body))
				fmt.Println(errs)

				data, ok := resp.Data.(*SignAddData)
				fmt.Println(data)

				if ok && resp.No == 0 {
					total++
				}
			}
		}
		if total > 0 {
			content := "一共签到 " + strconv.Itoa(total) + " 个吧"
			notification.SendPushPlus("【"+appName+"】签到成功", content)
		}
	}
}

// OneKeySignIn 一键签到，仅限 7 级以上的贴吧
func OneKeySignIn(tbs string) {
	// login := GetTbs()
	// if login.IsLogin == 0 {
	// 	notification.SendPushPlus("【"+appName+"】签到失败", "未登录，请更新 cookie 后再执行")
	// 	return
	// }

	client := NewClient()

	args := fiber.AcquireArgs()
	args.Set("ie", "utf-8")
	args.Set("tbs", tbs)
	client.Form(args)
	fiber.ReleaseArgs(args)

	req := client.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.SetRequestURI(host + apis["one_key_sign"])

	if err := client.Parse(); err != nil {
		panic(err)
	}

	var resp Response
	resp.Data = new(OneKeySignInData)

	_, body, errs := client.Struct(&resp)
	fmt.Println(string(body))
	fmt.Println(errs)
	// fmt.Println(resp.Data)
	data, ok := resp.Data.(*OneKeySignInData)

	if ok && resp.No == 0 {
		content := "已经签到 " + strconv.Itoa(data.SignedForumAmount) +
			" 个吧，未签到 " + strconv.Itoa(data.UnsignedForumAmount) +
			" 个吧，本次签到共加经验：" + strconv.Itoa(data.GradeNoVip)
		notification.SendPushPlus("【"+appName+"】签到成功", content)
	} else {
		content := resp.Error
		switch resp.No {
		case 403:
			content += "，请更新 cookie 后再执行。"
		}
		notification.SendPushPlus("【"+appName+"】签到失败", content)
	}
}
