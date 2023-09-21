package td_service

import "api/src/config"

func Delete(user string) (err error) {
	query := `DELETE FROM public.user
				WHERE username = ?;`
	if err = config.DB.Exec(query, user).Error; err != nil {
		return err
	}
	return nil
}
