package echo

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/echo"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func NewEchoHandler() http.Handler {
	e := echo.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)
	template.AddComp(chartjs.NewChart())

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJson(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(e); err != nil {
		panic(err)
	}

	e.GET("/admin", func(context echo.Context) error {
		engine.Content(context, func(ctx interface{}) (types.Panel, error) {
			return datamodel.GetContent()
		})
		return nil
	})

	return e
}
