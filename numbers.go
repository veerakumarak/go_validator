package go_validator

import(
	"fmt"
)

type number interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint | uint16 | uint32 | uint64 | float32 | float64
}

type numberRule[T number] struct {
	value T
	kind string
	name string
	rules []*Rule
}

func (i *numberRule[number]) add(validation Validation, message string) *numberRule[number] {
	if i.rules == nil {
		i.rules = make([]*Rule, 0)
	}
	i.rules = append(i.rules, &Rule{validation, message})
	return i
}

func (i *numberRule[number]) Gte(min number)  *numberRule[number] {
	return i.add(func() bool { return i.value >= min}, "is not greater than equal to " + fmt.Sprintf("%v", min))
}

func (i *numberRule[number]) Lte(max number)  *numberRule[number] {
	return i.add(func() bool { return i.value <= max}, "is not less than equal to " + fmt.Sprintf("%v", max))
}

func (i *numberRule[number]) Gt(min number)  *numberRule[number] {
	return i.add(func() bool { return i.value > min}, "is not greater than " + fmt.Sprintf("%v", min))
}

func (i *numberRule[number]) Lt(max number)  *numberRule[number] {
	return i.add(func() bool { return i.value < max}, "is not less than " + fmt.Sprintf("%v", max))
}

func (i *numberRule[number]) Eq(value number) *numberRule[number] {
	return i.add(func() bool { return i.value == value}, "is not equal to " + fmt.Sprintf("%v", value))
}

func (i *numberRule[number]) Ne(value number)  *numberRule[number] {
	return i.add(func() bool { return i.value != value}, "is equal to " + fmt.Sprintf("%v", value))
}

func (i *numberRule[number]) Validate() []error {
	return Validate(i.rules, i.kind, fmt.Sprintf("%v", i.value))
}


func Int(value int) *numberRule[int] {
	return &numberRule[int]{value: value, kind: "int"}
}
func Int8(value int8) *numberRule[int8] {
	return &numberRule[int8]{value: value, kind: "int8"}
}
func Int16(value int16) *numberRule[int16] {
	return &numberRule[int16]{value: value, kind: "int16"}
}
func Int32(value int32) *numberRule[int32] {
	return &numberRule[int32]{value: value, kind: "int32"}
}
func Int64(value int64) *numberRule[int64] {
	return &numberRule[int64]{value: value, kind: "int64"}
}
func Uint(value uint) *numberRule[uint] {
	return &numberRule[uint]{value: value, kind: "uint"}
}
func Uint8(value uint8) *numberRule[uint8] {
	return &numberRule[uint8]{value: value, kind: "uint8"}
}
func Uint16(value uint16) *numberRule[uint16] {
	return &numberRule[uint16]{value: value, kind: "uint16"}
}
func Uint32(value uint32) *numberRule[uint32] {
	return &numberRule[uint32]{value: value, kind: "uint32"}
}
func Uint64(value uint64) *numberRule[uint64] {
	return &numberRule[uint64]{value: value, kind: "uint64"}
}
func Float32(value float32) *numberRule[float32] {
	return &numberRule[float32]{value: value, kind: "float32"}
}
func Float64(value float64) *numberRule[float64] {
	return &numberRule[float64]{value: value, kind: "float64"}
}
