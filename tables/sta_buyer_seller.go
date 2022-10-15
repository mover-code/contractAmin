package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetStaBuyerSellerTable(ctx *context.Context) table.Table {

	staBuyerSeller := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "exchange"))

	info := staBuyerSeller.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("User_address", "user_address", db.Varchar)
	info.AddField("Type", "type", db.Int)
	info.AddField("Tranfer_amount", "tranfer_amount", db.Int)
	info.AddField("Sta_time", "sta_time", db.Bigint)
	info.AddField("Deleted", "deleted", db.Tinyint)
	info.AddField("Create_time", "create_time", db.Bigint)
	info.AddField("Update_time", "update_time", db.Bigint)

	info.SetTable("sta_buyer_seller").SetTitle("StaBuyerSeller").SetDescription("StaBuyerSeller")

	formList := staBuyerSeller.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("User_address", "user_address", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Int, form.Number)
	formList.AddField("Tranfer_amount", "tranfer_amount", db.Int, form.Number)
	formList.AddField("Sta_time", "sta_time", db.Bigint, form.Number)
	formList.AddField("Deleted", "deleted", db.Tinyint, form.Number)
	formList.AddField("Create_time", "create_time", db.Bigint, form.Number)
	formList.AddField("Update_time", "update_time", db.Bigint, form.Number)

	formList.SetTable("sta_buyer_seller").SetTitle("StaBuyerSeller").SetDescription("StaBuyerSeller")

	return staBuyerSeller
}
