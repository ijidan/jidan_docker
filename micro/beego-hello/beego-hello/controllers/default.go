package controllers

import (
	"beego-hello/lib"
	"github.com/astaxie/beego"
	"os"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	host, _ := os.Hostname()
	ip := c.Ctx.Input.IP()
	remoteAddr := c.Ctx.Request.RemoteAddr
	realIp := c.Ctx.Request.Header.Get("X-Real-IP")
	publicIp := lib.GetPublicIpAddr()
	interfaceIp := lib.GetInterfaceIps()
	localIp := lib.GetLocalIp()
	proxy := c.Ctx.Input.Proxy()

	nginxHost := c.Ctx.Request.Header.Get("Host")
	httpHost := c.Ctx.Request.Header.Get("Http-Host")
	XForwardedFor := c.Ctx.Request.Header.Get("X-Forwarded-For")
	HTTP_X_FORWARDED_FOR := c.Ctx.Request.Header.Get("HTTP_X_FORWARDED_FOR")

	c.Data["host"] = host
	c.Data["httpHost"] = httpHost
	c.Data["ip"] = ip
	c.Data["remoteAddr"] = remoteAddr
	c.Data["realIp"] = realIp
	c.Data["publicIp"] = publicIp
	c.Data["interfaceIp"] = interfaceIp
	c.Data["localIp"] = localIp
	c.Data["proxy"] = proxy
	c.Data["nginxHost"] = nginxHost
	c.Data["XForwardedFor"] = XForwardedFor
	c.Data["HTTP_X_FORWARDED_FOR"] = HTTP_X_FORWARDED_FOR
	c.TplName = "index.html"
}
