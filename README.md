# zkPoD-contract

Smart contracts for *zkPoD Decentralized Exchange*.

## ABI & Wrapper

### ABI JSON File

- [*zkPoDExchange.abi*](abi/zkPoDExchange.abi)

### Go Wrapper

- [*zkPoDExchange.go*](abi/zkPoDExchange.go)

## Usage & Development

### Build

```
git clone
npm install
truffle compile
```

### Generate ABI

```
npm run genABI
```

### Generate Go Wrapper

```
npm run genWrapper
```

### Test

```
npm run test
```

### Lint

```
npm run lint:contracts
```

### Deploy On Ropsten

```
export mnemonic=YOUR_SUPER_SECRET_PRIVATE_KEY_HERE
truffle migrate --network ropsten
```

## Live On Ethereum Networks

### Ropsten Testnet

- *PublicVar* - [0x51D24a878716A99cFBD4894272A233187903BB84](https://ropsten.etherscan.io/address/0x51D24a878716A99cFBD4894272A233187903BB84)
- *zkPoDExchange* - [0xC57C51d78Ec58371218e5856afDC9D05F8C2cBd4](https://ropsten.etherscan.io/address/0xC57C51d78Ec58371218e5856afDC9D05F8C2cBd4)

