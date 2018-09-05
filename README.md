[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://godoc.org/github.com/BTWhite/go-btw-photon)
[![Go Report Card](https://goreportcard.com/badge/github.com/BTWhite/go-btw-photon?1)](https://goreportcard.com/report/github.com/BTWhite/go-btw-photon)

# BitWhite Photon
BitWhite is an efficient, flexible, and safe decentralized application platform designed to provide effortless development of 
decentralized applications. The root implementation was written in JavaScript. But in the future we plan to launch BTW on Golang and 
leave JS as an additional implementation.

### Photon
Photon - is a new Protocol based on asynchronous graphs, designed to solve problems with scalability and speed reduction of transactions.

The Protocol allows ordinary clients not to synchronize the entire blockchain, but only their history connected directly to their account 
and history, which leads to evidence of the legality of transactions that change the balance of the sender. This way we achieve a linear
transaction rate when chains are filled with their owners.

Thanks to the branching of the network, we are able to get rid of blocks, replacing them with snapshots. Snapshot - a Snapshot of the
network every 10 minutes, unlike blocks, they do not store information about transactions, but only record the balances of accounts at
the moment. Thus, we do not care how many transactions the sender will make during this time, because as a result only the final balance
will be in the picture. But despite the asynchronous operation of transactions, each member of the network is obliged to synchronize 
the Snapshots, supplemented with information about voting for delegates and freezing wallets Validators.

In order that the chains could not be changed, the Protocol received delegates and validators synchronizing the history of each chain. 
Despite the fact that the transaction is considered to be fully confirmed after the recording of delegates, you can consider it 
confirmed instantly, as to make a transfer, the sender must find several validators who will freeze the money from their personal 
accounts in advance and in case of fraud, the Protocol will reimburse the loss to the victim by taking the required amount from the 
validators.

To become a validator and split the сommission between yourself and the delegates enough to freeze the wallet to certain snapshot and 
then start the node with a special key. But keep in mind that freezing your coins will require a standard fee for any transaction and 
you must have some capital in order to be more popular, because the more the transfer amount the more you need to find validators' 
reserve coins.

A Photon is a new era for our community. We plan to fundamentally change the Protocol in order for the network to expand the transaction
speed as the number of nodes increases and also change the programming language to GO.

### Why GO?
It is difficult to find a language that would be both fast and simple enough for DAPP developers. Golang is the best option, because it 
runs tens of times faster than JavaScript and was developed by Google for its not very experienced programmers. Go is good not only with
a rich set of built-in libraries, but also by imposing a single style and recommendations for good code, just read about gofmt.

## JSON RPC
When you start the node, you start the http server (by default `8080` port), you can send `RPC` requests using `HTTP POST`. To do this, 
send a special object by `http://localhost:8080/jsonrpc/` (or your any `IP` and `port`).

Ping request:
```json
{
  "id": 3,
  "method": "ping",
  "params": "beep"
}
```
Ping response:
```json
{
  "id": 3,
  "result": "bup"
}
```
In the event of an error, instead of the result field, the error field returns:
```json
{
  "id": 3,
  "error":{
    "code": 0,
    "message": "what is beep? :/"
  }
}
```
Get information on available methods in [RPC-METHODS.md](https://github.com/BTWhite/go-btw-photon/blob/master/docs/RPC-METHODS.md)

## Contribution
You can also participate in writing go-btw, for this you need to get acquainted with the main 
[repository](https://github.com/BTWhite/BTWChain) and create their own pull requests. We adhere to the 
[standards of Golang](https://golang.org/doc/effective_go.html) and therefore you need to read about 
[formatting](https://golang.org/doc/effective_go.html#formatting) and [commentary](https://golang.org/doc/effective_go.html#commentary).
We will be happy if someone from the community will participate in the development. 

### Pull Request
* Optionally, create a new branch in your fork, although this is not a mandatory requirement.
* Change the `README.md` information about the interface changes if this is required.
* In any commit add a prefix with the name of the changed packages. For example: `chain, types: Fix beep-boop bug`
* Specify meaningful names of commits. Bad example: `chain: Update chain.go`

