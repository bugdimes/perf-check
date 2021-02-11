var start = new Date()

var fs = require('fs');

const data = fs.readFileSync('./angularjs.js', {encoding:'utf8', flag:'r'});

final_data = data.replace("var", "let");

fs.writeFileSync("result_node.js", final_data); 

var end = new Date() - start

console.log('time consumed:', end)

// 16ms