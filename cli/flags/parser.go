package flags

import (
	"regexp"
	"strconv"
)

var flagKeyRegex = regexp.MustCompile(`-{1,2}(.+)`)

type FlagValue interface {
	String() string
	Set(string) error
}

type intValue int

func (i *intValue) String() string {
	return strconv.Itoa(int(*i))
}

func (i *intValue) Set(value string) error {
	val, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	*i = intValue(val)
	return nil
}

type stringValue string

func (s *stringValue) Set(value string) error {
	*s = stringValue(value)
	return nil
}

func (s *stringValue) String() string {
	return string(*s)
}

type boolValue bool

func (b *boolValue) Set(value string) error {
	val, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}

	*b = boolValue(val)
	return nil
}

func (b *boolValue) String() string {
	return strconv.FormatBool(bool(*b))
}

func ParseFlags(args []string) []*Flag {
	var flags []*Flag
	for i, arg := range args[1:] {
		key := flagKeyRegex.FindString(args[i-1])
		if key != "" {
			if !flagKeyRegex.MatchString(arg) {
				_, err := strconv.Atoi(arg)
				if err != nil {
					flags = append(flags, NewFlag[*stringValue](key, arg))
					continue
				}

				flags = append(flags, NewFlag[*intValue](key, arg))
				continue
			}

			flags = append(flags, NewFlag[*boolValue](key, "true"))
		}
	}
	return flags
}