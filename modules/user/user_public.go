package user

import (
	"github.com/jinzhu/gorm"
	"github.com/robfig/revel"
)

func GenerateNewUserId(db *gorm.DB) int64 {
	var unique bool
	var tries int
	for !unique {
		if tries >= 10 {
			revel.ERROR.Println("tried", tries, "times to generate a unqiue id")
			return 0
		}
		id := generateRandId()
		u := GetUserBasicByUserId(db, id)
		if u == nil {
			return id
		} else {
			tries++
		}
	}
	panic("shouldn't get here")
	return -1
}
