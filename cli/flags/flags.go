package flags

type Flag struct {
	Key   string
	Value FlagValue
}

func NewFlag[T FlagValue](key string, value string) *Flag {
	var val T
	val.Set(value)
	return &Flag{
		Key: key,
		Value: val,
	}
}