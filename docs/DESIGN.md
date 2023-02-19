# Developers Guide

This section explains some design considerations around how the chat application is implemented using cosmos-sdk and ignite CLI.

## Design

### Structure of a message

For sending a chat application, we would need to determine how to represent the message. For this application it is quite straighforward,
as our Message will have only these few attributes:
- senderAddress: wallet address of sender
- receiverAddress: wallet address of receiver
- body: body of the message
- id: unique KSUID to identify the message

Just these fields will be sufficient for representing a message in this simple application.

### Storing a message upon creation of message

Given that we will be storing the messages in a KV store (IAVL store to be precise), we need to ensure when storing the messages
we are able to cater for multiple read access patterns, namely:
- Users are able to list all messages sent to their wallets
- Users are able to list all messages sent from their wallets

To cater for such read patterns, we chose to store the same message 2 times in the following way:
- we store the message under the prefix `/Messages/Sender/<Sender_wallet_Address>`
- we store the same message under the prefix `/Messages/Receiver/<Receiver_wallet_Address>` as well

By storing the same message twice, we facilitate both use cases described above. For example:
- alice sends bob a message X
- the application stores the message twice:
  - once in `/Messages/Sender/<alice_address>`
  - once in `/Messages/Receiver/<bob_address>`
- bob sends alice a message Y
- the application stores the message twice:
    - once in `/Messages/Sender/<bob_address>`
    - once in `/Messages/Receiver/<alice_address>`
- If alice wants to see how many messages she sent, she can just look up from the prefix `/Messages/Sender/<alice_address>` and
she will see message X
- Same if bob wants to see how many messages he received, he can just look up from the prefix `/Messages/Receiver/<bob_address>` and he
will see message X
- If bob wants to see how many messages he sent, he can just look up from the prefix `/Messages/Sender/<bob_address>` and he 
will see message Y

So by storing the same message twice in different prefixes, it caters for multiple use cases where users both want to see all
messages sent to their wallets and from their wallets, especially in a KV store and not something like SQL environment.
Note that when storing in the chain, for the key, we are using the message ID and for the value, it's the entire serialized message.
By doing so, if a user wants to read all messages sent to his/her wallet, they can just query the `/Messages/Sender/<wallet_address>` prefix only 
with pagination, as there might be many messages that is sent to their wallet, and they might not want to load all at one go.

### Retrieving messages from the store

As mentioned above, if a user wants to list all messages sent to his/her wallet, they can just query the `/Messages/Receiver/<wallet_address>` prefix, and
for all messages sent from their wallets, they can just query the `/Messages/Sender/<wallet_address>` prefix.

Note that there is encryption at rest using RSA, so the messages retrieved from the chain are encrypted already, and when
retrieving the messages they will still be encrypted as well. The intention is the client side will decrypt it using the private
key.

### Encryption Design

One thing this chat application tries to do is message encryption.

All messages are encrypted using RSA, an asymmetric . In our example, we will be encrypting the messages on the chat application
via a public key registered by the user (which is public and thus can be shared), and the user can decrypt the message in
the client side/frontend using the private key.

The key for each user is stored in the `/Messages/EncryptionKey/` prefix and the key is the wallet address. The key is only
retrieved during the sending of messages to encrypt the message. Messages are encrypted and then base64 encoded for easier storage
and transmission.

Note that one limitation with this design is that it is not truly end-to-end encrypted. For it to be end-to-end encrypted,
the encryption has to be done at client side as well. However, as I wanted to try out doing encryption of the message at rest, 
I chose to do it at the server side instead, otherwise without a proper frontend it will be hard to do the encryption.

## Code Structure

Most code implementation are found in `x/cosmosmessenger/keeper`. Some other code are boilerplate from scaffolding, and some
testing utility code can be found in `testutil`. 
- `x/cosmosmessenger/keeper/encryption.go`: handles encryption logic
- `x/cosmosmessenger/keeper/encryption_store.go`: handles saving and reading of encryption keys in the KV store
- `x/cosmosmessenger/keeper/msg_server_create_message.go`: handle creating of message into the chain
- `x/cosmosmessenger/keeper/msg_server_register_wallet_key.go`: handle saving of public key into the chain and handle errors
- `x/cosmosmessenger/keeper/msg_store.go`: handles saving and reading of messages in the KV store
- `x/cosmosmessenger/keeper/query_show_received_messages.go`: handles showing of messages received for wallet
- `x/cosmosmessenger/keeper/query_show_sent_messages.go`: handles showing of messages sent for wallet
