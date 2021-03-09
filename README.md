# gomicroservice

Eine Art Sandbox eines Go Microservices

Um das Ganze ausführen zu können Braucht ihr natürlich Go, bekommt man hier her:

https://golang.org/

Ihr solltet Go noch zu eurer Path-Variable hinzufügen, um das Terminal benutzen zu können: 

In Windows:

 - Rechtsclick "Dieser PC" (File Explorer)
 - Eigenschaften
 - Erweiterte Systemeinstellungen
 - Umgebungsvariablen
 - Bei Systemvariablen zu Path navigieren
 - Bearbeiten
 - Neu
 - Den \bin Ordner eurer Go installation angeben (z.B. C:\Go\bin)
 - Neues Terminal aufmachen und "go version" ausführen, um zu sehehn obs funktioniert hat

Als IDE ist zu empfehlen Entweder VSCode mit der Go Extension, GoLand von JetBrains, oder IntelliJ mit go Extension

Um das ganze auszuführen navigiert ihr zum Ordner mit der "main.go" und führt "go run main.go" aus

Dann passiert scheinbar gar nix.

Die Anwendung kann zum Stand auf master einen HTTP-Server starten der auf zwei messages hört oder sauber runterfährt.

Heißt, um zu sehen was das ganze macht wenns läuft muss man per curl, Postman oder auf anderem Wege HTTP-Nachrichten auf Port 9090 schicken.

Bei einer Nachricht an localhost:9090 antwortet die Anwendung mit "Message Received!". 
Bei einer Nachricht an localhost:9090/goodbye erhält der Sender die Nachricht "Goodbye" zurück.
Bei Beenden der Anwendung wird der Server gestoppt indem er aufhört neue Anfragen entgegenzunehmen, die Existierenden abarbeitet und dann herunterfährt.

-> erkennbar an den Konsolenausgaben: "product-api2021/03/09 16:42:36 Recieved terminate, graceful shutdown interrupt
                                       product-api2021/03/09 16:42:36 http: Server closed"
