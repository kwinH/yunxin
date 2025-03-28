package yunxin

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

const (
	updateUinfoPoint = neteaseBaseURL + "/user/updateUinfo.action"
	getUinfoPoint    = neteaseBaseURL + "/user/getUinfos.action"
)

//UpdateUinfo 更新用户名片
/**
 * @param accid 网易云通信ID，最大长度32字符，必须保证一个APP内唯一（只允许字母、数字、半角下划线_、@、半角点以及半角-组成，不区分大小写，会统一小写处理，请注意以此接口返回结果中的accid为准）。
 * @param name 网易云通信ID昵称，最大长度64字符，用来PUSH推送时显示的昵称
 * @param icon 网易云通信ID头像URL，第三方可选填，最大长度1024
 * @param sign 用户签名，最大长度256字符
 * @param email 用户email，最大长度64字符
 * @param birth 用户生日，最大长度16字符
 * @param mobile 用户mobile，最大长度32字符
 * @param gender 用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
 * @param ex 用户名片扩展字段，最大长度1024字符，用户可自行扩展，建议封装成JSON字符串
 */
func (c *ImClient) UpdateUinfo(u *ImUser) (res *ResponseBase, err error) {
	if u.ID == "" {
		return nil, errors.New("accid cannot be empty")
	}
	if len(u.ID) > 32 {
		return nil, errors.New("accid exceeds 32 characters")
	}

	param := map[string]string{"accid": u.ID}

	if len(u.Name) > 0 {
		param["name"] = u.Name
	}
	if len(u.IconURL) > 0 {
		param["icon"] = u.IconURL
	}
	if len(u.Sign) > 0 {
		param["sign"] = u.Sign
	}
	if len(u.Email) > 0 {
		param["email"] = u.Email
	}
	if len(u.Birthday) > 0 {
		param["birth"] = u.Birthday
	}
	if len(u.Mobile) > 0 {
		param["mobile"] = u.Mobile
	}
	if len(u.Extension) > 0 {
		param["ex"] = u.Extension
	}
	if u.Gender == 1 || u.Gender == 2 {
		param["gender"] = strconv.Itoa(u.Gender)
	}

	_, respBody, err := c.Curl("POST", updateUinfoPoint, param)

	if err != nil {
		return
	}

	res = new(ResponseBase)
	err = json.Unmarshal(respBody, res)
	if err != nil {
		return
	}

	if res.Code != 200 {
		return nil, fmt.Errorf("code:%d, desc:%s", res.Code, res.Desc)
	}

	return
}

func (c *ImClient) GetUinfo(accids []string, muteStatus bool) (res *Response[[]Uinfo], err error) {
	accidsStr, err := MarshalToString(accids)
	if err != nil {
		return
	}
	param := map[string]string{"accids": accidsStr, "muteStatus": strconv.FormatBool(muteStatus)}
	_, respBody, err := c.Curl("POST", getUinfoPoint, param)

	if err != nil {
		return
	}

	aux := new(struct {
		*Response[[]Uinfo]
		Uinfos *[]Uinfo `json:"uinfos"`
	})

	err = json.Unmarshal(respBody, aux)
	if err != nil {
		return nil, err
	}

	res = aux.Response
	res.Data = aux.Uinfos

	if res.Code != 200 {
		return res, fmt.Errorf("code:%d, desc:%s", res.Code, res.Desc)
	}

	return
}
