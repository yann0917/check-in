package tieba

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"

	"github.com/imroc/req"
	"github.com/yann0917/check-in/global"
	"github.com/yann0917/check-in/notification"
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

func NewClient() *global.Client {
	return global.NewClient(global.Config.TieBa.Cookie, referer)
}

// GetTbs get tbs
func GetTbs() (resp TbsResp) {
	client := NewClient()
	response, err := client.Get(apis["tbs"], client.Headers)
	if err != nil {
		notification.SendPushPlus("【"+appName+"】签到失败", err.Error())
	}
	err = response.ToJSON(&resp)
	return
}

// GetForumList get 贴吧列表
func GetForumList() (tbs string, list []Forum) {

	client := NewClient()
	resp, err := client.Get(host, client.Headers)
	if err != nil {
		notification.SendPushPlus("【"+appName+"】签到失败", err.Error())
		return
	}

	body := resp.Bytes()
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

				param := req.Param{
					"ie":  "utf-8",
					"kw":  forum.ForumName,
					"tbs": tbs,
				}

				resp, err := client.Post(host+apis["sign_add"], param)
				if err != nil {
					notification.SendPushPlus("【"+appName+"】签到失败", err.Error())
					return
				}

				var result Response
				result.Data = new(SignAddData)
				err = resp.ToJSON(&result)
				if err != nil {
					log.Println(err)
					return
				}

				log.Println(result)

				_, ok := result.Data.(*SignAddData)

				if ok && result.No == 0 {
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
	param := req.Param{
		"ie":  "utf-8",
		"tbs": tbs,
	}

	resp, err := client.Post(host+apis["one_key_sign"], param)
	if err != nil {
		notification.SendPushPlus("【"+appName+"】签到失败", err.Error())
		return
	}

	var result Response
	err = resp.ToJSON(&result)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)

	data, ok := result.Data.(*OneKeySignInData)

	if ok && result.No == 0 {
		content := "已经签到 " + strconv.Itoa(data.SignedForumAmount) +
			" 个吧，未签到 " + strconv.Itoa(data.UnsignedForumAmount) +
			" 个吧，本次签到共加经验：" + strconv.Itoa(data.GradeNoVip)
		notification.SendPushPlus("【"+appName+"】签到成功", content)
	} else {
		content := result.Error
		switch result.No {
		case 403:
			content += "，请更新 cookie 后再执行。"
		}
		notification.SendPushPlus("【"+appName+"】签到失败", content)
	}
}
