CURDIR=$(dirname $(readlink -f $0))/

go build -o $CURDIR../build $CURDIR../cmd/app

echo
echo "==> Results:"
ls -hl $CURDIR../build/