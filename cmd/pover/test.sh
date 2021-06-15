#!/bin/sh

if date -n | cmd/pover/pover
then
	echo 'date sent ok'
else
	echo 'date failed'
fi

if /bin/dd if=/dev/urandom of=/dev/stdout count=1024 | cmd/pover/pover
then
	echo 'max length message ok'
else
	echo 'max length message failed'
fi

# we expect pover to print a warning to standard error but send a truncated message anyway.
if /bin/dd if=/dev/urandom of=/dev/stdout count=2048 | cmd/pover/pover
then
	echo 'too long message ok'
else
	echo 'too long message failed'
fi

if ! echo | cmd/pover/pover
then
	echo 'blank message ok'
else
	echo 'blank message failed'
fi

if ! cmd/pover/pover -f /dev/null
then
	echo "empty config ok"
else
	echo "empty config failed"
fi

badconfig=`mktemp`
echo 'badconfig' > $badconfig
if ! cmd/pover/pover -f $badconfig
then
	echo "bad config ok"
else
	echo "bad config failed"
fi
rm $badconfig
