package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 100)
	fmt.Println("Literal inteiro é ", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é ", i1)
	fmt.Println("O tipo de i1 é ", reflect.TypeOf(i1))

	//representa um múltiplo de 32bits
	var i2 rune = 'a'
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais(float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo doliteral 49.99 é", reflect.TypeOf(49.99))

	//booleans
	bo := true
	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(bo)

	//strings
	s1 := "Olá meu nome é Felipe"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é ", len(s1))

	//string com multiplas linhas	(interpolação)	\n
	s2 := `Olá	
	meu nome é
	Felipe`
	fmt.Println(s2)

	//characteres	(unicode)
	s3 := 'a' //'a' é um rune	(charactere)	(unicode)	(int32)	(uint16)
	fmt.Println("O tipo de s3 é ", reflect.TypeOf(s3))
	fmt.Println(s3)
	fmt.Println(reflect.TypeOf(s3)) //(uint16)	(int32)	(rune)	(charactere)

}
