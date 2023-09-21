package td_service

import (
	"api/src/config"
	"api/src/model"
)

func PutService(u *model.UserModel) (err error) {
	query := `UPDATE public.user
	SET  first_name = ?, last_name = ?
	WHERE username = ?;`
	if err = config.DB.Exec(query, u.FirstName, u.LastName, u.Username).Error; err != nil {
		return err
	}
	return nil
}
