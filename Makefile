
VERSION 	:= $(shell date +%Y%m%d)
DIST_51_TGZ 	:= luaman-5.1.tgz
DIST_52_TGZ 	:= luaman-5.2.tgz
PRUTS 		:= pruts.nl:/home/ico/websites/www.zevv.nl/play/code/luaman

all: 
	./go 5.1
	./go 5.2

clean:
	rm -f */man3/*

dist:
	tar -zcvf $(DIST_51_TGZ) 5.1/man3
	tar -zcvf $(DIST_52_TGZ) 5.2/man3

upload: dist
	scp *tgz $(PRUTS)


