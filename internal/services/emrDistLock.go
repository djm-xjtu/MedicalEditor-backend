package services

import (
	"editor-backend/internal/entities"
	"log"
	"strings"
)

const EXPIRE_INTERVAL = 60

func TryLock(mzghxh, owner, lockTime string) (bool, string, string) {
	err := entities.AddLock(mzghxh, owner, lockTime)
	if err != nil {
		log.Println(err)
		distLock, _ := entities.GetLock(mzghxh)
		a := strings.Replace(distLock.LockTime, "T", " ", 1)
		b := strings.Replace(a, "Z", "", 1)
		return false, distLock.Owner, b
	}

	return true, "", ""
}

func Unlock(mzghxh, owner string) {
	entities.DeleteLock(mzghxh, owner)
}

// func TryUpdateLock(mzghxh, owner, lockTime string) bool {
// 	err := entities.UpdateLock(mzghxh, owner, lockTime)

// }
