# Krok 1: Budowanie aplikacji
FROM golang:1.23 AS builder

# Ustawienie katalogu roboczego
WORKDIR /app

# Skopiowanie plików projektu
COPY app .

# Pobranie zależności
RUN go mod download

# Kompilacja aplikacji jako statyczna binarka
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ise-mock

# Krok 2: Tworzenie finalnego obrazu scratch
FROM scratch

# Skopiowanie skompilowanej aplikacji
COPY --from=builder /app/ise-mock /ise-mock

# Ustawienie portu jako argumentu (można nadpisać przy uruchamianiu kontenera)
ENV SERVER_ADDRESS=0.0.0.0
ENV SERVER_PORT=8080

# Ustawienie domyślnego portu
EXPOSE 8080

# Uruchomienie aplikacji
CMD ["/ise-mock", "-s", "$SERVER_ADDRESS", "-p", "$SERVER_PORT"]