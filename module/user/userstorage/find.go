package userstorage

import (
	"programming_fresher/common"
)

func (s *userMySQL) FindUserStorage(conditions map[string]interface{}) (*common.UserInfo, error) {
	var data common.UserInfo
	db := s.db.Table(common.UserInfo{}.TableName())

	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, common.ErrorDBNotFound(err)
	}

	return &data, nil
}
