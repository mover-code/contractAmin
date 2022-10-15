package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetStaNftTable(ctx *context.Context) table.Table {

	staNft := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "exchange"))

	info := staNft.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Address", "address", db.Varchar)
	info.AddField("Count_nft", "count_nft", db.Int)
	info.AddField("Reviewed_nft", "reviewed_nft", db.Int)
	info.AddField("Newly_added_nft", "newly_added_nft", db.Int)
	info.AddField("Change_hands_nft", "change_hands_nft", db.Int)
	info.AddField("Sta_time", "sta_time", db.Bigint)
	info.AddField("Deleted", "deleted", db.Tinyint)
	info.AddField("Create_time", "create_time", db.Bigint)
	info.AddField("Update_time", "update_time", db.Bigint)

	info.SetTable("sta_nft").SetTitle("StaNft").SetDescription("StaNft")

	formList := staNft.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Address", "address", db.Varchar, form.Text)
	formList.AddField("Count_nft", "count_nft", db.Int, form.Number)
	formList.AddField("Reviewed_nft", "reviewed_nft", db.Int, form.Number)
	formList.AddField("Newly_added_nft", "newly_added_nft", db.Int, form.Number)
	formList.AddField("Change_hands_nft", "change_hands_nft", db.Int, form.Number)
	formList.AddField("Sta_time", "sta_time", db.Bigint, form.Number)
	formList.AddField("Deleted", "deleted", db.Tinyint, form.Number)
	formList.AddField("Create_time", "create_time", db.Bigint, form.Number)
	formList.AddField("Update_time", "update_time", db.Bigint, form.Number)

	formList.SetTable("sta_nft").SetTitle("StaNft").SetDescription("StaNft")

	return staNft
}
