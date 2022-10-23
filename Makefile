default:
	go build -o hostloc *.go

amd64linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hostloc
	tar zcvf hostloc-check-in-linux-amd64.tar.gz hostloc example.json README.md

arm64linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o hostloc
	tar zcvf hostloc-check-in-linux-arm64.tar.gz hostloc example.json README.md

amd64windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o hostloc.exe
	tar zcvf hostloc-check-in-windows-amd64.tar.gz hostloc.exe example.json README.md

amd64mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o hostloc
	tar zcvf hostloc-check-in-macOS-amd64.tar.gz hostloc example.json README.md

all:
	make amd64linux
	make arm64linux
	make amd64windows
	make amd64mac

clean:
	rm hostloc* -r
