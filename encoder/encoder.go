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


func (e *Encoder) encodeField(typeId descriptor.FieldDescriptorProto_Type, repeated bool, val interface{}) ([]byte, error) {
	if repeated {
		stringSliceVal, ok := val.([]string)
		if !ok {
			return nil, fmt.Errorf("val should be a slice of strings, got: %T", val)
		}
		res := make([]byte, 0)
		for _, v := range stringSliceVal {
			fieldData, err := e.encodeField(typeId, false, v)
			if err != nil {
				return nil, err
			}
			res = append(res, fieldData...)
		}
		res = append(proto.EncodeVarint(uint64(len(res))), res...)

		return res, nil
	}
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

	res := make([]byte, 0)

	for _, f := range fields {
		tf := e.findFieldByNumber(tp, f.Number)
		if tf == nil {
			return nil, fmt.Errorf("no field with number: %v", f.Number)
		}

		wireType, err := wireByType(tf.typeId, tf.repeated)
		if err != nil {
			return nil, err
		}


		bytes, err := e.encodeField(tf.typeId, tf.repeated, f.Val)
		if err != nil {
			return nil, fmt.Errorf("encode field %v error: %v", tf.name, err)
		}

		prefix := proto.EncodeVarint(uint64(tf.number << 3 | int32(wireType)))

		res =  append(res, prefix...)
		res =  append(res, bytes...)
	}

	return res, nil
}

func wireByType(typeId descriptor.FieldDescriptorProto_Type, repeated bool) (uint8, error) {
	if repeated {
		return 2, nil
	}
	switch typeId {
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return 0, nil
	default:
		return 0, fmt.Errorf("wire type for %v does not implemented", typeId)
	}

}