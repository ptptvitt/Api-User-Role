package storageuser

import (
	"context"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

func (s *sqlStore) UpdateUser(ctx context.Context,
	data *usermodel.UserUpdate,
	conditions map[string]interface{}) error {
	db := s.db

	if err := db.Where(conditions).Table("user_role").
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
