package yunxin

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	sendMsgPoint            = neteaseBaseURL + "/msg/sendMsg.action"
	sendBatchMsgPoint       = neteaseBaseURL + "/msg/sendBatchMsg.action"
	sendBatchAttachMsgPoint = neteaseBaseURL + "/msg/sendBatchAttachMsg.action"
	messageRecallPoint      = neteaseBaseURL + "/msg/recall.action"
)

const (
	//MsgTypeText 文本消息
	MsgTypeText = iota
	//MsgTypeImage 图片消息
	MsgTypeImage
	//MsgTypeVoice 语音消息
	MsgTypeVoice
	//MsgTypeVideo 视频消息
	MsgTypeVideo
	// MsgTypeLocation 地理位置消息
	MsgTypeLocation
	MsgTypeFile = 6
	// MsgTypeTips  提示消息
	MsgTypeTips = 10
)

// SendTextMessage 发送文本消息,消息内容最长5000
func (c *ImClient) SendTextMessage(fromID, toID string, msg *TextMessage, opt *ImSendMessageOption) (messageRes *Response[Message], err error) {
	bd, err := MarshalToString(msg)
	if err != nil {
		return
	}
	return c.SendMessage(fromID, toID, bd, 0, MsgTypeText, opt)
}

func (c *ImClient) SendLocationMessage(fromID, toID string, msg *LocationMessage, opt *ImSendMessageOption) (messageRes *Response[Message], err error) {
	bd, err := MarshalToString(msg)
	if err != nil {
		return
	}
	return c.SendMessage(fromID, toID, bd, 0, MsgTypeLocation, opt)
}

func (c *ImClient) SendFileMessage(fromID, toID string, msg *FileMessage, opt *ImSendMessageOption) (messageRes *Response[Message], err error) {
	bd, err := MarshalToString(msg)
	if err != nil {
		return
	}
	return c.SendMessage(fromID, toID, bd, 0, MsgTypeFile, opt)
}

func (c *ImClient) SendTipMessage(fromID, toID string, msg *TextMessage, opt *ImSendMessageOption) (messageRes *Response[Message], err error) {
	bd, err := MarshalToString(msg)
	if err != nil {
		return
	}
	return c.SendMessage(fromID, toID, bd, 0, MsgTypeTips, opt)
}

// SendBatchTextMessage 批量发送文本消息
func (c *ImClient) SendBatchTextMessage(fromID string, toIDs []string, msg *TextMessage, opt *ImSendMessageOption) (res *ResponseBase, err error) {
	bd, err := MarshalToString(msg)
	if err != nil {
		return nil, err
	}

	return c.SendBatchMessage(fromID, bd, toIDs, MsgTypeText, opt)
}

// SendBatchImageMessage 批量发送图片
func (c *ImClient) SendBatchImageMessage(fromID string, toIDs []string, msg *ImageMessage, opt *ImSendMessageOption) (res *ResponseBase, err error) {
	bd, err := MarshalToString(msg)
	if err != nil {
		return nil, err
	}

	return c.SendBatchMessage(fromID, bd, toIDs, MsgTypeImage, opt)
}

// SendBatchVoiceMessage .
func (c *ImClient) SendBatchVoiceMessage(fromID string, toIDs []string, msg *VoiceMessage, opt *ImSendMessageOption) (res *ResponseBase, err error) {
	bd, err := MarshalToString(msg)
	if err != nil {
		return nil, err
	}

	return c.SendBatchMessage(fromID, bd, toIDs, MsgTypeVoice, opt)
}

// SendBatchVideoMessage .
func (c *ImClient) SendBatchVideoMessage(fromID string, toIDs []string, msg *VideoMessage, opt *ImSendMessageOption) (res *ResponseBase, err error) {
	bd, err := MarshalToString(msg)
	if err != nil {
		return nil, err
	}

	return c.SendBatchMessage(fromID, bd, toIDs, MsgTypeVideo, opt)
}

//SendMessage 发送普通消息
/**
 * @param fromID 发送者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param toID ope==0是表示accid即用户id，ope==1表示tid即群id
 * @param ope 0：点对点个人消息，1：群消息（高级群），其他返回414
 * @param msgType 0 表示文本消息,1 表示图片，2 表示语音，3 表示视频，4 表示地理位置信息，6 表示文件，100 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
 * @param body 最大长度5000字符，为一个JSON串
 */
