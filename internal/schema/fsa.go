// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
package schema

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type Fsa struct {
	// This field is used as the key.
	Label string `json:"label"`

	Latitude float32 `json:"latitude"`

	Longitude float32 `json:"longitude"`

	Message_id string `json:"message_id"`

	Created_at int64 `json:"created_at"`

	Updated_at int64 `json:"updated_at"`
	// ISO 8601 datetime. Format: YYYY-MM-DD HH:MM:SS z
	Timestamp string `json:"timestamp"`
}

const FsaAvroCRC64Fingerprint = "\xaf\x9e\xa3\x94\xf1Tm,"

func NewFsa() *Fsa {
	return &Fsa{}
}

func DeserializeFsa(r io.Reader) (*Fsa, error) {
	t := NewFsa()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeFsaFromSchema(r io.Reader, schema string) (*Fsa, error) {
	t := NewFsa()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeFsa(r *Fsa, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Label, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Latitude, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Longitude, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Message_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Created_at, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Updated_at, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Timestamp, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Fsa) Serialize(w io.Writer) error {
	return writeFsa(r, w)
}

func (r *Fsa) Schema() string {
	return "{\"fields\":[{\"doc\":\"This field is used as the key.\",\"name\":\"label\",\"type\":\"string\"},{\"default\":0,\"name\":\"latitude\",\"type\":\"float\"},{\"default\":0,\"name\":\"longitude\",\"type\":\"float\"},{\"default\":\"\",\"name\":\"message_id\",\"type\":\"string\"},{\"name\":\"created_at\",\"type\":\"long\"},{\"name\":\"updated_at\",\"type\":\"long\"},{\"default\":\"\",\"doc\":\"ISO 8601 datetime. Format: YYYY-MM-DD HH:MM:SS z\",\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"com.flipp.fadmin.Fsa\",\"type\":\"record\"}"
}

func (r *Fsa) SchemaName() string {
	return "com.flipp.fadmin.Fsa"
}

func (_ *Fsa) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Fsa) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Fsa) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Fsa) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Fsa) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Fsa) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Fsa) SetString(v string)   { panic("Unsupported operation") }
func (_ *Fsa) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Fsa) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Label}
	case 1:
		return &types.Float{Target: &r.Latitude}
	case 2:
		return &types.Float{Target: &r.Longitude}
	case 3:
		return &types.String{Target: &r.Message_id}
	case 4:
		return &types.Long{Target: &r.Created_at}
	case 5:
		return &types.Long{Target: &r.Updated_at}
	case 6:
		return &types.String{Target: &r.Timestamp}
	}
	panic("Unknown field index")
}

func (r *Fsa) SetDefault(i int) {
	switch i {
	case 1:
		r.Latitude = 0
		return
	case 2:
		r.Longitude = 0
		return
	case 3:
		r.Message_id = ""
		return
	case 6:
		r.Timestamp = ""
		return
	}
	panic("Unknown field index")
}

func (r *Fsa) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Fsa) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Fsa) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Fsa) Finalize()                        {}

func (_ *Fsa) AvroCRC64Fingerprint() []byte {
	return []byte(FsaAvroCRC64Fingerprint)
}
