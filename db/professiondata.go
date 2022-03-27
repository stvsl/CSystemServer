package Sql

import (
	"errors"
	"time"
)

// Professiondata 专业机构检测结果
type Professiondata struct {
	ID             string    `gorm:"column:ID" json:"id"`                           // 节点编号
	COMMUNITY      string    `gorm:"column:COMMUNITY" json:"community"`             // 检测机构名称
	COMID          string    `gorm:"column:COMID" json:"comId"`                     // 检测机构的账户ID
	TIME           time.Time `gorm:"column:TIME" json:"time"`                       // 记录时间
	STATUS         int       `gorm:"column:STATUS" json:"status"`                   // 检测结果状态
	DESCRIPTION    string    `gorm:"column:DESCRIPTION" json:"description"`         // 描述
	PhDirectHigh   float32   `gorm:"column:PH_DIRECT_HIGH" json:"phdirecthigh"`     // PH直接排放上限
	PhDirectLow    float32   `gorm:"column:PH_DIRECT_LOW" json:"phdirectlow"`       // PH直接排放下限
	PhIndirectHigh float32   `gorm:"column:PH_INDIRECT_HIGH" json:"phindirecthigh"` // PH间接排放上限
	PhIndirectLow  float32   `gorm:"column:PH_INDORECT_LOW" json:"phindorectlow"`   // PH间接排放下限
	CODDirect      float32   `gorm:"column:COD_DIRECT" json:"coddirect"`            // 化学需氧量直接
	CODIndirect    float32   `gorm:"column:COD_INDIRECT" json:"codindirect"`        // 化学需氧量间接
	TPDirect       float32   `gorm:"column:TP_DIRECT" json:"tpdirect"`              // 总磷直接
	TPIndirect     float32   `gorm:"column:TP_INDIRECT" json:"tpindirect"`          // 总磷间接
	TNDirect       float32   `gorm:"column:TN_DIRECT" json:"tndirect"`              // 总氮直接
	IPIndirect     float32   `gorm:"column:IP_INDIRECT" json:"ipindirect"`          // 总氮间接
	ANDirect       float32   `gorm:"column:AN_DIRECT" json:"andirect"`              // 氨氮直接
	ANINDirect     float32   `gorm:"column:AN_INDIRECT" json:"anindirect"`          // 氨氮间接
	OCCDirect      float32   `gorm:"column:OCC_DIRECT" json:"occdirect"`            // 石油类直接
	OCCIndirect    float32   `gorm:"column:OCC_INDIRECT" json:"occindirect"`        // 石油类间接
	FSCDirectT     float32   `gorm:"column:FSC_DIRECT_T" json:"fscdirectt"`         // 悬浮物直接T
	FSCIndirectT   float32   `gorm:"column:FSC_INDIRECT_T" json:"fscindirectt"`     // 悬浮物间接T
	FSCDirectO     float32   `gorm:"column:FSC_DIRECT_O" json:"fscdirecto"`         // 悬浮物直接O
	FSCIndirectO   float32   `gorm:"column:FSC_INDIRECT_O" json:"fscindirecto"`     // 悬浮物间接O
	SADirect       float32   `gorm:"column:SA_DIRECT" json:"sadirect"`              // 硫化物直接
	SAIndirect     float32   `gorm:"column:SA_INDIRECT" json:"saindirect"`          // 硫化物间接
	FDirect        float32   `gorm:"column:F_DIRECT" json:"fdirect"`                // 氟化物直接
	FIndirect      float32   `gorm:"column:F_INDIRECT" json:"findirect"`            // 氟化物间接
	Cu             float32   `gorm:"column:CU" json:"cu"`                           // 总铜
	Zn             float32   `gorm:"column:ZN" json:"zn"`                           // 总锌
	Sn             float32   `gorm:"column:SN" json:"sn"`                           // 总锡
	Sb             float32   `gorm:"column:SB" json:"sb"`                           // 总锑
	Hg             float32   `gorm:"column:HG" json:"hg"`                           // 总汞
	Cd             float32   `gorm:"column:CD" json:"cd"`                           // 总镉
	Pb             float32   `gorm:"column:PB" json:"pb"`                           // 总铅
	As             float32   `gorm:"column:AS" json:"as"`                           // 总砷
	Cr6            float32   `gorm:"column:CR6" json:"cr6"`                         // 六价铬
	Gc             float32   `gorm:"column:GC" json:"gs"`                           // 气体浓度
	Density        float32   `gorm:"column:DENSITY" json:"density"`                 // 浊度
	Conductivity   float32   `gorm:"column:CONDUCTIVITY" json:"conductivity"`       // 电导率
	Mc             float32   `gorm:"column:MC" json:"mc"`                           // 重金属
	Sc             float32   `gorm:"column:SC" json:"sc"`                           // 溶解性固体
	Toc            float32   `gorm:"column:TOC" json:"toc"`                         // 总有机碳
	BOD5Direct     float32   `gorm:"column:BOD5_DIRECT" json:"bod5direct"`          // 五日生化需氧量直接
	BOD5Indirect   float32   `gorm:"column:BOD5_INDIRECT" json:"bod5indirect"`      // 五日生化需氧量间接
	BOD            float32   `gorm:"column:BOD" json:"bod"`                         // 生物需氧量
	PDirect        float32   `gorm:"column:P_DIRECT" json:"pdirect"`                // 总磷直接
	Bc             float32   `gorm:"column:BC" json:"bc"`                           // 细菌总数
	Slc            float32   `gorm:"column:SLC" json:"slc"`                         // 大肠杆菌数
	COLORDirect    float32   `gorm:"column:COLOR_DIRECT" json:"colordirect"`        // 色度直接
	COLORIndirect  float32   `gorm:"column:COLOR_INDIRECT" json:"colorindirect"`    // 色度间接
	AFDirect       float32   `gorm:"column:AF_DIRECT" json:"afdirect"`              // 动植物油直接
	AFINDirect     float32   `gorm:"column:AF_INDIRECT" json:"afindirect"`          // 动植物油间接
	CLDirect       float32   `gorm:"column:CL_DIRECT" json:"cldirect"`              // 氯离子直接
	CLIndirect     float32   `gorm:"column:CL_INDIRECT" json:"clindirect"`          // 氯离子间接
	PINDirect      float32   `gorm:"column:P_INDIRECT" json:"pindirect"`            // 总磷间接
	Cr             float32   `gorm:"column:CR" json:"cr"`                           // 总铬
	Ton            float32   `gorm:"column:TON" json:"ton"`                         // 单位
}

