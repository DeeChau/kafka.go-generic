// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
package schema

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type State struct {
	// A two character identifier for the states name. This field is used as the key.
	Abbreviation string `json:"abbreviation"`
	// The full name of the State.
	Name string `json:"name"`

	Message_id string `json:"message_id"`
	// ISO 8601 datetime. Format: YYYY-MM-DD HH:MM:SS z
	Timestamp string `json:"timestamp"`
}

const StateAvroCRC64Fingerprint = "C\x8c\xa4U\xec\x0eEO"

func NewState() *State {
	return &State{}
}

func DeserializeState(r io.Reader) (*State, error) {
	t := NewState()
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

func DeserializeStateFromSchema(r io.Reader, schema string) (*State, error) {
	t := NewState()

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

func writeState(r *State, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Abbreviation, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Message_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Timestamp, w)
	if err != nil {
		return err
	}
	return err
}

func (r *State) Serialize(w io.Writer) error {
	return writeState(r, w)
}

func (r *State) Schema() string {
	return "{\"fields\":[{\"doc\":\"A two character identifier for the states name. This field is used as the key.\",\"name\":\"abbreviation\",\"type\":\"string\"},{\"doc\":\"The full name of the State.\",\"name\":\"name\",\"type\":\"string\"},{\"default\":\"\",\"name\":\"message_id\",\"type\":\"string\"},{\"default\":\"\",\"doc\":\"ISO 8601 datetime. Format: YYYY-MM-DD HH:MM:SS z\",\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"com.flipp.fadmin.State\",\"type\":\"record\"}"
}

func (r *State) SchemaName() string {
	return "com.flipp.fadmin.State"
}

func (_ *State) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *State) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *State) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *State) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *State) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *State) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *State) SetString(v string)   { panic("Unsupported operation") }
func (_ *State) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *State) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Abbreviation}
	case 1:
		return &types.String{Target: &r.Name}
	case 2:
		return &types.String{Target: &r.Message_id}
	case 3:
		return &types.String{Target: &r.Timestamp}
	}
	panic("Unknown field index")
}

func (r *State) SetDefault(i int) {
	switch i {
	case 2:
		r.Message_id = ""
		return
	case 3:
		r.Timestamp = ""
		return
	}
	panic("Unknown field index")
}

func (r *State) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *State) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *State) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *State) Finalize()                        {}

func (_ *State) AvroCRC64Fingerprint() []byte {
	return []byte(StateAvroCRC64Fingerprint)
}