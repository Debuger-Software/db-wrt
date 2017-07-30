#!/bin/bash
eval `cat .config | grep CONFIG_VERSION_`

NEV_VER=$((CONFIG_VERSION_CODE + 1))
CONFVER=$(dialog --title "KONTROLA WERSJI" --clear --inputbox "Potwierdź numer rev" 10 40 $NEV_VER 3>&1 1>&2 2>&3)
#HSNAME=$CONFIG_VERSION_DIST"_v"$CONFIG_VERSION_NUMBER"-r"$CONFVER
sed -i "s/CONFIG_VERSION_CODE=.*/CONFIG_VERSION_CODE=\"$CONFVER\"/" .config
#sed -i "s/option hostname.*/option hostname '$HSNAME'/" files/etc/config/system
echo Numer rev ustawiony na: $CONFVER
MODE=$(dialog --title "MAKE" --menu "Co robimy ?" 10 50 3 1 Kompilacja 2 Konfiguracja 3 Czyszczenie 4 TEST 3>&1 1>&2 2>&3)
case $MODE in
	1) #KOMPILACJA
		echo KOMPILACJA
		MAKEPARAMS=$(dialog --title "GADATLIWY MAKE" --clear --inputbox "Potwierdź parametry make" 10 40 "-j3 V=1" 3>&1 1>&2 2>&3)
		dialog --title "KOMPILACJA W TOKU" --infobox "Rozpoczynanie kompilacji..." 3 40
		time colormake $MAKEPARAMS
			dialog --title "Are you fuc**ng ready ?" --yes-label "DAWAJ !!!" --no-label "Peniam..." --yesno "Flashujemy nowy soft ?" 5 35
				case $? in
					0)
					IMG_NAME="db-wrt-$CONFIG_VERSION_NUMBER-$CONFVER-ar71xx-generic-tl-wr740n-v4-squashfs-sysupgrade.bin"
					IMG_PATH="/home/debuger/lede/bin/targets/ar71xx/generic/$IMG_NAME"
					if [ -f $IMG_PATH ]; then
						sshpass -p 'koalapaint7' scp put $IMG_PATH root@192.168.1.1:/tmp
						sshpass -p 'koalapaint7' ssh root@192.168.1.1 sysupgrade /tmp/$IMG_NAME
						clear && echo "SYSUPGRADE IN PROGRESS.... WAIT FOR DEVICE REBOOT"
						./dev-ready.sh
					else
						dialog --title "COMPILATION FUCK OFF !!!" --infobox "Kompilacja nie wyszła, nie ma obrazu $IMG_NAME" 4 75
					fi
						;;
					*)clear;;
				esac;;
	2) #KONFIG
		dialog --infobox "Uruchamianie panelu konfiguracji..." 3 40
		make menuconfig
		clear
		dialog --title "TESTUJEMY ?" --yes-label "NO BA!" --no-label "NI CHU..." --yesno "KOMPILOWAĆ Z NOWYMI USTAWIENIAMI ?" 5 38
			case $? in
				0) echo TAK
				MAKEPARAMS=$(dialog --title "GADATLIWY MAKE" --clear --inputbox "Potwierdź parametry make" 10 40 "-j3 V=1" 3>&1 1>&2 2>&3)
				dialog --title "KOMPILACJA W TOKU" --infobox "Rozpoczynanie kompilacji..." 3 40
				time colormake $MAKEPARAMS
				
				dialog --title "Are you fuc**ng ready ?" --yes-label "DAWAJ !!!" --no-label "Peniam..." --yesno "Flashujemy nowy soft ?" 5 35
				case $? in
					0)
					IMG_NAME="db-wrt-$CONFIG_VERSION_NUMBER-$CONFVER-ar71xx-generic-tl-wr740n-v4-squashfs-sysupgrade.bin"
					IMG_PATH="/home/debuger/lede/bin/targets/ar71xx/generic/$IMG_NAME"
					if [ -f $IMG_PATH ]; then
						sshpass -p 'koalapaint7' scp put $IMG_PATH root@192.168.1.1:/tmp
						sshpass -p 'koalapaint7' ssh root@192.168.1.1 sysupgrade /tmp/$IMG_NAME
						clear && echo "SYSUPGRADE IN PROGRESS.... WAIT FOR DEVICE REBOOT"
						./dev-ready.sh
					else
						dialog --title "COMPILATION FUCK OFF !!!" --infobox "Kompilacja nie wyszła, nie ma obrazu $IMG_NAME" 4 75
					fi
						;;
					*)clear;;
				esac;;
	
	
				*) echo NIE;;
			esac;;
	3) #SPRZATANIE
		echo SPRZATANIE
		dialog --infobox "Sprzątam bajzel po poprzedniej kompilacji..." 3 50
		make clean
		clear;;

	4) #PARAM TEST
		IMG_NAME="dbgr-os-$CONFIG_VERSION_NUMBER-$CONFVER-ar71xx-generic-tl-wr740n-v4-squashfs-sysupgrade.bin"
		IMG_PATH="/home/debuger/lede/bin/targets/ar71xx/generic/$IMG_NAME"
#		sshpass -p 'koalapaint7' scp put $IMG_PATH root@192.168.1.1:/tmp
		echo PARAMS TEST: $IMG_PATH;;

	*) #CANCEL
		echo CANCEL;;
esac




#colormake -j V=1



#for ((i = 0 ; i <= 100 ; i+=20)); do sleep 1; echo $i"\n\n"; done | while read -r line; do rrr=$rrr$line; dialog --infobox $rrr 0 0; done

