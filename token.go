package yunxin

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

const (
	createImUserPoint = neteaseBaseURL + "/user/create.action"
	refreshTokenPoint = neteaseBaseURL + "/user/refreshToken.action"
)

//CreateImUser 创建网易云通信ID
/**
 * @param accid 网易云通信ID，最大长度32字符，必须保证一个APP内唯一（只允许字母、数字、半角下划线_、@、半角点以及半角-组成，不区分大小写，会统一小写处理，请注意以此接口返回结果中的accid为准）。
 * @param name 网易云通信ID昵称，最大长度64字符，用来PUSH推送时显示的昵称
 * @param props json属性，第三方可选填，最大长度1024字符
 * @param icon 网易云通信ID头像URL，第三方可选填，最大长度1024
 * @param token 网易云通信ID可以指定登录token值，最大长度128字符，并更新，如果未指定，会自动生成token，并在创建成功后返回
 * @param sign 用户签名，最大长度256字符
 * @param email 用户email，最大长度64字符
 * @param birth 用户生日，最大长度16字符
 * @param mobile 用户mobile，最大长度32字符
 * @param gender 用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
 * @param ex 用户名片扩展字段，最大长度1024字符，用户可自行扩展，建议封装成JSON字符串
 */
func (c *ImClient) CreateImUser(u *ImUser) (res *Response[TokenInfo], err error) {
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
	if len(u.Propertys) > 0 {
		param["props"] = u.Propertys
	}
	if len(u.IconURL) > 0 {
		param["icon"] = u.IconURL
	}
	if len(u.Token) > 0 {
		param["token"] = u.Token
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

	if u.Gender < 0 || u.Gender > 2 {
		return nil, errors.New("invalid gender value")
	}
	param["gender"] = strconv.Itoa(u.Gender) // 允许传递0

	_, respBody, err := c.Curl("POST", createImUserPoint, param)
	if err != nil {
		return nil, err
	}

	aux := new(struct {
		*Response[TokenInfo]
		Info *TokenInfo `json:"info"`
	})

	err = json.Unmarshal(respBody, aux)
	if err != nil {
		return nil, err
	}

	res = aux.Response
	res.Data = aux.Info

	if res.Code != 200 {
		return res, fmt.Errorf("code:%d, desc:%s", res.Code, res.Desc)
	}

	return
}

// RefreshToken 更新并获取新token
// 参数：
//   accid - 网易云通信ID，最大长度32字符，必须保证一个APP内唯一
// 返回值：
//   *Response[TokenInfo] - 包含新令牌的响应对象
//   error - 错误信息（如果有）

func (c *ImClient) RefreshToken(accid string) (res *Response[TokenInfo], err error) {
	if accid == "" {
		return nil, errors.New("必须指定网易云通信ID")
	}

	param := map[string]string{"accid": accid}

	_, respBody, err := c.Curl("POST", refreshTokenPoint, param)

	if err != nil {
		return nil, err
	}

	aux := new(struct {
		*Response[TokenInfo]
		Info *TokenInfo `json:"info"`
	})

	err = json.Unmarshal(respBody, aux)
	if err != nil {
		return nil, err
	}

	res = aux.Response
	res.Data = aux.Info

	if res.Code != 200 {
		return res, fmt.Errorf("code:%d, desc:%s", res.Code, res.Desc)
	}

	return
}
