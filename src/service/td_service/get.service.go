package td_service

import (
	"api/src/config"
	"api/src/model"
)

func GetService(u *[]model.UserModel) (err error) {
	query := `SELECT	u.username AS username
					, u.first_name AS first_name
					, u.last_name AS last_name
				FROM public.user u`
	if err = config.DB.Raw(query).Scan(&u).Error; err != nil {
		return err
	}
	return nil
}
