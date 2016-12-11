package classfile

type MemberInfo struct {
    cp *ConstantPool
    accessFlags uint16
    nameIndex uint16
    descriptorIndex uint16
    AttributeTable
}

func readMembers(reader *ClassReader, cp *ConstantPool) []*MemberInfo{
    memberCount := reader.readUint16()
    members := make([]*MemberInfo,memberCount)

    for i := range members {
        members[i] = &MemberInfo{cp:cp}
        members[i].read(reader)
    }
    return members
}

func (self *MemberInfo) read(reader *ClassReader) {
    self.accessFlags = reader.readUint16()
    self.nameIndex = reader.readUint16()
    self.descriptorIndex = reader.readUint16()
    self.attributes = readAttributes(reader,self.cp)
}

func (self *MemberInfo) Name() string {
    return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
    return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) Signature() string {
    signatureAttr := self.SignatureAttribute()
    if signatureAttr != nil {
        return signatureAttr.Signature()
    }
    return ""
}
