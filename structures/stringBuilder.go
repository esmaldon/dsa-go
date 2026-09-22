package structures

type stringBuilder struct {
	array []string
}

func NewStringBuilder() stringBuilder {
	return stringBuilder{}
}

func (sb *stringBuilder) append(s string) {
	sb.array = append(sb.array, s)
}

func (sb *stringBuilder) toString() string {
	string := ""
	for _, s := range sb.array {
		string = string + s
	}
	return string
}
