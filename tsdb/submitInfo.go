package influxdb

import (
	"encoding/json"
	"fmt"
)

//节点提交信息定义
type SubmitInfo struct {
	// 节点ID
	NodeId string `json:"NodeId"`
	// 校验值
	CheckValue string `json:"CheckValue"`
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
	SolidsConcentration float64 `json:"SolidsConcentration"`
	// 悬浮性固体
	FloatingSolidsConcentration float64 `json:"FloatingSolidsConcentration"`
	// 总氮
	TotalNitrogen float64 `json:"TotalNitrogen"`
	// 总磷
	TotalPhosphorus float64 `json:"TotalPhosphorus"`
	// 总有机碳
	TotalOrganicCarbon float64 `json:"TotalOrganicCarbon"`
	// 生物需氧量
	BiologicalOxygenDemand float64 `json:"BiologicalOxygenDemand"`
	// 化学需氧量
	ChemicalOxygenDemand float64 `json:"ChemicalOxygenDemand"`
	//细菌总数
	BacteriaCount int `json:"BacteriaCount"`
	// 大肠杆菌数
	StaphylococcusCount int `json:"StaphylococcusCount"`
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

// //示例json
// var sampleJson = `{
// 	"NodeId":"1",
// 	"CheckValue":"123456",
// 	"GasConcentration":0.1,
// 	"Temperature":0.2,
// 	"PH":0.3,
// 	"Density":0.4,
// 	"Conductivity":0.5,
// 	"OxygenConcentration":0.6,
// 	"MetalConcentration":0.7,
// 	"SolidsConcentration":0.8,
// 	"FloatingSolidsConcentration":0.9,
// 	"TotalNitrogen":1.0,
// 	"TotalPhosphorus":1.1,
// 	"TotalOrganicCarbon":1.2,
// 	"BiologicalOxygenDemand":1.3,
// 	"ChemicalOxygenDemand":1.4,
// 	"BacteriaCount":1,
// 	"StaphylococcusCount":2
// }`

//校验节点提交信息
func (submitInfo SubmitInfo) Check() bool {
	// 用于返回结果
	var result bool = true
	//submitInfo的checkValue的第一位是否为1且GasConcentration为0
	if submitInfo.CheckValue[0] == '1' && submitInfo.GasConcentration == 0 {
		result = false
	}
	//submitInfo的checkValue的第二位是否为1且Temperature为0
	if submitInfo.CheckValue[1] == '1' && submitInfo.Temperature == 0 {
		result = false
	}
	//submitInfo的checkValue的第三位是否为1且PH为0
	if submitInfo.CheckValue[2] == '1' && submitInfo.PH == 0 {
		result = false
	}
	//submitInfo的checkValue的第四位是否为1且Density为0
	if submitInfo.CheckValue[3] == '1' && submitInfo.Density == 0 {
		result = false
	}
	//submitInfo的checkValue的第五位是否为1且Conductivity为0
	if submitInfo.CheckValue[4] == '1' && submitInfo.Conductivity == 0 {
		result = false
	}
	//submitInfo的checkValue的第六位是否为1且OxygenConcentration为0
	if submitInfo.CheckValue[5] == '1' && submitInfo.OxygenConcentration == 0 {
		result = false
	}
	//submitInfo的checkValue的第七位是否为1且MetalConcentration为0
	if submitInfo.CheckValue[6] == '1' && submitInfo.MetalConcentration == 0 {
		result = false
	}
	//submitInfo的checkValue的第八位是否为1且SolidsConcentration为0
	if submitInfo.CheckValue[7] == '1' && submitInfo.SolidsConcentration == 0 {
		result = false
	}
	//submitInfo的checkValue的第九位是否为1且FloatingSolidsConcentration为0
	if submitInfo.CheckValue[8] == '1' && submitInfo.FloatingSolidsConcentration == 0 {
		result = false
	}
	//submitInfo的checkValue的第十位是否为1且TotalNitrogen为0
	if submitInfo.CheckValue[9] == '1' && submitInfo.TotalNitrogen == 0 {
		result = false
	}
	//submitInfo的checkValue的第十一位是否为1且TotalPhosphorus为0
	if submitInfo.CheckValue[10] == '1' && submitInfo.TotalPhosphorus == 0 {
		result = false
	}
	//submitInfo的checkValue的第十二位是否为1且TotalOrganicCarbon为0
	if submitInfo.CheckValue[11] == '1' && submitInfo.TotalOrganicCarbon == 0 {
		result = false
	}
	//submitInfo的checkValue的第十三位是否为1且BiologicalOxygenDemand为0
	if submitInfo.CheckValue[12] == '1' && submitInfo.BiologicalOxygenDemand == 0 {
		result = false
	}
	//submitInfo的checkValue的第十四位是否为1且ChemicalOxygenDemand为0
	if submitInfo.CheckValue[13] == '1' && submitInfo.ChemicalOxygenDemand == 0 {
		result = false
	}
	//submitInfo的checkValue的第十五位是否为1且BacteriaCount为0
	if submitInfo.CheckValue[14] == '1' && submitInfo.BacteriaCount == 0 {
		result = false
	}
	//submitInfo的checkValue的第十六位是否为1且StaphylococcusCount为0
	if submitInfo.CheckValue[15] == '1' && submitInfo.StaphylococcusCount == 0 {
		result = false
	}
	return result
}

// 打印输出
func (submitInfo SubmitInfo) Print() {
	fmt.Printf("SubmitInfo: %+v\n", submitInfo)
}

// 数据写入到时序数据库
func (submitInfo SubmitInfo) WriteToDB() {
	Write(&submitInfo)
}
