CURDIR=$(dirname $(readlink -f $0))/

go build -o $CURDIR../build/go-commerce $CURDIR../cmd/app

echo
echo "==> Results:"
ls -hl $CURDIR../build/