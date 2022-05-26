package entities

import (
	"editor-backend/internal/database"
	"log"
)

type EmrDistLock struct {
	Mzghxh   string // 门诊挂号序号 AbcDef abc_def
	Owner    string
	LockTime string
}

func (EmrDistLock) TableName() string {
	return "emr_dist_lock"
}

func AddLock(mzghxh, owner, lockTime string) error {
	distLock := EmrDistLock{
		Mzghxh:   mzghxh,
		Owner:    owner,
		LockTime: lockTime,
	}

	db := database.DB
	if err := db.Create(&distLock).Error; err != nil {
		return err
	}

	return nil
}

func GetLock(mzghxh string) (EmrDistLock, error) {
	distLock := EmrDistLock{
		Mzghxh: mzghxh,
	}

	db := database.DB
	if err := db.Where(&distLock).First(&distLock).Error; err != nil {
		return EmrDistLock{}, err
	}

	return distLock, nil
}

func DeleteLock(mzghxh, owner string) {
	db := database.DB
	err := db.Where("owner = ?", owner).Delete(&EmrDistLock{Mzghxh: mzghxh})
	log.Printf("%v\n", err)

}

func UpdateLock(mzghxh, owner, lockTime string) error {
	db := database.DB
	if err := db.Model(&EmrDistLock{Mzghxh: mzghxh}).Where("owner = ?", owner).Update("lock_time", lockTime).Error; err != nil {
		return err
	}

	return nil
}
