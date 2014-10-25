#!/usr/bin/lua

fd_out = nil

version = arg[1]

function w(...)
	fd_out:write(string.format(...) .. "\n")
end

fd = io.open(version .. "/manual.html")
html = fd:read("*a")
fd:close()



while true do
	name, html = html:match('<a name=".-"><code>(.-)</code></a></h3>(.*)')
	
	html = html:gsub("&sect;", "#")
	html = html:gsub("&lt;", "<")
	html = html:gsub("&gt;", ">")
	html = html:gsub("&amp;", "&")
	html = html:gsub("&#8211;", "-")
	html = html:gsub("&nbsp;", " ")

	print(name)

	os.execute("mkdir -p " .. version .. "/man3")
	fd_out = io.open(version .. "/man3/" .. name .. ".3", "w")
	-- hardcoded to Lua-5.2 release date
	w('.Dd December 16, 2011')
	w('.Dt ' .. name:upper() .. ' 3')
	w('.Os')
	w('.Sh NAME')
	w('.Nm ' .. name)
	-- XXX: w('.Nd ' .. description)

	api = nil
	api, tmp = html:match('^<p>\n<span class="apii">(.-)</span>(.*)')
	if api then
		api = api:gsub("<em>(.-)</em>", "%1")
		html = tmp
	end

	proto, html = html:match("<pre>(.-)</pre>(.*)")
	w('.Sh SYNOPSIS')
	for l, m, o in proto:gmatch("([^ ]+)[ ]+([^%( ]+)[ ]*%(([^\n]+)") do
		w('.Ft ' .. l)
		w('.Fo ' .. m)
		for arg in o:gmatch("([^,%)]+)[,%) ]+") do
			w('.Fa "' .. arg .. '"')
		end
		w('.Fc')
	-- XXX: replace (type arg, type2 arg2, ...., typen argn)
	-- with "type arg" "type2 arg2" ... "typen argn"
	end
--	XXX: not sure what this is
--	w('.br\n')
--	if api then
--		w('%s\n', api)
--		w('.br\n')
--	end
	
	desc, html = html:match("(.-)<h3>(.*)")
	w('.Sh DESCRIPTION')

	if not desc then break end

	desc = desc:gsub("<hr>", "")
	desc = desc:gsub("</?a ?[^>]*>", "")
	desc = desc:gsub("<p>", "\n.Pp\n")
	desc = desc:gsub("</p>", "")
	desc = desc:gsub("<ul>", ".Bl -bullet -offset indent")
	desc = desc:gsub("</ul>", ".El")
	desc = desc:gsub("<li><b><code>([^<]+)</code>: </b>", "\n.It\n.Dl %1 ")
	desc = desc:gsub("<li><b>'<code>(.-)</code>': </b>", "\n.It\n.Dl '%1' ")
	desc = desc:gsub("<li>(.-)</li>", "\n.It\n%1")
	desc = desc:gsub("</li>", "")
	desc = desc:gsub("'<code>(.-)</code>' *", "\n.Sq %1\n")
	desc = desc:gsub("<code>(.-)</code> *", "\n.Dl %1\n")
	--desc = desc:gsub("<b>(.-)</b> *", "\n.B %1")
	--desc = desc:gsub("<em>(.-)</em> *", "\n.B %1\n")
	desc = desc:gsub("<pre>(.-)</pre>", ".Bd -literal\n%1\n.Ed\n")
	desc = desc:gsub("\n+", "\n")
	--desc = desc:gsub("'", "\\(cq")
	--desc = desc:gsub("\n%.", "\n\\.")

	w("%s", desc)
	
end




-- vi: ft=lua ts=3 sw=3

