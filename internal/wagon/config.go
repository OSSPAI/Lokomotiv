package wagon

type ConfigType int8

const (
	ConfigTypeUnknown ConfigType = iota
	ConfigTypeString
	ConfigTypeInt
	ConfigTypeFloat
	ConfigTypeByte
	ConfigTypeBool
	ConfigTypeEnum
)

type ConfigStringParams struct {
	MinLength int
	MaxLength int
}

type ConfigIntParams struct {
	Min int
	Max int
}

type ConfigFloatParams struct {
	Min float64
	Max float64
}

type ConfigByteParams struct{}

type ConfigBoolParams struct{}

type ConfigEnumValue int8

type ConfigEnumParamsOption struct {
	Name  string
	Value ConfigEnumValue
}

type ConfigEnumParams struct {
	Options []ConfigEnumParamsOption
}

type ConfigDescriptionItem struct {
	Name         string
	ID           string
	Type         ConfigType
	DefaultValue *ConfigValue

	StringParams *ConfigStringParams
	IntParams    *ConfigIntParams
	FloatParams  *ConfigFloatParams
	ByteParams   *ConfigByteParams
	BoolParams   *ConfigBoolParams
	EnumParams   *ConfigEnumParams
}

type ConfigValue struct {
	ID   string
	Type ConfigType

	StringValue string
	IntValue    int
	FloatValue  float64
	ByteValue   []byte
	BoolValue   bool
	EnumValue   ConfigEnumValue
}

type Config []*ConfigValue
type ConfigDescription []*ConfigDescriptionItem

// --------------------------
// Config description methods
// --------------------------

func NewConfigStringDescription(name string, id string, params *ConfigStringParams) *ConfigDescriptionItem {
	return &ConfigDescriptionItem{
		Name:         name,
		ID:           id,
		Type:         ConfigTypeString,
		StringParams: params,
	}
}

func NewConfigIntDescription(name string, id string, params *ConfigIntParams) *ConfigDescriptionItem {
	return &ConfigDescriptionItem{
		Name:      name,
		ID:        id,
		Type:      ConfigTypeInt,
		IntParams: params,
	}
}

func NewConfigFloatDescription(name string, id string, params *ConfigFloatParams) *ConfigDescriptionItem {
	return &ConfigDescriptionItem{
		Name:        name,
		ID:          id,
		Type:        ConfigTypeFloat,
		FloatParams: params,
	}
}

func NewConfigByteDescription(name string, id string, params *ConfigByteParams) *ConfigDescriptionItem {
	return &ConfigDescriptionItem{
		Name:       name,
		ID:         id,
		Type:       ConfigTypeByte,
		ByteParams: params,
	}
}

func NewConfigBoolDescription(name string, id string, params *ConfigBoolParams) *ConfigDescriptionItem {
	return &ConfigDescriptionItem{
		Name:       name,
		ID:         id,
		Type:       ConfigTypeBool,
		BoolParams: params,
	}
}

func NewConfigEnumDescription(name string, id string, params *ConfigEnumParams) *ConfigDescriptionItem {
	return &ConfigDescriptionItem{
		Name:       name,
		ID:         id,
		Type:       ConfigTypeEnum,
		EnumParams: params,
	}
}

func (d ConfigDescription) GetDescription(id string) *ConfigDescriptionItem {
	if d == nil {
		return nil
	}

	for _, description := range d {
		if description != nil && description.ID == id {
			return description
		}
	}

	return nil
}

func (d *ConfigDescriptionItem) GetDefaultValue() *ConfigValue {
	if d == nil {
		return nil
	}

	return d.DefaultValue
}

// --------------------
// Config value methods
// --------------------

func NewConfigStringValue(id string, value string) *ConfigValue {
	return &ConfigValue{
		ID:          id,
		Type:        ConfigTypeString,
		StringValue: value,
	}
}

func NewConfigIntValue(id string, value int) *ConfigValue {
	return &ConfigValue{
		ID:       id,
		Type:     ConfigTypeInt,
		IntValue: value,
	}
}

func NewConfigFloatValue(id string, value float64) *ConfigValue {
	return &ConfigValue{
		ID:         id,
		Type:       ConfigTypeFloat,
		FloatValue: value,
	}
}

func NewConfigByteValue(id string, value []byte) *ConfigValue {
	return &ConfigValue{
		ID:        id,
		Type:      ConfigTypeByte,
		ByteValue: value,
	}
}

func NewConfigBoolValue(id string, value bool) *ConfigValue {
	return &ConfigValue{
		ID:        id,
		Type:      ConfigTypeBool,
		BoolValue: value,
	}
}

func NewConfigEnumValue(id string, value ConfigEnumValue) *ConfigValue {
	return &ConfigValue{
		ID:        id,
		Type:      ConfigTypeEnum,
		EnumValue: value,
	}
}

func (v *ConfigValue) IsString() bool {
	return v.Type == ConfigTypeString
}

func (v *ConfigValue) IsInt() bool {
	return v.Type == ConfigTypeInt
}

func (v *ConfigValue) IsFloat() bool {
	return v.Type == ConfigTypeFloat
}

func (v *ConfigValue) IsByte() bool {
	return v.Type == ConfigTypeByte
}

func (v *ConfigValue) IsBool() bool {
	return v.Type == ConfigTypeBool
}

func (v *ConfigValue) IsEnum() bool {
	return v.Type == ConfigTypeEnum
}

func (v *ConfigValue) GetString() *string {
	if !v.IsString() {
		return nil
	}

	return &v.StringValue
}

func (v *ConfigValue) GetInt() *int {
	if !v.IsInt() {
		return nil
	}

	return &v.IntValue
}

func (v *ConfigValue) GetFloat() *float64 {
	if !v.IsFloat() {
		return nil
	}

	return &v.FloatValue
}

func (v *ConfigValue) GetByte() []byte {
	if !v.IsByte() {
		return nil
	}

	return v.ByteValue
}

func (v *ConfigValue) GetBool() *bool {
	if !v.IsBool() {
		return nil
	}

	return &v.BoolValue
}

func (v *ConfigValue) GetEnum() *ConfigEnumValue {
	if !v.IsEnum() {
		return nil
	}

	return &v.EnumValue
}

func (c Config) GetValue(id string) *ConfigValue {
	if c == nil {
		return nil
	}

	for _, configItem := range c {
		if configItem != nil && configItem.ID == id {
			return configItem
		}
	}

	return nil
}

func (c Config) GetValueWithDefault(id string, defaultValue *ConfigValue) *ConfigValue {
	if c == nil {
		return nil
	}

	value := c.GetValue(id)

	if value == nil {
		value = defaultValue
	}

	return value
}
