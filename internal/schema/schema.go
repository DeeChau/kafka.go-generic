package schema

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type AvroSchemaConstraint interface {
	Fsa | FsaKey | State | StateKey
}

type AvroSchemaStruct[S any] interface {
	Serialize(w io.Writer) error
	Schema() string
	SchemaName() string
	Get(i int) types.Field
	SetBoolean(v bool)
	SetInt(v int32)
	SetLong(v int64)
	SetFloat(v float32)
	SetDouble(v float64)
	SetBytes(v []byte)
	SetString(v string)
	SetUnionElem(v int64)
	AppendMap(key string) types.Field 
	AppendArray() types.Field 
	Finalize()
	AvroCRC64Fingerprint() []byte
	SetDefault(i int)
	NullField(i int)
	*S
}


// Schema SerDe
func DeserializeFromSchema[S AvroSchemaConstraint, PT AvroSchemaStruct[S]](r io.Reader, schema string) (*S, error) {
	t := PT(NewSchemaStruct[S]())

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

func NewSchemaStruct[S any]() *S {
	return new(S)
}
