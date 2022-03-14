package influxdb

import (
	"encoding/json"
	"fmt"

	"stvsljl.com/stvsl/Sql"
)

//节点提交信息定义
type SubmitInfo struct {
	// 节点ID
	NodeId string
	// 气体浓度
	GasConcentration float64 `json:"GasConcentration"`
	// 温度
	Temperature float64 `json:"Temperature"`
	// PH
	PH float64 `json:"PH"`
	// 浊度
	Density float64 `json:"Density"`
	// 电导率
	Conductivity float64 `json:"Conductivity"`
	// 含氧量
	OxygenConcentration float64 `json:"OxygenConcentration"`
	// 重金属
	MetalConcentration float64 `json:"MetalConcentration"`
	// 溶解性固体
	SC float64 `json:"SC"`
	// 悬浮性固体
	FSC float64 `json:"FSC"`
	// 总氮
	TN float64 `json:"TN"`
	// 总磷
	TP float64 `json:"TP"`
	// 总有机碳
	TOC float64 `json:"TOC"`
	// 生物需氧量
	BOD float64 `json:"BOD"`
	// 化学需氧量
	COD float64 `json:"COD"`
	// 细菌总数
	BC int64 `json:"BC"`
	// 大肠杆菌数
	SLC int64 `json:"SLC"`
}

//解析json，使用json数据实例化SubmitInfo
func (submitInfo SubmitInfo) DecodeJSON(jsonstr string) error {
	err := json.Unmarshal([]byte(jsonstr), &submitInfo)
	return err
}

//将数据转吗成json返回
func (submitInfo SubmitInfo) EncodeJSON() (string, error) {
	jsonstr, err := json.Marshal(submitInfo)
	return string(jsonstr), err
}

//校验节点提交信息
func (submitInfo SubmitInfo) Check() bool {
	var s Sql.NodeInformation
	c, err := s.GetCheck(submitInfo.NodeId)
	if err != nil {
		return false
	}
	var result bool = true
	//submitInfo的checkValue的第一位是否为1且GasConcentration为0
	if c[0] == '1' && submitInfo.GasConcentration == 0 {
		result = false
	}
	//submitInfo的checkValue的第二位是否为1且Temperature为0
	if c[1] == '1' && submitInfo.Temperature == 0 {
		result = false
	}
	//submitInfo的checkValue的第三位是否为1且PH为0
	if c[2] == '1' && submitInfo.PH == 0 {
		result = false
	}
	//submitInfo的checkValue的第四位是否为1且Density为0
	if c[3] == '1' && submitInfo.Density == 0 {
		result = false
	}
	//submitInfo的checkValue的第五位是否为1且Conductivity为0
	if c[4] == '1' && submitInfo.Conductivity == 0 {
		result = false
	}
	//submitInfo的checkValue的第六位是否为1且OxygenConcentration为0
	if c[5] == '1' && submitInfo.OxygenConcentration == 0 {
		result = false
	}
	//submitInfo的checkValue的第七位是否为1且MetalConcentration为0
	if c[6] == '1' && submitInfo.MetalConcentration == 0 {
		result = false
	}
	//submitInfo的checkValue的第八位是否为1且SC为0
	if c[7] == '1' && submitInfo.SC == 0 {
		result = false
	}
	//submitInfo的checkValue的第九位是否为1且FSC为0
	if c[8] == '1' && submitInfo.FSC == 0 {
		result = false
	}
	//submitInfo的checkValue的第十位是否为1且TN为0
	if c[9] == '1' && submitInfo.TN == 0 {
		result = false
	}
	//submitInfo的checkValue的第十一位是否为1且TP为0
	if c[10] == '1' && submitInfo.TP == 0 {
		result = false
	}
	//submitInfo的checkValue的第十二位是否为1且TOC为0
	if c[11] == '1' && submitInfo.TOC == 0 {
		result = false
	}
	//submitInfo的checkValue的第十三位是否为1且BOD为0
	if c[12] == '1' && submitInfo.BOD == 0 {
		result = false
	}
	//submitInfo的checkValue的第十四位是否为1且COD为0
	if c[13] == '1' && submitInfo.COD == 0 {
		result = false
	}
	//submitInfo的checkValue的第十五位是否为1且BC为0
	if c[14] == '1' && submitInfo.BC == 0 {
		result = false
	}
	//submitInfo的checkValue的第十六位是否为1且SLC为0
	if c[15] == '1' && submitInfo.SLC == 0 {
		result = false
	}
	return result
}

// 打印输出
func (submitInfo SubmitInfo) Print() {
	fmt.Printf("SubmitInfo: %+v\n", submitInfo)
}

// 数据写入到时序数据库
func (submitInfo SubmitInfo) WriteToDB() error {
	err := Write(&submitInfo)
	return err
}
