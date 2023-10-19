#!/usr/bin/env bash

merlind query bank balances $(merlind keys show -a validator --keyring-backend=test)
merlind tx mint burn-tokens 500000000stake --keyring-backend=test --from=validator --chain-id=testing --home=$HOME/.merlind/ --yes  --broadcast-mode=block
merlind query bank balances $(merlind keys show -a validator --keyring-backend=test)
