package plugins

type Accumulator interface{
	Add(name string, value interface{}, tags map[string]string)
}