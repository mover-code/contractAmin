package tables

import (
    "github.com/GoAdminGroup/go-admin/context"
    "github.com/GoAdminGroup/go-admin/modules/db"
    "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
    "github.com/GoAdminGroup/go-admin/template/types"
    "github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetContract.
func GetContract(ctx *context.Context) (contractTable table.Table) {

    contractTable = table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

    // connect your custom connection
    // authorsTable = table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

    info := contractTable.GetInfo()
    info.AddField("ID", "id", db.Int).FieldSortable()
    info.AddField("合约地址", "addr", db.Varchar).FieldDisplay(
        func(model types.FieldModel) interface{} {
            // href := `<a class="btn btn-sm btn-primary grid-refresh" href="/admin/dao?contract=` + model.Value + `"><i class="fa fa-play"></i> ` + model.Value + `</a>`
            href := `<a class="btn btn-sm grid-refresh" target="_blank" href="/admin/web3?contract=` + model.Value + `"><i class="fa fa-play"></i> ` + model.Value + `</a>`
            return href
        })
    info.AddField("rpc", "chainRpc", db.Varchar)
    info.AddField("abi", "abi", db.Varchar).FieldHide()
    info.AddField("scan", "scan", db.Varchar)
    info.AddField("链id", "chainId", db.Varchar)

    info.SetTable("contracts").SetTitle("智能合约列表").SetDescription("智能合约信息")

    formList := contractTable.GetForm()
    formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
    formList.AddField("合约地址", "addr", db.Varchar, form.Text)
    formList.AddField("rpc", "chainRpc", db.Varchar, form.Text)
    formList.AddField("scan", "scan", db.Varchar, form.Text)
    formList.AddField("abi", "abi", db.Varchar, form.Text)
    formList.AddField("链id", "chainId", db.Varchar, form.Text)

    formList.SetTable("contracts").SetTitle("智能合约编辑").SetDescription("智能合约")

    return
}
