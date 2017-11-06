package encoder;

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/proto"
	"strconv"
	"errors"
)

type Field struct {
	Number int32
	Val interface{}
}

type Encoder struct {
	types map[string]*typeInfo
}

type typeInfo struct {
	fields []*fieldInfo
}

type fieldInfo struct {
	name string
	number int32
	typeId  descriptor.FieldDescriptorProto_Type
}

func (e *Encoder) findFieldByNumber(tp *typeInfo, number int32) *fieldInfo {
	for _, tf := range tp.fields {
		if tf.number == number {
			return tf
		}
	}
	return nil
}


func (e *Encoder) encodeField(typeId descriptor.FieldDescriptorProto_Type, val interface{}) ([]byte, error){
	switch typeId {
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		strVal, ok := val.(string)
		if !ok {
			return nil, fmt.Errorf("val should be a string, got: %T", val)

		}
		valInt32, err := strconv.ParseInt(strVal, 10, 32)
		if err != nil {
			return nil, err
		}
		return proto.EncodeVarint(uint64(valInt32)), nil
	default:
		return nil, errors.New("not implemented")
	}
}
func (e *Encoder) Encode(typeName string, fields []*Field) ([]byte, error) {
	tp, ok := e.types[typeName]
	if !ok {
		return nil, fmt.Errorf("no such type: %v", typeName)
	}
	for _, f := range fields {
		tf := e.findFieldByNumber(tp, f.Number)
		if tf == nil {
			return nil, fmt.Errorf("no field with number: %v", f.Number)
		}

		wireType, err := wireByType(tf.typeId)
		if err != nil {
			return nil, err
		}

		prefix := byte(tf.number) << 3 | wireType //TODO: big field number

		bytes, err := e.encodeField(tf.typeId, f.Val)
		if err != nil {
			return nil, fmt.Errorf("encode field %v error: %v", tf.name, err)
		}
		return append([]byte{prefix}, bytes...), nil
	}

	return nil, fmt.Errorf("not implemented")

}

func wireByType(typeId descriptor.FieldDescriptorProto_Type) (uint8, error) {
	switch typeId {
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return 0, nil
	default:
		return 0, fmt.Errorf("wire type for %v does not implemented", typeId)
	}

}