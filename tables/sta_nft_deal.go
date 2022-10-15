package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetStaNftDealTable(ctx *context.Context) table.Table {

	staNftDeal := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "exchange"))

	info := staNftDeal.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Address", "address", db.Varchar)
	info.AddField("Total_flow", "total_flow", db.Decimal)
	info.AddField("Gas_money", "gas_money", db.Decimal)
	info.AddField("Platform_money", "platform_money", db.Decimal)
	info.AddField("Now_profit", "now_profit", db.Decimal)
	info.AddField("Sales_volume", "sales_volume", db.Int)
	info.AddField("Sum_money", "sum_money", db.Decimal)
	info.AddField("Sta_time", "sta_time", db.Bigint)
	info.AddField("Deleted", "deleted", db.Tinyint)
	info.AddField("Create_time", "create_time", db.Bigint)
	info.AddField("Update_time", "update_time", db.Bigint)

	info.SetTable("sta_nft_deal").SetTitle("StaNftDeal").SetDescription("StaNftDeal")

	formList := staNftDeal.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Address", "address", db.Varchar, form.Text)
	formList.AddField("Total_flow", "total_flow", db.Decimal, form.Text)
	formList.AddField("Gas_money", "gas_money", db.Decimal, form.Text)
	formList.AddField("Platform_money", "platform_money", db.Decimal, form.Text)
	formList.AddField("Now_profit", "now_profit", db.Decimal, form.Text)
	formList.AddField("Sales_volume", "sales_volume", db.Int, form.Number)
	formList.AddField("Sum_money", "sum_money", db.Decimal, form.Text)
	formList.AddField("Sta_time", "sta_time", db.Bigint, form.Number)
	formList.AddField("Deleted", "deleted", db.Tinyint, form.Number)
	formList.AddField("Create_time", "create_time", db.Bigint, form.Number)
	formList.AddField("Update_time", "update_time", db.Bigint, form.Number)

	formList.SetTable("sta_nft_deal").SetTitle("StaNftDeal").SetDescription("StaNftDeal")

	return staNftDeal
}
