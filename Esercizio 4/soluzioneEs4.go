/*

TESTO ESERCIZIO
--------------------

Vi è dato un programma (che trovate su Moodle: tunnelBug.go)
che simuli la seguente situazione: Ci sono due gruppi di palline G1 e G2
in due luoghi diversi L1 e L2 uniti da un tunnel.
In L1 e in L2 ci sono due persone P1 e P2.

La persona P1 vuole lanciare tutte le palline in G1 da L1 a L2,
e viceversa P2 vuole lanciare lanciare le palline in G2 da L2 a L1.

Il tunnel è stretto, ci può passare solo una pallina alla volta.
Se due palline vengono lanciate nel tunnel contemporaneamente,
tornano al punto di partenza (immediatamente).
Una pallina attraversa il tunnel in un secondo (time.Sleep(time.Second)).

Una persona non può lanciare una pallina finché quella che ha lanciato
precedentemente non è arrivata a destinazione o
non ha incontrato una pallina che andava in senso contrario.

Ci sono due gruppi di palline e due routine che lanciano
le palline da un capo all’altro.
Le routine attendono un tempo casuale
(time.Sleep(time.Duration(rand.Intn(2))*time.Second))
prima di lanciare una nuova palla. Le routine finiscono
quando nel relativo gruppo le palline finiscono.

● Debuggate i deadlock! Ci sono errori nel codice.
Fate esperimenti, vedete cosa succede, ripercorrete la logica
e rendete mutualmente esclusive la parti di codice
che devono avvenire una alla volta.
Possono esserci soluzioni diverse e non mi aspetto
che raggiungiate tutti la stessa.

● Per vedere il funzionamento del programma,
inserite delle stampe che vi mostrino in che ordine girano i thread
e aggiungete degli sleep per forzare un certo ordine dell’esecuzione.

● Non ci sono commenti nel code. Dovete capire voi cosa fa.

*/



package main

// Pacchetti importati
import "fmt"
import "time"
import "math/rand"


// Struttura Gruppo
type Gruppo struct {
    nome string
    nPalline int
}

// Struttura Tunnel
type Tunnel struct {
    libero bool
}

// Verifica se ci sono ancora palline da lanciare
func transumanza(g Gruppo, t chan Tunnel, c1 chan int, c2 chan int){

  for g.nPalline > 0{

    // Aspetta un tempo variabile e le lancia
    time.Sleep(time.Duration(rand.Intn(2))*time.Second)
    mandaPersona(&g, t, c1, c2)
   }
}

func mandaPersona(g *Gruppo, t chan Tunnel, c1 chan int, c2 chan int){

  //Prende il tunnel e controlla se è libero
  tunnel := <- t
  if tunnel.libero {
    // Essendo libero lo setta a occupato
    tunnel.libero = false
    t <- tunnel
    fmt.Println(g.nome, "qui")
    // Channel che mi indica se il tunnel è libero o occupato
    c2 <- 1


    select{

      // Mi da lo scontro tra le palline se vengono lanciate entrambe nel
      // tunnel prima che una delle 2 concluda l'attraversamento
      case x := <- c1:
        fmt.Println("scontro ", g.nome)
        x = x-x

      // Attraversamento di una pallina nel tunnel
      case <-time.After(time.Second):

        // Mi indica che la pallina ha attraversato il tunnel
        <- c2

        // Mi dice che il tunnel è nuovamente libero
        tunnel := <- t
        tunnel.libero = true
        t <- tunnel

        // Mi dice che la pallina è passata e quante me ne rimangono da lanciare
        fmt.Println("passato")
        g.nPalline = g.nPalline - 1
        fmt.Println("rimangono ", g.nPalline, " nel gruppo ", g.nome,)



      }
      // Sleep per tentare di forzare il deadlock
      //time.Sleep(time.Duration(4)*time.Second)

   } else{

     // Sleep per tentare di forzare il deadlock
     //time.Sleep(time.Duration(4)*time.Second)

     select{

     // Indica che il tunnel è ancora occupato e che se lancio ora
     // le palline si scontreranno
     case <- c2:
       c1 <- 1
       t <- tunnel

      // Indica che il tunnel è libero e posso tirare la pallina senza problemi
      default:

        tunnel.libero = true
        t <- tunnel

      }
      // Sleep per tentare di forzare il deadlock
      //time.Sleep(time.Duration(4)*time.Second)
  }
}

func main() {

  // Prende il tempo
  rand.Seed(time.Now().UnixNano())

  // Crea i 2 gruppi di palline
  gruppo1 := Gruppo{"destra", 5}
  gruppo2 := Gruppo{"sinistra", 5}

  // Crea i 2 channel di controllo
  c1 := make(chan int, 1)
  c2 := make(chan int, 1)

  // Crea il channel nel quale lanciare le palline
  // e lo inizializzo a vuoto
  tunnelChannel := make(chan Tunnel, 1)
  tunnel := Tunnel{true}
  tunnelChannel <- tunnel

  // Crea i 2 thread che lanciano le palline
  // sinistro e destro
  go transumanza(gruppo1, tunnelChannel, c1, c2)
  go transumanza(gruppo2, tunnelChannel, c1, c2)

  // Ferma il programma in caso di stallo
  time.Sleep(time.Minute)
}
