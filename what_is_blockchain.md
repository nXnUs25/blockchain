# What is blockchain ?

**Blockchain** is a decentralized and distributed digital ledger technology that records transactions across multiple computers in a way that makes them tamper-resistant and secure. It is the underlying technology behind cryptocurrencies like Bitcoin, but its applications extend far beyond digital currencies. Here are some key characteristics and concepts associated with blockchain:

1. **Decentralization**: Blockchain operates on a decentralized network of computers, often referred to as nodes. This means there is no central authority, such as a bank or government, that controls the network. Each node has a copy of the entire blockchain, ensuring transparency and redundancy.

2. **Immutable Ledger**: Once a transaction is added to the blockchain, it is extremely difficult to alter or delete. This immutability is achieved through cryptographic hashing and consensus mechanisms.

3. **Transparency**: All transactions recorded on the blockchain are visible to anyone with access to the network. This transparency is a key feature, especially in applications where trust and verification are essential.

4. **Cryptography**: Cryptography is used to secure the data and control access to the blockchain. Private keys are used to sign transactions, and public keys are used to verify them.

5. **Consensus Mechanisms**: To validate and add new transactions to the blockchain, nodes in the network must agree on the validity of those transactions. Popular consensus mechanisms include Proof of Work (PoW) and Proof of Stake (PoS).

6. **Smart Contracts**: Smart contracts are self-executing contracts with the terms of the agreement directly written into code. They automatically execute when predefined conditions are met, without the need for intermediaries.

7. **Use Cases**: Blockchain has a wide range of use cases beyond cryptocurrencies, including supply chain management, identity verification, voting systems, healthcare records, and more. It can be applied to any scenario where transparency, security, and decentralization are important.

8. **Public vs. Private Blockchains**: Public blockchains, like Bitcoin and Ethereum, are open to anyone, while private blockchains are restricted to specific entities. Private blockchains are often used by businesses for internal purposes.

9. **Tokens and Cryptocurrencies**: Many blockchains have their own native tokens, often used as a means of exchange or to access the network's features. Bitcoin and Ether (Ethereum's native token) are examples of cryptocurrencies.

Blockchain technology has the potential to revolutionize various industries by reducing fraud, enhancing transparency, and streamlining processes. However, it's important to note that while blockchain is highly secure, it's not entirely immune to vulnerabilities, and the implementation and governance of blockchain systems can vary significantly.


## Blockchain structure 

### Block Structure

A block in a blockchain typically comprises various components, organized in a specific format. Here's an overview of the common components found in the structure of a block:

**Block Header:**
- *Version Number*: This is a numerical value that indicates the version of the blockchain protocol being used.
- *Timestamp*: The timestamp records when the block was created, often in Unix time format.
- *Nonce*: A random number that, when hashed with other block header data, must meet specific criteria defined by the consensus mechanism (e.g., Proof of Work) to mine or validate the block.
- *Merkle Root*: A hash of all the transactions in the block, forming a Merkle tree. This root hash is used to efficiently verify the transactions contained in the block.

**Previous Block's Hash**: A reference to the cryptographic hash of the previous block in the blockchain. This link ensures the chronological order of blocks in the blockchain.

**Transaction Data**: This section contains the actual transactions that the block is recording. These transactions can include data about cryptocurrency transfers, smart contract executions, or other relevant information.

**Additional Data (optional)**: Some blockchain protocols allow for an optional field where additional data can be included, depending on the specific use case or requirements.

The structure of a block may also include other metadata and details specific to the blockchain's design and protocol. The block structure can vary between different blockchain platforms, but the core components listed above are commonly found in most blockchain block structures.


```
+---------- Block ----------+        +---------- Block ----------+         +---------- Block ----------+        +---------- Block ----------+
|                           |        |                           |         |                           |        |                           |
| prev hash     timestamp   | -----> | prev hash     timestamp   | ----->  | prev hash     timestamp   | -----> | prev hash     timestamp   | 
|                           |  OOO   |                           |  OOO    |                           |  OOO   |                           |
| nonce         transaction |        | nonce         transaction |         | nonce         transaction |        | nonce         transaction |
|                           |        |                           |         |                           |        |                           |
+---------------------------+        +---------------------------+         +---------------------------+        +---------------------------+
```