package csv_to_json

import (
	"encoding/csv"
	"strings"

	"github.com/OSSPAI/Lokomotiv/internal/wagon"
)

const (
	CSV2JSONCellDelimeterParam = "cellDelimeter"
	CSV2JSONFormatParam        = "format"
)

type CSVToJSONWagon struct {
	configDescription wagon.ConfigDescription
}

func NewCSVToJSONWagon() wagon.Wagon {
	return &CSVToJSONWagon{
		configDescription: wagon.ConfigDescription{
			{
				Name:         "Cell delimeter",
				ID:           CSV2JSONCellDelimeterParam,
				Type:         wagon.ConfigTypeString,
				DefaultValue: wagon.NewConfigStringValue(CSV2JSONCellDelimeterParam, ","),
			},
			{
				Name:         "Format",
				ID:           CSV2JSONFormatParam,
				Type:         wagon.ConfigTypeEnum,
				DefaultValue: wagon.NewConfigEnumValue(CSV2JSONFormatParam, 1),
				EnumParams: &wagon.ConfigEnumParams{
					Options: []wagon.ConfigEnumParamsOption{
						{
							Name:  "Array of dictionaries",
							Value: 0,
						},
						{
							Name:  "Array of arrays",
							Value: 1,
						},
					},
				},
			},
		},
	}
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

func (w *CSVToJSONWagon) GetConfigDescription() wagon.ConfigDescription {
	return w.configDescription
}

func (w *CSVToJSONWagon) Run(data *wagon.WagonData, config wagon.Config) (*wagon.WagonData, error) {
	delimeter := config.GetValueWithDefault(
		CSV2JSONCellDelimeterParam,
		w.GetConfigDescription().GetDescription(CSV2JSONCellDelimeterParam).GetDefaultValue(),
	).GetString()

	reader := csv.NewReader(strings.NewReader(data.DataString))

	if delimeter != nil {
		reader.Comma = rune((*delimeter)[0])
	}

	_, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
