**主要功能**
---
- [1] 将采集过来的数据，统一转换成实现了IOValue接口的数据类型

- [2] 目前支持的数据类型有NULL、Bool、Int、Float、String和Byte数组类型

- [3] 同时支持interface{}原生的数据类型转IOValue类型，也支持IOValue类型转原生数据类型

- [4] 在包中还定义了采集过来的数据转换的GoMetric结构体，还有上传协议往下写的GoWriteMetric结构体

**TODO**
---
在GoField结构体，可能要实现UnmarshalJSON和MarshalJSON函数，来实现IOValue字段的JSON反序列化和序列化功能