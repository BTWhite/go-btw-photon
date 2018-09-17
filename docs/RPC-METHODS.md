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

## Snapshot
### `snapshot.get`
Gets snapshot data by Id or Height

| Param      | Required | Type |Description |
|:----------:|:----------:|:----------:|-------------|
|`id`|[one of the possible]|string|Id snapshot|
|`height`|[one of the possible]|height|Height snapshot|
```json
{
	"id":"e6662654523c5726db7b9e61347dc9f7cfa513271c8be2e2e8b3be4e6f2edb48"
}
```
```json
{
	"id": 0,
	"result": {
		"version": 0,
		"id": "e6662654523c5726db7b9e61347dc9f7cfa513271c8be2e2e8b3be4e6f2edb48",
		"height": 15,
		"previousSnapShot": "c1a435f71b5b876c6e93f4b55d7a852a7c55b37f41538bc15defe362c723df67",
		"generatorPublicKey": "02a09c2e80f2701fb277ccb29f6919c97cc4743b7e7d984685cd2e31769f1871",
		"votes": null,
		"balances": [
			{
				"addr": "B06",
				"blnc": 2500000000000,
				"ltx": "b8921b0bb99ba2ec69b1585c72219f5f97b131487b67ec144d4dd57ab978cc1f"
			},
			{
				"addr": "B0AoqtQfSZyfCex8q4fGwMAFEQbcDt4mZfP",
				"blnc": 5995000000000000,
				"ltx": "b8921b0bb99ba2ec69b1585c72219f5f97b131487b67ec144d4dd57ab978cc1f"
			}
		],
		"timestamp": 1536580477,
		"signaturess": null,
		"signature": "592cef2c41de9eca543cfd110dd32d22b4c93d8e05c7f86b247efaca4473fd4b361938d83623c93b84fa2eb4afc79f5e7506b98b5cc172e0bb49a6d4048f4405"
	},
	"peer": {
		"ip": "127.0.0.1",
		"port": 8080
	}
}
```

### `snapshot.last`
Gets last snapshot data

| Param      | Required | Type |Description |
|:----------:|:----------:|:----------:|-------------|
```json
{}
```
```json
{
	"id": 0,
	"result": {
		"version": 0,
		"id": "e6662654523c5726db7b9e61347dc9f7cfa513271c8be2e2e8b3be4e6f2edb48",
		"height": 15,
		"previousSnapShot": "c1a435f71b5b876c6e93f4b55d7a852a7c55b37f41538bc15defe362c723df67",
		"generatorPublicKey": "02a09c2e80f2701fb277ccb29f6919c97cc4743b7e7d984685cd2e31769f1871",
		"votes": null,
		"balances": [
			{
				"addr": "B06",
				"blnc": 2500000000000,
				"ltx": "b8921b0bb99ba2ec69b1585c72219f5f97b131487b67ec144d4dd57ab978cc1f"
			},
			{
				"addr": "B0AoqtQfSZyfCex8q4fGwMAFEQbcDt4mZfP",
				"blnc": 5995000000000000,
				"ltx": "b8921b0bb99ba2ec69b1585c72219f5f97b131487b67ec144d4dd57ab978cc1f"
			}
		],
		"timestamp": 1536580477,
		"signaturess": null,
		"signature": "592cef2c41de9eca543cfd110dd32d22b4c93d8e05c7f86b247efaca4473fd4b361938d83623c93b84fa2eb4afc79f5e7506b98b5cc172e0bb49a6d4048f4405"
	},
	"peer": {
		"ip": "127.0.0.1",
		"port": 8080
	}
}
```

### `snapshot.list`
Gets list snapshots

| Param      | Required | Type |Description |
|:----------:|:----------:|:----------:|-------------|
|`offset`|-|int|Constant displacement of snapshots (default: 0)|
|`limit`|-|int|Maximum number of returned snapshots (default: 20)|
```json
{
	"limit":3,
	"offset": 2
}
```
```json
{
	"id": 0,
	"result": [
		{
			"version": 0,
			"id": "60a5c0899060611cec7eb4c860af972bc8cc9fa93a606293a6909822d55f40f1",
			"height": 7,
			"previousSnapShot": "8bf6f764b497faddb330abf8a1dfbda3a43db4aa99690ce3f7e9a74d9d8eb618",
			"generatorPublicKey": "02a09c2e80f2701fb277ccb29f6919c97cc4743b7e7d984685cd2e31769f1871",
			"votes": null,
			"balances": null,
			"timestamp": 1536580397,
			"signaturess": null,
			"signature": "827834c93c9428c82281f277ebeda5a939d61e2d5561f00265d33bdd7a508421dba885d6c3e82bf5185762c8e2a78a22470d7d15c8bf06b1ceb008b48f127303"
		},
		{
			"version": 0,
			"id": "8962ee6fb51c46b0b85dec31128baa87f5ead50c59f23ff33987f371bb76690a",
			"height": 8,
			"previousSnapShot": "60a5c0899060611cec7eb4c860af972bc8cc9fa93a606293a6909822d55f40f1",
			"generatorPublicKey": "02a09c2e80f2701fb277ccb29f6919c97cc4743b7e7d984685cd2e31769f1871",
			"votes": null,
			"balances": null,
			"timestamp": 1536580407,
			"signaturess": null,
			"signature": "a2ba216eb674e2d21eba567c15953f8e5f9effd1a0a1a458f012dd3c019905ebfe7dc88e76e966d844f77233b47824e57104aa20dad2fda23068923a3f7efd02"
		},
		{
			"version": 0,
			"id": "eca6cb003d1f8f905dcf622a4a76a8bf850a3119120b1225808ad50b9206d709",
			"height": 9,
			"previousSnapShot": "8962ee6fb51c46b0b85dec31128baa87f5ead50c59f23ff33987f371bb76690a",
			"generatorPublicKey": "02a09c2e80f2701fb277ccb29f6919c97cc4743b7e7d984685cd2e31769f1871",
			"votes": null,
			"balances": null,
			"timestamp": 1536580417,
			"signaturess": null,
			"signature": "900bdebf4236bcc14589f701c16939464db76c287fa58583d6eaa6b201beaec969a00394f603404109630e6c5b90db72708e531ca08dc4b98f24c22067454d0e"
		}
	],
	"peer": {
		"ip": "127.0.0.1",
		"port": 8080
	}
}
```
