#!/usr/bin/env bash

merlind tx nftstaking register-nft-staking --from validator --nft-identifier "identifier3" --nft-metadata "metadata" --reward-address "pop1snktzg6rrncqtct3acx2vz60aak2a6fke3ny3c" --reward-weight 1000 --chain-id=testing --home=$HOME/.merlind --keyring-backend=test --broadcast-mode=block --yes
merlind tx nftstaking set-nft-type-perms NFT_TYPE_DEFAULT SET_SERVER_ACCESS --from=validator --chain-id=testing --home=$HOME/.merlind --keyring-backend=test --broadcast-mode=block --yes
merlind tx nftstaking set-access-info $(merlind keys show -a validator --keyring-backend=test) server1#chan1#chan2,server2#chan3 --from=validator --chain-id=testing --home=$HOME/.merlind --keyring-backend=test --broadcast-mode=block --yes

merlind query bank balances pop1uef5c6tx7vhjyhfumhzdhvwkepshcmljyv4wh4
merlind query nftstaking access-infos
merlind query nftstaking access-info $(merlind keys show -a validator --keyring-backend=test)
merlind query nftstaking all-nfttype-perms
merlind query nftstaking has-permission $(merlind keys show -a validator --keyring-backend=test) aaa
merlind query nftstaking nfttype-perms aaa
merlind query nftstaking staking aaa
merlind query nftstaking stakings
merlind query nftstaking stakings_by_owner $(merlind keys show -a validator --keyring-backend=test)

