package mini

import (
	. "github.com/jacktea/wxproxy/common"
	"github.com/kataras/iris/v12"
)

//设置小程序服务域名
func (this *MiniAction) ModifyDomain(c iris.Context) {
	this.postTransparentJson(c, SET_MINI_SVRDOMAIN)
}

//设置小程序业务域名
func (this *MiniAction) Setwebviewdomain(c iris.Context) {
	this.postTransparentJson(c, SET_MINI_WEBDOMAIN)
}
