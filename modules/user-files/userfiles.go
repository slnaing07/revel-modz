package userfiles

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

// a table to store agglomeration of Id's of users and data sets, meta info, etc
type UserFileInfo struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt time.Time

	// keys to the data table
	UserId int64 `sql:"not null"`
	FileId int64 `sql:"not null;unique"`

	Name   string
	Path   string
	Type   string
	Folder bool
	Size   int
}

type UserFile struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt time.Time

	// keys to the data table
	UserId int64 `sql:"not null"`
	FileId int64 `sql:"not null;unique"`

	Content []byte
}

type UserDataPermissions struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt time.Time

	// keys to the data table
	UserId int64 `sql:"not null"`
	FileId int64 `sql:"not null;unique"`

	// need to use RBAC
	// or
	// maybe PSQL has RBAC at the table level
	// if so, each user could have own table(s) DB(s)
	Permissions int64
}

type UserDataWire struct {
	Name    string
	Path    string
	Type    string
	Folder  bool
	Size    int
	Content string
}

func AddUserFile(db *gorm.DB, udt *UserFileInfo, content []byte) error {

	lastId, err := getLastFileIdByUserId(db, udt.UserId)
	if err != nil {
		revel.TRACE.Println(err)
		return err
	}
	nextId := lastId + 1
	revel.INFO.Println("Next DS id: ", nextId)
	udt.FileId = nextId

	err = db.Save(udt).Error
	if err != nil {
		revel.TRACE.Println(err)
		return err
	}

	ds := &UserFile{
		UserId:  udt.UserId,
		FileId:  udt.FileId,
		Content: content,
	}

	err = db.Save(ds).Error
	if err != nil {
		revel.TRACE.Println(err)
		return err
	}

	return err
}

func getLastFileByUserId(db *gorm.DB, uId int64) (int64, error) {
	var udt UserFileInfo
	err := db.Where(&UserFileInfo{UserId: uId}).Order("file_id desc").First(&udt).Error
	if err == gorm.RecordNotFound {
		return 0, nil
	}
	if err != nil {
		revel.TRACE.Println(err)
		return 0, err
	}
	return udt.FileId, nil
}

func GetUserFileInfos(db *gorm.DB, uId int64) ([]UserFileInfo, error) {
	var err error
	var infos []UserFileInfo
	err = db.Where(&UserFileInfo{UserId: uId}).Order("file_id").Find(&infos).Error
	if err != nil {
		revel.ERROR.Println(err)
		return nil, err
	}
	return infos, nil
}

func GetUserFileById(db *gorm.DB, uId, dsId int64) (ds *UserFile, err error) {
	err = db.Where(&UserFileInfo{UserId: uId, FileId: dsId}).First(ds).Error
	if err != nil {
		revel.TRACE.Println(err)
		return nil, err
	}
	return
}
