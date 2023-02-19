# How To Start

This document contains instructions on how to set up and use this basic chat application module
built using Ignite and Cosmos SDK. This messenger allows for the following:
- send messages to other people registered on the chain, including yourself
- read all messages you have sent to others
- read all messages others have sent to you
- encryption of messages between you and others. All messages sent and received are encrypted
using RSA2048, and you will need to decrypt the messages yourself using the private key, while the server
will use the public key to encrypt all messages.

## Pre-requisites

You will need the following to start:
- [Ignite CLI](https://ignite.com/cli)
- Go 1.19

## Instructions

This section contains the instructions to set up the chat application and 

### Step 1: Clone this repo

Clone the repository from GitHub. Then change directory to the repo.
```shell
git clone git@github.com:coby9241/cosmos-messenger.git
cd cosmos-messenger
```

### Step 2: Start up the chain

Assuming you have already installed ignite CLI, boot up the chain by running:
```shell
ignite chain serve
```

You should then see the chain boot up with the following:
```shell
  Blockchain is running

  👤 alice's account address: cosmos1h0rgynu730qadyy46qaf60spjlx753325m3cdd
  👤 bob's account address: cosmos1ndkhum7mdmnys9jqpcplqdf4hzajy80eqtf3ug

  🌍 Tendermint node: http://0.0.0.0:26657
  🌍 Blockchain API: http://0.0.0.0:1317
  🌍 Token faucet: http://0.0.0.0:4500

  ⋆ Data directory: /<path/to/repo>/.cosmos-messenger
  ⋆ App binary: /<path/to/repo>/go/bin/cosmos-messengerd

  Press the 'q' key to stop serve
```

### Step 3: Generate public and private key pair for encryption and register it

Next step, you will need generate your own RSA key pair and register your public key to the application.

You can either use an online tool like [CryptoTools](https://cryptotools.net/rsagen#:~:text=To%20generate%20a%20key%20pair,take%20up%20to%20several%20minutes.) to
generate one, or more securely, use openssl on your computer.

First generate the private key to a file called `private.pem`
```shell
openssl genrsa -out private.pem 2048
```
Then generate the public key to a file called `public.pem`
```shell
openssl rsa -in private.pem -outform PEM -pubout -out public.pem
```

Alternatively, you can use the public and private key provided in `examples` folder of this repo for testing purposes only.

Next you will need to register it to the server. First remove the newlines from the public key and encode it in base64. This can be done via
```shell
cat public.pem | base64 | tr -d \\n
```
Next run the command to register it in the server. Example to register a new public key for alice:
```shell
/path/to/cosmos-messengerd tx cosmosmessenger register-wallet-key '<output_of_previous_command>' --from alice
```
Accept the transaction fee by typing `y` and pressing Enter/Return.
The path to the chain binary can be found when running the `ignite chain serve` command.
Note that both sender and receiver must have their public key registered before they can send a message to each other.

### Step 4: Send a message using chain's binary

Now that you have registered the public key, you can now send a message to other people.

To send a message `xxx` to Bob from alice for example, run the following command:
```shell
/path/to/cosmos-messengerd tx cosmosmessenger create-message <bobs_wallet_address> 'xxx' --from alice
```
Accept the transaction fee by typing `y` and pressing Enter/Return.
Note that you need to specify bob's wallet address rather than the account name. If either alice or bob's public key is not
registered, it will result in an error.

### Step 5: Showing all sent messages and received messages

To show all sent messages for alice for example, use the following command:
```shell
/path/to/cosmos-messengerd q cosmosmessenger show-sent-messages <alice_wallet_address>
```
This will show all messages sent by alice. An example output will be as follows:
```shell
$ /path/to/cosmos-messengerd q cosmosmessenger show-sent-messages cosmos1h0rgynu730qadyy46qaf60spjlx753325m3cdd
messages:
- body: LYvVKJidEAWLLOUkVDYjgXotzJbHrEYEaQ8nOrBJoeZG1NaiXm7uYI9Zz6hBCkWqWHGZvWn2L43BMNSUWbvcroQ7bJ3EXxGbZuWWP8mB+Bs5CBJLXIEGFIKwQUpkZNn9QoLk4Wsq4UovBm/XBRBnyVuM1uN4lA3tlodfkfxu/kV6bDijNBwPcOdpcyxfc68W0UXxW3x2WDBpQI8U49vRxEVWcJYY5dOUyOeNG4SmpVippAddl4HTMpdnarayFOEQtyCvYGCPGbV+9R/UCTgNXEeZMIk9LBOKoI2Kz/EpEcKSh5BL5yJcfSsMfJiWx6sv73NUsWkbvr9daH9cHjTGvg==
  id: 2Lwsi5TyNZFgQJWIBIV8zUGEiV5
  receiverAddress: cosmos1h0rgynu730qadyy46qaf60spjlx753325m3cdd
  senderAddress: cosmos1h0rgynu730qadyy46qaf60spjlx753325m3cdd
pagination:
  next_key: null
  total: "1"
```
Note that the `body` of the messages are all encrypted and base64 encoded, and can only be decrypted using the
private key. This is the same for showing all received messages as well. 
To show all received messages for alice for example, use the following command:
```shell
/path/to/cosmos-messengerd q cosmosmessenger show-received-messages <alice_wallet_address>
```
In a real world scenario, the messages should be decrypted in the frontend and by the user using the private key that they
have by the frontend client, hence the `body` of the message from both queries are encrypted + encoded and does not seem to make sense.
To allow you to see the actual message, we have created a simple python script that can be found in `examples/decrypt.py` of 
this repo. 

By replacing the `private_key` with the private key you have generated, and `encoded_str` with the `body` attribute from the output
of `show-received-messages/show-sent-messages`, you can decrypt and see the actual message.
```shell
$ python3 examples/decrypt.py
xxx
```
This will print out the actual message that was sent.
Note that the script will require you to have both:
- python 3, and
- pycryptodome

installed. You can install pycryptodome by running `pip3 install pycryptodome`

