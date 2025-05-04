package csv_to_json

import "github.com/OSSPAI/Lokomotiv/internal/wagon"

type CSVToJSONWagon struct {
}

func NewCSVToJSONWagon() wagon.Wagon {
	return &CSVToJSONWagon{}
}

func (W *CSVToJSONWagon) GetType() wagon.Type {
	return wagon.TypeDataFormat
}

func (w *CSVToJSONWagon) GetName() string {
	return "CSV to JSON"
}

func (w *CSVToJSONWagon) GetDescription() string {
	return "Convert CSV to JSON"
}

func (w *CSVToJSONWagon) GetInputType() wagon.DataType {
	return wagon.DataTypeString
}

func (w *CSVToJSONWagon) GetOutputType() wagon.DataType {
	return wagon.DataTypeJSON
}

func (w *CSVToJSONWagon) GetConfigDescription() []wagon.ConfigDescription {
	return []wagon.ConfigDescription{
		{
			Name:         "Cell delimeter",
			ID:           "cellDelimeter",
			Type:         wagon.ConfigTypeString,
			DefaultValue: wagon.NewConfigStringValue("cellDelimeter", ","),
		},
		{
			Name:         "Row delimeter",
			ID:           "rowDelimeter",
			Type:         wagon.ConfigTypeString,
			DefaultValue: wagon.NewConfigStringValue("rowDelimeter", "\\r\\n"),
		},
	}
}

func (w *CSVToJSONWagon) Run(data *wagon.WagonData, config []*wagon.ConfigValue) *wagon.WagonData {
	return nil
}
