
# RUN
go run cmd/weatherapp/main.go

# Weather APP

**WeatherAPI:** Docs | http://Open-Meteo.com\
**Spis miast z długością i szerokością geo:** Poland Cities Database | http://Simplemaps.com\
Producer-consumer pattern: https://jenkov.com/tutorials/java-concurrency/producer-consumer.html 

Rzeczy do poruszenia: współdzielona pamięć (mutex), ograniczenia w strzałach do api (semaphore)


Cel projektu:
Mając dane ~172 (api prosi o max 10k strzałów per dzień) polskie miasta w pliku pl172.json (na dysku obok zadania) przeanalizować dla nich pogodę z ostatniego półrocza.
W ramach analizy chcemy się dowiedzieć:
Chcemy poznać miasto z największą średnią temperaturą
Chcemy poznać miasto gdzie najczęściej występuje mgła (weather code 45) i miasto gdzie jest najczęściej czyste niebo (weather code 0)
Program na wejściu dostaje plik json z informacjami o miastach.
Wynikiem pracy programu powinien być plik results.json, w którym zawarte zostaną wyżej wymienione informacje.

Program powinien być w stanie zliczyć czas swojego działania, żeby móc porównać pod względem wydajności kolejne etapy. Wystarczy zwykłe liczenie czasu z biblioteką time

Etapy:

Rozwiązujemy problem jednowątkowo (jeśli następny wydałby się za trudny na początek, można skipnąć ten etap)
Tworzymy program, który posiada 1go producera i 1go consumera.
Program powinien nadal mieć 1go producera, ale konfigurowalną liczbę consumerów.
Program posiada konfigurowalną liczbę producerów i consumerów.
Etap poprzedni, ale tylko k producerów jednocześnie może strzelać do API pogodowego.
Etap poprzedni, ale wynik chcemy dostać po maksymalnym czasie x, gdzie x jest konfigurowalne.
    zwrocic stan po danym czasie
Etap poprzedni ale analizę robimy dla kilku krajów. (na pewno Paweł chcę żeby to zrobił)

Etapy oddajemy po kolei, po przyklepaniu przez prowadzącego lecimy z kolejnym. Fajnie by było przed rozpoczęciem pisania etapu zaproponować jak będzie on rozwiązywany (wystarczy ustnie przy okazji jak się widzimy, chciałbym wiedzieć przed pisaniem że jest wszystko jasne). Jeśli coś jest niedoprecyzowane to elementem zadania jest również dopytanie o szczegóły, które nie są oczywiste.
