const fs = require('fs');
const contract = JSON.parse(fs.readFileSync('./build/contracts/zkPoDExchange.json', 'utf8'));
// console.log(JSON.stringify(contract.abi));
let path = 'abi/zkPoDExchange.abi';
fs.writeFileSync(path, JSON.stringify(contract.abi, null, "\t"));

console.log(`ABI extracted and output file wrote to: ${path}`);