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

	fd_out = io.open(version .. "/man3/" .. name .. ".3", "w")
	w('.TH %s 3 "2012" "Lua ' .. version .. '" "Lua ' ..version .. ' manual"', name:upper())
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
	desc = desc:gsub("</?a ?[^>]*>", "")
	desc = desc:gsub("<p>", "\n.P\n")
	desc = desc:gsub("</p>", "")
	desc = desc:gsub("</?ul>", "")
	desc = desc:gsub("<li><b><code>([^<]+)</code>: </b>", "\n.IP \\[bu]\n%1 ")
	desc = desc:gsub("<li><b>'<code>(.-)</code>': </b>", "\n.IP \\[bu]\n'%1' ")
	desc = desc:gsub("<li>(.-)</li>", "\n.IP \\[bu]\n%1")
	desc = desc:gsub("</li>", "")
	desc = desc:gsub("<code>(.-)</code> *", "\n.I %1\n")
	desc = desc:gsub("<b>(.-)</b> *", "\n.B %1")
	desc = desc:gsub("<em>(.-)</em> *", "\n.B %1\n")
	desc = desc:gsub("<pre>(.-)</pre>", "\n.nf\n%1.fi\n")
	desc = desc:gsub("\n+", "\n")
	desc = desc:gsub("'", "\\(cq")
	desc = desc:gsub("\n%.", "\n\\.")

	w("%s", desc)
	
end




-- vi: ft=lua ts=3 sw=3

