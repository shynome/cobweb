package tables

import (
	"bytes"
	"fmt"
	"html/template"

	_ "embed"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/shynome/cobweb/models"
	"github.com/shynome/cobweb/v2ray"
	"github.com/v2fly/v2ray-core/v4/common/uuid"
)

//go:embed v2ray-user-share-col.html
var shareColTpl string
var shareCol, _ = template.New("").Parse(shareColTpl)

func GetV2rayUsersTable(ctx *context.Context) table.Table {

	v2 := v2ray.FromContext(ctx.Request.Context())
	orm := models.GetORM()

	v2rayUsers := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := v2rayUsers.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Integer).
		FieldFilterable()
	info.AddField("Username", "username", db.Text)
	info.AddField("备注", "remark", db.Text)
	info.AddField("Uuid", "uuid", db.Text).
		FieldDisplay(func(value types.FieldModel) interface{} {
			var tpl bytes.Buffer
			remark, ok := value.Row["remark"].(string)
			if !ok || remark == "" {
				remark = value.Row["username"].(string)
			}
			if err := shareCol.Execute(&tpl, map[string]string{
				"uuid":   value.Value,
				"remark": remark,
			}); err != nil {
				return err
			}
			return tpl.String()
		})
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("禁用", "deleted_at", db.Datetime).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value != "" {
				return `<span style="color: gray">已禁用</span>`
			}
			return `<span style="color: green">正常</span>`
		})
	info.AddActionButton("禁用", action.Ajax("/v2ray_user/disable", func(ctx *context.Context) (success bool, msg string, data interface{}) {
		orm := models.GetORM()
		id := ctx.FormValue("id")
		fmt.Println(id)
		u := models.V2rayUser{}
		if err := orm.Where("id", id).Delete(&u).Error; err != nil {
			return false, "禁用失败", ""
		}
		return true, "禁用成功", ""
	}))
	info.AddActionButton("启用", action.Ajax("/v2ray_user/enable", func(ctx *context.Context) (success bool, msg string, data interface{}) {
		orm := models.GetORM()
		id := ctx.FormValue("id")
		u := models.V2rayUser{}
		if err := orm.Unscoped().Model(&u).Where("id", id).Update("deleted_at", nil).Error; err != nil {
			return false, "启用失败", ""
		}
		return true, "启用成功", ""
	}))

	info.SetPreDeleteFn(func(ids []string) (err error) {
		var users []models.V2rayUser
		if err = orm.Find(&users, ids).Error; err != nil {
			return
		}
		for _, u := range users {
			v2.RemoveUser(u.Username)
		}
		return
	})

	info.SetTable("v2ray_users").SetTitle("V2rayUsers").SetDescription("V2rayUsers")

	formList := v2rayUsers.GetForm()
	formList.AddField("Id", "id", db.Integer, form.Default).
		FieldDisableWhenCreate().
		FieldNotAllowEdit()
	formList.AddField("Username", "username", db.Text, form.Text).
		FieldMust()
	_uuid := uuid.New()
	formList.AddField("Uuid", "uuid", db.Text, form.Text).
		FieldMust().
		FieldDefault(_uuid.String())
	formList.AddField("remark", "remark", db.Text, form.Text)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldNowWhenInsert().
		FieldNotAllowEdit()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldNow().
		FieldNowWhenUpdate().
		FieldNotAllowEdit()

	formList.SetPostHook(func(values form2.Values) (err error) {
		if values.IsInsertPost() {
			user := models.V2rayUser{
				Username: values.Get("username"),
				Uuid:     values.Get("uuid"),
			}
			return v2.AddUser(user)
		} else if values.IsUpdatePost() {
			user := models.V2rayUser{}
			if err = orm.Unscoped().Where("id", values.Get("id")).Find(&user).Error; err != nil {
				return
			}
			// 删除有可能失败
			_ = v2.RemoveUser(user.Username)
			// 用户已被删除, 不用加回来了
			if user.DeletedAt.Valid {
				return
			}
			if err = v2.AddUser(user); err != nil {
				return
			}
		}
		return nil
	})

	formList.SetTable("v2ray_users").SetTitle("V2rayUsers").SetDescription("V2rayUsers")

	return v2rayUsers
}
