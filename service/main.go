package main

import (
	"FirstService/service/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// definition eigener Handler fuer API Verbindungen zum Server
	// hh = Hello Handler
	// gh = goodbye Handler
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// Ein Mux Funktioniert wie ein router
	//-> routed die API anfragen ausgehend von der Signatur an verschiedene Handler
	// sm = ServeMux

	sm := http.NewServeMux()

	//Eine API Signatur gehoert immer zu einem Handler, welcher diese bearbeitet

	// hier werden die API signaturen den vorher definierten Handlern zugewiesen
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)


	// Definition eines simplen HTTP Servers in GO
	// s = Server
	// Hier koennen verschiedene Eigenschaften des Servers definiert werden
	// Fuer uns sind z.B. Performanz relevant, da die verpflichtende umsetzung von TLS
	// andernfalls sehr die Serverperformance beeintraechtigt
	// Zusaetzlich ist fuer uns wichtig uns vor DOS-Attacken oder "toten" Clients zu schuetzen,
	// indem Beschraenkungen bei den Verbindungen festgelegt werden

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}

	// Server wird eingeschaltet
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Definiert einen neuen Kanal fuer Signale, der in diesem Fall die Signale des Betriebssystems bekommt
	sigChan := make(chan os.Signal)
	// im definierten Signal Channel landen dadurch die Interrupt/Kill Signale des Betriebssystems
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Im Fall, dass der Channel ein Signal erhaelt,
	// wird der Server Heruntergefahren und das erhaltene Signal gelogged
	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	// Deadline context unter dem der Server gestoppt wird
	tc, _ := context.WithTimeout(context.Background(), 30 *time.Second)

	// Server nimmt keine neuen Anfragen mehr entgegen, bearbeitet die verbleibenden und schaltet sich dann ab
	// -> besser als abrupter stopp
	s.Shutdown(tc)
}
