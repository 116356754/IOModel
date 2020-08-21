/**
 * @Author: wangky
 * @Description: 
 * @Version: 1.0.0
 * @Date: 2020-08-20 16:20
 */
package IOModel

//质量戳暂时分好与坏
const (
	QualityBad    uint8 = 0x00
	QualityGood   uint8 = 0x01
	QualityUnkown uint8 = 0x02
)

//采集记录的数据结构
type GoField struct {
	Id        string  `json:"id"`        //每个点的唯一id
	Status    uint8   `json:"status"`    //每个点的质量戳
	Val       IOValue `json:"val"`       //每个点的数据类型可能不一样
	Timestamp int64   `json:"timestamp"` //每个点得到的时间戳
}

//采集端的数据模型，由设备和下属点构成
type GoMetric struct {
	DeviceId  string             //设备的id
	Status    uint8              //表示设备的状态
	Fields    map[string]GoField //设备下面的记录集
	Timestamp int64              //表示最后一次更新的时间
}

//反写操作的数据结构
type GoWriteMetric struct {
	DeviceId string             //需要写入的设备的id
	Fields   map[string]IOValue //需要写入的点id和对应的值
}
