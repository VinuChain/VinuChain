package gossip

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math"
)

var M = decimal.NewFromInt(1000000)
var Ro, _ = decimal.NewFromString("4.201037667e-24")
var OneDecimal = decimal.NewFromInt(1)
var TwoDecimal = decimal.NewFromInt(2)
var EDecimal = decimal.NewFromFloat(math.E)

func GetCurrentQuotaBalanceForUser(blockNumber int64, address common.Address) (quota decimal.Decimal) {
	avgSum, _ := mockSumAvgQuotasInBlocks(blockNumber, blockNumber-74)
	userSum, _ := mockSumUserQuotasInBlocks(blockNumber, blockNumber-74, common.Address{})
	quota = avgSum.Sub(userSum)

	return quota
}

func QuotaForIBLockCalculateForAddress(avgQuotaPerBlock, stakedTokensAddress decimal.Decimal) (quota decimal.Decimal, err error) {
	L, err := calculateNetworkLoadParameter(avgQuotaPerBlock)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return M.Mul(
			OneDecimal.Sub(
				TwoDecimal.Div(OneDecimal.Add(
					EDecimal.Pow(
						L.Mul(stakedTokensAddress).Mul(Ro),
					),
				)),
			),
		),
		err
}

func getUsedQuotaBalanceForAddress(left, right int64, address common.Address) (userQuotasInBlocksSegment []decimal.Decimal, err error) {
	return mockGetUsedQuotaBalanceForAddressOnSegment(left, right, address), nil
}

func calculateNetworkLoadParameter(avgQuotaPerBlock decimal.Decimal) (loadParameter decimal.Decimal, err error) {
	unknownConst1Decimal, _ := decimal.NewFromString("1050000")
	unknownConst2Decimal, _ := decimal.NewFromString("2100000")

	unknownConst3Decimal := decimal.NewFromFloat(8.260667775706495)
	unknownConst4Decimal := decimal.NewFromFloat(1.6949794096275418)

	tenDecimal := decimal.NewFromInt(10)
	const09Decimal := decimal.NewFromFloat(0.9)

	f1 := func() decimal.Decimal {
		// TODO: test it

		// 2 - e^(8.25067775706495*e - 0.9 * (avgQuotaPerBlock - unknownConst1Decimal))
		return TwoDecimal.Sub(
			EDecimal.Pow(
				unknownConst3Decimal.Mul(EDecimal).Sub(const09Decimal.Mul(avgQuotaPerBlock.Sub(unknownConst1Decimal))),
			),
		)

	}

	f2 := func() decimal.Decimal {
		// TODO: test it

		// e^1.69497940962754188*e - 10 * (unknownConst2Decimal - avgQuotaPerBlock)
		return EDecimal.Pow(
			unknownConst4Decimal.Mul(EDecimal).Sub(tenDecimal.Mul(unknownConst2Decimal.Sub(avgQuotaPerBlock))),
		).
			Sub(const09Decimal)
	}

	// avgQuotaPerBloc <= unknownConst1Decimal
	if avgQuotaPerBlock.LessThanOrEqual(unknownConst1Decimal) {
		return decimal.NewFromString("1")
	}

	// avgQuotaPerBloc > unknownConst1Decimal && avgQuotaPerBloc <= unknownConst2Decimal
	if avgQuotaPerBlock.GreaterThan(unknownConst1Decimal) &&
		avgQuotaPerBlock.LessThanOrEqual(unknownConst2Decimal) {
		return f1(), nil
	}

	// avgQuotaPerBloc > unknownConst2Decimal
	if avgQuotaPerBlock.GreaterThan(unknownConst2Decimal) {
		return f2(), nil
	}

	return loadParameter, nil
}

func mockGetUsedQuotaBalanceForAddressOnSegment(left, right int64, address common.Address) (userQuotasInBlocksSegment []decimal.Decimal) {
	for i := left; i <= right; i++ {
		userQuotasInBlocksSegment = append(userQuotasInBlocksSegment, decimal.NewFromInt(30000).Sub(decimal.NewFromInt(i)))
	}
	return userQuotasInBlocksSegment
}

func mockGetUsedQuotaBalanceForAddress(blockNumber int64, address common.Address) (userQuotaInBlocks decimal.Decimal) {
	return decimal.NewFromInt(30000).Sub(decimal.NewFromInt(blockNumber))
}

func mockSumUserQuotasInBlocks(left, right int64, address common.Address) (sum decimal.Decimal, err error) {
	for i := left; i < right; i++ {
		sum = sum.Add(mockGetUsedQuotaBalanceForAddress(i, address))
	}

	return sum, nil
}

func mockSumAvgQuotasInBlocks(left, right int64) (sum decimal.Decimal, err error) {
	const amountOfUsersWithQuotas = 10

	for i := 0; i < amountOfUsersWithQuotas; i++ {
		for j := left; j < right; j++ {
			sum = sum.Add(mockGetUsedQuotaBalanceForAddress(j, common.Address{}))
		}
	}

	return sum, nil
}
