package influxdb

import (
	"context"
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// 通讯channel

func Write(submit *SubmitInfo) error {
	// err := ConnectTest
	// if err != nil {
	// 	return err()
	// }
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
	query := `from(bucket: "stvsljc")
	|> range(start: -1h) 
	|> filter(fn: (r) => r["_measurement"] == "data")
	|> filter(fn: (r) => r["ID"] == "CX0000001")
	|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	//获取查询表结果
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return "查询失败", err
	}
	//总是在最后关闭客户端
	defer client.Close()
	// 打印result
	for result.Next() {
		record := result.Record()
		// fmt.Printf("%v: %v value: %v\n", record.Time(), record.ValueByKey("StaphylococcusCount"), record.Value())
		fmt.Printf("%v: BacteriaCount %v ", record.Time(), record.ValueByKey("BacteriaCount"))
		fmt.Printf(" BiologicalOxygenDemand:%v", record.ValueByKey("BiologicalOxygenDemand"))
		fmt.Printf(" ChemicalOxygenDemand:%v", record.ValueByKey("ChemicalOxygenDemand"))
		fmt.Printf(" Conductivity:%v", record.ValueByKey("Conductivity"))
		fmt.Printf(" Density:%v", record.ValueByKey("Density"))
		fmt.Printf(" FloatingSolidsConcentration:%v", record.ValueByKey("FloatingSolidsConcentration"))
		fmt.Printf(" GasConcentration:%v", record.ValueByKey("GasConcentration"))
		fmt.Printf(" MetalConcentration:%v", record.ValueByKey("MetalConcentration"))
		fmt.Printf(" O2Concentration:%v", record.ValueByKey("OxygenConcentration"))
		fmt.Printf(" PH:%v", record.ValueByKey("PH"))
		fmt.Printf(" SolidsConcentration:%v", record.ValueByKey("SolidsConcentration"))
		fmt.Printf(" TotalNitrogen:%v", record.ValueByKey("TotalNitrogen"))
		fmt.Printf(" TotalOrganicCarbon:%v", record.ValueByKey("TotalOrganicCarbon"))
		fmt.Printf(" TotalPhosphorus:%v", record.ValueByKey("TotalPhosphorus"))
		fmt.Printf("\n")
		fmt.Println()
	}
	return "查询成功", nil
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
