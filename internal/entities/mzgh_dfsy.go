package entities

import "editor-backend/internal/database"

type MzghDfsy struct {
	Mzghxh string // 门诊挂号序号 AbcDef abc_def
	Xm     string // 姓名
	Xb     string // 性别
	Cssj   string // 出生时间
	Nl     string // 年龄
	Kb     string // 科别
	YbDept string // 医保xx
	Ghfs   string // 挂号方式
	Df     string // 大夫
	Ghje   int    // 挂号金额
	Zlje   int    // 治疗金额
	Je     int    // 金额
	Rq     string // 日期
	Cdno   string // 病人编号
	Fylb   string // 费用类别
	Tag    string // ？？
	Qy     string // ？？
	CdJqm  string //
	Sfzhm  string // 身份证号码
	Class  string // 家庭地址
	Tel    string // 电话号码
	Sxbl   int    // ???
	Lry    string // 录入员
	Zdlb   string // 诊断类别
}

func (MzghDfsy) TableName() string {
	return "MZGH_DFSY"
}

func GetMzghDfsy(mzghxh string) (MzghDfsy, error) {
	m := MzghDfsy {
		Mzghxh: mzghxh,
	}

	db := database.DB
	// select * from MZGH_DFSY where mzghxh = .. limit 1
	if err := db.Where(&m).First(&m).Error; err != nil {
		return MzghDfsy{}, err
	}

	return m, nil
}
