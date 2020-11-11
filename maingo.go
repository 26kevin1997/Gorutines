package main

import (
	"fmt"
    "time"
    "strconv"
)
func proceso(id uint64 ,c1 chan string)  {
	i := uint64(0);
    for {
		f := "Id:"
		f += strconv.FormatUint(id, 10)
        f += " "
        
		g  := strconv.FormatUint(i, 10)
		f += g
		c1 <- f
		i = i + 1
		time.Sleep(time.Millisecond * 500)
    }
}

func main()  {
    canal1 := make(chan string)
    canal2 := make(chan string)
    var id uint64
    id = 0
    mostrar :=0
	cont := 1
	
	for cont != 0 {
        var c int
        if mostrar == 0{
            fmt.Println("1)Agregar proceso\n2)Mostrar procesos\n3)Eliminar proceso\n4)Exit")
            fmt.Scanln(&c)
        }
        switch c{
            case 1:
                id += 1
                go proceso(id,canal1)
                go func (){
                    for{
                        <-canal1
                    }
                }()
            case 2:
                go func ()  {
                    for {
                        select {
                        case msg1 := <- canal1:
                            if msg1 !=  "2"{
                                fmt.Println(msg1)
                            }
                        case msg2 := <-canal2:
                            msg2 = "1"
                            msg2 = msg2+"2"
                            <-canal1
                            return                    
                        }
                    }
                }()
                var input string
                fmt.Scanln(&input)
                canal2 <- input
           
                case 3:
                    fmt.Println("Eliminar")

                default:
                    fmt.Println("Fin")
                        cont = 0
                                  
        }
    }
}