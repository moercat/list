package list

import (
	"database/sql/driver"
	"github.com/spf13/cast"
	"sort"
)

// 	包 list 用来解决 go 中 slice 切片函数操作方法过少的问题.
//	Package list is used to solve the problem of too few slice functions in go.
// 	通过实现 python 中 pop remove 等方法来提高可用性
//	Improve usability by implementing methods like pop remove in python
type List struct {
	value  *[]string
	length int
}

// NewList converts a interface to List.
func NewList(va interface{}) List {
	val := cast.ToStringSlice(va)
	l := len(val)
	return List{
		value:  &val,
		length: l,
	}
}

func NilList(va interface{}) List {
	var val []string
	return List{
		value:  &val,
		length: 0,
	}
}

func (d List) Copy() List {
	d.ensureInitialized()
	return List{
		value:  &(*d.value),
		length: d.length,
	}
}

func NewStrSlice(va interface{}) *[]string {
	val := cast.ToStringSlice(va)
	return &val
}

// New returns a new fixed-point decimal, value * 10 ^ length.
func New(value int64, length int) List {
	return List{
		value:  NewStrSlice(value),
		length: length,
	}
}

// Abs returns the absolute value of the string slice.
func (d List) Abs() List {
	d.ensureInitialized()
	d2Value := make([]string, 0, len(*d.value))
	for _, v := range *d.value {
		if val := cast.ToInt(v); val < 0 {
			d2Value = append(d2Value, cast.ToString(-val))
		} else {
			d2Value = append(d2Value, v)
		}
	}

	return List{
		value:  &d2Value,
		length: d.length,
	}
}

// Add returns d + d2.
func (d List) Add(d2 List) List {
	rdv := *d.value
	rdv2 := *d2.value
	l2 := d.length + d2.length
	d3Value := append(rdv, rdv2...)
	return List{
		value:  &d3Value,
		length: l2,
	}
}

// Sub returns d - d2.
func (d List) Sub(d2 List) List {
	rdv := *d.value
	rdv2 := *d2.value
	d2Value := append(rdv, rdv2...)
	d2Map := make(map[string]bool)
	for _, v := range d2Value {
		if !d2Map[v] {
			d2Map[v] = true
		}
	}
	l2 := len(rdv) + len(rdv2)
	d3Value := make([]string, 0, l2)
	for k, _ := range d2Map {
		d3Value = append(d3Value, k)
	}

	return List{
		value:  &d3Value,
		length: len(d3Value),
	}
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (d List) Equal(d2 List) bool {

	s1 := *d.value
	s2 := *d2.value
	if len(s1) != len(s2) {
		return false
	}
	for i, n := range s1 {
		if n != s2[i] {
			return false
		}
	}
	return true
}

// Equals is deprecated, please use Equal method instead
func (d List) Equals(d2 List) bool {
	return d.Equal(d2)
}

// Length returns the length
func (d List) Length() int {
	return d.length
}

// Int returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
func (d List) Int() []int {
	d.ensureInitialized()
	dValue := *d.value
	return cast.ToIntSlice(dValue)
}

func (d List) Bool() []bool {
	dValue := *d.value
	return cast.ToBoolSlice(dValue)
}

// String returns the string representation of the decimal
func (d List) String() []string {
	return d.string()
}

// Value implements the driver.Valuer interface for database serialization.
func (d List) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d List) string() []string {
	return cast.ToStringSlice(*d.value)
}

func (d *List) ensureInitialized() {
	if d.value == nil {
		d.value = new([]string)
	}
}

func (d List) Min() int {

	d2 := *d.value
	d3 := cast.ToIntSlice(d2)
	sort.Ints(d3)

	return d3[0]
}

// Max returns the largest List that was passed in the arguments.
func (d List) Max() int {

	d2 := *d.value
	d3 := cast.ToIntSlice(d2)
	sort.Ints(d3)

	return d3[len(d3)-1]
}

// Sum returns the combined total of the provided first and rest Decimals
func (d List) Sum() int {
	total := 0
	for _, item := range *d.value {
		total += cast.ToInt(item)
	}

	return total
}

func In(fat, sub interface{}) bool {
	_, ok := InI(fat, sub)
	return ok
}

func Index(fat, idx interface{}) int {
	index, _ := InI(fat, idx)
	return index
}

func Extend(fat, sub interface{}) interface{} {
	fats := cast.ToStringSlice(fat)
	subs := cast.ToStringSlice(sub)

	fats = append(fats, subs...)

	return fats
}

func Pop(fat interface{}, idx int) interface{} {
	fats := cast.ToStringSlice(fat)
	if idx < 0 {
		idx = len(fats) + idx
	}

	return append(fats[:idx], fats[(idx+1):]...)
}

func Remove(fat interface{}, value interface{}) interface{} {
	fats := cast.ToStringSlice(fat)
	str := cast.ToString(value)
	for i, v := range fats {
		if v == str {
			fat = append(fats[:i], fats[(i+1):]...)
		}

	}

	return fat
}

func Append(fat interface{}, value interface{}) interface{} {
	fats := cast.ToStringSlice(fat)
	str := cast.ToString(value)

	return append(fats, str)
}

func Insert(fat interface{}, idx int, value interface{}) interface{} {
	fats := cast.ToStringSlice(fat)
	str := cast.ToString(value)

	res := append(fats[:idx], str)
	fat = append(res, fats[idx:]...)

	return append(res, fats[idx:]...)
}

func Count(fat interface{}, value interface{}) (count int) {
	fats := cast.ToStringSlice(fat)
	str := cast.ToString(value)

	for _, v := range fats {
		if v == str {
			count += 1
		}
	}

	return count
}

func InI(fat interface{}, sub interface{}) (int, bool) {
	fats := cast.ToStringSlice(fat)
	s := cast.ToString(sub)

	for i, v := range fats {
		if v == s {
			return i, true
		}
	}

	return -1, false
}

func Equal(fat, sub interface{}) bool {
	s1 := cast.ToStringSlice(fat)
	s2 := cast.ToStringSlice(sub)
	if len(s1) != len(s2) {
		return false
	}
	for i, n := range s1 {
		if n != s2[i] {
			return false
		}
	}
	return true
}
