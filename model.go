package IOValue

const (
	QUALITY_BAD           uint32 = 0x00
	QUALITY_UNCERTAIN     uint32 = 0x01
	QUALITY_NOT_CONNECTED uint32 = 0x02
	QUALITY_NOT_RESPONSE  uint32 = 0x03
	QUALITY_GOOD          uint32 = 0xc0
)

//采集记录的数据结构
type GoRecord struct {
	Id        string  `json:"id"`
	Status    uint32  `json:"status"`
	Val       IOValue `json:"val"`
	Timestamp int64   `json:"timestamp"`
}

//反写操作的数据结构
type GoWriteRecord struct {
	TagId    string `json:"id"`
	TagValue string `json:"val"`
}

//反写操作结果的数据结构
type GoWriteResult struct {
	TagId    string `json:"id"`
	TagValue string `json:"val"`
	Status   bool   `json:"status"`
}

//func (r *GoRecord) UnmarshalJSON(data []byte) error {
//	decoder := json.NewDecoder(strings.NewReader(utils.BytesToString(data)))
//	// UseNumber causes the Decoder to unmarshal a number into an interface{} as a Number instead of as a float64.
//	decoder.UseNumber()
//	var record interface{}
//	if err := decoder.Decode(&record); err != nil {
//		fmt.Println("error:", err)
//		return err
//	}
//	m := record.(map[string]interface{})
//
//	var idOK, valOK bool
//	r.Id, idOK = m["id"].(string)
//	if !idOK {
//		return fmt.Errorf("record id %v is not string", m["id"])
//	}
//
//	var status uint32
//	iStatus, statusOk := m["status"].(json.Number)
//	if !statusOk {
//		//return fmt.Errorf("record id %s status %v is not Number", m["id"], m["status"])
//		var temp string
//		temp, statusOk = m["status"].(string)
//		if !statusOk {
//			return fmt.Errorf("record id %s status %v is not Number or IOString", m["id"], m["status"])
//		}
//		if strings.EqualFold(temp, "Good") {
//			status = QUALITY_GOOD
//		} else {
//			status = QUALITY_BAD
//		}
//	} else {
//		tmp, err := iStatus.Int64()
//		if err != nil {
//			return err
//		}
//		status = uint32(tmp)
//	}
//
//	iTimeStamp, tsOk := m["timestamp"].(json.Number)
//	if !tsOk {
//		return fmt.Errorf("record id %s timestamp %v is not Number", m["id"], m["timestamp"])
//	}
//	timeStamp, err := iTimeStamp.Int64()
//	if err != nil {
//		return err
//	}
//
//	r.Val, valOK = m["val"].(string)
//	if !valOK {
//		return fmt.Errorf("record id %s, value %v is not string", m["id"], m["val"])
//	}
//
//	r.Status = status
//	r.Timestamp = timeStamp
//
//	return nil
//}
//
////自定义的序列化，上传的modbus等驱动的上传临时结构体
//func (r *GoRecord) MarshalJSON() ([]byte, error) {
//	status := "Good"
//	if r.Status != QUALITY_GOOD {
//		status = "Error"
//	}
//	control := struct {
//		Id        string `json:"id"`
//		Status    string `json:"status"`
//		Val       string `json:"val"`
//		Timestamp int64  `json:"timestamp"`
//	}{
//		Id:        r.Id,
//		Status:    status,
//		Val:       r.Val.IOString(),
//		Timestamp: r.Timestamp,
//	}
//	return json.Marshal(control)
//}

//报警记录
type AlarmHistory struct {
	EventName     string      //事件名称
	AlarmType     string      //是报警触发还是报警解除 //AlarmEvent;alarmRecover
	TagId         string      //tag点
	CurrentStatus uint32      //发生报警是的质量
	CurrentValue  interface{} //发生报警时的值
	Content       string      //报警描述
	OcuurTime     string      //发生时的时间
}

//命名管道查询的返回结果
type GoPipeRecord struct {
	DeviceCode string     `json:"deviceCode"`
	Val        []GoRecord `json:"val"`
}

//GoTagExtInfo 点的非通信信息
type GoTagExtInfo struct {
	Id         string `json:"id"`
	DeviceCode string `json:"deviceCode"`
	TagCode    string `json:"tagCode"`
	ObjectType string `json:"objectType"`
}

//GoDeviceStatus 设备的上下线通知
type GoDeviceStatus struct {
	DeviceCode string //设备的标识
	Status     bool   //是否在线
}
