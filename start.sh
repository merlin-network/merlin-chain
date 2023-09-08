#!/bin/bash

rm -rf $HOME/.fanfuryd/

cd $HOME

fanfuryd init --chain-id=testing testing --home=$HOME/.fanfuryd
fanfuryd keys add validator --keyring-backend=test --home=$HOME/.fanfuryd
fanfuryd add-genesis-account $(fanfuryd keys show validator -a --keyring-backend=test --home=$HOME/.fanfuryd) 100000000000ufury,100000000000stake --home=$HOME/.fanfuryd
fanfuryd gentx validator 500000000stake --keyring-backend=test --home=$HOME/.fanfuryd --chain-id=testing
fanfuryd collect-gentxs --home=$HOME/.fanfuryd

VALIDATOR=$(fanfuryd keys show -a validator --keyring-backend=test --home=$HOME/.fanfuryd)

sed -i '' -e 's/"owner": ""/"owner": "'$VALIDATOR'"/g' $HOME/.fanfuryd/config/genesis.json
sed -i '' -e 's/"voting_period": "172800s"/"voting_period": "20s"/g' $HOME/.fanfuryd/config/genesis.json
sed -i '' -e 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOME/.fanfuryd/config/app.toml 
sed -i '' -e 's/enable = false/enable = true/g' $HOME/.fanfuryd/config/app.toml 
sed -i '' -e 's/cors_allowed_origins = \[\]/cors_allowed_origins = ["*"]/g' $HOME/.fanfuryd/config/config.toml 
jq '.app_state.gov.voting_params.voting_period = "20s"'  $HOME/.fanfuryd/config/genesis.json > temp.json ; mv temp.json $HOME/.fanfuryd/config/genesis.json;

fanfuryd start --home=$HOME/.fanfuryd

# git checkout v1.3.0
# go install ./cmd/fanfuryd
# sh start.sh
# fanfuryd tx gov submit-proposal software-upgrade "v1.4.0" --upgrade-height=12 --title="title" --description="description" --from=validator --keyring-backend=test --chain-id=testing --home=$HOME/.fanfuryd/ --yes --broadcast-mode=block --deposit="100000000stake"
# fanfuryd tx gov vote 1 Yes --from=validator --keyring-backend=test --chain-id=testing --home=$HOME/.fanfuryd/ --yes  --broadcast-mode=block
# fanfuryd query gov proposals
# git checkout ica_controller
# go install ./cmd/fanfuryd
# fanfuryd start
# fanfuryd query interchain-accounts controller params
