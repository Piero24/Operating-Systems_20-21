/*

TESTO ESERCIZIO
--------------------

Scrivete un programma che simuli l’ordinazione, la cottura e l’uscita dei piatti
in un ristorante.
10 clienti ordinano contemporaneamente i loro piatti.
In cucina vengono preparati in un massimo di 3 alla volta, essendoci solo 3 fornelli.
Il tempo necessario per preparare ogni piatto è fra i 4 e i 6 secondi.
Dopo che un piatto viene preparato, viene portato fuori da un cameriere,
che impiega 3 secondi a portarlo fuori. Ci sono solamente 2 camerieri nel ristorante.

● Creare la strutture Piatto e Cameriere col relativo campo “nome”.

● Creare le funzioni ordina che aggiunge il piatto a un buffer di piatti da fare;
  creare la function cucina che cucina ogni piatto e lo mette in lista per essere consegnato;
  creare la function consegna che fa uscire un piatto dalla cucina.

● Ogni cameriere può portare solo un piatto alla volta.

● Usate buffered channels per svolgere il compito.

● Attenzione: se per cucinare un piatto lo mandate nel buffer fornello di capienza 3
  e lo ritirate dopo 3 secondi, non è detto che ritiriate lo stesso piatto
  che avete messo sul fornello. Tenetelo in memoria. Ovviamente la vostra soluzione
  potrebbe differire dalla mia e questo hint potrebbe non servirvi.

*/



package main

// Pacchetti importati
import "fmt"
import "time"
import "math/rand"

// Struttura Piatto
type Piatto struct {
  nome string
}

// Struttura Cameriere
type Cameriere struct {
    nome string
    piatto Piatto
}

// Prende gli ordini dei clienti e li aggiunge al channel delle prenotazioni
// da mandare alla cucina
func ordine(prenotazioni chan Piatto, clienteNumero int) {

    fmt.Println(clienteNumero, "sta per ordinare")
    prenotazioni <- Piatto{fmt.Sprint("",clienteNumero)}

}



// Prepara gi ordini sui fornelli
func cucina(prenotazioni chan Piatto, pronto chan Piatto) {

    for {

        inPreparazione := <- prenotazioni
        // Tempo impiengato per la preparazione del piatto
        time.Sleep((time.Duration(rand.Int31n(6 - 4) + 4)) * time.Second)
        pronto <- inPreparazione
    }


}


// Il cameriere prende il piatto e lo porta al cliente aggiungendolo al channel
// dei piatti consegnati
func consegna(cameriere Cameriere, pronto chan Piatto, consegnato chan Piatto) {

    for {

        cameriere.piatto = <- pronto
        // Tempo per portare il piatto dalla cucina al tavolo
        time.Sleep(time.Duration(3) * time.Second)
        consegnato <- cameriere.piatto

    }
}



func main() {

    // N° clienti nel locale
    nClienti := 10

    rand.Seed(time.Now().UnixNano())

    // Channel dei piatti prenotati dai clienti
    prenotazioni := make(chan Piatto, nClienti)
    // Channel dei piatti già cucinati e pronti per essere presi in carico dal cameriere
    pronto := make(chan Piatto, nClienti)
    // Channel dei piatti consegnati ai clienti
    consegnato := make(chan Piatto, nClienti)

    // Camerieri del ristorante
    cameriere1 := Cameriere{nome: fmt.Sprint("cam1")}
    cameriere2 := Cameriere{nome: fmt.Sprint("cam1")}

    // 3 fornelli a disposizione per cucinare di conseguenza 3 thread per la preparazione
    // dei piatti in cucina
    for i := 0; i < 3; i++ {

        go cucina(prenotazioni, pronto)
    }

    // Ordini effettuati (anche in simmultanea) dei clienti
    for i := 0; i < nClienti; i++ {

        go ordine(prenotazioni, i)

    }

    // 2 Thread (uno per cameriere) per la consegna dei piatti ai clienti
    go consegna(cameriere1, pronto, consegnato)
    go consegna(cameriere2, pronto, consegnato)

    // Stampa a schermo i piatti consegnati ai clienti
    for i := 0; i < nClienti; i++ {

        fmt.Println(<- consegnato, "consegnato al cliente")
    }

}
