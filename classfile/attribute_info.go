package classfile

var (
	_attrDeprecated = &DeprecatedAttribute{}
	_attrSynthetic = &SyntheticAttribute{}
)

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader,cp *ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo,attributesCount)
	for i := range attributes{
		attributes[i] = readAttribute(reader,cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader,cp *ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrLen := reader.readUint32()
	attrName := cp.getUtf8(attrNameIndex)
	attrInfo := newAttributeInfo(attrName,cp)
	if attrInfo == nil {
		attrInfo = &UnparsedAttribute{
			name: attrName,
			length: attrLen,
		}
	}

	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string,cp *ConstantPool) AttributeInfo {
	switch attrName {
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{}
	case "Code":
		return &CodeAttribute{cp:cp}
	case "ConstantValue":
		return nil
	case "Deprecated":
		return _attrDeprecated
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{cp:cp}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTypeTableAttribute{}
	case "Signature":
		return &SignatureAttribute{cp:cp}
	case "SourceFile":
		return &SourceFileAttribute{cp:cp}
	case "Synthetic":
		return _attrSynthetic
	default:
		return nil
	}
}


