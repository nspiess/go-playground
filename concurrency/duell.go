package concurrency

import (
	"crypto/rand"
	"fmt"
	"math/big"
)
import "time"
import "sync"

func shoot(a, b chan int, name string) { // Funktion, die 2 mal als goroutine aufgerufen wird
	var t = randomRangeValue(1, 3)                 // zufälliger Wert zwischen 1 und 3
	time.Sleep(time.Nanosecond * time.Duration(t)) // zufälliges, durch t bestimmtes Warten
	a <- 1                                         // Den Wert eins in den Channel schreiben, und ...
	<-b                                            // ... aus dem anderen Channel lesen (blockt)
	fmt.Println(name)                              // Wenn die andere Instanz geschrieben hat, kommen wir aus dem Block und schreiben unseren Namen
	return
}

func Duell() {
	var wg sync.WaitGroup

	aToB := make(chan int, 1) // Channel für die eine Richtung mit int als Typ
	bToA := make(chan int, 1) // Channel für die gegensätzliche Richtung mit int als Typ

	wg.Go(func() { // Die Funktion Go ist neu und erlaubt es eine Funktion zu instanziieren und den Wait Zähler gleichzeitig zu erhöhen
		shoot(aToB, bToA, "d'Artagnant") // Funktion die aufgerufen wird mit spezifischen Aufruf-Argumenten
	})

	wg.Go(func() { shoot(bToA, aToB, "Rochefort") })

	wg.Wait() // Wait Call, der blockiert, bis sämtliche (hier 2) Routinen der Gruppe terminiert sind
}

func randomRangeValue(min, max int) int {
	a, _ := rand.Int(rand.Reader, big.NewInt(100))
	return int(a.Int64())%(max-min+1) + min
}
