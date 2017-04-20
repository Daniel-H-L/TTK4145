# Exercise 9
## Task 1
1. Prioritet i sanntidssammenheng avhenger av midlertidige ressurskrav, ikke hvor viktig oppgaven er for at systemet skal oppnå/opprettholde riktig funksjonalitet. Oppgaver blir tilegnet prioritet for å kunne avgjøre hvilke oppgaver som skal gjennomføres først dersom flere ber om de samme ressursene samtidig. 
2. For at en planlegger skal kunne brukes på sanntidssystemer må den være forutsigbar og vi må kunne benytte tester for å kunne verifisere at alle oppgaver blir fullført innenfor deadline. 

## Task 2
1. Uten arv:
| Task\Time | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10| 11| 12| 13| 14|
|-----------|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---| 
| a | | | | |E| | | | | | |Q|V|E| | 
| b | | |E|V| |V|E|E|E| | | | | | | 
| c |E|Q| | | | | | | |Q|Q| | | |E|

2. Med arv:
| Task\Time | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10| 11| 12| 13| 14|
|-----------|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---| 
| a | | | | |E| | |Q| |V|E| | | | | 
| b | | |E|V| | | | |V| | |E|E|E| | 
| c |E|Q| | | |Q|Q| | | | | | | |E|

## Task 3
1. Priority inversion: når en oppgave er avhengig av ressurser som benyttes av en oppgave med lavere prioritet. 
Unbounded priority inversion: en oppgave med høy prioritet kan ende opp med å vente uendelig lenge på at en ressurs skal bli frigjort av en oppgave med lavere prioritet.
2. Prioritetsarv unngår ikke deadlocks.

## Task 4
1. (11.2.4) Simple task model
* The application is assumed to consist of a fixed set of tasks. Reasonable in theory, maybe a challenge in practice. 
* All tasks are periodic, with known periods. Realistic. 
* The tasks are completely independent of each other. Can be difficult to realize. 
* All system overheads, context-switching times and so on are ignored (that is, assumed to have zero cost). 
* All tasks have deadlines equal to their periods. Sounds fair. 
* All tasks have fixed worst-case execution times. Is this possible?
* No task contains any internal suspension points. Can be difficult to realize.
* All tasks execute on a single processor. 

2. 
|Task | Utilization | Total |
|-----|-------------|-------|
| a | 0.3   |
| b | 0.33  |
| c | 0.25  |
|   | 0.8833|

From equation (11.1): 3*(2^(1/3) - 1) = 0.7798 > 0.693 
The test fails, so the taks set may not be schedulable. 

3. Response time analysis: 
* Task c:
w0 = 5 => Rc = 5 <= 20, ok
* Task b:
w0 = 10
w1 = 10 + ceil(10/20)*5 = 15
w2 = 10 + ceil(15/20)*5 = 15
=> Rb = 15 <= 30, ok
* Task a:
w0 = 15
w1 = 15 + ceil(15/30)*10 + ceil(15/20)*5 = 15 + 10 + 5 = 30
w2 = 15 + ceil(30/30)*10 + ceil(30/20)*5 = 15 + 10 + 10 = 35
w3 = 15 + ceil(35/30)*10 + ceil(35/20)*5 = 15 + 20 + 10 = 45
w4 = 15 + ceil(45/30)*10 + ceil(45/20)*5 = 15 + 20 + 15 = 50
w5 = 15 + ceil(50/30)*10 + ceil(50/20)*5 = 15 + 20 + 15 = 50
=> Ra = 50 <= 50, ok






