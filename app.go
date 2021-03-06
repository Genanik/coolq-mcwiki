package main

import (
	"net/url"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

//go:generate cqcfg -c .
// cqp: 名称: McWikiDirect
// cqp: 版本: 1.0.0:0
// cqp: 作者: BaiMeow
// cqp: 简介: 贼水的URL拼接
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "cn.miaoscraft.McWikiDirect" // TODO: 修改为这个插件的ID
	cqp.GroupMsg = onGroupMsg
	cqp.PrivateMsg = onPrivateMsg
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	if len(msg) < 6 || msg[:5] != "wiki " {
		return 0
	}
	keyword := msg[5:]
	cqp.SendPrivateMsg(fromQQ, getURL(keyword))
	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup int64, fromQQ int64, fromAnonymous string, msg string, font int32) int32 {
	if len(msg) < 6 || msg[:5] != "wiki " {
		return 0
	}
	keyword := msg[5:]
	cqp.SendGroupMsg(fromGroup, getURL(keyword))
	return 0
}

func getURL(keyword string) string {
	return "官方Wiki:\n" +
		"https://minecraft-zh.gamepedia.com/index.php?search=" + url.QueryEscape(keyword) + "&title=Special:%E6%90%9C%E7%B4%A2&go=%E5%89%8D%E5%BE%80" + "\n" +
		"镜像Wiki:\n" +
		"https://searchwiki.biligame.com/mc/index.php?search=" + url.QueryEscape(keyword)
}
