package com

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
)

//说明：每个寄存器为2字节数据，每次查询设定查询寄存器的数量和起始地址。寄存器的起始地址为0x0000，代表寄存器0x40001，每次查询最多查询100个寄存器。
//如从机地址为01，查询40001寄存器，主机发送：
//byte1		从机地址			01
//byte2		功能码			03
//byte3  	起始地址高字节		00
//byte4		起始地址低字节		00
//byte5		寄存器数量高字节	00
//byte6		寄存器数量低字节	01
//byte7		CRC校验低字节		CRCL
//byte8		CRC校验高字节		CRCH
//查询40001～40050寄存器：
//主机发送：
//byte1	从机地址			01
//byte2	功能码			03
//byte3	起始地址高字节		00
//byte4	起始地址低字节		00
//byte5	寄存器数量高字节	00
//byte6	寄存器数量低字节	32
//byte7	CRC校验低字节		CRCL
//byte8	CRC校验高字节		CRCH
//从机回复：
//byte1	从机地址			01
//byte2	功能码				03
//byte3	应答字节数			2*n
//byte4	第一个寄存器高字节	00
//byte5	第一个寄存器低字节	（0～9）
//……
//byte2n+2第n个寄存器高字节	00
//byte2n+3第n个寄存器低字节	（0～9）
//byte2n+4	CRC校验低字节		CRCL
//byte2n+5	CRC校验高字节		CRCH

// Com采集传输对象
type ComCollection struct {
	// com传输io
	reader io.ReadWriteCloser
	//寄存器数量
	registerNum int
	//发送请求指令
	requestCommand []byte
	// 定义chan 进行数据传输
	comDataChan chan string
	// 数据分发接口
	dataTransfer DataTransfer
}

func New(iorwc io.ReadWriteCloser, requestCommand []byte) *ComCollection {
	//获取后两位标示寄存器长度字节
	registerBuffer := requestCommand[len(requestCommand)-2:]
	//字节补齐
	registerBuffer = append([]byte{0x00, 0x00}, registerBuffer...)
	registerNum := int(binary.BigEndian.Uint32(registerBuffer))
	//计算数量
	//fmt.Println(resultBuffer)
	fmt.Printf("registerNum:%d \n", registerNum)
	com := &ComCollection{
		reader:         iorwc,
		registerNum:    registerNum,
		requestCommand: requestCommand,
		comDataChan:    make(chan string),
	}
	return com
}

//发送指令到采集服务消息
func (c *ComCollection) send() {
	command := CRC(c.requestCommand)
	// 写入串口命令
	beego.BeeLogger.Info("写入的指令 %x", command)
	resultBuffer := new(bytes.Buffer)
	for {
		_, err := c.reader.Write([]byte(command))

		if err != nil {
			beego.BeeLogger.Error("send command error", err)
			os.Exit(1)
		} else {
			c.receive(resultBuffer)
			resultBuffer.Reset()
		}
	}
}

// 接收com口消息
func (c *ComCollection) receive(resultBuffer *bytes.Buffer) {
	for {
		//安全期间每次取一个
		buffer := make([]byte, 1)
		num, err := c.reader.Read(buffer)
		if err != nil {
			beego.BeeLogger.Error("read data from com error", err)
			os.Exit(1)
		}
		if num > 0 {
			resultBuffer.Write(buffer)
		} else {
			break
		}
	}
	beego.BeeLogger.Info("---------------------")
	if resultBuffer.Len() > 0 {
		beego.BeeLogger.Info("++++++++++++++++++++")
		//发送数据到通道
		c.comDataChan <- resultBuffer.String()
	}

}

//消息转移
func (c *ComCollection) transfer() {
	for {
		select {
		case v := <-c.comDataChan:
			{
				fmt.Println(string(v))
				//1、数据校验成功后进行第二步操作
				//2、数据转换成二进制分发
				if c.dataTransfer != nil {
					c.dataTransfer.Transfer(make([]byte, 128))
				}
			}
		}
	}
}

func (c *ComCollection) Main() {
	//发送采集指令并采集数据
	go c.send()

	//发送消息
	go c.transfer()
}

// 关闭资源
func (c *ComCollection) Exit() {
	if c.reader != nil {
		c.reader.Close()
	}
	if c.comDataChan != nil {
		close(c.comDataChan)
	}
}
