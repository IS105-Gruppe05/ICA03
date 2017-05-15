# is105-ica03

Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.
- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor kan det vise at det er bare 5/8 som har lastet opp noe. 


Repository for IS-105 experiments with encoding of signs and symbols.

![Alt Text](url)
Ica 3. Gruppe 5.
Oppgave 1: ascii_main.go
a)
![Bilde1](https://i.gyazo.com/7a8bcacfacde867cd4e9cf5c352a68fc.png)
Ved å sjekke tabellen på ascii-code.com fant vi ut at de første 31 karakterene i ASCII-tabellen er kontroll-karakterer som ikke kan printes ut. Et eksempel er #24 i lista, hex 18 som er Cancel.
![Bilde2](https://i.gyazo.com/f0177cd8b6d6bcea6349d63eb33bba9b.png)
Her er bilde av de 16 første karaktere i ASCII tabellen. 

b) ![Bilde3](https://i.gyazo.com/83e515b1437e4428055e6ce9979cd1f9.png)
Her kan vi se at tegn som “ og mellomrom også har sine egne ASCII-verdier, som vi kan bruke for å hardkode en setning. I dette tilfellet er det standard ASCII, som vil si at vi ikke trenger å encode setningen noe mer utover dette. 

c) ascii_test.go
Hver karakter har en verdi. I standard ASCII er dette fra 1-127, og utvidet ASCII er over 127. Vi lagde en test som kontrollerer om verdien på karakterene er under 127, siden alle karakterene i funksjonen er i standard ASCII så består testen.
VI har brukt “range” koden slik den kan finne ut av hvor den er i lista og lagrer den lista den har gått igjennom så langt. Hadde vi bare brukt en for-løkke ville koden ikke visst hvor den var  i lista. 
![Bilde4](https://i.gyazo.com/06b39d9d8780ab306edb79cb4832a91d.png)
Vi prøvde å sette inn tegn som ikke var ascii, men da fikk vi bare tilbakemelding om at dette var “Feil”, og at det ikke var riktige ascii tegn. Det er en feilkode som vi selv har kodet inn for å hjelpe oss i å forstå hva som  er feil i koden.









Oppgave 2:isoMain.go
a)

