package encoder

import (
	"testing"
	"bytes"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/komly/grpcd/encoder/fixtures"
	 "github.com/golang/protobuf/proto"
)

func TestEncodeSimple(t *testing.T) {
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


func TestEncodeSimple2(t *testing.T) {
	e := Encoder{
		types: map[string]*typeInfo{
			".TestMessage": {
				fields: []*fieldInfo{
					{
						name: "first",
						number: 1,
						typeId: descriptor.FieldDescriptorProto_TYPE_INT32,
					},
					{
						name: "second",
						number: 2,
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
		{
			Number: 2,
			Val: "150",
		},
	})
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}

	m := &fixtures.TestMessage{
		A: 150,
		B: 150,
	}

	pData, err := proto.Marshal(m)

	if bytes.Compare(pData, data) != 0 {
		t.Fatalf("bytes does not equal, expected: %+v, got: %+v", pData, data)
	}
}


func TestEncodeRepeatedField(t *testing.T) {
	e := Encoder{
		types: map[string]*typeInfo{
			".TestMessage": {
				fields: []*fieldInfo{
					{
						name: "first",
						number: 1,
						typeId: descriptor.FieldDescriptorProto_TYPE_INT32,
						repeated: true,
					},

				},
			},
		},
	}
	data, err := e.Encode(".TestMessage", []*Field{
		{
			Number: 1,
			Val: []interface{}{"1", "2", "3"},
		},

	})
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}

	m := &fixtures.TestMessageWithRepeated{
		A: []int32{1, 2, 3},
	}

	pData, err := proto.Marshal(m)

	if bytes.Compare(pData, data) != 0 {
		t.Fatalf("bytes does not equal, expected: %+v, got: %+v", pData, data)
	}
}


func TestEncodeRepeatedFieldEmpty(t *testing.T) {
	e := Encoder{
		types: map[string]*typeInfo{
			".TestMessage": {
				fields: []*fieldInfo{
					{
						name: "first",
						number: 1,
						typeId: descriptor.FieldDescriptorProto_TYPE_INT32,
						repeated: true,
					},

				},
			},
		},
	}
	data, err := e.Encode(".TestMessage", []*Field{
		{
			Number: 1,
			Val: []interface{}{},
		},

	})
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}

	m := &fixtures.TestMessageWithRepeated{
		A: []int32{},
	}

	pData, err := proto.Marshal(m)

	if bytes.Compare(pData, data) != 0 {
		t.Fatalf("bytes does not equal, expected: %+v, got: %+v", pData, data)
	}
}


func TestEncodeInnerStruct(t *testing.T) {
	e := Encoder{
		types: map[string]*typeInfo{
			".Inner": {
				fields: []*fieldInfo{
					{
						name: "first",
						number: 1,
						typeId: descriptor.FieldDescriptorProto_TYPE_INT32,
					},

				},
			},
			".TestMessage": {
				fields: []*fieldInfo{
					{
						name: "first",
						number: 1,
						typeId: descriptor.FieldDescriptorProto_TYPE_MESSAGE,
						typeName: ".Inner",
					},

				},
			},
		},
	}
	data, err := e.Encode(".TestMessage", []*Field{
		{
			Number: 1,
			Val: []*Field{
				{
					Number: 1,
					Val: "150",
				},
			},
		},

	})
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}

	m := &fixtures.TestMessageWithInner{
		Inner: &fixtures.InnerMessage{
			A: 150,
		},
	}

	pData, err := proto.Marshal(m)

	if bytes.Compare(pData, data) != 0 {
		t.Fatalf("bytes does not equal, expected: %+v, got: %+v", pData, data)
	}
}

func TestEncodeRepeatedInnerStruct(t *testing.T) {
	e := Encoder{
		types: map[string]*typeInfo{
			".Inner": {
				fields: []*fieldInfo{
					{
						name: "first",
						number: 1,
						typeId: descriptor.FieldDescriptorProto_TYPE_INT32,
					},

				},
			},
			".TestMessage": {
				fields: []*fieldInfo{
					{
						name: "first",
						number: 1,
						typeId: descriptor.FieldDescriptorProto_TYPE_MESSAGE,
						typeName: ".Inner",
						repeated: true,
					},

				},
			},
		},
	}
	var val []interface{}

	msg1 := []*Field{
		{
			Number: 1,
			Val:    "1",
		},
	}
	msg2 := []*Field{
		{
			Number: 1,
			Val:    "1",
		},
	}
	val = append(val, msg1, msg2)

	data, err := e.Encode(".TestMessage", []*Field{
		{
			Number: 1,
			Val:    val,
		},
	})
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}

	m := &fixtures.TestRepeatedInner{
		Inner: []*fixtures.InnerMessage{
			{
				A: 1,
			},
			{
				A: 1,
			},
		},
	}

	pData, err := proto.Marshal(m)

	if bytes.Compare(pData, data) != 0 {
		t.Fatalf("bytes does not equal, expected: %+v, got: %+v", pData, data)
	}
}
