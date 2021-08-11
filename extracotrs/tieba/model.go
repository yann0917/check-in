package tieba

type Response struct {
	No    int         `json:"no"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

type OneKeySignInData struct {
	SignedForumAmount         int           `json:"signedForumAmount"`         // 已签到数量
	SignedForumAmountFail     int           `json:"signedForumAmountFail"`     // 签到失败数量
	UnsignedForumAmount       int           `json:"unsignedForumAmount"`       // 未签到数量
	VipExtraSignedForumAmount int           `json:"vipExtraSignedForumAmount"` // 成为超级会员可多签
	GradeNoVip                int           `json:"gradeNoVip"`                // 本次签到共加经验
	GradeVip                  int           `json:"gradeVip"`                  // 超级会员最多可加经验
	ForumList                 []OneKeyForum `json:"forum_list"`
}

type TbsResp struct {
	Tbs     string `json:"tbs"`
	IsLogin int    `json:"is_login"`
}

type OneKeyForum struct {
	ForumId      int    `json:"forum_id"`
	ForumName    string `json:"forum_name"`
	IsSignIn     int    `json:"is_sign_in"`
	LevelId      int    `json:"level_id"`
	ContSignNum  int    `json:"cont_sign_num"`
	LoyaltyScore struct {
		NormalScore int `json:"normal_score"`
		HighScore   int `json:"high_score"`
	} `json:"loyalty_score"`
}

type Forum struct {
	UserId    int    `json:"user_id"`
	ForumId   int    `json:"forum_id"`
	ForumName string `json:"forum_name"`
	IsLike    int    `json:"is_like"`
	IsBlack   int    `json:"is_black"`
	LikeNum   int    `json:"like_num"`
	IsTop     int    `json:"is_top"`
	InTime    int    `json:"in_time"`
	LevelId   int    `json:"level_id"`
	LevelName string `json:"level_name"`
	CurScore  int    `json:"cur_score"`
	ScoreLeft int    `json:"score_left"`
	IsSign    int    `json:"is_sign"` // 0-未签到，1-已签到
}

type SignAddData struct {
	Errno       int    `json:"errno"`
	Errmsg      string `json:"errmsg"`
	SignVersion int    `json:"sign_version"`
	IsBlock     int    `json:"is_block"`
	Finfo       struct {
		ForumInfo struct {
			ForumId   int    `json:"forum_id"`
			ForumName string `json:"forum_name"`
		} `json:"forum_info"`
		CurrentRankInfo struct {
			SignCount int `json:"sign_count"`
		} `json:"current_rank_info"`
	} `json:"finfo"`
	Uinfo struct {
		UserId           int `json:"user_id"`
		IsSignIn         int `json:"is_sign_in"`
		UserSignRank     int `json:"user_sign_rank"`
		SignTime         int `json:"sign_time"`
		ContSignNum      int `json:"cont_sign_num"`
		TotalSignNum     int `json:"total_sign_num"`
		CoutTotalSingNum int `json:"cout_total_sing_num"`
		HunSignNum       int `json:"hun_sign_num"`
		TotalResignNum   int `json:"total_resign_num"`
		IsOrgName        int `json:"is_org_name"`
	} `json:"uinfo"`
}