func (c *ImClient) SendMessage(fromID, toID, body string, ope, msgType int, opt *ImSendMessageOption) (messageRes *Response[Message], err error) {
	param := map[string]string{"from": fromID}

	param["ope"] = strconv.Itoa(ope)
	param["to"] = toID
	param["type"] = strconv.Itoa(msgType)
	param["body"] = body

	if opt != nil {
		param["antispam"] = strconv.FormatBool(opt.Antispam)

		if opt.AntispamCustom != nil {
			param["antispamCustom"], _ = MarshalToString(opt.AntispamCustom)
		}

		if opt.Option != nil {
			param["option"], _ = MarshalToString(opt.Option)
		}

		if len(opt.Pushcontent) > 0 {
			param["pushcontent"] = opt.Pushcontent
		}

		if len(opt.Payload) > 0 {
			param["payload"] = opt.Payload
		}

		if len(opt.Extension) > 0 {
			param["ext"] = opt.Extension
		}

		if opt.ForcePushList != nil {
			param["forcepushlist"], _ = MarshalToString(opt.ForcePushList)
		}

		if len(opt.ForcePushContent) > 0 {
			param["forcepushcontent"] = opt.ForcePushContent
		}
		param["forcepushall"] = strconv.FormatBool(opt.ForcePushAll)
		if len(opt.Bid) > 0 {
			param["bid"] = opt.Bid
		}
	}

	_, respBody, err := c.Curl("POST", sendMsgPoint, param)

	if err != nil {
		return nil, err
	}

	messageRes = new(Response[Message])
	err = json.Unmarshal(respBody, &messageRes)
	if err != nil {
		return nil, err
	}

	if messageRes.Code != 200 {
		err = fmt.Errorf("code:%v,desc:%s", messageRes.Code, messageRes.Desc)
		return nil, err
	}

	return messageRes, nil
}

//SendBatchMessage 批量发送点对点普通消息
/**
 * @param fromID 发送者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param toIDs ["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），限500人
 * @param msgType 0 表示文本消息,1 表示图片，2 表示语音，3 表示视频，4 表示地理位置信息，6 表示文件，100 自定义消息类型
 */
func (c *ImClient) SendBatchMessage(fromID, body string, toIDs []string, msgType int, opt *ImSendMessageOption) (res *ResponseBase, err error) {
	param := map[string]string{"fromAccid": fromID}

	to, err := MarshalToString(toIDs)
	if err != nil {
		return nil, err
	}
	param["toAccids"] = to
	param["type"] = strconv.Itoa(msgType)
	param["body"] = body

	if opt != nil {
		if opt.Option != nil {
			param["option"], _ = MarshalToString(opt.Option)
		}

		if len(opt.ForcePushContent) > 0 {
			param["forcepushcontent"] = opt.ForcePushContent
		}

		if len(opt.Payload) > 0 {
			param["payload"] = opt.Payload
		}

		if len(opt.Extension) > 0 {
			param["ext"] = opt.Extension
		}

		if len(opt.Bid) > 0 {
			param["bid"] = opt.Bid
		}
	}

	_, respBody, err := c.Curl("POST", sendBatchMsgPoint, param)

	if err != nil {
		return nil, err
	}

	res = new(ResponseBase)
	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return res, fmt.Errorf("code:%d, desc:%s", res.Code, res.Desc)
	}

	return res, nil
}

//SendBatchAttachMsg 批量发送点对点自定义系统通知
/**
 * @param fromID 发送者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param toIDs ["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），限500人
 * @param attach 自定义通知内容，第三方组装的字符串，建议是JSON串，最大长度4096字符
 */
func (c *ImClient) SendBatchAttachMsg(fromID, attach string, toIDs []string, opt *ImSendAttachMessageOption) (res *ResponseBase, err error) {
	param := map[string]string{"fromAccid": fromID}

	to, err := MarshalToString(toIDs)
	if err != nil {
		return nil, err
	}

	param["toAccids"] = to
	param["attach"] = attach
	if opt != nil {
		if len(opt.Pushcontent) > 0 {
			param["pushcontent"] = opt.Pushcontent
		}

		if len(opt.Payload) > 0 {
			param["payload"] = opt.Payload
		}

		if len(opt.Sound) > 0 {
			param["sound"] = opt.Payload
		}

		if opt.Save == 1 || opt.Save == 2 {
			param["save"] = strconv.Itoa(opt.Save)
		}

		if opt.Option != nil {
			param["option"], _ = MarshalToString(opt.Option)
		}
	}

	_, respBody, err := c.Curl("POST", sendBatchAttachMsgPoint, param)

	if err != nil {
		return nil, err
	}

	res = new(ResponseBase)
	err = json.Unmarshal(respBody, res)
	if err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return res, fmt.Errorf("code:%d, desc:%s", res.Code, res.Desc)
	}

	return res, nil
}

//RecallMessage 消息撤回
/**
 * @param deleteMsgid 要撤回消息的msgid
 * @param timetag 要撤回消息的创建时间
 * @param fromID 发消息的accid
 * @param toID 如果点对点消息，为接收消息的accid,如果群消息，为对应群的tid
 * @param msgtype 7:表示点对点消息撤回，8:表示群消息撤回，其它为参数错误
 */
func (c *ImClient) RecallMessage(deleteMsgid, timetag, fromID, toID string, msgtype int) (res *ResponseBase, err error) {
	param := map[string]string{"from": fromID, "to": toID, "type": strconv.Itoa(msgtype), "timetag": timetag, "deleteMsgid": deleteMsgid, "msg": "."}

	_, respBody, err := c.Curl("POST", messageRecallPoint, param)

	if err != nil {
		return nil, err
	}

	res = new(ResponseBase)
	err = json.Unmarshal(respBody, res)
	if err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return nil, fmt.Errorf("code:%d, desc:%s", res.Code, res.Desc)
	}

	return res, nil
}
