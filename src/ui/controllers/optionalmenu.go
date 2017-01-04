package controllers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	dao "github.com/vmware/harbor/src/common/daomongo"
	"github.com/vmware/harbor/src/common/models"
	"github.com/vmware/harbor/src/common/utils/log"
)

// OptionalMenuController handles request to /optional_menu
type OptionalMenuController struct {
	BaseController
}

// Get renders optional menu, Admin user has "Add User" menu
func (omc *OptionalMenuController) Get() {
	sessionUserID := omc.GetSession("userId")

	var hasLoggedIn bool
	var allowAddNew bool

	var isAdminForLdap bool
	var allowSettingAccount bool

	if sessionUserID != nil {
		hasLoggedIn = true
		userID := sessionUserID.(bson.ObjectId)
		u, err := dao.GetUser(models.User{UserID: userID})
		if err != nil {
			log.Errorf("Error occurred in GetUser, error: %v", err)
			omc.CustomAbort(http.StatusInternalServerError, "Internal error.")
		}
		if u == nil {
			log.Warningf("User was deleted already, user id: %v, canceling request.", userID)
			omc.CustomAbort(http.StatusUnauthorized, "")
		}
		omc.Data["Username"] = u.Username

		if u.Username == "admin" {
			isAdminForLdap = true
		}

		if omc.AuthMode == "db_auth" || isAdminForLdap {
			allowSettingAccount = true
		}

		isAdmin, err := dao.IsAdminRole(userID)
		if err != nil {
			log.Errorf("Error occurred in IsAdminRole: %v", err)
			omc.CustomAbort(http.StatusInternalServerError, "")
		}

		if isAdmin && omc.AuthMode == "db_auth" {
			allowAddNew = true
		}
	}
	omc.Data["AddNew"] = allowAddNew
	omc.Data["SettingAccount"] = allowSettingAccount
	omc.Data["HasLoggedIn"] = hasLoggedIn
	omc.TplName = "optional-menu.htm"
	omc.Render()

}
