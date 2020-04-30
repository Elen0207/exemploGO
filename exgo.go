package main

import (
    "fmt"
    "math/rand"
)//biblioteca
//0 , 999999999999
func geraVetor()[100]int{
    var vetor [100]int

    for i := 0 ; i < 100 ; i++ {  //:= significa que o valor ainda não foi definido
        vetor[i]  = int((rand.Float32() * 1000) + 1)  // parse para ,
    }
    return vetor
}
func  main(){
    var vet [100]int
    var qtdPares, qtdImpares int

    vet = geraVetor()
    fmt.Println(vet)

    for i := 0 ; i < 100 ; i++ {
        if vet[i] % 2 ==0 {
            qtdPares++
        } else{
            qtdImpares
        }
    }
    fmt.Println("Total pares: ", qtdPares)
	fmt.Println("Total impares: ", qtdImpares)

}