// 添加记录
func (professiondata Professiondata) Add() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	return db.Create(&professiondata).Error
}

// 修改记录
func (professiondata Professiondata) Update() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	return db.Save(&professiondata).Error
}

// 删除记录
func (professiondata Professiondata) Delete() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	return db.Delete(&professiondata).Error
}

func (professiondata Professiondata) GetByID(id string) (Professiondata, error) {
	db, i, err := GetDB()
	if err != nil {
		return professiondata, err
	}
	defer Release(i)
	// 查询最新的一条记录
	err = db.Where("id = ?", id).Last(&professiondata).Error
	return professiondata, err
}

func (professiondata Professiondata) GetByIDs(ids []string) (*[]Professiondata, error) {
	db, i, err := GetDB()
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("ids is empty")
	}
	defer Release(i)
	var professiondatas []Professiondata
	err = db.Where("id in (?)", ids).Find(&professiondatas).Error
	return &professiondatas, err
}

func (professiondata Professiondata) GetAllByIDs(ids []string, startTime time.Time, endTime time.Time) (*[]Professiondata, error) {
	db, i, err := GetDB()
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("ids is empty")
	}
	defer Release(i)
	var professiondatas []Professiondata
	// 查询全部符合的记录
	err = db.Where("id in (?) and TIME >= ? and TIME <= ?", ids, startTime, endTime).Find(&professiondatas).Error
	return &professiondatas, err
}

// 查询从XX年XX月XX日到XX年XX月XX日id为XXX的多个节点的记录并返回长度
func (professiondata Professiondata) GetByTime(id []string, startTime, endTime string) ([]*Professiondata, int, error) {
	db, i, err := GetDB()
	if err != nil {
		return nil, 0, err
	}
	defer Release(i)
	var professiondatas []*Professiondata
	// 依次查询每个id的记录
	for _, v := range id {
		var temp []Professiondata
		err = db.Where("id = ?", v).Where("time >= ?", startTime).Where("time <= ?", endTime).Find(&temp).Error
		if err != nil {
			return nil, 0, err
		}
		for _, v := range temp {
			professiondatas = append(professiondatas, &v)
		}
	}
	return professiondatas, len(professiondatas), nil
}
