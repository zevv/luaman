
LUAREF_URL ?= http://www.lua.org/work/doc/manual.html

all:
	lua5.3 ./luaref2man $(LUAREF_URL)

clean:
	rm -rf man3
