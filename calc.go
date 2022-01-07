package main

import (
  "fmt"
  "time"
)

func main() {
  var numero int
  var numero2 int
  var selec int
  fmt.Println("\n[~] Calculadora en Go")
  fmt.Println("[~] Elije una opcion")
  fmt.Println("[1] Sumar")
  fmt.Println("[2] Resta")
  fmt.Println("[3] Multiplicacion")
  fmt.Println("[4] Dividir")
  fmt.Print("[~] Opcion: ")
  fmt.Scanf("%d\n",&selec)
  if selec == 1 {
    fmt.Print("\n[+] Ingresa el primer numero: ")
    fmt.Scanf("%d\n",&numero)
    fmt.Print("\n[+] Ingresa el segundo numero: ")
    fmt.Scanf("%d\n",&numero2)
    fmt.Println("[+] Sumando...")
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("[~] Resultado: ", numero + numero2)
  } else if selec == 2 {
    fmt.Print("\n[-] Ingresa el primer numero: ")
    fmt.Scanf("%d\n",&numero)
    fmt.Print("\n[-] Ingresa el segundo numero: ")
    fmt.Scanf("%d\n",&numero2)
    fmt.Println("[~] Restando...")
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("[~] Resultado: ", numero - numero2)
  } else if selec == 3 {
    fmt.Print("\n[X] Ingresa el primer numero: ")
    fmt.Scanf("%d\n",&numero)
    fmt.Print("\n[X] Ingresa el segundo numero: ")
    fmt.Scanf("%d\n",&numero2)
    fmt.Println("[~] Multiplicando...")
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("[~] Resultado: ", numero * numero2)
  } else if selec == 4{
    fmt.Print("\n[X] Ingresa el primer numero: ")
    fmt.Scanf("%d\n",&numero)
    fmt.Print("\n[X] Ingresa el segundo numero: ")
    fmt.Scanf("%d\n",&numero2)
    fmt.Println("[~] Dividiendo...")
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("[~] Resultado: ", numero / numero2)
  } else {
    fmt.Print("[~] Error opcion invalida intenta de nuevo.")
  }
}
