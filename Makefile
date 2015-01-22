
LUAREF_URL ?= http://www.lua.org/manual/5.3/manual.html

all:
	lua5.3 ./luaref2man $(LUAREF_URL)

clean:
	rm -rf man3

dist:
	tar -zcf /tmp/luaman-5.3.tgz man3
