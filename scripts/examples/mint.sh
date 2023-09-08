#!/usr/bin/env bash

fanfuryd query bank balances $(fanfuryd keys show -a validator --keyring-backend=test)
fanfuryd tx mint burn-tokens 500000000stake --keyring-backend=test --from=validator --chain-id=testing --home=$HOME/.fanfuryd/ --yes  --broadcast-mode=block
fanfuryd query bank balances $(fanfuryd keys show -a validator --keyring-backend=test)
