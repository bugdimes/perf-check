import time
start = time.time()

resultant_file = open('result_python.js', 'w')

with open('angularjs.js') as file:
    for each_line in file:
        resultant_file.write(each_line.replace('var', 'let'))

resultant_file.close()

print(time.time()-start)

# 22ms
