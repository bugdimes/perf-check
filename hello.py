import time
start = time.time()

resultant_file = open('result_python.js', 'w')

with open('angularjs.js') as file:
    data = file.read()

final_data = data.replace('var', 'let')
resultant_file.write(final_data)

resultant_file.close()

print((time.time()-start)*1000)

# 10ms
