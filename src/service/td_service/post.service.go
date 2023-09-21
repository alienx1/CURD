package td_service

import (
	"api/src/config"
	"api/src/model"
)

func PostService(u *model.UserModel) (err error) {
	query := `INSERT INTO public.user (username, first_name, last_name)
	VALUES (?, ?, ?)`
	if err = config.DB.Exec(query, u.Username, u.FirstName, u.LastName).Error; err != nil {
		return err
	}
	return nil
}
