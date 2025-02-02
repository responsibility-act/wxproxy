package mini

import (
	"fmt"
	. "github.com/jacktea/wxproxy/common"
	"github.com/jacktea/wxproxy/service"
	"github.com/jacktea/wxproxy/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// 为授权的小程序帐号上传小程序代码
func (this *MiniAction) Commit(c iris.Context) {
	this.postTransparentJson(c, UPLOAD_MINI_COMMITCODE)
}

// 获取体验小程序的体验二维码
func (this *MiniAction) GetQrCode(c iris.Context) {
	this.getTransparentJson(c, GET_MINI_QRCODE)
}

// 体验小程序二维码
func (this *MiniAction) GetQrCodeEx(c iris.Context) {
	componentAppid := c.Params().Get("componentAppid")
	appid := c.Params().Get("appid")
	path := c.FormValue("path")
	force, _ := strconv.ParseBool(c.FormValueDefault("force", "true"))
	resp := this.MiniSvr.GetQrCode(componentAppid, appid, path, force)
	if resp.IsSuccess() {
		url := fmt.Sprintf("%s://%s%s", utils.Scheme(c.Request()), c.Host(), c.Path()) + "/" + resp.Url
		resp.Url = strings.Replace(url, "getqrcodeex", "prevqrcode", 1)
	}
	c.JSON(resp)
}

// 小程序二维码
//func (this *MiniAction) GetWxQrCode(c iris.Context)  {
//	componentAppid := c.Params().Get("componentAppid")
//	appid := c.Params().Get("appid")
//	page := c.FormValue("page")
//	scene := c.FormValue("scene")
//	force,_ := strconv.ParseBool(c.FormValueDefault("force","false"))
//	resp := this.MiniSvr.GetWxQrCode(componentAppid,appid,page,scene,force)
//	if resp.IsSuccess() {
//		url := fmt.Sprintf("%s://%s%s",utils.Scheme(c.Request()),c.Host(),c.Path()) + "/" + resp.Url
//		resp.Url = strings.Replace(url,"getwxqrcode","prevqrcode",1)
//	}
//	c.JSON(resp)
//}

// 小程序二维码
func (this *MiniAction) GetWXACode(c iris.Context) {
	componentAppid := c.Params().Get("componentAppid")
	appid := c.Params().Get("appid")
	path := c.FormValue("path")
	force, _ := strconv.ParseBool(c.FormValueDefault("force", "false"))
	resp := this.MiniSvr.GetWxACode(componentAppid, appid, path, force)
	if resp.IsSuccess() {
		url := fmt.Sprintf("%s://%s%s", utils.Scheme(c.Request()), c.Host(), c.Path()) + "/" + resp.Url
		resp.Url = strings.Replace(url, "getwxacode", "prevqrcode", 1)
	}
	c.JSON(resp)
}

// 小程序二维码
func (this *MiniAction) GetWXACodeUnlimit(c iris.Context) {
	componentAppid := c.Params().Get("componentAppid")
	appid := c.Params().Get("appid")
	page := c.FormValue("page")
	scene := c.FormValue("scene")
	force, _ := strconv.ParseBool(c.FormValueDefault("force", "false"))
	resp := this.MiniSvr.GetWxACodeUnlimit(componentAppid, appid, page, scene, force)
	if resp.IsSuccess() {
		url := fmt.Sprintf("%s://%s%s", utils.Scheme(c.Request()), c.Host(), c.Path()) + "/" + resp.Url
		resp.Url = strings.Replace(url, "getwxacodeunlimit", "prevqrcode", 1)
	}
	c.JSON(resp)
}

func (this *MiniAction) PrevQrCode(c iris.Context) {
	componentAppid := c.Params().Get("componentAppid")
	appid := c.Params().Get("appid")
	fName := c.Params().Get("fName")
	down, _ := c.URLParamBool("down")
	b, err := this.MiniSvr.DownLoadQrCode(componentAppid, appid, fName)
	defer b.Close()
	if err != nil {
		c.JSON(service.NewServerErrorResp(err))
		return
	} else {
		if down {
			c.ContentType(context.ContentBinaryHeaderValue)
			c.Header("Content-Disposition", "attachment; filename=qrcode.jpg")
		} else {
			c.ContentType("image/jpeg")
		}
		io.Copy(c.ResponseWriter(), b)
	}
}

func (this *MiniAction) Preview(c iris.Context) {
	componentAppid := c.Params().Get("componentAppid")
	appid := c.Params().Get("appid")
	path := c.URLParam("path")
	force, err := c.URLParamBool("force")
	if err != nil {
		force = false
	}
	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		c.JSON(service.NewServerErrorResp(err))
		return
	}
	ret := this.MiniSvr.MiniPreview(componentAppid, appid, string(data), path, force)
	if ret.IsSuccess() {
		url := fmt.Sprintf("%s://%s%s", utils.Scheme(c.Request()), c.Host(), c.Path()) + "/" + ret.Url
		ret.Url = strings.Replace(url, "preview", "prevqrcode", 1)
	}
	c.JSON(ret)
}

// 获取授权小程序帐号的可选类目
func (this *MiniAction) GetCategory(c iris.Context) {
	this.getTransparentJson(c, GET_MINI_CATEGORY)
}

// 获取小程序的第三方提交代码的页面配置
func (this *MiniAction) GetPage(c iris.Context) {
	this.getTransparentJson(c, GET_MINI_PAGE)
}

// 将第三方提交的代码包提交审核
func (this *MiniAction) SubmitAudit(c iris.Context) {
	this.postTransparentJson(c, SUBMIT_MINI_AUDIT)
}

// 查询某个指定版本的审核状态
func (this *MiniAction) QueryAuditStatus(c iris.Context) {
	this.postTransparentJson(c, QUERY_MINI_AUDITSTATUS)
}

// 查询最新一次提交的审核状态
func (this *MiniAction) QueryLastAuditStatus(c iris.Context) {
	this.getTransparentJson(c, QUERY_MINI_LASTAUDITSTATUS)
}

// 发布已通过审核的小程序
func (this *MiniAction) DoRelease(c iris.Context) {
	this.postTransparentJson(c, DO_MINI_RELEASE)
}

// 修改小程序线上代码的可见状态
func (this *MiniAction) ChangeVisitStatus(c iris.Context) {
	this.postTransparentJson(c, CHANGE_MINI_VISITSTATUS)
}

// 小程序版本回退
func (this *MiniAction) RevertCodeRelease(c iris.Context) {
	this.getTransparentJson(c, DO_MINI_REVERTCODERELEASE)
}

// 查询当前设置的最低基础库版本及各版本用户占比
func (this *MiniAction) QueryWeAppSupportVersion(c iris.Context) {
	this.postTransparentJson(c, GET_MINI_WEAPPSUPPORTVERSION)
}

// 设置最低基础库版本
func (this *MiniAction) SetMinWeAppSupportVersion(c iris.Context) {
	this.postTransparentJson(c, SET_MINI_WEAPPSUPPORTVERSION)
}
