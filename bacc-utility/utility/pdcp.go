package utility

import (
	"encoding/binary"
	"fmt"
	"io"
)

// pdxp package struct
type Package struct {
	VER    byte    //版本
	MID    [2]byte //任务标识
	SID    [4]byte //信源地址
	DID    [4]byte // 新宿地址
	BID    [4]byte //数据标识
	NO     uint32  //包序号
	FLAG   byte    //数据处理标识
	RETAIN [4]byte //保留字
	DATE   [2]byte //发送日期
	TIME   [4]byte //发送时标
	L      uint16  // 数据域长度
	DATA   []byte  // 数据域
}

func (p *Package) Pack(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, &p.VER)
	err = binary.Write(writer, binary.BigEndian, &p.MID)
	err = binary.Write(writer, binary.BigEndian, &p.SID)
	err = binary.Write(writer, binary.BigEndian, &p.DID)
	err = binary.Write(writer, binary.BigEndian, &p.BID)
	err = binary.Write(writer, binary.BigEndian, &p.NO)
	err = binary.Write(writer, binary.BigEndian, &p.FLAG)
	err = binary.Write(writer, binary.BigEndian, &p.RETAIN)
	err = binary.Write(writer, binary.BigEndian, &p.DATE)
	err = binary.Write(writer, binary.BigEndian, &p.TIME)
	err = binary.Write(writer, binary.BigEndian, &p.L)
	err = binary.Write(writer, binary.BigEndian, &p.DATA)
	return err
}
func (p *Package) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.VER)
	err = binary.Read(reader, binary.BigEndian, &p.MID)
	err = binary.Read(reader, binary.BigEndian, &p.SID)
	err = binary.Read(reader, binary.BigEndian, &p.DID)
	err = binary.Read(reader, binary.BigEndian, &p.BID)
	err = binary.Read(reader, binary.BigEndian, &p.NO)
	err = binary.Read(reader, binary.BigEndian, &p.FLAG)
	err = binary.Read(reader, binary.BigEndian, &p.RETAIN)
	err = binary.Read(reader, binary.BigEndian, &p.DATE)
	err = binary.Read(reader, binary.BigEndian, &p.TIME)
	err = binary.Read(reader, binary.BigEndian, &p.L)
	p.DATA = make([]byte, p.L)
	err = binary.Read(reader, binary.BigEndian, &p.DATA)
	return err
}

func (p *Package) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s",
		p.VER,
		p.MID,
		p.SID,
		p.DID,
		p.BID,
		p.NO,
		p.FLAG,
		p.RETAIN,
		p.DATE,
		p.TIME,
		p.L,
		p.DATA,
	)
}
