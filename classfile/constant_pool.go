package classfile

import (
    "fmt"
)

type ConstantPool struct {
    cf *ClassFile
    cpInfos []ConstantInfo
}

func (self *ConstantPool) read(reader *ClassReader) {
    cpCount := int(reader.readUint16())
    self.cpInfos = make([]ConstantInfo,cpCount)

    for i:=1; i < cpCount ; i++ {
        self.cpInfos[i] = readConstantInfo(reader,self)

        switch self.cpInfos[i].(type) {
        case *ConstantLongInfo,*ConstantDoubleInfo:
            i++
        }
    }
}

func (self *ConstantPool) Infos() []ConstantInfo {
    return self.cpInfos
}

func (self *ConstantPool) getConstantInfo(index uint16) ConstantInfo{
    cpInfo := self.cpInfos[index]
    if cpInfo == nil {
        panic(fmt.Errorf("Bad constant pool index: %v!",index))
    }

    return cpInfo
}

func (self *ConstantPool) getNameAndType(index uint16) (name,_type string){
    ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
    name = self.getUtf8(ntInfo.nameIndex)
    _type = self.getUtf8(ntInfo.descriptorIndex)
    return
}

func (self *ConstantPool) getClassName(index uint16) string {
    classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
    return self.getUtf8(classInfo.nameIndex)
}

func (self *ConstantPool) getUtf8(index uint16) string {
    utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
    return utf8Info.str
}
