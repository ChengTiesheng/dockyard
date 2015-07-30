package middleware

import (
	"fmt"

	"github.com/Unknwon/macaron"

	"github.com/containerops/wrench/setting"
)

//TBD:codes as below should be updated when user config management is ready
func setRespHeaders() macaron.Handler {
	return func(ctx *macaron.Context) {
		ctx.Resp.Header().Set("Content-Type", "application/json;charset=UTF-8")
		ctx.Resp.Header().Set("X-Docker-Registry-Standalone", "True")                                         //Standalone
		ctx.Resp.Header().Set("X-Docker-Registry-Version", setting.Version)                                   //Version
		ctx.Resp.Header().Set("X-Docker-Registry-Config", setting.RunMode)                                    //Config
		ctx.Resp.Header().Set("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\",Token", "containerops.me")) //docker V2
		ctx.Resp.Header().Set("Docker-Distribution-API-Version", "registry/2.0")                              //docker V2
	}
}
