start = Time.now

file = File.open("angularjs.js")

file_data = file.read

file_data.gsub! 'var', 'let'

File.write("result_ruby.js", file_data)

finish = Time.now

puts (finish - start) * 1000