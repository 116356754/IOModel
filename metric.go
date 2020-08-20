/**
 * @Author: wangky
 * @Description: 
 * @Version: 1.0.0
 * @Date: 2020-08-20 16:20
 */
package IOModel

//采集记录的数据结构
type GoField struct {
	Id        string  `json:"id"`
	Status    uint8  `json:"status"`
	Val       IOValue `json:"val"`
	Timestamp int64   `json:"timestamp"`
}

//采集端的数据模型，由设备和下属点构成
type GoMetric struct {
	DeviceId  string             //设备的id
	Status    bool               //表示设备的状态
	Fields    map[string]GoField //设备下面的记录集
	Timestamp int64              //表示最后一次更新的时间
}

//反写操作的数据结构
type GoWriteRecord struct {
	TagId    string `json:"id"`
	TagValue IOValue `json:"val"`
}