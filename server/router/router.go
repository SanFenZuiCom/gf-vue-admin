package router

import (
	extra "gf-vue-admin/router/extra"
	"gf-vue-admin/router/internal"
	system "gf-vue-admin/router/system"
	workflow "gf-vue-admin/router/workflow"
	"github.com/gogf/gf/frame/g"
)

var Routers = new(routers)

type routers struct{}

func (r *routers) Init() {
	public := g.Server().Group("")
	{ // 无需鉴权中间件
		system.NewBaseGroup(public).Init()
	}
	private := g.Server().Group("").Middleware(internal.Middleware.JwtAuth, internal.Middleware.CasbinRbac)
	{ // 需要Jwt鉴权, casbin鉴权
		system.NewApiRouter(private).Init()
		system.NewMenuRouter(private).Init()
		system.NewEmailRouter(private).Init()
		system.NewAdminRouter(private).Init()
		system.NewSystemRouter(public).Init()
		system.NewCasbinRouter(private).Init()
		system.NewGenerateRouter(private).Init()
		system.NewAuthorityRouter(private).Init()
		system.NewDictionaryRouter(private).Init()
		system.NewJwtBlacklistRouter(private).Init()
		system.NewOperationRecordRouter(private).Init()
		system.NewDictionaryDetailRouter(private).Init()

		extra.NewFileRouter(public).Init()
		extra.NewExcelRouter(private).Init()
		extra.NewSimpleUploaderRouter(private).Init()

		workflow.NewWorkflowRouter(private).Init()
	}
}
