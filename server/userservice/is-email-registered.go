package userservice

import "log"

//IsEmailRegistered checks whether email is already registered or not
func (db *Service) IsEmailRegistered(email string) bool {
	var count int32
	err := db.QueryRow(
		"SELECT count(*) as counter FROM user WHERE email=?",
		email,
	).Scan(&count)
	log.Println(count)
	return err == nil && count > 0
}
