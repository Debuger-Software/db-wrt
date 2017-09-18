#!/bin/bash


function updatevariables() {
	eval `cat $HOME/db-wrt/.config | egrep -i 'config_(target_(board|subtarget|profile)=|version_(dist|number|code)=)'`;
	typeset -l IMG_NAME="$CONFIG_VERSION_DIST-$CONFIG_VERSION_NUMBER-$CONFIG_VERSION_CODE-$CONFIG_TARGET_BOARD-$CONFIG_TARGET_SUBTARGET-${CONFIG_TARGET_PROFILE#*_}-squashfs-sysupgrade.bin";
	typeset -l BIN_PATH="$HOME/db-wrt/bin/targets/$CONFIG_TARGET_BOARD/$CONFIG_TARGET_SUBTARGET/";
	IMG_PATH="$BIN_PATH$IMG_NAME";
}


function image_exists() { return $(![[ -e $IMG_PATH ]]); }




updatevariables

echo "Test for: $IMG_PATH";
if ( image_exists ); then echo "OK"; fi