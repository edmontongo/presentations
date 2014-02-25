require 'coderay'
File.open('coderay.html','w') { |f| f.write(
	CodeRay.scan_file('iteration.pl', :perl, {coderay_line_numbers: ''}).page(:line_numbers => nil)
) }
system("chromium-browser coderay.html")
