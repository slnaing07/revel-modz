package maillist

import (
	"github.com/jinzhu/gorm"
)

func AddUser(db *gorm.DB, uId int64, email, list string) error {
	mu := &MaillistUser{
		UserId: uId,
		Email:  email,
		List:   list,
	}
	err := addUser(db, mu)
	return err
}

func DeleteUser(db *gorm.DB, uId int64) error {
	return deleteUserById(db, uId)
}

func GetUserById(db *gorm.DB, uId int64) (*MaillistUser, error) {
	return getUserById(db, uId)
}

func GetUserByEmail(db *gorm.DB, email string) (*MaillistUser, error) {
	return getUserByEmail(db, email)
}

func GetAllUsers(db *gorm.DB) ([]MaillistUser, error) {
	return getAllUsers(db)
}

func GetUsersByList(db *gorm.DB, list string) ([]MaillistUser, error) {
	return getUsersByList(db, list)
}

func SaveDraft(db *gorm.DB, author, list, subject, textbody, htmlbody string) error {
	mm := &MaillistMessage{
		Author:    author,
		List:      list,
		Subject:   subject,
		PlainBody: textbody,
		HtmlBody:  htmlbody,
	}
	err := saveDraftMessage(db, mm)
	return err
}

func DeleteDraft(db *gorm.DB, subject string) error {
	return deleteDraftMessage(db, subject)
}

func SendMessage(db *gorm.DB, subject string) error {
	return sendMessage(db, subject)
}

func GetDraft(db *gorm.DB, subject string) (*MaillistMessage, error) {
	return getDraftMessage(db, subject)
}

func GetAllDraftMessages(db *gorm.DB) ([]MaillistMessage, error) {
	var mms []MaillistMessage
	err := db.Where(&MaillistMessage{Draft: true}).Find(&mms).Error
	if err != nil {
		return nil, err
	}
	return mms, nil
}

func GetAllMessages(db *gorm.DB) ([]MaillistMessage, error) {
	var mms []MaillistMessage
	err := db.Find(&mms).Error
	if err != nil {
		return nil, err
	}
	return mms, nil
}
