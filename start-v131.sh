#!/bin/bash

rm -rf $HOME/.fanfuryd/

cd $HOME

fanfuryd131 init --chain-id=testing testing --home=$HOME/.fanfuryd
fanfuryd131 keys add validator --keyring-backend=test --home=$HOME/.fanfuryd
fanfuryd131 add-genesis-account $(fanfuryd131 keys show validator -a --keyring-backend=test --home=$HOME/.fanfuryd) 100000000000ufury,100000000000stake --home=$HOME/.fanfuryd
fanfuryd131 gentx validator 500000000stake --keyring-backend=test --home=$HOME/.fanfuryd --chain-id=testing
fanfuryd131 collect-gentxs --home=$HOME/.fanfuryd

VALIDATOR=$(fanfuryd131 keys show -a validator --keyring-backend=test --home=$HOME/.fanfuryd)

sed -i '' -e 's/"owner": ""/"owner": "'$VALIDATOR'"/g' $HOME/.fanfuryd/config/genesis.json
sed -i '' -e 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOME/.fanfuryd/config/app.toml 
sed -i '' -e 's/enable = false/enable = true/g' $HOME/.fanfuryd/config/app.toml 
sed -i '' -e 's/cors_allowed_origins = \[\]/cors_allowed_origins = ["*"]/g' $HOME/.fanfuryd/config/config.toml 
sed -i '' 's/"voting_period": "172800s"/"voting_period": "20s"/g' $HOME/.fanfuryd/config/genesis.json

fanfuryd131 start --home=$HOME/.fanfuryd