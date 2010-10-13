#!/usr/bin/lua

fd_out = nil

function w(...)
	fd_out:write(string.format(...) .. "\n")
end

fd = io.open("manual.html")
html = fd:read("*a")
fd:close()



while true do
	name, html = html:match('<a name=".-"><code>(.-)</code></a></h3>(.*)')

	print(name)


	fd_out = io.open("man3/" .. name .. ".3", "w")
	w('.TH %s 3 "2010" "Lua 5.1" "Lua 5.1 manual"', name:upper())
	w('.SH NAME')
	w(name)

	api = nil
	api, tmp = html:match('^<p>\n<span class="apii">(.-)</span>(.*)')
	if api then
		api = api:gsub("<em>(.-)</em>", "%1")
		html = tmp
	end

	proto, html = html:match("<pre>(.-)</pre>(.*)")
	w('.SH SYNOPSIS')
	for l in proto:gmatch("([^\n]+)") do
		w('.B " %s', l)
		w('.br')
	end
	w('.br\n')
	if api then
		w('%s\n', api)
		w('.br\n')
	end
	
	desc, html = html:match("(.-)<h3>(.*)")
	w('.SH DESCRIPTION')

	if not desc then break end

	desc = desc:gsub("<hr>", "")
	desc = desc:gsub("&sect;", "#")
	desc = desc:gsub("&nbsp;", " ")
	desc = desc:gsub("<a.->(.-)</a>", "%1")
	desc = desc:gsub("<p>", "\n.br\n")
	desc = desc:gsub("</?ul>", "")
	desc = desc:gsub("<li>(.-)</li>", "%1")
	desc = desc:gsub("<b>(.-)</b> *", "\n.B %1\n")
	desc = desc:gsub("<em>(.-)</em> *", "\n.B %1\n")
	desc = desc:gsub("<pre>(.-)</pre>", ".nf%1.fi")
	desc = desc:gsub("<code>(.-)</code> *", "\n.I %1\n")
	desc = desc:gsub("\n\n", "\n")

	w("%s", desc)
	
end




-- vi: ft=lua ts=3 sw=3

