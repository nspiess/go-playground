package concurrency

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) { // Dies ist die Worker Funktion, jeder Worker simuliert mit dem Sleep Arbeit die getan wird
	for j := range jobs {
		fmt.Printf("worker %d startet Job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d stoppt Job %d\n", id, j)
		results <- j * 2
	}
	fmt.Printf("worker %d beendet\n", id)
}

func Worker() {
	const numJobs = 5                  // Anzahl der Jobs insgesamt
	jobs := make(chan int, numJobs)    // Channel für die Jobs
	results := make(chan int, numJobs) // Channel für die Resultate

	for w := 1; w <= 3; w++ { // 3 Worker werden gestartet und mit den Referenzen auf die Channels versorgt
		go worker(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // schliesst den Channel und gibt worker frei, die auf channel hören
	// listener wissen, dass nichts mehr kommt, weil channel entsprechend leer wird

	for a := 1; a <= numJobs; a++ {
		<-results
	}

	fmt.Println("Alle Jobs fertig")
}
