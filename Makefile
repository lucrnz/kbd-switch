all: kbd-switch
.PHONY: install clean

kbd-switch:
	go build

install:
	cp -v kbd-switch /usr/bin/kbd-switch

clean:
	test -f kbd-switch && rm kbd-switch
