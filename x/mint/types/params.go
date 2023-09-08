package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyMintDenom                            = []byte("MintDenom")
	KeyGenesisBlockProvisions               = []byte("GenesisBlockProvisions")
	KeyReductionPeriodInBlocks              = []byte("ReductionPeriodInBlocks")
	KeyReductionFactor                      = []byte("ReductionFactor")
	KeyPoolAllocationRatio                  = []byte("PoolAllocationRatio")
	KeyDeveloperRewardsReceiver             = []byte("DeveloperRewardsReceiver")
	KeyMintingRewardsDistributionStartBlock = []byte("MintingRewardsDistributionStartBlock")
	KeyUsageIncentiveAddress                = []byte("UsageIncentiveAddress")
	KeyGrantsProgramAddress                 = []byte("GrantsProgramAddress")
	KeyTeamReserveAddress                   = []byte("TeamReserveAddress")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams returns new mint module parameters initialized to the given values.
func NewParams(
	mintDenom string, genesisBlockProvisions sdk.Dec,
	ReductionFactor sdk.Dec, reductionPeriodInBlocks int64, distrProportions DistributionProportions,
	weightedDevRewardsReceivers []MonthlyVestingAddress, MintingRewardsDistributionStartBlock int64,
) Params {
	return Params{
		MintDenom:                            mintDenom,
		GenesisBlockProvisions:               genesisBlockProvisions,
		ReductionPeriodInBlocks:              reductionPeriodInBlocks,
		ReductionFactor:                      ReductionFactor,
		DistributionProportions:              distrProportions,
		WeightedDeveloperRewardsReceivers:    weightedDevRewardsReceivers,
		MintingRewardsDistributionStartBlock: MintingRewardsDistributionStartBlock,
	}
}

func addressTable() map[string]string {
	addressJSON := `{
		"adress1": "fury1zyakv8ny9p5esrpv3rgls707rd9anjzl047cyd",
		"adress2": "fury10rp3k6jh8nxmrvdxaf6vwcv6z0ad6p7asfjuds",
		"adress3": "fury16n36a4xryrcaf4vtk9nuqq0lzrs3qkmj5nj64k",
		"adress4": "fury1s8qa7466v6pnc7mqhntnzx0kukf3nl52y0mlvq",
		"adress5": "fury1xwmdtmhmtx0vsz6vd6yjn6z26rwj6c592a32pd",
		"adress6": "fury1g7ryul6kv8wv7p032c3shede4yps74h0jq7f4s",
		"adress7": "fury184vpdnt4pzkz70ery009l9ac5p8sel7sud8gfa",
		"adress8": "fury17nqxtdm7nrj0ne0jumkkmsjghxytdv8lj566c0",
		"adress9": "fury1843zuxx87tfy0rxlfv4ulgvxqt5kk3jjyhp4zp",
		"adress10": "fury1f9dkjdelh3nnmpkahztdxpt2vas5a8jx2dphfc",
		"adress11": "fury1hv4wp790e47y4aw2rrk4s0e35ta4nfrzf8kqrq",
		"adress12": "fury10tm5wcdkvvzyhmjd44aeg4r7zlfpwyufplhlnj",
		"adress13": "fury15vc2563rxqzulsjzt89ugyqae063ezrfedts4m",
		"adress14": "fury1v47dvyflzgatgdul52vgxy6fv8rlgmtwr0d9yr",
		"adress15": "fury1wue905yydrxfysqq6ewgpsx0zdsdspn0zaz2dk",
		"adress16": "fury1mq05ml4zmg4eus96k72re3n06ghuz0tx7h2yfj",
		"adress17": "fury1negrycg7hsaumjedue8my9xhr688guav4xqsu4",
		"adress18": "fury1pv6n9f8eml89rmzxnuzz70936hm60a39vsgk02",
		"adress19": "fury1nm50zycnm9yf33rv8n6lpks24usxzahk80qa5g",
		"adress20": "fury1znhgcje2np5v34nk7j7t4jes4f8mu6alh5t0ju",
		"adress21": "fury1uwr8dn8h3qsrwt2pew57r577qhzk9w5wcv97py",
		"adress22": "fury18cgtgz6q7ly4suukk744cjep4uxhm5z30u4qp5",
		"adress23": "fury1jtz7h88hzufhwz4pnwaagv6j7czddcz6xkjdg4",
		"adress24": "fury12dgvzxvd339paqvu83vx9wq36j0w3zyxzm064u",
		"adress25": "fury1nra74gcsqy88m9xe5r6jpgyfr6w7zj39axghxc",
		"adress26": "fury1shfq05pu5x8lwm4rng44v7qt888hg78wm2krke",
		"adress27": "fury1l3ggmanvvmm3ph66tw04gdpyd0qwm7pkquv26u",
		"adress28": "fury10y7y3rrmawsfx7n57qxjst6gd6zreuqpnqx7xl"
	}`

	var addressMap map[string]string
	json.Unmarshal([]byte(addressJSON), &addressMap)
	return addressMap
}

func parseMonthlyVesting() []MonthlyVestingAddress {
	records := [][]string{}
	lines := strings.Split(vestingStr, "\n")
	for _, line := range lines {
		records = append(records, strings.Split(line, ","))
	}

	addressMap := addressTable()
	vAddrs := []MonthlyVestingAddress{}
	for _, addr := range records[0] {
		vAddrs = append(vAddrs, MonthlyVestingAddress{
			Address:        addressMap[addr],
			MonthlyAmounts: []sdk.Int{},
		})
	}

	for _, line := range records[1:] {
		for index, amountStr := range line {
			amountDec := sdk.MustNewDecFromStr(amountStr)
			amountInt := amountDec.Mul(sdk.NewDec(1000_000)).TruncateInt()
			vAddrs[index].MonthlyAmounts = append(vAddrs[index].MonthlyAmounts, amountInt)
		}
	}

	return vAddrs
}

