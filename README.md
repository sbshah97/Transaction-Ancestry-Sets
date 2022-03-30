# Transaction-Ancestry-Sets
Building a basic transaction ancestry set calculator.

## Task 1
Fetches all transactions for a block 680000

## APIs to use

### Step 1
* GET /block-height/:height

* Request
curl --location --request GET 'https://blockstream.info/api/block-height/680000'

* Response
000000000000000000076c036ff5119e5a5a74df77abf64203473364509f7732


* GET /block/:hash/txs[/:start_index]

* Request
curl --location --request GET 'https://blockstream.info/api/block/000000000000000000076c036ff5119e5a5a74df77abf64203473364509f7732/txs'

* Response
```json
[
    {
        "txid": "6587550e410fa1815cf180decc03ef84dcddd19478a081097bd2090c5e85b4b6",
        "version": 1,
        "locktime": 4278917605,
        "vin": [
            {
                "txid": "0000000000000000000000000000000000000000000000000000000000000000",
                "vout": 4294967295,
                "prevout": null,
                "scriptsig": "0340600a0463f77f602f706f6f6c696e2e636f6d2ffabe6d6d4e4176dc3ab61820d9dde06fde2d60832bbfd022a3c3dd2fe59d88f95a0b08fa0100000000000000c5398f003bff87935185713123d0946c12904517cd0069396a0000000000",
                "scriptsig_asm": "OP_PUSHBYTES_3 40600a OP_PUSHBYTES_4 63f77f60 OP_PUSHBYTES_47 706f6f6c696e2e636f6d2ffabe6d6d4e4176dc3ab61820d9dde06fde2d60832bbfd022a3c3dd2fe59d88f95a0b08fa OP_PUSHBYTES_1 00 OP_0 OP_0 OP_0 OP_0 OP_0 OP_0 OP_RETURN_197 OP_PUSHBYTES_57 <push past end>",
                "witness": [
                    "0000000000000000000000000000000000000000000000000000000000000000"
                ],
                "is_coinbase": true,
                "sequence": 4294967295
            }
        ],
        "vout": [
            {
                "scriptpubkey": "76a914f5da28aa5ed75c2e850a8b998e92a5bec005561e88ac",
                "scriptpubkey_asm": "OP_DUP OP_HASH160 OP_PUSHBYTES_20 f5da28aa5ed75c2e850a8b998e92a5bec005561e OP_EQUALVERIFY OP_CHECKSIG",
                "scriptpubkey_type": "p2pkh",
                "scriptpubkey_address": "1PQwtwajfHWyAkedss5utwBvULqbGocRpu",
                "value": 907377760
            },
            {
                "scriptpubkey": "6a24b9e11b6df46989f4b5ef750e551a34cdac2c75f711712ed4457084698bf049c378b7e6f3",
                "scriptpubkey_asm": "OP_RETURN OP_PUSHBYTES_36 b9e11b6df46989f4b5ef750e551a34cdac2c75f711712ed4457084698bf049c378b7e6f3",
                "scriptpubkey_type": "op_return",
                "value": 0
            },
            {
                "scriptpubkey": "6a24aa21a9eda6b9213aaceb28a8297f060ac8058111b60450eac3832b65c6336e88719d9066",
                "scriptpubkey_asm": "OP_RETURN OP_PUSHBYTES_36 aa21a9eda6b9213aaceb28a8297f060ac8058111b60450eac3832b65c6336e88719d9066",
                "scriptpubkey_type": "op_return",
                "value": 0
            },
            {
                "scriptpubkey": "6a2952534b424c4f434b3a9bad0a4663489369d5b539327730c77246be3c3c56ab61fd95f40c2f00321e70",
                "scriptpubkey_asm": "OP_RETURN OP_PUSHBYTES_41 52534b424c4f434b3a9bad0a4663489369d5b539327730c77246be3c3c56ab61fd95f40c2f00321e70",
                "scriptpubkey_type": "op_return",
                "value": 0
            }
        ],
        "size": 362,
        "weight": 1340,
        "fee": 0,
        "status": {
            "confirmed": true,
            "block_height": 680000,
            "block_hash": "000000000000000000076c036ff5119e5a5a74df77abf64203473364509f7732",
            "block_time": 1618999138
        }
    }
]
```

## Run code
go run .

[
    txId: A
    vIn {
        txId: B
    }
    taxId: C
        vIn {
            txId: E,
            txId: A
    }
    txId: B
        vIn {
            txId: D
            txId: E
        }

]

B is parent of A

map[string]bool

A
|
B
|
C

Direct + Indirect parents = DP(A) + DP(B) + DP(C)

dip_map[string]int

