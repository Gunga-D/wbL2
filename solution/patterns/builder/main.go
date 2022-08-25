package main

import "fmt"

type StringBuilder struct {
	kernel string
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{}
}

func (sb *StringBuilder) Append(element string) *StringBuilder {
	sb.kernel = sb.kernel + element
	return sb
}

func (sb *StringBuilder) Insert(index int, element string) *StringBuilder {
	sb.kernel = sb.kernel[0:index] + element + sb.kernel[index:]
	return sb
}

func (sb StringBuilder) String() string {
	return sb.kernel
}

func main() {
	stringBuilder := NewStringBuilder().Append("Hello").Append("!").Insert(5, " World")
	fmt.Println(stringBuilder)
}
