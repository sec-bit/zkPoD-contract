# zkPoD-contract

Smart contracts for *zkPoD Decentralized Exchange*.

## Related zkPoD projects

- [zkPoD-node](https://github.com/sec-bit/zkPoD-node) Node application written in Golang for sellers (Alice) and buyers (Bob). It deals with communication, smart contract calling, data transferring, and other zkPoD protocol interactions.
- [zkPoD-lib](https://github.com/sec-bit/zkPoD-lib) zkPoD core library written in C++ shipping with Golang bindings.

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

- *PublicVar* - [0xCf06Ff284988ebDc00752bD5E87AA78f853357d7](https://ropsten.etherscan.io/address/0xCf06Ff284988ebDc00752bD5E87AA78f853357d7)
- *zkPoDExchange* - [0x6BB3240C7a0850E99aEab96e210cbbD8b7B609D0](https://ropsten.etherscan.io/address/0x6BB3240C7a0850E99aEab96e210cbbD8b7B609D0)

