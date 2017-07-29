#!/bin/bash


dialog --sleep 40 --title "FLASHOWANIE PAMIĘCI MTD W TOKU . . ." --infobox "" 1 36




online=0
(
while [ $online -lt 10 ]
	do ping -c 1 -W 1 192.168.1.1 >/dev/null
	case $? in
		0) online=$(($online+1));;
		2) online=0 && sleep 1;;
		*) online=0;;
	esac
	echo $(($online*10))
	sleep 1
done
) | dialog --title "OCZEKIWANIE NA ROUTER" --beep-after --gauge "\n\nPróba nawiązania stabilnego(!) połączenia z routerem..." 10 60 0

clear
echo "Exit with online: $online"
	
