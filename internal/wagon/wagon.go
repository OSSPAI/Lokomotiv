package wagon

type Type int8

const (
	TypeUnknown Type = iota
	TypeDataFormat
)

type DataType int8

const (
	DataTypeUnknown DataType = iota
	DataTypeJSON
	DataTypeString
)

type WagonData struct {
	Type DataType

	DataJSON   string
	DataString string
}

type Wagon interface {
	GetType() Type
	GetName() string
	GetDescription() string
	GetInputType() DataType
	GetOutputType() DataType
	GetConfigDescription() ConfigDescription

	Run(data *WagonData, config Config) (*WagonData, error)
}
