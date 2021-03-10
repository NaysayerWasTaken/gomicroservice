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
	ph := handlers.NewProducts(l)

	// Ein Mux Funktioniert wie ein router
	//-> routed die API anfragen ausgehend von der Signatur an verschiedene Handler
	// sm = ServeMux

	sm := http.NewServeMux()

	//Eine API Signatur gehoert immer zu einem Handler, welcher diese bearbeitet

	// hier werden die API signaturen den vorher definierten Handlern zugewiesen
	sm.Handle("/", ph)

	// Definition eines simplen HTTP Servers in GO
	// s = Server
	// Hier koennen verschiedene Eigenschaften des Servers definiert werden
	// Fuer uns sind z.B. Performanz relevant, da die verpflichtende umsetzung von TLS
	// andernfalls sehr die Serverperformance beeintraechtigt
	// Zusaetzlich ist fuer uns wichtig uns vor DOS-Attacken oder "toten" Clients zu schuetzen,
	// indem Beschraenkungen bei den Verbindungen festgelegt werden

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Server wird eingeschaltet
	go func() {

		l.Println("Starting Server on Port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Definiert einen neuen Kanal fuer Signale, der in diesem Fall die Signale des Betriebssystems bekommt
	c := make(chan os.Signal)
	// im definierten Signal Channel landen dadurch die Interrupt/Kill Signale des Betriebssystems
	signal.Notify(c, os.Interrupt)

	// Im Fall, dass der Channel ein Signal erhaelt,
	// wird der Server Heruntergefahren und das erhaltene Signal gelogged
	sig := <-c
	l.Println("Recieved terminate, graceful shutdown", sig)

	// Deadline context unter dem der Server gestoppt wird
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Server nimmt keine neuen Anfragen mehr entgegen, bearbeitet die verbleibenden und schaltet sich dann ab
	// -> besser als abrupter stopp
	s.Shutdown(tc)
}
