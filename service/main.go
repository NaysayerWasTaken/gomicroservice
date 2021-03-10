package main

import (
	"FirstService/service/handlers"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	//Definition eines Loggers fuer die "product-api"
	// l = Logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Initialisierung eines neuen Handlers fuer produkte, der den "globalen" Logger verwendet
	// Ein Handler ist zustaendig fuer die bearbeitung von requests bestimmter Art
	//(Hier alle die mit produkten zu tun haben)
	// ph = product handler
	ph := handlers.NewProduct(l)

	// EIn Mux kann man sich vorstellen als Router, der HTTP Requests weiterleitet
	// sm = serve Mux
	sm := mux.NewRouter()

	// Hier werden fuer jeden relevante HTTP Request Typ ein eigener Subrouter erstellt

	// Subrouter fuer die GET Methode, legt fest, dass die eingehenden GET Requests am Pfad "/" von der zugehoerigen
	//Methode im ProductHandler bearbeitet werden
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	// Subrouter fuer die PUT Methode, legt fest, dass die eingehenden PUT Requests am Pfad "/(beliebige valide id)"
	//von der zugehoerigen Methode im ProductHandler bearbeitet werden
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	// Subrouter fuer die POST Methode, legt fest, dass die eingehenden POST Requests am Pfad "/"
	//von der zugehoerigen Methode im ProductHandler bearbeitet werden
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProducts)
	postRouter.Use(ph.MiddlewareValidateProduct)

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
		ReadTimeout:  5 * time.Second,
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
	signal.Notify(c, os.Kill)

	// Im Fall, dass der Channel ein Signal erhaelt,
	// wird der Server Heruntergefahren und das erhaltene Signal gelogged
	sig := <-c
	l.Println("Recieved terminate, graceful shutdown", sig)

	// Deadline context unter dem der Server gestoppt wird
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Server nimmt keine neuen Anfragen mehr entgegen, bearbeitet die verbleibenden und schaltet sich dann ab
	// -> besser als abrupter stopp
	s.Shutdown(ctx)
}
