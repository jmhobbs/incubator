package device

type Type uint8

const (
	Temperature Type = iota
	Humidity
)

type Measurement struct {
	Type   Type
	Source string
	Value float64
}
