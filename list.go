package list

import (
	"database/sql/driver"
	"github.com/spf13/cast"
	"sort"
	"strings"
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

func NewStrSlice(va interface{}) *[]string {
	val := cast.ToStringSlice(va)
	return &val
}

func (d List) Min() int {

	d2 := *d.value
	d3 := cast.ToIntSlice(d2)
	sort.Ints(d3)

	if d.length > 0 {
		return d3[0]
	}

	return -1
}

func (d List) Max() int {

	d2 := *d.value
	d3 := cast.ToIntSlice(d2)
	sort.Ints(d3)

	if d.length > 0 {
		return d3[len(d3)-1]
	}

	return -1
}

func (d List) Sum() int {
	total := 0
	for _, item := range *d.value {
		total += cast.ToInt(item)
	}

	return total
}

func (d List) Set() List {
	rdv := *d.value
	d2Map := make(map[string]bool)
	for _, v := range rdv {
		if !d2Map[v] {
			d2Map[v] = true
		}
	}

	d3Value := make([]string, 0, d.length)
	for k, _ := range d2Map {
		d3Value = append(d3Value, k)
	}

	return List{
		value:  &d3Value,
		length: len(d3Value),
	}
}

func (d List) Abs() List {
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

func (d List) Copy() List {
	return List{
		value:  &(*d.value),
		length: d.length,
	}
}

func (d List) Pop(idx int) List {
	fats := *d.value
	if idx < 0 {
		idx = d.length + idx
	}
	if idx > d.length {
		return d
	}

	*d.value = append(fats[:idx], fats[(idx+1):]...)
	return d
}

func (d List) Extend(sub interface{}) List {

	subs := cast.ToStringSlice(sub)
	*d.value = append(*d.value, subs...)

	return d
}

func (d List) Dup(d2 interface{}) List {
	dv := cast.ToStringSlice(d2)
	rdv := *d.value
	rdv2 := dv
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

func (d List) In(sub interface{}) bool {
	_, ok := inI(d.value, sub)
	return ok
}

func (d List) Remove(value interface{}) List {
	fats := *d.value
	str := cast.ToString(value)
	for i, v := range fats {
		if v == str {
			*d.value = append(fats[:i], fats[(i+1):]...)
		}
	}

	return d
}

func (d List) Append(value interface{}) List {
	fats := *d.value
	str := cast.ToString(value)
	*d.value = append(fats, str)
	return d
}

func (d List) Equal(d2 interface{}) bool {
	dv := cast.ToStringSlice(d2)
	s1 := *d.value
	s2 := dv
	if d.length != len(s2) {
		return false
	}
	for i, n := range s1 {
		if n != s2[i] {
			return false
		}
	}
	return true
}

func (d List) Index(sub interface{}) int {
	index, _ := inI(d.value, sub)
	return index
}

func (d List) Insert(idx int, value interface{}) List {
	fats := *d.value
	str := cast.ToString(value)

	if idx > d.length {
		return d
	}

	res := append(fats[:idx], str)
	*d.value = append(res, fats[idx:]...)

	return d
}

func (d List) Count(value interface{}) (count int) {
	fats := *d.value
	str := cast.ToString(value)

	for _, v := range fats {
		if v == str {
			count += 1
		}
	}

	return
}

// Length returns the length
func (d List) Length() int {
	return d.length
}

func (d List) IntSlice() []int {
	dValue := *d.value
	return cast.ToIntSlice(dValue)
}

func (d List) BoolSlice() []bool {
	dValue := *d.value
	return cast.ToBoolSlice(dValue)
}

func (d List) StringSlice() []string {
	return d.string()
}

// String 修改输出结果为数组形式
func (d List) String() string {
	v := *d.value
	return "[" + strings.Join(v, " ") + "]"
}

func (d List) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d List) string() []string {
	return cast.ToStringSlice(*d.value)
}

func (d List) ensureInitialized() {
	if d.value == nil {
		d.value = new([]string)
	}
}

func inI(fat *[]string, sub interface{}) (int, bool) {
	s := cast.ToString(sub)

	for i, v := range *fat {
		if v == s {
			return i, true
		}
	}

	return -1, false
}
