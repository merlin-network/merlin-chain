#!/usr/bin/env bash

fanfuryd tx nftstaking register-nft-staking --from validator --nft-identifier "identifier3" --nft-metadata "metadata" --reward-address "pop1snktzg6rrncqtct3acx2vz60aak2a6fke3ny3c" --reward-weight 1000 --chain-id=testing --home=$HOME/.fanfuryd --keyring-backend=test --broadcast-mode=block --yes
fanfuryd tx nftstaking set-nft-type-perms NFT_TYPE_DEFAULT SET_SERVER_ACCESS --from=validator --chain-id=testing --home=$HOME/.fanfuryd --keyring-backend=test --broadcast-mode=block --yes
fanfuryd tx nftstaking set-access-info $(fanfuryd keys show -a validator --keyring-backend=test) server1#chan1#chan2,server2#chan3 --from=validator --chain-id=testing --home=$HOME/.fanfuryd --keyring-backend=test --broadcast-mode=block --yes

fanfuryd query bank balances pop1uef5c6tx7vhjyhfumhzdhvwkepshcmljyv4wh4
fanfuryd query nftstaking access-infos
fanfuryd query nftstaking access-info $(fanfuryd keys show -a validator --keyring-backend=test)
fanfuryd query nftstaking all-nfttype-perms
fanfuryd query nftstaking has-permission $(fanfuryd keys show -a validator --keyring-backend=test) aaa
fanfuryd query nftstaking nfttype-perms aaa
fanfuryd query nftstaking staking aaa
fanfuryd query nftstaking stakings
fanfuryd query nftstaking stakings_by_owner $(fanfuryd keys show -a validator --keyring-backend=test)

