package list

import (
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"testing"
	"time"
)

func RandomStringSlice() (str []string) {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rand.Intn(999); i++ {
		str = append(str, cast.ToString(rand.Int()))
	}

	return str
}

func TestList_Abs(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Abs())

}

func TestList_Append(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Append(rand.Int()))
}

func TestList_BoolSlice(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).BoolSlice())
}

func TestList_Copy(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Copy())
}

func TestList_Count(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Count(1))
}

func TestList_Dup(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Dup(RandomStringSlice))
}

func TestList_Equal(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Equal(str))
	fmt.Println(NewList(str).Equal(RandomStringSlice))
}

func TestList_Extend(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Extend(str))
	fmt.Println(NewList(str).Extend(RandomStringSlice))
}

func TestList_In(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).In(rand.Int()))
}

func TestList_Index(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Index(rand.Int()))
}

func TestList_Insert(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Insert(rand.Int(), rand.Int()))
}

func TestList_IntSlice(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).IntSlice())
}

func TestList_Length(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Length())
}

func TestList_Max(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Max())
}

func TestList_Min(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Min())
}

func TestList_Pop(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Pop(rand.Int()))
}

func TestList_Remove(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Remove(rand.Int()))
}

func TestList_Set(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Set())
}

func TestList_String(t *testing.T) {

	str := RandomStringSlice()

	fmt.Println(NewList(str).String())

}

func TestList_StringSlice(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).StringSlice())
}

func TestList_Sum(t *testing.T) {
	str := RandomStringSlice()

	fmt.Println(NewList(str).Sum())
}
