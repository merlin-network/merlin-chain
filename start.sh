#!/bin/bash

rm -rf $HOME/.merlind/

cd $HOME

merlind init --chain-id=testing testing --home=$HOME/.merlind
merlind keys add validator --keyring-backend=test --home=$HOME/.merlind
merlind add-genesis-account $(merlind keys show validator -a --keyring-backend=test --home=$HOME/.merlind) 100000000000umerlin,100000000000stake --home=$HOME/.merlind
merlind gentx validator 500000000stake --keyring-backend=test --home=$HOME/.merlind --chain-id=testing
merlind collect-gentxs --home=$HOME/.merlind

VALIDATOR=$(merlind keys show -a validator --keyring-backend=test --home=$HOME/.merlind)

sed -i '' -e 's/"owner": ""/"owner": "'$VALIDATOR'"/g' $HOME/.merlind/config/genesis.json
sed -i '' -e 's/"voting_period": "172800s"/"voting_period": "20s"/g' $HOME/.merlind/config/genesis.json
sed -i '' -e 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOME/.merlind/config/app.toml 
sed -i '' -e 's/enable = false/enable = true/g' $HOME/.merlind/config/app.toml 
sed -i '' -e 's/cors_allowed_origins = \[\]/cors_allowed_origins = ["*"]/g' $HOME/.merlind/config/config.toml 
jq '.app_state.gov.voting_params.voting_period = "20s"'  $HOME/.merlind/config/genesis.json > temp.json ; mv temp.json $HOME/.merlind/config/genesis.json;

merlind start --home=$HOME/.merlind

# git checkout v1.3.0
# go install ./cmd/merlind
# sh start.sh
# merlind tx gov submit-proposal software-upgrade "v1.4.0" --upgrade-height=12 --title="title" --description="description" --from=validator --keyring-backend=test --chain-id=testing --home=$HOME/.merlind/ --yes --broadcast-mode=block --deposit="100000000stake"
# merlind tx gov vote 1 Yes --from=validator --keyring-backend=test --chain-id=testing --home=$HOME/.merlind/ --yes  --broadcast-mode=block
# merlind query gov proposals
# git checkout ica_controller
# go install ./cmd/merlind
# merlind start
# merlind query interchain-accounts controller params
