package wagon

type ConfigType int8

const (
	ConfigTypeString ConfigType = iota
	ConfigTypeInt
	ConfigTypeFloat
	ConfigTypeByte
	ConfigTypeBool
)

type ConfigDescription struct {
	Name         string
	ID           string
	Type         ConfigType
	DefaultValue *ConfigValue
}

type ConfigValue struct {
	ID   string
	Type ConfigType

	StringValue string
	IntValue    int
	FloatValue  float64
	ByteValue   []byte
	BoolValue   bool
}

// --------------------------
// Config description methods
// --------------------------

func NewConfigStringDescription(name string, id string) *ConfigDescription {
	return &ConfigDescription{
		Name: name,
		ID:   id,
		Type: ConfigTypeString,
	}
}

func NewConfigIntDescription(name string, id string) *ConfigDescription {
	return &ConfigDescription{
		Name: name,
		ID:   id,
		Type: ConfigTypeInt,
	}
}

func NewConfigFloatDescription(name string, id string) *ConfigDescription {
	return &ConfigDescription{
		Name: name,
		ID:   id,
		Type: ConfigTypeFloat,
	}
}

func NewConfigByteDescription(name string, id string) *ConfigDescription {
	return &ConfigDescription{
		Name: name,
		ID:   id,
		Type: ConfigTypeByte,
	}
}

func NewConfigBoolDescription(name string, id string) *ConfigDescription {
	return &ConfigDescription{
		Name: name,
		ID:   id,
		Type: ConfigTypeBool,
	}
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
