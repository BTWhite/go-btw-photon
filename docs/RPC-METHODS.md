## Transactions
### `tx.get`
Gets complete information about the transaction by its ID.

| Param      | Required | Type |Description |
|:----------:|:----------:|:----------:|-------------|
|`id`|+|`string`|Id transaction|

```json
{
	"id": "ecd23f34f831d40958a63d62aaf72dc375ad349e5d89306953ad7437b3c05c61"
}
```
```json
{
	"id": 0,
	"result": {
		"id": "ecd23f34f831d40958a63d62aaf72dc375ad349e5d89306953ad7437b3c05c61",
		"senderPublicKey": "0000000000000000000000000000000000000000000000000000000000000000",
		"senderId": null,
		"recipientId": "B0AoqtQfSZyfCex8q4fGwMAFEQbcDt4mZfP",
		"amount": 6000000000000000,
		"fee": 0,
		"signature": "0000000000000000000000000000000000000000000000000000000000000000",
		"timestamp": 1535500800,
		"nonce": 0,
		"height": 1,
		"chain": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		"previousTx": null
	}
}
```
### `tx.create`
Creates new tx (unsafe method).

| Param      | Required | Type |Description |
|:----------:|:----------:|:----------:|-------------|
|`secret`|+|string|Your secret by account|
|`address`|+|string|Recipient address|
|`ammount`|+|int|Ammount in `100000000` = `1` coin|

```json
{
	"secret": "HelloSubject",
	"address": "B0DAO",
	"amount":1500000000
}
```
```json
{
	"id": 0,
	"result": "573763c5934efc899c0f7eee1ee6cce7ceeb120686bc1fd9450a2853ce5c6fb9"
}
```

## Account
### `acc.open`
Gets account information by publicKey (safe)

| Param      | Required | Type |Description |
|:----------:|:----------:|:----------:|-------------|
|`publicKey`|+|string|Your publicKey by account|

```json
{
	"publicKey": "02a09c2e80f2701fb277ccb29f6919c97cc4743b7e7d984685cd2e31769f1871"
}
```
```json
{
	"id": 0,
	"result": {
		"address": "B0AoqtQfSZyfCex8q4fGwMAFEQbcDt4mZfP",
		"publicKey": "02a09c2e80f2701fb277ccb29f6919c97cc4743b7e7d984685cd2e31769f1871",
		"balance": 5999997000000000
	}
}
```
