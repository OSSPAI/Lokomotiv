package wagon

type Type int8

const (
	TypeDataFormat Type = iota
)

type DataType int8

const (
	DataTypeJSON DataType = iota
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
	GetConfigDescription() []ConfigDescription

	Run(data *WagonData, config []*ConfigValue) *WagonData
}
