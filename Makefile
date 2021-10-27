# Compilation
mc-whitelist-bot: go.sum main.go */**/*.go .env */**/**/*.go
	CGO_ENABLED=0 ./scripts/build.sh

mc-whitelist-bot.exe : go.sum main.go */**/*.go .env */**/**/*.go
	env CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 ./scripts/build.sh

# fetch dépendencies
go.sum : go.mod
	go get
	go mod tidy

.PHONY: install deb docs win all clean manuals
# require write rights for $(DESTDIR)
install : mc-whitelist-bot
	mkdir -p $(DESTDIR)/usr/sbin/
	cp ./mc-whitelist-bot $(DESTDIR)/usr/sbin/mc-whitelist-bot

all: deb win linux manuals
	./scripts/publish.sh
	
deb :
	./scripts/deb.sh
win :
	./scripts/win.sh
linux : 
	./scripts/linux.sh

docs : */**/*.go
	mkdir -p docs
	godoc-static --destination ./docs .

manuals :
	./scripts/manuals.sh

clean : 
	rm -f mc-whitelist-bot go.sum
	rm -rf package/
	rm -rf publish/