package encoder

import (
	"testing"
	"bytes"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func TestEncode(t *testing.T) {
	e := Encoder{
		types: map[string]*typeInfo{
			".TestMessage": {
				fields: []*fieldInfo{
					{
						name: "first",
						number: 1,
						typeId: descriptor.FieldDescriptorProto_TYPE_INT32,
					},
				},
			},
		},
	}
	data, err := e.Encode(".TestMessage", []*Field{
		{
			Number: 1,
			Val: "150",
		},
	})
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}
	expected := []byte{0x08, 0x96, 0x01}

	if bytes.Compare(expected, data) != 0 {
		t.Fatalf("bytes does not equal, expected: %+v, got: %+v", expected, data)
	}
}
