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


Zum Programm an sich:


Um das ganze auszuführen navigiert ihr zum Ordner mit der "main.go" und führt "go run main.go" aus

Dann passiert scheinbar gar nix.

Die Anwendung kann zum Stand auf master einen HTTP-Server starten der auf zwei messages hört oder sauber runterfährt.

Heißt, um zu sehen was das ganze macht wenns läuft muss man per curl, Postman oder auf anderem Wege HTTP-Nachrichten auf Port 9090 schicken.

- GET: localhost:9090 
        -> dadurch bekommt man alle (gehardcodeten vorhandenen Produkte als antwort)
- PUT: localhost:9090/1 {"name": "Neues Produkt"}
        -> dadurch wird der name des Produkts mit der ID 1 zu "Neues Produkt"
- POST: localhost:9090 {"name": "Neues Produkt", "price": "1.20", "sku": "abc-abc-abc"}
        -> dadurch wird ein Produkt mit den in der JSON definierten Parametern erstellt
  (Über eine weitere GET Request lassen sich die Änderungen überprüfen)
