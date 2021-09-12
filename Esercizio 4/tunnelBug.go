package main



import "fmt"
import "time"
import "math/rand"



type Gruppo struct {
    nome string
    nPalline int
}

type Tunnel struct {
    libero bool
}

func transumanza(g Gruppo, t chan Tunnel, c1 chan int){

  for g.nPalline > 0{


    time.Sleep(time.Duration(rand.Intn(2))*time.Second)
    mandaPersona(&g,t,c1)

  }
}

func mandaPersona(g *Gruppo, t chan Tunnel, c1 chan int){

    tunnel := <- t

    if tunnel.libero {

        tunnel.libero = false
        t <- tunnel
        fmt.Println("qui")
        fmt.Println(g.nome)

        select{

          case <- c1:
               fmt.Println("qui")
			         fmt.Println("scontro ", g.nome)
                // GO si arrabbia se non usate una varabile...

		      case <-time.After(time.Second):

                tunnel := <- t
                tunnel.libero = true
                t <- tunnel
                fmt.Println("passato")
                g.nPalline = g.nPalline - 1
                fmt.Println("rimangono ", g.nPalline, " nel gruppo ", g.nome)
	    }

        //time.Sleep(time.Duration(4)*time.Second)
    } else{
        //time.Sleep(time.Duration(4)*time.Second)
        c1 <- 1
        tunnel.libero = true
        t <- tunnel
        //time.Sleep(time.Duration(4)*time.Second)
    }

}

func main() {

    rand.Seed(time.Now().UnixNano())
    gruppo1 := Gruppo{"destra",5}
    gruppo2 := Gruppo{"sinistra",5}

    c1 := make(chan int,1)

    tunnelChannel := make(chan Tunnel)
    tunnel := Tunnel{true}
    tunnelChannel <- tunnel

    go transumanza(gruppo1,tunnelChannel,c1)
    go transumanza(gruppo2,tunnelChannel,c1)

	time.Sleep(time.Minute)
}
