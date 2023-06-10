package main

import (
	"errors"
	"fmt"
	"strconv"
	"try"
)

func Reverse(v string) (string, error) {
	if len(v) == 4 {
		return "", errors.New("no 4 letter words")
	}
	runes := []rune(v)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}

func RandomValidation(in try.Result[string]) try.Result[string] {
	return try.IfOk(in, func(in try.Ok[string]) try.Result[string] {
		if in.Value == "hi" {
			return try.Fail[string](errors.New("no hi. only zulu"))
		}
		return in
	})
}

func dataSetExample() {
	data := []string{"2", "4444", "44", "44fas", "ih"}

	step1 := try.Map(try.To(Reverse))(try.PassSlice(data))
	step2 := try.Map(RandomValidation)(step1)
	step3 := try.Map(try.To(strconv.Atoi))(step2)
	pass, fail := try.Collect(step3)

	fmt.Println("PASS:")
	for _, e := range pass {
		fmt.Println("  ", e)
	}
	fmt.Println("Fail:")
	for _, e := range fail {
		fmt.Println("  ", e.Error())
	}
}

func singleValueExample(data string) {
	step1 := try.To(Reverse)(try.Pass(data))
	step2 := RandomValidation(step1)
	step3 := try.To(strconv.Atoi)(step2)

	switch res := step3.(type) {
	case try.Ok[int]:
		fmt.Println("I did it!:", res.Value)
	case try.Error:
		fmt.Println("I failed:", res.Value)
	}
}

func main() {
	dataSetExample()
	singleValueExample("44")
	singleValueExample("ih")
}
