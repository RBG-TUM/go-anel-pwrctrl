# go-anel-pwrctrl
go wrapper for ip power outlets (ANEL-Elektronik NET-PwrCtrl)

## How the API works
The Power outlets api is very rudimentary. It's status is reported through a csv like (; seperated) string similar to this:

|Device type|Device name|Device IP|Device netmask|Gateway?|MAC-addr|?|?|?|?|Outlet name 1|Outlet name 2|...||| | | | | | | | outlet status 1| outlet status 2 |... | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | |  
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|NET-PWRCTRL_04.6|NET-CONTROL |123.4.5.6|255.255.255.0|123.4.5.254|00:01:02:03:04:05|06 | |H | |Rote Lampe (Rec)|Dell Monitor|Nr. 3|Nr. 4|Nr. 5|Nr. 6|Nr. 7|Nr. 8| | |1 |0 |0 |0 |0 |0 |0 |0 | | |0 |0 |0 |1 |1 |1 |1 |1 | | |an vom Browser|aus vom Browser|aus vom Browser|Anfangsstatus|Anfangsstatus|Anfangsstatus|Anfangsstatus|Anfangsstatus| | | | | | | | | | |end|NET - Power Control|

Index 20+i tells us whether outlet i is on or not.