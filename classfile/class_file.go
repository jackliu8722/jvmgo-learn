package classfile

type ClassFile struct {
    minorVersion uint16
    majorVersion uint16
    constantPool *ConstantPool
    accessFlags uint16
    thisClass uint16
    superClass uint16
    interfaces []uint16
    fields []*MemberInfo
    methods []*MemberInfo
    AttributeTable
}

func (self *ClassFile) read(reader *ClassReader) {
    self.readAndCheckMagic(reader)
    self.readVersions(reader)
    self.readConstantPool(reader)
    self.accessFlags = reader.readUint16()
    self.thisClass = reader.readUint16()
    self.superClass = reader.readUint16()
    self.interfaces = reader.readUint16s()
    self.fields = readMembers(reader,self.constantPool)
    self.methods = readMembers(reader,self.constantPool)
    self.attributes = readAttributes(reader,self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
    magic := reader.readUint32()
    if magic != 0xCAFEBABE {
        panic("Bad magic!")
    }
}

func (self *ClassFile) readVersions(reader *ClassReader) {
    self.minorVersion = reader.readUint16()
    self.majorVersion = reader.readUint16()

    //todo check versions
}

func (self *ClassFile) readConstantPool(reader *ClassReader){
    self.constantPool = &ConstantPool{cf : self}
    self.constantPool.read(reader)
}

func (self *ClassFile) AccessFlags() uint16 {
    return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo{
    return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo{
    return self.methods
}

func (self *ClassFile) ClassName() string {
    return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SupperClassName() string {
    if self.superClass != 0 {
        return self.constantPool.getClassName(self.superClass)
    }
    return ""
}

func (self *ClassFile) InterfaceClass() []string {
    interfaceNames := make([]string,len(self.interfaces))
    for i,cpIndex := range self.interfaces {
        interfaceNames[i] = self.constantPool.getClassName(cpIndex)
    }
    return interfaceNames
}
