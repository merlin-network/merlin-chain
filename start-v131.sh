#!/bin/bash

rm -rf $HOME/.merlind/

cd $HOME

merlind131 init --chain-id=testing testing --home=$HOME/.merlind
merlind131 keys add validator --keyring-backend=test --home=$HOME/.merlind
merlind131 add-genesis-account $(merlind131 keys show validator -a --keyring-backend=test --home=$HOME/.merlind) 100000000000umerlin,100000000000stake --home=$HOME/.merlind
merlind131 gentx validator 500000000stake --keyring-backend=test --home=$HOME/.merlind --chain-id=testing
merlind131 collect-gentxs --home=$HOME/.merlind

VALIDATOR=$(merlind131 keys show -a validator --keyring-backend=test --home=$HOME/.merlind)

sed -i '' -e 's/"owner": ""/"owner": "'$VALIDATOR'"/g' $HOME/.merlind/config/genesis.json
sed -i '' -e 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOME/.merlind/config/app.toml 
sed -i '' -e 's/enable = false/enable = true/g' $HOME/.merlind/config/app.toml 
sed -i '' -e 's/cors_allowed_origins = \[\]/cors_allowed_origins = ["*"]/g' $HOME/.merlind/config/config.toml 
sed -i '' 's/"voting_period": "172800s"/"voting_period": "20s"/g' $HOME/.merlind/config/genesis.json

merlind131 start --home=$HOME/.merlind