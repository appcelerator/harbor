// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/goharbor/harbor/src/common/models"
)

// AddUserOrg add UserOrg to the database.
func AddUserOrg(userOrg *models.UserOrg) (int64, error) {
	now := time.Now()
	userOrg.CreationTime = now
	userOrg.UpdateTime = now
	return GetOrmer().Insert(userOrg)
}

// GetUserOrg returns UserOrg by user ID.
func GetUserOrg(userID int) (*models.UserOrg, error) {
	q := models.UserOrg{UserID: userID}
	err := GetOrmer().Read(&q, "UserID")
	if err == orm.ErrNoRows {
		return nil, nil
	}
	return &q, err
}

// UpdateUserOrg update the UserOrg object.
func UpdateUserOrg(userOrg *models.UserOrg) error {
	userOrg.UpdateTime = time.Now()
	_, err := GetOrmer().Update(userOrg)
	return err
}

// UpsertUserOrg insert or update the UserOrg object.
func UpsertUserOrg(userOrg *models.UserOrg) error {
	uo, err := GetUserOrg(userOrg.UserID)
	if err != nil {
		return err
	}
	if uo == nil {
		_, err := AddUserOrg(userOrg)
		return err
	}
	return UpdateUserOrg(userOrg)
}
