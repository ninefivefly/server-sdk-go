package sdk

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

// RCMsg RCMsg接口
type RCMsg interface {
	toString() (string, error)
}

// TXTMsg 消息
type TXTMsg struct {
	Content string `json:"content"`
	Extra   string `json:"extra"`
}

// InfoNtf 消息
type InfoNtf struct {
	Message string `json:"message"`
	Extra   string `json:"extra"`
}

// VCMsg 消息
type VCMsg struct {
	Content  string `json:"content"`
	Extra    string `json:"extra"`
	Duration string `json:"duration"`
}

// IMGTextMsg 消息
type IMGTextMsg struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Extra    string `json:"extra"`
	Duration string `json:"duration"`
	URL      string `json:"url"`
}

// FileMsg 消息
type FileMsg struct {
	Name    string `json:"name"`
	Size    string `json:"size"`
	Type    string `json:"type"`
	FileURL string `json:"fileUrl"`
}

// LBSMsg 消息
type LBSMsg struct {
	Content   string  `json:"content"`
	Extra     string  `json:"extra"`
	POI       string  `json:"poi"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ProfileNtf 消息
type ProfileNtf struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
	Extra     string `json:"extra"`
}

// CMDNtf 消息
type CMDNtf struct {
	Name string `json:"operation"`
	Data string `json:"data"`
}

// CMDMsg 消息
type CMDMsg struct {
	Name string `json:"operation"`
	Data string `json:"data"`
}

// ContactNtf 消息
type ContactNtf struct {
	Operation    string `json:"operation"`
	SourceUserID string `json:"sourceUserId"`
	TargetUserID string `json:"targetUserId"`
	Message      string `json:"message"`
	Extra        string `json:"extra"`
}

// GrpNtf 消息
type GrpNtf struct {
	OperatorUserID string `json:"operatorUserId"`
	Operation      string `json:"operation"`
	Data           string `json:"data"`
	Message        string `json:"message"`
	Extra          string `json:"extra"`
}

// DizNtf 消息
type DizNtf struct {
	Type      int    `json:"type"`
	Extension string `json:"extension"`
	Operation string `json:"operation"`
}

// TemplateMsgContent 消息模版
type TemplateMsgContent struct {
	TargetID    string
	Data        map[string]string
	PushContent string
}

// MsgContent 消息
type MsgContent struct {
	Content string `json:"content"`
	Extra   string `json:"extra"`
}

// MentionedInfo Mentioned
type MentionedInfo struct {
	Type        int      `json:"type"`
	UserIDs     []string `json:"userIdList"`
	PushContent string   `json:"mentionedContent"`
}

// MentionMsgContent MentionMsgContent
type MentionMsgContent struct {
	Content       string        `json:"content"`
	MentionedInfo MentionedInfo `json:"mentionedinfo"`
}

// History History
type History struct {
	URL string `json:"url"`
}

// ToString MsgContent
func (msg *MsgContent) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString TXTMsg
func (msg *TXTMsg) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString InfoNtf
func (msg *InfoNtf) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString VCMsg
func (msg *VCMsg) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString IMGTextMsg
func (msg *IMGTextMsg) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString FileMsg
func (msg *FileMsg) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString LBSMsg
func (msg *LBSMsg) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString ProfileNtf
func (msg *ProfileNtf) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString CMDNtf
func (msg *CMDNtf) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString CMDMsg
func (msg *CMDMsg) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString ContactNtf
func (msg *ContactNtf) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString GrpNtf
func (msg *GrpNtf) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString DizNtf
func (msg *DizNtf) toString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// PrivateSend 发送单聊消息方法（一个用户向另外一个用户发送消息，单条消息最大 128k。每分钟最多发送 6000 条信息，每次发送用户上限为 1000 人，如：一次发送 1000 人时，示为 1000 条消息。）
/*
 *@param  senderID:发送人用户 ID。
 *@param  targetID:接收用户 ID。
 *@param  objectName:发送的消息类型。
 *@param  msg:消息内容。
 *@param  pushContent:定义显示的 Push 内容，如果 objectName 为融云内置消息类型时，则发送后用户一定会收到 Push 信息。如果为自定义消息，则 pushContent 为自定义消息显示的 Push 内容，如果不传则用户不会收到 Push 通知。
 *@param  pushData:针对 iOS 平台为 Push 通知时附加到 payload 中，Android 客户端收到推送消息时对应字段名为 pushData。
 *@param  count:针对 iOS 平台，Push 时用来控制未读消息显示数，只有在 toUserId 为一个用户 Id 的时候有效。
 *@param  verifyBlacklist:是否过滤发送人黑名单列表，0 表示为不过滤、 1 表示为过滤，默认为 0 不过滤。
 *@param  isPersisted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行存储，0 表示为不存储、 1 表示为存储，默认为 1 存储消息。
 *@param  isCounted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行未读消息计数，0 表示为不计数、 1 表示为计数，默认为 1 计数，未读消息数增加 1。
 *@param  isIncludeSender:发送用户自已是否接收消息，0 表示为不接收，1 表示为接收，默认为 0 不接收。
 *
 *@return error
 */
func (rc *RongCloud) PrivateSend(senderID, targetID, objectName string, msg RCMsg,
	pushContent, pushData string, count, verifyBlacklist, isPersisted, isCounted, isIncludeSender int) error {
	if senderID == "" {
		return RCErrorNew(1002, "Paramer 'senderID' is required")
	}

	if targetID == "" {
		return RCErrorNew(1002, "Paramer 'targetID' is required")
	}

	req := httplib.Post(RONGCLOUDURI + "/message/private/publish." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("toUserId", targetID)
	req.Param("objectName", objectName)

	msgstr, err := msg.toString()
	if err != nil {
		return err
	}
	req.Param("content", msgstr)
	req.Param("pushData", pushData)
	req.Param("pushContent", pushContent)
	req.Param("count", strconv.Itoa(count))
	req.Param("verifyBlacklist", strconv.Itoa(verifyBlacklist))
	req.Param("isPersisted", strconv.Itoa(isPersisted))
	req.Param("isCounted", strconv.Itoa(isCounted))
	req.Param("isIncludeSender", strconv.Itoa(isIncludeSender))

	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// PrivateMSend 发送单聊消息方法（一个用户向另外多个用户发送消息，单条消息最大 128k。每分钟最多发送 6000 条信息，每次发送用户上限为 1000 人，如：一次发送 1000 人时，示为 1000 条消息。）
/*
 *@param  senderID:发送人用户 ID。
 *@param  targetsID:接收用户 ID。
 *@param  objectName:发送的消息类型。
 *@param  msg:消息内容。
 *@param  pushContent:定义显示的 Push 内容，如果 objectName 为融云内置消息类型时，则发送后用户一定会收到 Push 信息。如果为自定义消息，则 pushContent 为自定义消息显示的 Push 内容，如果不传则用户不会收到 Push 通知。
 *@param  pushData:针对 iOS 平台为 Push 通知时附加到 payload 中，Android 客户端收到推送消息时对应字段名为 pushData。
 *@param  count:针对 iOS 平台，Push 时用来控制未读消息显示数，只有在 toUserId 为一个用户 Id 的时候有效。
 *@param  verifyBlacklist:是否过滤发送人黑名单列表，0 表示为不过滤、 1 表示为过滤，默认为 0 不过滤。
 *@param  isPersisted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行存储，0 表示为不存储、 1 表示为存储，默认为 1 存储消息。
 *@param  isCounted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行未读消息计数，0 表示为不计数、 1 表示为计数，默认为 1 计数，未读消息数增加 1。
 *@param  isIncludeSender:发送用户自已是否接收消息，0 表示为不接收，1 表示为接收，默认为 0 不接收。
 *
 *@return error
 */
func (rc *RongCloud) PrivateMSend(senderID string, targetsID []string, objectName string, msg RCMsg,
	pushContent, pushData string, count, verifyBlacklist, isPersisted, isCounted, isIncludeSender int) error {
	if senderID == "" {
		return RCErrorNew(1002, "Paramer 'senderID' is required")
	}

	if len(targetsID) == 0 {
		return RCErrorNew(1002, "Paramer 'targetID' is required")
	}

	req := httplib.Post(RONGCLOUDURI + "/message/private/publish." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	for _, v := range targetsID {
		req.Param("toUserId", v)
	}
	req.Param("objectName", objectName)

	msgstr, err := msg.toString()
	if err != nil {
		return err
	}

	req.Param("content", msgstr)
	req.Param("pushData", pushData)
	req.Param("pushContent", pushContent)
	req.Param("count", strconv.Itoa(count))
	req.Param("verifyBlacklist", strconv.Itoa(verifyBlacklist))
	req.Param("isPersisted", strconv.Itoa(isPersisted))
	req.Param("isCounted", strconv.Itoa(isCounted))
	req.Param("isIncludeSender", strconv.Itoa(isIncludeSender))

	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// PrivateRecall 撤回单聊消息方法
/*
*
*@param  senderID:发送人用户 ID。
*@param  targetID:接收用户 ID，可以实现向多人发送消息，每次上限为 1000 人。
*@param  uID:消息的唯一标识，各端 SDK 发送消息成功后会返回 uID。
*@param  sentTime:消息的发送时间，各端 SDK 发送消息成功后会返回 sentTime。
*@param  conversationType:会话类型，二人会话是 1 、群组会话是 3 。
*
*@return error
 */
func (rc *RongCloud) PrivateRecall(senderID, targetID, uID string, sentTime, conversationType int) error {
	if senderID == "" {
		return RCErrorNew(1002, "Paramer 'senderID' is required")
	}

	if targetID == "" {
		return RCErrorNew(1002, "Paramer 'targetID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/message/recall." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("targetId", targetID)
	req.Param("messageUID", uID)
	req.Param("sentTime", strconv.Itoa(sentTime))
	req.Param("conversationType", strconv.Itoa(conversationType))

	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// PrivateSendTemplate 向多个用户发送不同内容消息
/*
 *@param  senderID:发送人用户 ID。
 *@param  objectName:发送的消息类型。
 *@param  template:消息模版。
 *@param  content:数据内容，包含消息内容和接收者。
 *
 *@return error
 */
func (rc *RongCloud) PrivateSendTemplate(senderID, objectName string, template MsgContent, content []TemplateMsgContent) error {
	if senderID == "" {
		return RCErrorNew(1002, "Paramer 'senderID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/message/private/publish_template." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)

	var toUserIDs, push []string
	var values []map[string]string

	for _, v := range content {
		if v.TargetID == "" {
			return RCErrorNew(1002, "Paramer 'TargetID' is required")
		}
		toUserIDs = append(toUserIDs, v.TargetID)
		values = append(values, v.Data)
		push = append(push, v.PushContent)
	}

	bytes, err := json.Marshal(template)
	if err != nil {
		return err
	}

	param := map[string]interface{}{}
	param["fromUserId"] = senderID
	param["objectName"] = objectName
	param["content"] = string(bytes)
	param["toUserId"] = toUserIDs
	param["values"] = values
	param["pushContent"] = push
	param["verifyBlacklist"] = 0
	req.JSONBody(param)

	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// GroupSend 发送群组消息方法（以一个用户身份向群组发送消息，单条消息最大 128k.每秒钟最多发送 20 条消息，每次最多向 3 个群组发送，如：一次向 3 个群组发送消息，示为 3 条消息。）
/*
 *@param  senderID:发送人用户 ID 。
 *@param  targetID:接收群ID.
 *@param  objectName:消息类型
 *@param  msg:发送消息内容
 *@param  pushContent:定义显示的 Push 内容，如果 objectName 为融云内置消息类型时，则发送后用户一定会收到 Push 信息. 如果为自定义消息，则 pushContent 为自定义消息显示的 Push 内容，如果不传则用户不会收到 Push 通知。
 *@param  pushData:针对 iOS 平台为 Push 通知时附加到 payload 中，Android 客户端收到推送消息时对应字段名为 pushData。
 *@param  isPersisted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行存储，0 表示为不存储、 1 表示为存储，默认为 1 存储消息。
 *@param  isCounted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行未读消息计数，0 表示为不计数、 1 表示为计数，默认为 1 计数，未读消息数增加 1。
 *@param  isIncludeSender:发送用户自已是否接收消息，0 表示为不接收，1 表示为接收，默认为 0 不接收。
 *
 *@return error
 */
func (rc *RongCloud) GroupSend(senderID, targetID, objectName string, msg RCMsg,
	pushContent string, pushData string, isPersisted int, isCounted int, isIncludeSender int) error {
	if senderID == "" {
		return RCErrorNew(20013, "Paramer 'senderID' is required")
	}

	if targetID == "" {
		return RCErrorNew(20013, "Paramer 'senderID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/message/group/publish." + ReqType)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("toGroupId", targetID)
	req.Param("objectName", objectName)
	msgstr, err := msg.ToString()
	if err != nil {
		return err
	}
	req.Param("content", msgstr)
	req.Param("pushContent", pushContent)
	req.Param("pushData", pushData)
	req.Param("isPersisted", strconv.Itoa(isPersisted))
	req.Param("isCounted", strconv.Itoa(isCounted))
	req.Param("isIncludeSender", strconv.Itoa(isIncludeSender))
	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// GroupRecall 撤回群聊消息
/*
*@param  senderID:发送人用户 ID。
*@param  targetID:接收用户 ID，可以实现向多人发送消息，每次上限为 1000 人。
*@param  uID:消息的唯一标识，各端 SDK 发送消息成功后会返回 uID。
*@param  sentTime:消息的发送时间，各端 SDK 发送消息成功后会返回 sentTime。
*@param  conversationType:会话类型，二人会话是 1 、群组会话是 3 。
*
*@return error
 */
func (rc *RongCloud) GroupRecall(senderID, targetID, uID string, sentTime, conversationType int) error {
	if senderID == "" {
		return RCErrorNew(1002, "Paramer 'senderID' is required")
	}

	if targetID == "" {
		return RCErrorNew(1002, "Paramer 'targetID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/message/recall." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("targetId", targetID)
	req.Param("messageUID", uID)
	req.Param("sentTime", strconv.Itoa(sentTime))
	req.Param("conversationType", strconv.Itoa(conversationType))

	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// GroupSendMention 发送群组 @ 消息
/*
*@param  senderID:发送人用户 ID 。
*@param  targetID:接收群ID。
*@param  objectName:消息类型。
*@param  msg:发送消息内容。
*@param  pushContent:定义显示的 Push 内容，如果 objectName 为融云内置消息类型时，则发送后用户一定会收到 Push 信息. 如果为自定义消息，则 pushContent 为自定义消息显示的 Push 内容，如果不传则用户不会收到 Push 通知。
*@param  pushData:针对 iOS 平台为 Push 通知时附加到 payload 中，Android 客户端收到推送消息时对应字段名为 pushData。
*@param  isPersisted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行存储，0 表示为不存储、 1 表示为存储，默认为 1 存储消息。
*@param  isCounted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行未读消息计数，0 表示为不计数、 1 表示为计数，默认为 1 计数，未读消息数增加 1。
*@param  isIncludeSender:发送用户自已是否接收消息，0 表示为不接收，1 表示为接收，默认为 0 不接收。
*@param  isMentioned:是否为 @消息，0 表示为普通消息，1 表示为 @消息，默认为 0。当为 1 时 content 参数中必须携带 mentionedInfo @消息的详细内容。为 0 时则不需要携带 mentionedInfo。当指定了 toUserId 时，则 @ 的用户必须为 toUserId 中的用户。
*@param  contentAvailable:针对 iOS 平台，对 SDK 处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。 允许应用在收到通知后在后台运行一段代码，且能够马上执行，查看详细。1 表示为开启，0 表示为关闭，默认为 0
*
*@return error
 */
func (rc *RongCloud) GroupSendMention(senderID, targetID, objectName string, msg MentionMsgContent,
	pushContent, pushData string, isPersisted int, isCounted int, isIncludeSender, isMentioned, contentAvailable int) error {
	if senderID == "" {
		return RCErrorNew(20013, "Paramer 'senderID' is required")
	}

	if targetID == "" {
		return RCErrorNew(20013, "Paramer 'senderID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/message/group/publish." + ReqType)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("toGroupId", targetID)
	req.Param("objectName", objectName)
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req.Param("content", string(bytes))
	req.Param("pushContent", pushContent)
	req.Param("pushData", pushData)
	req.Param("isPersisted", strconv.Itoa(isPersisted))
	req.Param("isCounted", strconv.Itoa(isCounted))
	req.Param("isIncludeSender", strconv.Itoa(isIncludeSender))
	req.Param("isMentioned", strconv.Itoa(isMentioned))
	req.Param("contentAvailable", strconv.Itoa(contentAvailable))
	rep, err := req.Bytes()
	if err != nil {
		return err
	}
	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// ChatRoomSend 发送聊天室消息方法。（以一个用户身份向群组发送消息，单条消息最大 128k.每秒钟最多发送 20 条消息，每次最多向 3 个群组发送，如：一次向 3 个群组发送消息，示为 3 条消息。）
/*
*@param  senderID:发送人用户 ID 。
*@param  targetID:接收聊天室ID.
*@param  objectName:消息类型
*@param  msg:发送消息内容
*
*@return error
 */
func (rc *RongCloud) ChatRoomSend(senderID, targetID, objectName string, msg RCMsg) error {
	if senderID == "" {
		return RCErrorNew(20013, "Paramer 'senderID' is required")
	}

	if targetID == "" {
		return RCErrorNew(20013, "Paramer 'senderID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/message/chatroom/publish." + ReqType)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("toChatroomId", targetID)
	req.Param("objectName", objectName)
	msgstr, err := msg.toString()
	if err != nil {
		return err
	}
	req.Param("content", msgstr)

	rep, err := req.Bytes()
	if err != nil {
		return err
	}
	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// ChatRoomBroadcast 向应用内所有聊天室广播消息方法，此功能需开通 专属服务（以一个用户身份向群组发送消息，单条消息最大 128k.每秒钟最多发送 20 条消息。）
/*
*@param  senderID:发送人用户 ID 。
*@param  objectName:消息类型
*@param  msg:发送消息内容
*
*@return error
 */
func (rc *RongCloud) ChatRoomBroadcast(senderID, objectName string, msg RCMsg) error {
	if senderID == "" {
		return RCErrorNew(1002, "Paramer 'senderID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/message/chatroom/broadcast." + ReqType)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("objectName", objectName)
	msgstr, err := msg.toString()
	if err != nil {
		return err
	}
	req.Param("content", msgstr)

	rep, err := req.Bytes()
	if err != nil {
		return err
	}
	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// SystemSend 一个用户向一个或多个用户发送系统消息，单条消息最大 128k，会话类型为 SYSTEM。
/*
*@param  senderID:发送人用户 ID。
*@param  targetID:接收用户 ID。
*@param  objectName:发送的消息类型。
*@param  msg:消息。
*@param  pushContent:定义显示的 Push 内容，如果 objectName 为融云内置消息类型时，则发送后用户一定会收到 Push 信息。如果为自定义消息，则 pushContent 为自定义消息显示的 Push 内容，如果不传则用户不会收到 Push 通知。
*@param  pushData:针对 iOS 平台为 Push 通知时附加到 payload 中，Android 客户端收到推送消息时对应字段名为 pushData。
*@param  count:针对 iOS 平台，Push 时用来控制未读消息显示数，只有在 toUserId 为一个用户 Id 的时候有效。
*@param  isPersisted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行存储，0 表示为不存储、 1 表示为存储，默认为 1 存储消息。
*@param  isCounted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行未读消息计数，0 表示为不计数、 1 表示为计数，默认为 1 计数，未读消息数增加 1。
*
*@return error
 */
func (rc *RongCloud) SystemSend(senderID, targetID, objectName string, msg RCMsg,
	pushContent, pushData string, count, isPersisted, isCounted int) error {

	if senderID == "" {
		return RCErrorNew(1002, "Paramer 'senderID' is required")
	}

	if targetID == "" {
		return RCErrorNew(1002, "Paramer 'targetID' is required")
	}

	req := httplib.Post(RONGCLOUDURI + "/message/system/publish." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("toUserId", targetID)
	req.Param("objectName", objectName)
	msgstr, err := msg.toString()
	if err != nil {
		return err
	}

	req.Param("content", msgstr)
	req.Param("pushData", pushData)
	req.Param("pushContent", pushContent)
	req.Param("count", strconv.Itoa(count))
	req.Param("isPersisted", strconv.Itoa(isPersisted))
	req.Param("isCounted", strconv.Itoa(isCounted))

	rep, err := req.Bytes()
	if err != nil {
		return err
	}
	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// SystemBroadcast 给应用内所有用户发送消息方法，每小时最多发 2 次，每天最多发送 3 次（以一个用户身份向群组发送消息，单条消息最大 128k.每秒钟最多发送 20 条消息。）
/*
*@param  senderID:发送人用户 ID 。
*@param  objectName:消息类型
*@param  msg:发送消息内容
*
*@return error
 */
func (rc *RongCloud) SystemBroadcast(senderID, objectName string, msg RCMsg) error {
	if senderID == "" {
		return RCErrorNew(20013, "Paramer 'senderID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/message/broadcast." + ReqType)
	rc.FillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("objectName", objectName)
	msgstr, err := msg.toString()
	if err != nil {
		return err
	}
	req.Param("content", msgstr)

	rep, err := req.Bytes()
	if err != nil {
		return err
	}
	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// SystemSendTemplate 一个用户向一个或多个用户发送系统消息，单条消息最大 128k，会话类型为 SYSTEM
/*
*@param  senderID:发送人用户 ID。
*@param  objectName:发送的消息类型。
*@param  template:消息模版。
*@param  content:数据内容，包含消息内容和接收者。
*
*@return error
 */
func (rc *RongCloud) SystemSendTemplate(senderID, objectName string, template MsgContent, content []TemplateMsgContent) error {
	if senderID == "" {
		return errors.New("1002 Paramer 'senderID' is required")
	}
	req := httplib.Post(rc.RongCloudURI + "/message/system/publish_template." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)

	var toUserIDs, push []string
	var values []map[string]string

	for _, v := range content {
		if v.TargetID == "" {
			return errors.New("1002 Paramer 'TargetID' is required")
		}
		toUserIDs = append(toUserIDs, v.TargetID)
		values = append(values, v.Data)
		push = append(push, v.PushContent)
	}

	bytes, err := json.Marshal(template)
	if err != nil {
		return err
	}

	param := map[string]interface{}{}
	param["fromUserId"] = senderID
	param["objectName"] = objectName
	param["content"] = string(bytes)
	param["toUserId"] = toUserIDs
	param["values"] = values
	param["pushContent"] = push
	param["verifyBlacklist"] = 0

	req.JSONBody(param)
	rep, err := req.Bytes()

	if err != nil {
		return err
	}
	var code CodeResult

	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// HistoryGet 按小时获取历史消息日志文件 URL，包含小时内应用产生的所有消息，消息日志文件无论是否已下载，3 天后将从融云服务器删除
/*
*@param date:精确到小时，例如: 2018030210 表示获取 2018 年 3 月 2 日 10 点至 11 点产生的数据
*
*@return History error
 */
func (rc *RongCloud) HistoryGet(date string) (History, error) {
	req := httplib.Post(rc.RongCloudURI + "/message/history." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("date", date)
	rep, err := req.Bytes()
	if err != nil {
		return History{}, err
	}
	var code CodeResult
	var history History
	if err := json.Unmarshal(rep, &code); err != nil {
		return History{}, err
	}
	if err := json.Unmarshal(rep, &history); err != nil {
		return History{}, err
	}

	if code.Code != 200 {
		return History{}, RCErrorNew(code.Code, code.ErrorMessage)
	}
	return history, nil
}

// HistoryRemove 删除历史消息日志文件
/*
*@param date:精确到小时，例如: 2018030210 表示获取 2018 年 3 月 2 日 10 点至 11 点产生的数据
*
*@return error
 */
func (rc *RongCloud) HistoryRemove(date string) error {
	if date == "" {
		return RCErrorNew(1002, "Paramer 'date' is required")
	}
	req := httplib.Post(rc.RongCloudURI + "/message/history/delete." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("date", date)
	rep, err := req.Bytes()
	if err != nil {
		return err
	}
	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil

}
