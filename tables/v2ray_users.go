package tables

import (
	"bytes"
	"html/template"

	_ "embed"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
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
	info.AddField("Uuid", "uuid", db.Text).
		FieldDisplay(func(value types.FieldModel) interface{} {
			var tpl bytes.Buffer
			if err := shareCol.Execute(&tpl, map[string]string{"uuid": value.Value}); err != nil {
				return err
			}
			return tpl.String()
		})
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

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
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldNowWhenInsert().
		FieldNotAllowEdit()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldNow().
		FieldNowWhenUpdate().
		FieldNotAllowEdit()

	formList.SetPostHook(func(values form2.Values) error {
		user := models.V2rayUser{
			Username: values.Get("username"),
			Uuid:     values.Get("uuid"),
		}
		if values.IsInsertPost() {
			return v2.AddUser(user)
		} else if values.IsUpdatePost() {
			var err error
			if err = v2.RemoveUser(user.Username); err != nil {
				return err
			}
			if err = v2.AddUser(user); err != nil {
				return err
			}
		}
		return nil
	})

	formList.SetTable("v2ray_users").SetTitle("V2rayUsers").SetDescription("V2rayUsers")

	return v2rayUsers
}
