package com

//采集数据转移。调用者需要实现此接口进行采集数据分发
type DataTransfer interface {
	Transfer(data []byte)
}
