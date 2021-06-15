#!/bin/sh

while getopt 'f:' flag
do
	case $flag in
	f) config=$OPTARG ;;
	*) echo "unknown flag" $flag ;;
	esac
done

pover -f $config <<EOF
$NOTIFICATIONTYPE

$SERVICEDISPLAYNAME - $SERVICEDESC - on $HOSTDISPLAYNAME ($HOSTADDRESS) is $SERVICESTATE
Time: $LONGDATETIME
Output: $SERVICEOUTPUT
Comments: [$NOTIFICATIONAUTHORNAME] $NOTIFICATIONCOMMENT
EOF