// DefaultParams returns the default minting module parameters.
func DefaultParams() Params {
	return Params{
		MintDenom:               sdk.DefaultBondDenom,
		GenesisBlockProvisions:  sdk.NewDec(47000000),        //  300 million /  6307200 * 10 ^ 6
		ReductionPeriodInBlocks: 6307200,                     // 1 year - 86400 x 365 / 5
		ReductionFactor:         sdk.NewDecWithPrec(6666, 4), // 0.6666
		DistributionProportions: DistributionProportions{
			GrantsProgram:    sdk.NewDecWithPrec(10, 2), // 10%
			CommunityPool:    sdk.NewDecWithPrec(10, 2), // 10%
			UsageIncentive:   sdk.NewDecWithPrec(25, 2), // 25%
			Staking:          sdk.NewDecWithPrec(40, 2), // 40%
			DeveloperRewards: sdk.NewDecWithPrec(15, 2), // 15%
		},
		WeightedDeveloperRewardsReceivers:    parseMonthlyVesting(),
		UsageIncentiveAddress:                "fury1at6zkjpxleg8nd8u67542fprzgsev6jhxqhymx",
		GrantsProgramAddress:                 "fury1a28lq0usqrma2tn5t7vmdg3jnglh3v3qdvtsvs",
		TeamReserveAddress:                   "fury1efcnw3j074urqryseyx4weahr2p5at9l93xm6r",
		MintingRewardsDistributionStartBlock: 0,
	}
}

// Validate validates mint module parameters. Returns nil if valid,
// error otherwise
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateGenesisBlockProvisions(p.GenesisBlockProvisions); err != nil {
		return err
	}
	if err := validateReductionPeriodInBlocks(p.ReductionPeriodInBlocks); err != nil {
		return err
	}
	if err := validateReductionFactor(p.ReductionFactor); err != nil {
		return err
	}
	if err := validateDistributionProportions(p.DistributionProportions); err != nil {
		return err
	}

	if err := validateAddress(p.UsageIncentiveAddress); err != nil {
		return err
	}

	if err := validateAddress(p.GrantsProgramAddress); err != nil {
		return err
	}

	if err := validateAddress(p.TeamReserveAddress); err != nil {
		return err
	}

	if err := validateWeightedDeveloperRewardsReceivers(p.WeightedDeveloperRewardsReceivers); err != nil {
		return err
	}
	if err := validateMintingRewardsDistributionStartBlock(p.MintingRewardsDistributionStartBlock); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {

	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyGenesisBlockProvisions, &p.GenesisBlockProvisions, validateGenesisBlockProvisions),
		paramtypes.NewParamSetPair(KeyReductionPeriodInBlocks, &p.ReductionPeriodInBlocks, validateReductionPeriodInBlocks),
		paramtypes.NewParamSetPair(KeyReductionFactor, &p.ReductionFactor, validateReductionFactor),
		paramtypes.NewParamSetPair(KeyPoolAllocationRatio, &p.DistributionProportions, validateDistributionProportions),
		paramtypes.NewParamSetPair(KeyDeveloperRewardsReceiver, &p.WeightedDeveloperRewardsReceivers, validateWeightedDeveloperRewardsReceivers),
		paramtypes.NewParamSetPair(KeyUsageIncentiveAddress, &p.UsageIncentiveAddress, validateAddress),
		paramtypes.NewParamSetPair(KeyGrantsProgramAddress, &p.GrantsProgramAddress, validateAddress),
		paramtypes.NewParamSetPair(KeyTeamReserveAddress, &p.TeamReserveAddress, validateAddress),
		paramtypes.NewParamSetPair(KeyMintingRewardsDistributionStartBlock, &p.MintingRewardsDistributionStartBlock, validateMintingRewardsDistributionStartBlock),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateGenesisBlockProvisions(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("genesis block provision must be non-negative")
	}

	return nil
}

func validateReductionPeriodInBlocks(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("reduction period must be positive: %d", v)
	}

	return nil
}

func validateReductionFactor(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.GT(sdk.NewDec(1)) {
		return fmt.Errorf("reduction factor cannot be greater than 1")
	}

	if v.IsNegative() {
		return fmt.Errorf("reduction factor cannot be negative")
	}

	return nil
}

func validateDistributionProportions(i interface{}) error {
	v, ok := i.(DistributionProportions)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.GrantsProgram.IsNegative() {
		return errors.New("staking distribution ratio should not be negative")
	}

	if v.CommunityPool.IsNegative() {
		return errors.New("staking distribution ratio should not be negative")
	}

	if v.UsageIncentive.IsNegative() {
		return errors.New("community pool distribution ratio should not be negative")
	}

	if v.Staking.IsNegative() {
		return errors.New("staking distribution ratio should not be negative")
	}

	if v.DeveloperRewards.IsNegative() {
		return errors.New("developer rewards distribution ratio should not be negative")
	}

	totalProportions := v.GrantsProgram.Add(v.CommunityPool).Add(v.UsageIncentive).Add(v.Staking).Add(v.DeveloperRewards)

	if !totalProportions.Equal(sdk.NewDec(1)) {
		return errors.New("total distributions ratio should be 1")
	}

	return nil
}

func validateWeightedDeveloperRewardsReceivers(i interface{}) error {
	v, ok := i.([]MonthlyVestingAddress)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// fund community pool when rewards address is empty
	if len(v) == 0 {
		return nil
	}

	return nil
}

func validateMintingRewardsDistributionStartBlock(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("start block must be non-negative")
	}

	return nil
}

func validateAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	_, err := sdk.AccAddressFromBech32(v)

	return err
}
