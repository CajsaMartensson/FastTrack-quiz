# Quiz

## För att starta API't:
**Skriv följande i terminalen**: go run main.go

## För att sedan starta programmet:
**Skriv följande i terminalen**: go run cmd/cli/main.go quiz

## API Endpoints
* `GET /questions` - Hämtar alla frågor och svarsalternativ.
* `POST /submitAnswers` - Skickar in användarens svar och returnerar poäng samt statistik.

## Struktur
* `main.go`: Serverns startpunkt och HTTP-handlers.
* `server/`: Datamodeller och in-memory "databas".
* `cmd/`: Innehåller CLI-logiken och Cobra-kommandon.
