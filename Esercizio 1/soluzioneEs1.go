/*

TESTO ESERCIZIO
--------------------

Scrivete un programma che simuli una agenzia di viaggi che deve gestire
le prenotazioni per due diversi viaggi da parte di 7 clienti.
Ogni cliente fa una prenotazione per un viaggio in una
delle due mete disponibili (Spagna e Francia), ognuna delle quali ha
un numero minimo di partecipanti per essere confermata (rispettivamente 4 e 2).

● Creare la struttura Cliente col relativo campo “nome”.

● Creare la struttura Viaggio col rispettivo campo “meta”.

● Creare la function prenota, che prende come input una persona e
  che prenota uno a caso dei due viaggi.

● Creare una function stampaPartecipanti che alla fine del processo stampa
  quali viaggi sono confermati e quali persone vanno dove.

● Ogni persona può prenotarsi al viaggio contemporaneamente.

● Create tutte le classi e function che vi servono, ma mantenete
  la struttura data dalle due strutture e le due function che ho elencato sopra.

*/



package main



// Pacchetti importati
import "fmt"
import "time"
import "math/rand"



// Struttura Cliente
type Cliente struct {
  nome string
  viaggio Viaggio
}

// Struttura Viaggio
type Viaggio struct {
  meta string
}



// Permette ad ogni utente di prenotare il viaggio
func prenota(maxP int) {

  // Crea un channel di lunghezza pari al numero dei clienti
  prenotazione := make(chan Cliente, maxP)

  for i := 0; i < maxP; i++ {

    go sceltaDestinazione(Cliente{nome: fmt.Sprint("user_", i)}, prenotazione)

  }

  stampaPartecipanti(prenotazione, maxP)

}



// Stampa clienti catalogati per meta del viaggio
func stampaPartecipanti(prenotazione chan Cliente, maxP int) {

  var fString string
  var sString string
  fCount := 0
  sCount := 0

  for i := 0; i < maxP; i++ {

    // Fa un "pop" dell'ultimo elemento entrato nel channel
    // e lo assegna alla variabile persona
    persona := <- prenotazione

    // Verifica la destinazione del passeggero
    // E appende il nome alla stringa contenete tutti i nomi dello stesso volo
    if persona.viaggio.meta == "Francia" {

      fCount ++
      fString = fmt.Sprint(fString, persona.nome, "\n",)

    }else{

      sCount ++
      sString = fmt.Sprint(sString, persona.nome, "\n",)

    }

  }

  // Stampa le varie stringhe i voli che hanno un numero sufficiente
  // di clienti per partire (4 per la Spagna e 2 per la Francia)
  if sCount >= 4 {

    fmt.Println("E' stato raggiunto il numero minimo di clienti per il viaggio in Spagna.")
    fmt.Println("I clienti che hanno prenotato il viaggio in Spagna sono:")
    fmt.Println(sString)

  }
  if fCount >= 2 {

    fmt.Println("E' stato raggiunto il numero minimo di clienti per il viaggio in Francia.")
    fmt.Println("I clienti che hanno prenotato il viaggio in Francia sono:")
    fmt.Println(fString)

  }

}



// Avvia i singoli thread che scelgono la destinazione per ogni cliente
func sceltaDestinazione(persona Cliente, prenotazione chan Cliente) {

  var locDest string

  // Sceglie a random la destinazione del viaggio del passeggero
  // rand.Intn(n) parte da 0 fino a n-1 per questo motivo la meta' esatta e' 49
  if rand.Intn(100) < 49 {

    locDest = "Francia"

  } else {

    locDest = "Spagna"

  }

  // Asseggna alla persona il viaggio scelto
  persona.viaggio = Viaggio{meta: locDest}
  // Inserisce la struttura persona con il relativo campo viaggio compilato
  // nel channel creato nella function prenota
  prenotazione <- persona

}



// Inizio del programma
func main() {

  // "randomizza" maggiormente i numeri generati
  rand.Seed(time.Now().UnixNano())
  prenota(7)

}
