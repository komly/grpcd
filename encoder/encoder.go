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
	typeName string
	repeated bool
}

func (e *Encoder) findFieldByNumber(tp *typeInfo, number int32) *fieldInfo {
	for _, tf := range tp.fields {
		if tf.number == number {
			return tf
		}
	}
	return nil
}


func (e *Encoder) encodeField(typeId descriptor.FieldDescriptorProto_Type, typeName string, val interface{}) ([]byte, error) {
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
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		fieldsVal, ok := val.([]*Field)
		if !ok {
			return nil, fmt.Errorf("val should be a `[]*Field`, got: %T", val)
		}
		data, err := e.Encode(typeName, fieldsVal)
		if err != nil {
			return nil, err
		}
		res := make([]byte, 0)
		res = append(res, proto.EncodeVarint(uint64(len(data)))...)
		res = append(res, data...) // ???
		return res, nil
	default:
		return nil, errors.New("not implemented")
	}
}

func (e *Encoder) encodePackedRepeated(f *fieldInfo, val interface{}) ([]byte, error) {

	stringSliceVal, ok := val.([]interface{})
	if !ok {
		return nil, fmt.Errorf("val should be a slice of interface{}, got: %T", val)
	}

	if len(stringSliceVal) == 0 {
		return []byte{}, nil
	}

	res := proto.EncodeVarint(uint64(f.number << 3 | int32(2)))

	lres := make([]byte, 0)
	for _, v := range stringSliceVal {
		fieldData, err := e.encodeField(f.typeId, f.typeName, v)
		if err != nil {
			return nil, err
		}
		lres = append(lres, fieldData...)
	}
	res = append(res, proto.EncodeVarint(uint64(len(lres)))...)
	return append(res, lres...), nil
}

func (e *Encoder) encodeRepeated(f *fieldInfo, val interface{}) ([]byte, error) {
	stringSliceVal, ok := val.([]interface{})
	if !ok {
		return nil, fmt.Errorf("val should be a slice of interface{}, got: %T", val)
	}

	res := make([]byte, 0)
	for _, v := range stringSliceVal {
		wireType, err := wireByType(f.typeId)
		if err != nil {
			return nil, err
		}
		bytes, err := e.encodeField(f.typeId, f.typeName, v)
		if err != nil {
			return nil, err
		}
		res = append(res, proto.EncodeVarint(uint64(f.number << 3 | int32(wireType)))...)
		res = append(res, bytes...)
	}

	return res, nil
}


func (e *Encoder) encodeSingle(f *fieldInfo, val interface{}) ([]byte, error) {
	wireType, err := wireByType(f.typeId)
	if err != nil {
		return nil, err
	}

	bytes, err := e.encodeField(f.typeId, f.typeName, val)
	if err != nil {
		return nil, fmt.Errorf("encode field `%v` error: %v", f.name, err)
	}
	return append(proto.EncodeVarint(uint64(f.number << 3 | int32(wireType))), bytes...), nil
}

func (e *Encoder) Encode(typeName string, fields []*Field) ([]byte, error) {
	tp, ok := e.types[typeName]
	if !ok {
		return nil, fmt.Errorf("no such type: %v", typeName)
	}

	res := make([]byte, 0)

	for _, f := range fields {
		tf := e.findFieldByNumber(tp, f.Number)
		if tf == nil {
			return nil, fmt.Errorf("no field with number: %v", f.Number)
		}

		var (
			data []byte
			err error
		)

		if tf.repeated {
			if packedByType(tf.typeId) {
				data, err = e.encodePackedRepeated(tf, f.Val)
			} else {
				data, err = e.encodeRepeated(tf, f.Val)

			}
		} else {
			data, err = e.encodeSingle(tf, f.Val)
		}
		if err != nil {
			return nil, err
		}
		res = append(res, data...)
	}

	return res, nil
}

func packedByType(typeId descriptor.FieldDescriptorProto_Type) bool {
	switch typeId {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_BOOL,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_ENUM,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		return true
	default:
		return false
	}

}

func wireByType(typeId descriptor.FieldDescriptorProto_Type) (uint8, error) {
	switch typeId {
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return 0, nil
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		return 2, nil
	default:
		return 0, fmt.Errorf("wire type for %v does not implemented", typeId)
	}

}