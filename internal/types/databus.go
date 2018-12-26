package types

//Databus denotes the databus to use
type Databus struct {
	Type   string            `yaml:"type"`
	Config map[string]string `yaml:"config"`
}

//GetConfig get config with default value
func (d Databus) GetConfig(key, defaultValue string) string {
	val, ok := d.Config[key]
	if ok {
		return val
	}
	return defaultValue
}
