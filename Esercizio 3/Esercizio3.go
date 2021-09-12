/*

TESTO ESERCIZIO
--------------------

Scrivete un programma che simuli un lavoro fatto da tre operai,
ognuno dei quali deve usare un martello, un cacciavite e un trapano
per fare un lavoro.
Devono usare il cacciavite DOPO il trapano e il martello in un qualsiasi momento.
Ovviamente, possono fare solo un lavoro alla volta.
Gli attrezzi a disposizione sono: due trapani, un martello e un cacciavite,
quindi I tre operai devono aspettare di avere a disposizione
gli attrezzi per usarli. Modellate questa situazione minimizzando
il più possibile le attese.

● Creare la struttura Operaio col relativo campo “nome”.

● Creare la strutture Martello, Cacciavite e Trapano che devono
	essere “prese” dagli operai.

● Nelle function che creerete, inserite una stampa che dica quando
	l’operaio x ha preso l’oggetto y e quando ha finito di usarlo.

● Hint sulla logica: ogni operaio può avere solo un oggetto alla volta
	e ogni oggetto può essere in mano a un solo operaio.

● Per assicurarmi che ogni oggetto sia in mano a solo un operaio,
	posso mettere ogni operaio in un channel, e prima di cercare di
	prendere un oggetto...

*/
package main

// Pacchetti importati
import "fmt"
import "time"
import "math/rand"

// Struttura Operaio
type Operaio struct {
	nome string
	attrezzo1 Martello
	attrezzo2 Cacciavite
	attrezzo3 Trapano
	utilizzoMartello bool
	utilizzoCacciavite bool
	utilizzoTrapano bool
}

// Struttura Martello
type Martello struct {
	nome string
}

// Struttura Cacciavite
type Cacciavite struct {
	nome string
}

// Struttura Trapano
type Trapano struct {
	nome string
}


func lavoroOperaio(operaio Operaio, channelTrapano chan Trapano, channelMartello chan Martello, channelCacciavite chan Cacciavite) {


	for {

		// Se l'operaio ha utilizzato tutti gli attrezzi allora concludo
		if (operaio.utilizzoMartello == true) && (operaio.utilizzoCacciavite == true) && (operaio.utilizzoTrapano == true) {

			fmt.Println("")
			fmt.Println("l'operaio ", operaio.nome, "ha utilizzato tutti e 3 gli attrezzi")
			fmt.Println("")
			return
		}

		// Se l'operaio non ha utilizzato il Trapano
		if operaio.utilizzoTrapano != true {

			//fmt.Println("3")

			// Selezione in base ai casi possibili
			select {

				// Se il trapano è disponibile l'operaio lo prende dal channel
				case att := <- channelTrapano:
					fmt.Println(operaio.nome, "prende il trapano")
					// L'operaio ha utilizzato il Trapano
					operaio.utilizzoTrapano = true
					operaio.attrezzo3 = att
					time.Sleep(time.Duration(3) * time.Second)
					// Rimetto il trapano nel channel
					channelTrapano <- operaio.attrezzo3
					fmt.Println(operaio.nome, "ha finito con il trapano e lo rimette nel channel")

				// Se il trapano NON è disponibile l'operaio lo prende dal channel
				case att := <- channelMartello:
					fmt.Println(operaio.nome, "prende il martello")
					// L'operaio ha utilizzato il Martello
					operaio.utilizzoMartello = true
					operaio.attrezzo1 = att
					time.Sleep(time.Duration(3) * time.Second)
					// Rimetto il Martello nel channel
					channelMartello <- operaio.attrezzo1
					fmt.Println(operaio.nome, "ha finito con il martello e lo rimette nel channel")
			}

			// Se l'operaio ha usato il Trapano ma non il Cacciavite
		} else if (operaio.utilizzoTrapano == true) && (operaio.utilizzoCacciavite == false) {

			select {
				// Se il Cacciavite è disponibile l'operaio lo prende dal channel
				case att := <- channelCacciavite:
					fmt.Println(operaio.nome, "prende il cacciavite")
					// L'operaio ha utilizzato il Cacciavite
					operaio.utilizzoCacciavite = true
					operaio.attrezzo2 = att
					time.Sleep(time.Duration(3) * time.Second)
					// Rimetto il Cacciavite nel channel
					channelCacciavite <- operaio.attrezzo2
					fmt.Println(operaio.nome, "ha finito con il cacciavite e lo rimette nel channel")

				// Se il Cacciavite NON è disponibile l'operaio lo prende dal channel
				case att := <- channelMartello:
					fmt.Println(operaio.nome, "prende il martello")
					// L'operaio ha utilizzato il Martello
					operaio.utilizzoMartello = true
					operaio.attrezzo1 = att
					time.Sleep(time.Duration(3) * time.Second)
					// Rimetto il Martello nel channel
					channelMartello <- operaio.attrezzo1
					fmt.Println(operaio.nome, "ha finito con il martello e lo rimette nel channel")
			}

		} else {

			// L' operaio lo prende dal channel
			att := <- channelMartello
			fmt.Println(operaio.nome, "prende il martello")
			// L'operaio ha utilizzato il Martello
			operaio.utilizzoMartello = true
			operaio.attrezzo1 = att
			time.Sleep(time.Duration(3) * time.Second)
			// Rimetto il Martello nel channel
			channelMartello <- operaio.attrezzo1
			fmt.Println(operaio.nome, "ha finito con il martello e lo rimette nel channel")

		}
	}
}

// Inizio del programma
func main() {

	// Prende il tempo
	rand.Seed(time.Now().UnixNano())

	// Inizializzo il trapano e lo inserisco nel suo channel
	trapano := Trapano{fmt.Sprint("trapano")}
	channelTrapano := make(chan Trapano, 1)
	channelTrapano <- trapano

	// Inizializzo il Cacciavite e lo inserisco nel suo channel
	cacciavite := Cacciavite{fmt.Sprint("cacciavite")}
	channelCacciavite := make(chan Cacciavite, 1)
	channelCacciavite <- cacciavite

	// Inizializzo i Martelli e li inserisco nel loro channel
	martello1 := Martello{fmt.Sprint("martello1")}
	martello2 := Martello{fmt.Sprint("martello2")}
	channelMartello := make(chan Martello, 2)
	channelMartello <- martello1
	channelMartello <- martello2

	// Inizializzo i 3 operai
	operaio1 := Operaio{nome: fmt.Sprint("Operaio1"), utilizzoMartello: false, utilizzoCacciavite: false, utilizzoTrapano: false}
	operaio2 := Operaio{nome: fmt.Sprint("Operaio2"), utilizzoMartello: false, utilizzoCacciavite: false, utilizzoTrapano: false}
	operaio3 := Operaio{nome: fmt.Sprint("Operaio3"), utilizzoMartello: false, utilizzoCacciavite: false, utilizzoTrapano: false}

	// Faccio partire i thread (1 per ogni operaio)
	go lavoroOperaio(operaio1, channelTrapano, channelMartello, channelCacciavite)
	go lavoroOperaio(operaio2, channelTrapano, channelMartello, channelCacciavite)
	go lavoroOperaio(operaio3, channelTrapano, channelMartello, channelCacciavite)

	// Timer
	time.Sleep(time.Duration(50) * time.Second)


}
