package influxdb

import (
	"context"
	"fmt"
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
		AddField("SC", submit.SC).
		AddField("FSC", submit.FSC).
		AddField("TN", submit.TN).
		AddField("TP", submit.TP).
		AddField("TOC", submit.TOC).
		AddField("BOD", submit.BOD).
		AddField("COD", submit.COD).
		AddField("BC", submit.BC).
		AddField("SLC", submit.SLC).
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

// 查询从XXXX到XXXX期间内多个节点的数据(只获取最新数据)	 时间格式：2021-05-28T23:30:00Z
func Query(nodeid []string, startTime, endTime string) ([]SubmitInfo, error) {
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
			|> range(start: ` + startTime + `, stop: ` + endTime + `) 
			|> filter(fn: (r) => r["_measurement"] == "data")
			|> filter(fn: (r) => r["ID"] == "` + id + `")
			|> last()
			|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
		//获取查询表结果
		result, err := queryAPI.Query(context.Background(), query)
		if err != nil {
			fmt.Println("query error:", err)
			return nil, err
		}
		for result.Next() {
			record := result.Record()
			// 转换为结构体实例
			BC, _ := strconv.ParseInt(fmt.Sprint(record.ValueByKey("BC")), 10, 64)
			SLC, _ := strconv.ParseInt(fmt.Sprint(record.ValueByKey("SLC")), 10, 64)
			submitInfo := &SubmitInfo{
				NodeId:              id,
				GasConcentration:    record.ValueByKey("GasConcentration").(float64),
				Temperature:         record.ValueByKey("Temperature").(float64),
				PH:                  record.ValueByKey("PH").(float64),
				Density:             record.ValueByKey("Density").(float64),
				Conductivity:        record.ValueByKey("Conductivity").(float64),
				OxygenConcentration: record.ValueByKey("OxygenConcentration").(float64),
				MetalConcentration:  record.ValueByKey("MetalConcentration").(float64),
				SC:                  record.ValueByKey("SC").(float64),
				FSC:                 record.ValueByKey("FSC").(float64),
				TN:                  record.ValueByKey("TN").(float64),
				TP:                  record.ValueByKey("TP").(float64),
				TOC:                 record.ValueByKey("TOC").(float64),
				BOD:                 record.ValueByKey("BOD").(float64),
				COD:                 record.ValueByKey("COD").(float64),
				BC:                  BC,
				SLC:                 SLC,
			}
			suminfos = append(suminfos, *submitInfo)
		}
	}
	return suminfos, nil
}

func QueryAll(nodeid []string, startTime, endTime string) ([]SubmitInfoWithTime, error) {
	// 创建客户端
	client := influxdb2.NewClient("http://127.0.0.1:8086", "3TRkGcboToJkMfoPUwvuVvC0Rn1Tstzq5beZAjyv-MscinJA43c0ZIdW70eapXCB5MRTlwQ92VkDMs-Qs6yfDw==")
	//获取非阻塞写入客户端
	queryAPI := client.QueryAPI("stvsl-jc")
	defer client.Close()
	// 存储结果
	var suminfos []SubmitInfoWithTime
	// 查询每个节点的数据
	query := `from(bucket: "stvsljc")
			|> range(start: ` + startTime + `, stop: ` + endTime + `) 
			|> filter(fn: (r) => r["_measurement"] == "data")
			|> filter(fn: (r) => r["ID"] == "` + nodeid[0]

	for _, id := range nodeid[1:] {
		query += `" or r["ID"] == "` + id
	}
	query += `")`
	query += `|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
			  |> yield(name: "mean")`
	//获取查询表结果
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		fmt.Println("query error:", err)
		return nil, err
	}
	// 打印
	for result.Next() {
		record := result.Record()
		// 转换为结构体实例
		BC, _ := strconv.ParseInt(fmt.Sprint(record.ValueByKey("BC")), 10, 64)
		SLC, _ := strconv.ParseInt(fmt.Sprint(record.ValueByKey("SLC")), 10, 64)
		submitInfo := &SubmitInfoWithTime{
			NodeId:              record.ValueByKey("ID").(string),
			GasConcentration:    record.ValueByKey("GasConcentration").(float64),
			Temperature:         record.ValueByKey("Temperature").(float64),
			PH:                  record.ValueByKey("PH").(float64),
			Density:             record.ValueByKey("Density").(float64),
			Conductivity:        record.ValueByKey("Conductivity").(float64),
			OxygenConcentration: record.ValueByKey("OxygenConcentration").(float64),
			MetalConcentration:  record.ValueByKey("MetalConcentration").(float64),
			SC:                  record.ValueByKey("SC").(float64),
			FSC:                 record.ValueByKey("FSC").(float64),
			TN:                  record.ValueByKey("TN").(float64),
			TP:                  record.ValueByKey("TP").(float64),
			TOC:                 record.ValueByKey("TOC").(float64),
			BOD:                 record.ValueByKey("BOD").(float64),
			COD:                 record.ValueByKey("COD").(float64),
			BC:                  BC,
			SLC:                 SLC,
			Time:                fmt.Sprintf("%v", record.ValueByKey("_time")),
		}
		suminfos = append(suminfos, *submitInfo)
		// 跳过一次
		result.Next()
	}
	return suminfos, nil
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
