package handlers

import (
	"editor-backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 门诊医生登录
// @Param mzghxh body string true "门诊挂号序号"
// @Param dfname body string true "大夫姓名"
// @Param loginTime body string true "登陆时间"
// @Success 0
// @Router /outpatient/login [get]
func Login(c *gin.Context) {
	mzghxh := c.Query("mzghxh")
	dfname := c.Query("dfname")
	loginTime := c.Query("loginTime")
	log.Println(dfname)
	data, err := services.GetData(mzghxh)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"exist": false,
		})

		return
	}

	ok, lockBy, lockTime := services.TryLock(mzghxh, dfname, loginTime)
	log.Println(lockTime)
	resp := gin.H{
		"exist":    true,
		"data":     data,
		"readOnly": !ok,
	}
	if !ok {
		resp["owner"] = lockBy
		resp["lockTime"] = lockTime
	}

	c.JSON(http.StatusOK, resp)
}

func Logout(c *gin.Context) {
	mzghxh := c.Query("mzghxh")
	dfname := c.Query("dfname")
	log.Println(dfname)
	services.Unlock(mzghxh, dfname)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
