package influxdb

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func Write() {
	// 创建客户端
	// 您可以从 UI 中的"API 令牌选项卡"生成 API 令牌
	client := influxdb2.NewClient("http://127.0.0.1:8086", "3TRkGcboToJkMfoPUwvuVvC0Rn1Tstzq5beZAjyv-MscinJA43c0ZIdW70eapXCB5MRTlwQ92VkDMs-Qs6yfDw==")
	//获取非阻塞写入客户端
	writeAPI := client.WriteAPI("stvsl-jc", "stvsljc")
	//写入线路协议
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	//刷新写入
	writeAPI.Flush()
	//总是在最后关闭客户端
	defer client.Close()

}

func Query() {
	//创建客户端
	//您可以从ui中的“api令牌标签”生成api令牌
	client := influxdb2.NewClient("http://127.0.0.1:8086", "3TRkGcboToJkMfoPUwvuVvC0Rn1Tstzq5beZAjyv-MscinJA43c0ZIdW70eapXCB5MRTlwQ92VkDMs-Qs6yfDw==")
	// Get query client
	queryAPI := client.QueryAPI("stvsl-jc")

	query := `from(bucket:"stvsljc")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`

	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	// Iterate over query response
	for result.Next() {
		//更改组密钥时请注意
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		//访问数据
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	//检查是否有错误
	if result.Err() != nil {
		fmt.Printf("query parsing error: \n" + result.Err().Error())
	}
	//总是在最后关闭客户端
	defer client.Close()
}
