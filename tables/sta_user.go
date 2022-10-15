package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

func GetStaUserTable(ctx *context.Context) table.Table {

	staUser := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "exchange"))

	info := staUser.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Bigint).
		FieldFilterable()
	info.AddField("访问量", "login_amount", db.Int)
	// info.AddField("Deleted", "deleted", db.Tinyint)
	info.AddField("时间", "create_time", db.Bigint).FieldDisplay(func(model types.FieldModel) interface{} {
		return TimeToStr(model.Value)
	})
	// info.AddField("Update_time", "update_time", db.Bigint)
	// info.AddField("Sta_time", "sta_time", db.Bigint)

	info.SetTable("sta_user").SetTitle("访问量").SetDescription("访问量")

	// formList := staUser.GetForm()
	// formList.AddField("Id", "id", db.Bigint, form.Default)
	// formList.AddField("Login_amount", "login_amount", db.Int, form.Number)
	// formList.AddField("Deleted", "deleted", db.Tinyint, form.Number)
	// formList.AddField("Create_time", "create_time", db.Bigint, form.Number)
	// formList.AddField("Update_time", "update_time", db.Bigint, form.Number)
	// formList.AddField("Sta_time", "sta_time", db.Bigint, form.Number)

	// formList.SetTable("sta_user").SetTitle("StaUser").SetDescription("StaUser")

	return staUser
}
