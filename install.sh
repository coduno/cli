#!/bin/bash

LOCATION=${LOCATION:='/usr/local/bin'}
TMPFILE='/tmp/coduno'

if [ $(uname) != 'Linux' ]
then
	echo 'Your operating system was not recognized as Linux, refusing to install.'
	exit 1
fi

ARCH=386

if [ $(uname -m | grep 64) ]
then
	ARCH=amd64
fi

URL=https://drone.io/github.com/coduno/cli/files/linux-$ARCH/coduno

if [ ! -f $TMPFILE ]
then
	curl --output $TMPFILE $URL
	if [ $? ]
	then
		echo ''
		echo 'The command line client for Coduno was downloaded.'
	else
		echo "Error $? occured while downloading the binary."
		exit 2
	fi
else
	echo "The installer found $TMPFILE and assumes that this is the current"
	echo 'version of the command line client for Coduno. Download was skipped.'
fi


SHA=$(shasum $TMPFILE | cut -d' ' -f1)
MD5=$(md5sum $TMPFILE | cut -d' ' -f1)
SIZE=$(stat -c %s $TMPFILE)

cat << STOP

Now would be a good time to check integrity against

	https://drone.io/github.com/coduno/cli/files

	linux-$ARCH/coduno
	$SIZE bytes
	SHA $SHA
	MD5 $MD5

The installer will now attempt to move the downloaded
binary to

	$LOCATION

STOP

while true; do
	read -p "Is that okay? Please answer 'yes' or 'no': " ack < /dev/tty
	case $ack in
		'yes' )
			mv -v $TMPFILE $LOCATION/coduno
			chmod -v u+x $LOCATION/coduno
			break;;
		'no' )
			rm -vf $TMPFILE
			echo 'Aborting installation as requested.'
			exit;;
	esac
done
