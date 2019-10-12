build: bindata
	go build
	mkdir -v build
	mv -v feinschmecker build/
	cp -vR perl build/

bindata:
	go-bindata -pkg bindata -o bindata/bindata.go templates/...

clean:
	rm -rf bindata/
	rm -rf build/
