package influxdb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// 通讯channel
func Write(submit *SubmitInfo) error {
	// 创建客户端
	client := influxdb2.NewClient("http://127.0.0.1:8086", "3TRkGcboToJkMfoPUwvuVvC0Rn1Tstzq5beZAjyv-MscinJA43c0ZIdW70eapXCB5MRTlwQ92VkDMs-Qs6yfDw==")
	//获取非阻塞写入客户端
	writeAPI := client.WriteAPI("stvsl-jc", "stvsljc")
	// create point using fluent style
	p := influxdb2.NewPointWithMeasurement("data").
		// 将数据submit的数据逐个添加到point中
		AddTag("ID", submit.NodeId).
		AddField("GasConcentration", submit.GasConcentration).
		AddField("Temperature", submit.Temperature).
		AddField("PH", submit.PH).
		AddField("Density", submit.Density).
		AddField("Conductivity", submit.Conductivity).
		AddField("OxygenConcentration", submit.OxygenConcentration).
		AddField("MetalConcentration", submit.MetalConcentration).
		AddField("SolidsConcentration", submit.SolidsConcentration).
		AddField("FloatingSolidsConcentration", submit.FloatingSolidsConcentration).
		AddField("TotalNitrogen", submit.TotalNitrogen).
		AddField("TotalPhosphorus", submit.TotalPhosphorus).
		AddField("TotalOrganicCarbon", submit.TotalOrganicCarbon).
		AddField("BiologicalOxygenDemand", submit.BiologicalOxygenDemand).
		AddField("ChemicalOxygenDemand", submit.ChemicalOxygenDemand).
		AddField("BacteriaCount", submit.BacteriaCount).
		AddField("StaphylococcusCount", submit.StaphylococcusCount).
		SetTime(time.Now())
	// 异步写入点
	fmt.Println(p)
	writeAPI.WritePoint(p)
	//刷新写入
	fmt.Println("flush")
	writeAPI.Flush()
	//总是在最后关闭客户端
	defer client.Close()
	return nil
}

// 查询从XXXX到XXXX期间内多个节点的数据
func Query(nodeid []string, startTime, endTime string) (string, error) {
	// 创建客户端
	client := influxdb2.NewClient("http://127.0.0.1:8086", "3TRkGcboToJkMfoPUwvuVvC0Rn1Tstzq5beZAjyv-MscinJA43c0ZIdW70eapXCB5MRTlwQ92VkDMs-Qs6yfDw==")
	//获取非阻塞写入客户端
	queryAPI := client.QueryAPI("stvsl-jc")
	defer client.Close()
	// 存储结果
	var suminfos []SubmitInfo
	// 查询每个节点的数据
	for _, id := range nodeid {
		// 查询每个节点的数据
		query := `from(bucket: "stvsljc")
			|> range(start: -1h) 
			|> filter(fn: (r) => r["_measurement"] == "data")
			|> filter(fn: (r) => r["ID"] == "` + id + `")
			|> last()
			|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
		//获取查询表结果
		result, err := queryAPI.Query(context.Background(), query)
		if err != nil {
			return "查询失败", err
		}
		for result.Next() {
			record := result.Record()
			// 转换为结构体实例
			bacteriaCount, _ := strconv.ParseInt(fmt.Sprint(record.ValueByKey("BacteriaCount")), 10, 64)
			staphylococcusCount, _ := strconv.ParseInt(fmt.Sprint(record.ValueByKey("StaphylococcusCount")), 10, 64)
			submitInfo := &SubmitInfo{
				NodeId:                      id,
				GasConcentration:            record.ValueByKey("GasConcentration").(float64),
				Temperature:                 record.ValueByKey("Temperature").(float64),
				PH:                          record.ValueByKey("PH").(float64),
				Density:                     record.ValueByKey("Density").(float64),
				Conductivity:                record.ValueByKey("Conductivity").(float64),
				OxygenConcentration:         record.ValueByKey("OxygenConcentration").(float64),
				MetalConcentration:          record.ValueByKey("MetalConcentration").(float64),
				SolidsConcentration:         record.ValueByKey("SolidsConcentration").(float64),
				FloatingSolidsConcentration: record.ValueByKey("FloatingSolidsConcentration").(float64),
				TotalNitrogen:               record.ValueByKey("TotalNitrogen").(float64),
				TotalPhosphorus:             record.ValueByKey("TotalPhosphorus").(float64),
				TotalOrganicCarbon:          record.ValueByKey("TotalOrganicCarbon").(float64),
				BiologicalOxygenDemand:      record.ValueByKey("BiologicalOxygenDemand").(float64),
				ChemicalOxygenDemand:        record.ValueByKey("ChemicalOxygenDemand").(float64),
				BacteriaCount:               bacteriaCount,
				StaphylococcusCount:         staphylococcusCount,
			}
			suminfos = append(suminfos, *submitInfo)
		}
	}
	// 转换为json
	json, err := json.Marshal(suminfos)
	if err != nil {
		return "转换失败", err
	}
	fmt.Println(string(json))
	return string(json), nil
}

func Query2() error {
	//创建客户端
	client := influxdb2.NewClient("http://127.0.0.1:8086", "3TRkGcboToJkMfoPUwvuVvC0Rn1Tstzq5beZAjyv-MscinJA43c0ZIdW70eapXCB5MRTlwQ92VkDMs-Qs6yfDw==")
	// Get query client
	queryAPI := client.QueryAPI("stvsl-jc")

	query := `from(bucket:"stvsljc")
				|> range(start: -1h) 
				|> filter(fn: (r) => r._measurement == "data")`

	//获取查询表结果
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Panicln("query error:", err)
		return err
	}
	//检查是否有错误
	if result.Err() != nil {
		fmt.Printf("query parsing error: \n" + result.Err().Error())
	}
	//总是在最后关闭客户端
	defer client.Close()
	return nil
}

// 连接测试请求
func ConnectTest() error {
	//创建客户端
	client := influxdb2.NewClient("http://127.0.0.1:8086", "3TRkGcboToJkMfoPUwvuVvC0Rn1Tstzq5beZAjyv-MscinJA43c0ZIdW70eapXCB5MRTlwQ92VkDMs-Qs6yfDw==")
	// // 连接测试
	// _, err := client.Ping(client)
	// if err != nil {
	// 	log.Panicln("connect error:", err)
	// 	return err
	// }
	client.Close()
	return nil
}
