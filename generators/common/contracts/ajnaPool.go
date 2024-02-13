// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AjnaPoolMetaData contains all meta data concerning the AjnaPool contract.
var AjnaPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AddAboveAuctionPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddAboveAuctionPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountLTMinDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotClearable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotCleared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotCleared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotTakeable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionPriceGtBucketPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowerNotSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowerOk\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowerUnderCollateralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketBankruptcyBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotMergeToHigherPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DustAmountNotExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlashloanCallbackFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlashloanIncorrectBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlashloanUnavailableForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLP\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAllowancesInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LUPBelowHTP\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitIndexExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MoveToSameIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAuction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReservesAuction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PRBMathSD59x18__DivInputTooSmall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rAbs\",\"type\":\"uint256\"}],\"name\":\"PRBMathSD59x18__DivOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"x\",\"type\":\"int256\"}],\"name\":\"PRBMathSD59x18__Exp2InputTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"x\",\"type\":\"int256\"}],\"name\":\"PRBMathSD59x18__FromIntOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"x\",\"type\":\"int256\"}],\"name\":\"PRBMathSD59x18__FromIntUnderflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"x\",\"type\":\"int256\"}],\"name\":\"PRBMathSD59x18__LogInputTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PRBMathSD59x18__MulInputTooSmall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rAbs\",\"type\":\"uint256\"}],\"name\":\"PRBMathSD59x18__MulOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"x\",\"type\":\"int256\"}],\"name\":\"PRBMathSD59x18__SqrtNegativeInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"x\",\"type\":\"int256\"}],\"name\":\"PRBMathSD59x18__SqrtOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"}],\"name\":\"PRBMath__MulDivFixedPointOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"PRBMath__MulDivOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceBelowLUP\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoveDepositLockedByAuctionDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoveDepositLockedByAuctionDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReserveAuctionTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToSameOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferorNotApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDebtToCollateral\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAwarded\",\"type\":\"uint256\"}],\"name\":\"AddCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAwarded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lup\",\"type\":\"uint256\"}],\"name\":\"AddQuoteToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transferors\",\"type\":\"address[]\"}],\"name\":\"ApproveLPTransferors\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AuctionNFTSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"AuctionSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kicker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpForfeited\",\"type\":\"uint256\"}],\"name\":\"BucketBankruptcy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bondChange\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isReward\",\"type\":\"bool\"}],\"name\":\"BucketTake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kicker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAwardedTaker\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAwardedKicker\",\"type\":\"uint256\"}],\"name\":\"BucketTakeLPAwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"DecreaseLPAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBorrowed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralPledged\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lup\",\"type\":\"uint256\"}],\"name\":\"DrawDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Flashloan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"IncreaseLPAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"InterestUpdateFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bond\",\"type\":\"uint256\"}],\"name\":\"Kick\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableReservesRemaining\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBurnEpoch\",\"type\":\"uint256\"}],\"name\":\"KickReserveAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"LoanStamped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpRedeemedFrom\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAwardedTo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lup\",\"type\":\"uint256\"}],\"name\":\"MoveQuoteToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpRedeemed\",\"type\":\"uint256\"}],\"name\":\"RemoveCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lup\",\"type\":\"uint256\"}],\"name\":\"RemoveQuoteToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteRepaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralPulled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lup\",\"type\":\"uint256\"}],\"name\":\"RepayDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableReservesRemaining\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBurnEpoch\",\"type\":\"uint256\"}],\"name\":\"ReserveAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"ResetInterestRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"RevokeLPAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transferors\",\"type\":\"address[]\"}],\"name\":\"RevokeLPTransferors\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settledDebt\",\"type\":\"uint256\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bondChange\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isReward\",\"type\":\"bool\"}],\"name\":\"Take\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lp\",\"type\":\"uint256\"}],\"name\":\"TransferLP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"UpdateInterestRate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToAdd_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry_\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bucketLP_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry_\",\"type\":\"uint256\"}],\"name\":\"addQuoteToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bucketLP_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedAmount_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transferors_\",\"type\":\"address[]\"}],\"name\":\"approveLPTransferors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedTransferors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower_\",\"type\":\"address\"}],\"name\":\"auctionInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"kicker_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bondFactor_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondSize_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kickTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referencePrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"neutralPrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToCollateral_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"head_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"next_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prev_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower_\",\"type\":\"address\"}],\"name\":\"borrowerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bucketIndex_\",\"type\":\"uint256\"}],\"name\":\"bucketCollateralDust\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"bucketExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"exchangeRate_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"bucketInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrowerAddress_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"depositTake_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"bucketTake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnEventEpoch_\",\"type\":\"uint256\"}],\"name\":\"burnInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBurnEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debtInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"decreaseLPAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debt_\",\"type\":\"uint256\"}],\"name\":\"depositIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"depositScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"depositUpToIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositUtilization\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrowerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToBorrow_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitIndex_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralToPledge_\",\"type\":\"uint256\"}],\"name\":\"drawDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emasInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"increaseLPAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inflatorInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"npLimitIndex_\",\"type\":\"uint256\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kickReserveAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kicker_\",\"type\":\"address\"}],\"name\":\"kickerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lender_\",\"type\":\"address\"}],\"name\":\"lenderInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lpBalance_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositTime_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"npLimitIndex_\",\"type\":\"uint256\"}],\"name\":\"lenderKick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId_\",\"type\":\"uint256\"}],\"name\":\"loanInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loansInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"lpAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowance_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxLoan_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry_\",\"type\":\"uint256\"}],\"name\":\"moveQuoteToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromBucketLP_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBucketLP_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"movedAmount_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteTokenScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"removeCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"removedAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemedLP_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"removeQuoteToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"removedAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemedLP_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrowerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxQuoteTokenAmountToRepay_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmountToPull_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralReceiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"limitIndex_\",\"type\":\"uint256\"}],\"name\":\"repayDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountRepaid_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservesInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"}],\"name\":\"revokeLPAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transferors_\",\"type\":\"address[]\"}],\"name\":\"revokeLPTransferors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrowerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth_\",\"type\":\"uint256\"}],\"name\":\"settle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralSettled_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBorrowerSettled_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stampLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrowerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callee_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"take\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralTaken_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAmount_\",\"type\":\"uint256\"}],\"name\":\"takeReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAuctionsInPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalT0Debt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalT0DebtInAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"}],\"name\":\"transferLP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount_\",\"type\":\"uint256\"}],\"name\":\"withdrawBonds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AjnaPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AjnaPoolMetaData.ABI instead.
var AjnaPoolABI = AjnaPoolMetaData.ABI

// AjnaPool is an auto generated Go binding around an Ethereum contract.
type AjnaPool struct {
	AjnaPoolCaller     // Read-only binding to the contract
	AjnaPoolTransactor // Write-only binding to the contract
	AjnaPoolFilterer   // Log filterer for contract events
}

// AjnaPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AjnaPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AjnaPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AjnaPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AjnaPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AjnaPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AjnaPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AjnaPoolSession struct {
	Contract     *AjnaPool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AjnaPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AjnaPoolCallerSession struct {
	Contract *AjnaPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AjnaPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AjnaPoolTransactorSession struct {
	Contract     *AjnaPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AjnaPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AjnaPoolRaw struct {
	Contract *AjnaPool // Generic contract binding to access the raw methods on
}

// AjnaPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AjnaPoolCallerRaw struct {
	Contract *AjnaPoolCaller // Generic read-only contract binding to access the raw methods on
}

// AjnaPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AjnaPoolTransactorRaw struct {
	Contract *AjnaPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAjnaPool creates a new instance of AjnaPool, bound to a specific deployed contract.
func NewAjnaPool(address common.Address, backend bind.ContractBackend) (*AjnaPool, error) {
	contract, err := bindAjnaPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AjnaPool{AjnaPoolCaller: AjnaPoolCaller{contract: contract}, AjnaPoolTransactor: AjnaPoolTransactor{contract: contract}, AjnaPoolFilterer: AjnaPoolFilterer{contract: contract}}, nil
}

// NewAjnaPoolCaller creates a new read-only instance of AjnaPool, bound to a specific deployed contract.
func NewAjnaPoolCaller(address common.Address, caller bind.ContractCaller) (*AjnaPoolCaller, error) {
	contract, err := bindAjnaPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolCaller{contract: contract}, nil
}

// NewAjnaPoolTransactor creates a new write-only instance of AjnaPool, bound to a specific deployed contract.
func NewAjnaPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AjnaPoolTransactor, error) {
	contract, err := bindAjnaPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolTransactor{contract: contract}, nil
}

// NewAjnaPoolFilterer creates a new log filterer instance of AjnaPool, bound to a specific deployed contract.
func NewAjnaPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AjnaPoolFilterer, error) {
	contract, err := bindAjnaPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolFilterer{contract: contract}, nil
}

// bindAjnaPool binds a generic wrapper to an already deployed contract.
func bindAjnaPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AjnaPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AjnaPool *AjnaPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AjnaPool.Contract.AjnaPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AjnaPool *AjnaPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AjnaPool.Contract.AjnaPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AjnaPool *AjnaPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AjnaPool.Contract.AjnaPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AjnaPool *AjnaPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AjnaPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AjnaPool *AjnaPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AjnaPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AjnaPool *AjnaPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AjnaPool.Contract.contract.Transact(opts, method, params...)
}

// ApprovedTransferors is a free data retrieval call binding the contract method 0xd9606e08.
//
// Solidity: function approvedTransferors(address , address ) view returns(bool)
func (_AjnaPool *AjnaPoolCaller) ApprovedTransferors(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "approvedTransferors", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedTransferors is a free data retrieval call binding the contract method 0xd9606e08.
//
// Solidity: function approvedTransferors(address , address ) view returns(bool)
func (_AjnaPool *AjnaPoolSession) ApprovedTransferors(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _AjnaPool.Contract.ApprovedTransferors(&_AjnaPool.CallOpts, arg0, arg1)
}

// ApprovedTransferors is a free data retrieval call binding the contract method 0xd9606e08.
//
// Solidity: function approvedTransferors(address , address ) view returns(bool)
func (_AjnaPool *AjnaPoolCallerSession) ApprovedTransferors(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _AjnaPool.Contract.ApprovedTransferors(&_AjnaPool.CallOpts, arg0, arg1)
}

// AuctionInfo is a free data retrieval call binding the contract method 0x0448e51a.
//
// Solidity: function auctionInfo(address borrower_) view returns(address kicker_, uint256 bondFactor_, uint256 bondSize_, uint256 kickTime_, uint256 referencePrice_, uint256 neutralPrice_, uint256 debtToCollateral_, address head_, address next_, address prev_)
func (_AjnaPool *AjnaPoolCaller) AuctionInfo(opts *bind.CallOpts, borrower_ common.Address) (struct {
	Kicker           common.Address
	BondFactor       *big.Int
	BondSize         *big.Int
	KickTime         *big.Int
	ReferencePrice   *big.Int
	NeutralPrice     *big.Int
	DebtToCollateral *big.Int
	Head             common.Address
	Next             common.Address
	Prev             common.Address
}, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "auctionInfo", borrower_)

	outstruct := new(struct {
		Kicker           common.Address
		BondFactor       *big.Int
		BondSize         *big.Int
		KickTime         *big.Int
		ReferencePrice   *big.Int
		NeutralPrice     *big.Int
		DebtToCollateral *big.Int
		Head             common.Address
		Next             common.Address
		Prev             common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Kicker = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BondFactor = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BondSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.KickTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReferencePrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.NeutralPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DebtToCollateral = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Head = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.Next = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.Prev = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AuctionInfo is a free data retrieval call binding the contract method 0x0448e51a.
//
// Solidity: function auctionInfo(address borrower_) view returns(address kicker_, uint256 bondFactor_, uint256 bondSize_, uint256 kickTime_, uint256 referencePrice_, uint256 neutralPrice_, uint256 debtToCollateral_, address head_, address next_, address prev_)
func (_AjnaPool *AjnaPoolSession) AuctionInfo(borrower_ common.Address) (struct {
	Kicker           common.Address
	BondFactor       *big.Int
	BondSize         *big.Int
	KickTime         *big.Int
	ReferencePrice   *big.Int
	NeutralPrice     *big.Int
	DebtToCollateral *big.Int
	Head             common.Address
	Next             common.Address
	Prev             common.Address
}, error) {
	return _AjnaPool.Contract.AuctionInfo(&_AjnaPool.CallOpts, borrower_)
}

// AuctionInfo is a free data retrieval call binding the contract method 0x0448e51a.
//
// Solidity: function auctionInfo(address borrower_) view returns(address kicker_, uint256 bondFactor_, uint256 bondSize_, uint256 kickTime_, uint256 referencePrice_, uint256 neutralPrice_, uint256 debtToCollateral_, address head_, address next_, address prev_)
func (_AjnaPool *AjnaPoolCallerSession) AuctionInfo(borrower_ common.Address) (struct {
	Kicker           common.Address
	BondFactor       *big.Int
	BondSize         *big.Int
	KickTime         *big.Int
	ReferencePrice   *big.Int
	NeutralPrice     *big.Int
	DebtToCollateral *big.Int
	Head             common.Address
	Next             common.Address
	Prev             common.Address
}, error) {
	return _AjnaPool.Contract.AuctionInfo(&_AjnaPool.CallOpts, borrower_)
}

// BorrowerInfo is a free data retrieval call binding the contract method 0xca103d15.
//
// Solidity: function borrowerInfo(address borrower_) view returns(uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) BorrowerInfo(opts *bind.CallOpts, borrower_ common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "borrowerInfo", borrower_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// BorrowerInfo is a free data retrieval call binding the contract method 0xca103d15.
//
// Solidity: function borrowerInfo(address borrower_) view returns(uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolSession) BorrowerInfo(borrower_ common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.BorrowerInfo(&_AjnaPool.CallOpts, borrower_)
}

// BorrowerInfo is a free data retrieval call binding the contract method 0xca103d15.
//
// Solidity: function borrowerInfo(address borrower_) view returns(uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) BorrowerInfo(borrower_ common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.BorrowerInfo(&_AjnaPool.CallOpts, borrower_)
}

// BucketCollateralDust is a free data retrieval call binding the contract method 0x540c1433.
//
// Solidity: function bucketCollateralDust(uint256 bucketIndex_) pure returns(uint256)
func (_AjnaPool *AjnaPoolCaller) BucketCollateralDust(opts *bind.CallOpts, bucketIndex_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "bucketCollateralDust", bucketIndex_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BucketCollateralDust is a free data retrieval call binding the contract method 0x540c1433.
//
// Solidity: function bucketCollateralDust(uint256 bucketIndex_) pure returns(uint256)
func (_AjnaPool *AjnaPoolSession) BucketCollateralDust(bucketIndex_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.BucketCollateralDust(&_AjnaPool.CallOpts, bucketIndex_)
}

// BucketCollateralDust is a free data retrieval call binding the contract method 0x540c1433.
//
// Solidity: function bucketCollateralDust(uint256 bucketIndex_) pure returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) BucketCollateralDust(bucketIndex_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.BucketCollateralDust(&_AjnaPool.CallOpts, bucketIndex_)
}

// BucketExchangeRate is a free data retrieval call binding the contract method 0x16f8a463.
//
// Solidity: function bucketExchangeRate(uint256 index_) view returns(uint256 exchangeRate_)
func (_AjnaPool *AjnaPoolCaller) BucketExchangeRate(opts *bind.CallOpts, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "bucketExchangeRate", index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BucketExchangeRate is a free data retrieval call binding the contract method 0x16f8a463.
//
// Solidity: function bucketExchangeRate(uint256 index_) view returns(uint256 exchangeRate_)
func (_AjnaPool *AjnaPoolSession) BucketExchangeRate(index_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.BucketExchangeRate(&_AjnaPool.CallOpts, index_)
}

// BucketExchangeRate is a free data retrieval call binding the contract method 0x16f8a463.
//
// Solidity: function bucketExchangeRate(uint256 index_) view returns(uint256 exchangeRate_)
func (_AjnaPool *AjnaPoolCallerSession) BucketExchangeRate(index_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.BucketExchangeRate(&_AjnaPool.CallOpts, index_)
}

// BucketInfo is a free data retrieval call binding the contract method 0xa83de3ec.
//
// Solidity: function bucketInfo(uint256 index_) view returns(uint256, uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) BucketInfo(opts *bind.CallOpts, index_ *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "bucketInfo", index_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// BucketInfo is a free data retrieval call binding the contract method 0xa83de3ec.
//
// Solidity: function bucketInfo(uint256 index_) view returns(uint256, uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolSession) BucketInfo(index_ *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.BucketInfo(&_AjnaPool.CallOpts, index_)
}

// BucketInfo is a free data retrieval call binding the contract method 0xa83de3ec.
//
// Solidity: function bucketInfo(uint256 index_) view returns(uint256, uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) BucketInfo(index_ *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.BucketInfo(&_AjnaPool.CallOpts, index_)
}

// BurnInfo is a free data retrieval call binding the contract method 0x2c7b2e06.
//
// Solidity: function burnInfo(uint256 burnEventEpoch_) view returns(uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) BurnInfo(opts *bind.CallOpts, burnEventEpoch_ *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "burnInfo", burnEventEpoch_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// BurnInfo is a free data retrieval call binding the contract method 0x2c7b2e06.
//
// Solidity: function burnInfo(uint256 burnEventEpoch_) view returns(uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolSession) BurnInfo(burnEventEpoch_ *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.BurnInfo(&_AjnaPool.CallOpts, burnEventEpoch_)
}

// BurnInfo is a free data retrieval call binding the contract method 0x2c7b2e06.
//
// Solidity: function burnInfo(uint256 burnEventEpoch_) view returns(uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) BurnInfo(burnEventEpoch_ *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.BurnInfo(&_AjnaPool.CallOpts, burnEventEpoch_)
}

// CollateralAddress is a free data retrieval call binding the contract method 0x48d399e7.
//
// Solidity: function collateralAddress() pure returns(address)
func (_AjnaPool *AjnaPoolCaller) CollateralAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "collateralAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralAddress is a free data retrieval call binding the contract method 0x48d399e7.
//
// Solidity: function collateralAddress() pure returns(address)
func (_AjnaPool *AjnaPoolSession) CollateralAddress() (common.Address, error) {
	return _AjnaPool.Contract.CollateralAddress(&_AjnaPool.CallOpts)
}

// CollateralAddress is a free data retrieval call binding the contract method 0x48d399e7.
//
// Solidity: function collateralAddress() pure returns(address)
func (_AjnaPool *AjnaPoolCallerSession) CollateralAddress() (common.Address, error) {
	return _AjnaPool.Contract.CollateralAddress(&_AjnaPool.CallOpts)
}

// CollateralScale is a free data retrieval call binding the contract method 0xec0bdcfc.
//
// Solidity: function collateralScale() pure returns(uint256)
func (_AjnaPool *AjnaPoolCaller) CollateralScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "collateralScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralScale is a free data retrieval call binding the contract method 0xec0bdcfc.
//
// Solidity: function collateralScale() pure returns(uint256)
func (_AjnaPool *AjnaPoolSession) CollateralScale() (*big.Int, error) {
	return _AjnaPool.Contract.CollateralScale(&_AjnaPool.CallOpts)
}

// CollateralScale is a free data retrieval call binding the contract method 0xec0bdcfc.
//
// Solidity: function collateralScale() pure returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) CollateralScale() (*big.Int, error) {
	return _AjnaPool.Contract.CollateralScale(&_AjnaPool.CallOpts)
}

// CurrentBurnEpoch is a free data retrieval call binding the contract method 0x4ab1fc36.
//
// Solidity: function currentBurnEpoch() view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) CurrentBurnEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "currentBurnEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBurnEpoch is a free data retrieval call binding the contract method 0x4ab1fc36.
//
// Solidity: function currentBurnEpoch() view returns(uint256)
func (_AjnaPool *AjnaPoolSession) CurrentBurnEpoch() (*big.Int, error) {
	return _AjnaPool.Contract.CurrentBurnEpoch(&_AjnaPool.CallOpts)
}

// CurrentBurnEpoch is a free data retrieval call binding the contract method 0x4ab1fc36.
//
// Solidity: function currentBurnEpoch() view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) CurrentBurnEpoch() (*big.Int, error) {
	return _AjnaPool.Contract.CurrentBurnEpoch(&_AjnaPool.CallOpts)
}

// DebtInfo is a free data retrieval call binding the contract method 0x4d966198.
//
// Solidity: function debtInfo() view returns(uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) DebtInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "debtInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// DebtInfo is a free data retrieval call binding the contract method 0x4d966198.
//
// Solidity: function debtInfo() view returns(uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolSession) DebtInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.DebtInfo(&_AjnaPool.CallOpts)
}

// DebtInfo is a free data retrieval call binding the contract method 0x4d966198.
//
// Solidity: function debtInfo() view returns(uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) DebtInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.DebtInfo(&_AjnaPool.CallOpts)
}

// DepositIndex is a free data retrieval call binding the contract method 0x329d1a8b.
//
// Solidity: function depositIndex(uint256 debt_) view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) DepositIndex(opts *bind.CallOpts, debt_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "depositIndex", debt_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositIndex is a free data retrieval call binding the contract method 0x329d1a8b.
//
// Solidity: function depositIndex(uint256 debt_) view returns(uint256)
func (_AjnaPool *AjnaPoolSession) DepositIndex(debt_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.DepositIndex(&_AjnaPool.CallOpts, debt_)
}

// DepositIndex is a free data retrieval call binding the contract method 0x329d1a8b.
//
// Solidity: function depositIndex(uint256 debt_) view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) DepositIndex(debt_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.DepositIndex(&_AjnaPool.CallOpts, debt_)
}

// DepositScale is a free data retrieval call binding the contract method 0x3ab96ec5.
//
// Solidity: function depositScale(uint256 index_) view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) DepositScale(opts *bind.CallOpts, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "depositScale", index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositScale is a free data retrieval call binding the contract method 0x3ab96ec5.
//
// Solidity: function depositScale(uint256 index_) view returns(uint256)
func (_AjnaPool *AjnaPoolSession) DepositScale(index_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.DepositScale(&_AjnaPool.CallOpts, index_)
}

// DepositScale is a free data retrieval call binding the contract method 0x3ab96ec5.
//
// Solidity: function depositScale(uint256 index_) view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) DepositScale(index_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.DepositScale(&_AjnaPool.CallOpts, index_)
}

// DepositSize is a free data retrieval call binding the contract method 0x3fa8fdbb.
//
// Solidity: function depositSize() view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) DepositSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "depositSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositSize is a free data retrieval call binding the contract method 0x3fa8fdbb.
//
// Solidity: function depositSize() view returns(uint256)
func (_AjnaPool *AjnaPoolSession) DepositSize() (*big.Int, error) {
	return _AjnaPool.Contract.DepositSize(&_AjnaPool.CallOpts)
}

// DepositSize is a free data retrieval call binding the contract method 0x3fa8fdbb.
//
// Solidity: function depositSize() view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) DepositSize() (*big.Int, error) {
	return _AjnaPool.Contract.DepositSize(&_AjnaPool.CallOpts)
}

// DepositUpToIndex is a free data retrieval call binding the contract method 0xda7951a9.
//
// Solidity: function depositUpToIndex(uint256 index_) view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) DepositUpToIndex(opts *bind.CallOpts, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "depositUpToIndex", index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositUpToIndex is a free data retrieval call binding the contract method 0xda7951a9.
//
// Solidity: function depositUpToIndex(uint256 index_) view returns(uint256)
func (_AjnaPool *AjnaPoolSession) DepositUpToIndex(index_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.DepositUpToIndex(&_AjnaPool.CallOpts, index_)
}

// DepositUpToIndex is a free data retrieval call binding the contract method 0xda7951a9.
//
// Solidity: function depositUpToIndex(uint256 index_) view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) DepositUpToIndex(index_ *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.DepositUpToIndex(&_AjnaPool.CallOpts, index_)
}

// DepositUtilization is a free data retrieval call binding the contract method 0x3a0c8f07.
//
// Solidity: function depositUtilization() view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) DepositUtilization(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "depositUtilization")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositUtilization is a free data retrieval call binding the contract method 0x3a0c8f07.
//
// Solidity: function depositUtilization() view returns(uint256)
func (_AjnaPool *AjnaPoolSession) DepositUtilization() (*big.Int, error) {
	return _AjnaPool.Contract.DepositUtilization(&_AjnaPool.CallOpts)
}

// DepositUtilization is a free data retrieval call binding the contract method 0x3a0c8f07.
//
// Solidity: function depositUtilization() view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) DepositUtilization() (*big.Int, error) {
	return _AjnaPool.Contract.DepositUtilization(&_AjnaPool.CallOpts)
}

// EmasInfo is a free data retrieval call binding the contract method 0xe512c061.
//
// Solidity: function emasInfo() view returns(uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) EmasInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "emasInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// EmasInfo is a free data retrieval call binding the contract method 0xe512c061.
//
// Solidity: function emasInfo() view returns(uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolSession) EmasInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.EmasInfo(&_AjnaPool.CallOpts)
}

// EmasInfo is a free data retrieval call binding the contract method 0xe512c061.
//
// Solidity: function emasInfo() view returns(uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) EmasInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.EmasInfo(&_AjnaPool.CallOpts)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token_, uint256 ) view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) FlashFee(opts *bind.CallOpts, token_ common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "flashFee", token_, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token_, uint256 ) view returns(uint256)
func (_AjnaPool *AjnaPoolSession) FlashFee(token_ common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.FlashFee(&_AjnaPool.CallOpts, token_, arg1)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token_, uint256 ) view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) FlashFee(token_ common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AjnaPool.Contract.FlashFee(&_AjnaPool.CallOpts, token_, arg1)
}

// InflatorInfo is a free data retrieval call binding the contract method 0x063d829f.
//
// Solidity: function inflatorInfo() view returns(uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) InflatorInfo(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "inflatorInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// InflatorInfo is a free data retrieval call binding the contract method 0x063d829f.
//
// Solidity: function inflatorInfo() view returns(uint256, uint256)
func (_AjnaPool *AjnaPoolSession) InflatorInfo() (*big.Int, *big.Int, error) {
	return _AjnaPool.Contract.InflatorInfo(&_AjnaPool.CallOpts)
}

// InflatorInfo is a free data retrieval call binding the contract method 0x063d829f.
//
// Solidity: function inflatorInfo() view returns(uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) InflatorInfo() (*big.Int, *big.Int, error) {
	return _AjnaPool.Contract.InflatorInfo(&_AjnaPool.CallOpts)
}

// InterestRateInfo is a free data retrieval call binding the contract method 0x00cdcefb.
//
// Solidity: function interestRateInfo() view returns(uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) InterestRateInfo(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "interestRateInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// InterestRateInfo is a free data retrieval call binding the contract method 0x00cdcefb.
//
// Solidity: function interestRateInfo() view returns(uint256, uint256)
func (_AjnaPool *AjnaPoolSession) InterestRateInfo() (*big.Int, *big.Int, error) {
	return _AjnaPool.Contract.InterestRateInfo(&_AjnaPool.CallOpts)
}

// InterestRateInfo is a free data retrieval call binding the contract method 0x00cdcefb.
//
// Solidity: function interestRateInfo() view returns(uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) InterestRateInfo() (*big.Int, *big.Int, error) {
	return _AjnaPool.Contract.InterestRateInfo(&_AjnaPool.CallOpts)
}

// KickerInfo is a free data retrieval call binding the contract method 0x7323f853.
//
// Solidity: function kickerInfo(address kicker_) view returns(uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) KickerInfo(opts *bind.CallOpts, kicker_ common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "kickerInfo", kicker_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// KickerInfo is a free data retrieval call binding the contract method 0x7323f853.
//
// Solidity: function kickerInfo(address kicker_) view returns(uint256, uint256)
func (_AjnaPool *AjnaPoolSession) KickerInfo(kicker_ common.Address) (*big.Int, *big.Int, error) {
	return _AjnaPool.Contract.KickerInfo(&_AjnaPool.CallOpts, kicker_)
}

// KickerInfo is a free data retrieval call binding the contract method 0x7323f853.
//
// Solidity: function kickerInfo(address kicker_) view returns(uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) KickerInfo(kicker_ common.Address) (*big.Int, *big.Int, error) {
	return _AjnaPool.Contract.KickerInfo(&_AjnaPool.CallOpts, kicker_)
}

// LenderInfo is a free data retrieval call binding the contract method 0xa749f1a6.
//
// Solidity: function lenderInfo(uint256 index_, address lender_) view returns(uint256 lpBalance_, uint256 depositTime_)
func (_AjnaPool *AjnaPoolCaller) LenderInfo(opts *bind.CallOpts, index_ *big.Int, lender_ common.Address) (struct {
	LpBalance   *big.Int
	DepositTime *big.Int
}, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "lenderInfo", index_, lender_)

	outstruct := new(struct {
		LpBalance   *big.Int
		DepositTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DepositTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LenderInfo is a free data retrieval call binding the contract method 0xa749f1a6.
//
// Solidity: function lenderInfo(uint256 index_, address lender_) view returns(uint256 lpBalance_, uint256 depositTime_)
func (_AjnaPool *AjnaPoolSession) LenderInfo(index_ *big.Int, lender_ common.Address) (struct {
	LpBalance   *big.Int
	DepositTime *big.Int
}, error) {
	return _AjnaPool.Contract.LenderInfo(&_AjnaPool.CallOpts, index_, lender_)
}

// LenderInfo is a free data retrieval call binding the contract method 0xa749f1a6.
//
// Solidity: function lenderInfo(uint256 index_, address lender_) view returns(uint256 lpBalance_, uint256 depositTime_)
func (_AjnaPool *AjnaPoolCallerSession) LenderInfo(index_ *big.Int, lender_ common.Address) (struct {
	LpBalance   *big.Int
	DepositTime *big.Int
}, error) {
	return _AjnaPool.Contract.LenderInfo(&_AjnaPool.CallOpts, index_, lender_)
}

// LoanInfo is a free data retrieval call binding the contract method 0x8349d6be.
//
// Solidity: function loanInfo(uint256 loanId_) view returns(address, uint256)
func (_AjnaPool *AjnaPoolCaller) LoanInfo(opts *bind.CallOpts, loanId_ *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "loanInfo", loanId_)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LoanInfo is a free data retrieval call binding the contract method 0x8349d6be.
//
// Solidity: function loanInfo(uint256 loanId_) view returns(address, uint256)
func (_AjnaPool *AjnaPoolSession) LoanInfo(loanId_ *big.Int) (common.Address, *big.Int, error) {
	return _AjnaPool.Contract.LoanInfo(&_AjnaPool.CallOpts, loanId_)
}

// LoanInfo is a free data retrieval call binding the contract method 0x8349d6be.
//
// Solidity: function loanInfo(uint256 loanId_) view returns(address, uint256)
func (_AjnaPool *AjnaPoolCallerSession) LoanInfo(loanId_ *big.Int) (common.Address, *big.Int, error) {
	return _AjnaPool.Contract.LoanInfo(&_AjnaPool.CallOpts, loanId_)
}

// LoansInfo is a free data retrieval call binding the contract method 0x3884cd88.
//
// Solidity: function loansInfo() view returns(address, uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) LoansInfo(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "loansInfo")

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// LoansInfo is a free data retrieval call binding the contract method 0x3884cd88.
//
// Solidity: function loansInfo() view returns(address, uint256, uint256)
func (_AjnaPool *AjnaPoolSession) LoansInfo() (common.Address, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.LoansInfo(&_AjnaPool.CallOpts)
}

// LoansInfo is a free data retrieval call binding the contract method 0x3884cd88.
//
// Solidity: function loansInfo() view returns(address, uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) LoansInfo() (common.Address, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.LoansInfo(&_AjnaPool.CallOpts)
}

// LpAllowance is a free data retrieval call binding the contract method 0x483cd187.
//
// Solidity: function lpAllowance(uint256 index_, address spender_, address owner_) view returns(uint256 allowance_)
func (_AjnaPool *AjnaPoolCaller) LpAllowance(opts *bind.CallOpts, index_ *big.Int, spender_ common.Address, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "lpAllowance", index_, spender_, owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpAllowance is a free data retrieval call binding the contract method 0x483cd187.
//
// Solidity: function lpAllowance(uint256 index_, address spender_, address owner_) view returns(uint256 allowance_)
func (_AjnaPool *AjnaPoolSession) LpAllowance(index_ *big.Int, spender_ common.Address, owner_ common.Address) (*big.Int, error) {
	return _AjnaPool.Contract.LpAllowance(&_AjnaPool.CallOpts, index_, spender_, owner_)
}

// LpAllowance is a free data retrieval call binding the contract method 0x483cd187.
//
// Solidity: function lpAllowance(uint256 index_, address spender_, address owner_) view returns(uint256 allowance_)
func (_AjnaPool *AjnaPoolCallerSession) LpAllowance(index_ *big.Int, spender_ common.Address, owner_ common.Address) (*big.Int, error) {
	return _AjnaPool.Contract.LpAllowance(&_AjnaPool.CallOpts, index_, spender_, owner_)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token_) view returns(uint256 maxLoan_)
func (_AjnaPool *AjnaPoolCaller) MaxFlashLoan(opts *bind.CallOpts, token_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "maxFlashLoan", token_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token_) view returns(uint256 maxLoan_)
func (_AjnaPool *AjnaPoolSession) MaxFlashLoan(token_ common.Address) (*big.Int, error) {
	return _AjnaPool.Contract.MaxFlashLoan(&_AjnaPool.CallOpts, token_)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token_) view returns(uint256 maxLoan_)
func (_AjnaPool *AjnaPoolCallerSession) MaxFlashLoan(token_ common.Address) (*big.Int, error) {
	return _AjnaPool.Contract.MaxFlashLoan(&_AjnaPool.CallOpts, token_)
}

// PledgedCollateral is a free data retrieval call binding the contract method 0x307ee3b5.
//
// Solidity: function pledgedCollateral() view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) PledgedCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "pledgedCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgedCollateral is a free data retrieval call binding the contract method 0x307ee3b5.
//
// Solidity: function pledgedCollateral() view returns(uint256)
func (_AjnaPool *AjnaPoolSession) PledgedCollateral() (*big.Int, error) {
	return _AjnaPool.Contract.PledgedCollateral(&_AjnaPool.CallOpts)
}

// PledgedCollateral is a free data retrieval call binding the contract method 0x307ee3b5.
//
// Solidity: function pledgedCollateral() view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) PledgedCollateral() (*big.Int, error) {
	return _AjnaPool.Contract.PledgedCollateral(&_AjnaPool.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() pure returns(uint8)
func (_AjnaPool *AjnaPoolCaller) PoolType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() pure returns(uint8)
func (_AjnaPool *AjnaPoolSession) PoolType() (uint8, error) {
	return _AjnaPool.Contract.PoolType(&_AjnaPool.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() pure returns(uint8)
func (_AjnaPool *AjnaPoolCallerSession) PoolType() (uint8, error) {
	return _AjnaPool.Contract.PoolType(&_AjnaPool.CallOpts)
}

// QuoteTokenAddress is a free data retrieval call binding the contract method 0xbad34620.
//
// Solidity: function quoteTokenAddress() pure returns(address)
func (_AjnaPool *AjnaPoolCaller) QuoteTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "quoteTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteTokenAddress is a free data retrieval call binding the contract method 0xbad34620.
//
// Solidity: function quoteTokenAddress() pure returns(address)
func (_AjnaPool *AjnaPoolSession) QuoteTokenAddress() (common.Address, error) {
	return _AjnaPool.Contract.QuoteTokenAddress(&_AjnaPool.CallOpts)
}

// QuoteTokenAddress is a free data retrieval call binding the contract method 0xbad34620.
//
// Solidity: function quoteTokenAddress() pure returns(address)
func (_AjnaPool *AjnaPoolCallerSession) QuoteTokenAddress() (common.Address, error) {
	return _AjnaPool.Contract.QuoteTokenAddress(&_AjnaPool.CallOpts)
}

// QuoteTokenScale is a free data retrieval call binding the contract method 0x7b3f8655.
//
// Solidity: function quoteTokenScale() pure returns(uint256)
func (_AjnaPool *AjnaPoolCaller) QuoteTokenScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "quoteTokenScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteTokenScale is a free data retrieval call binding the contract method 0x7b3f8655.
//
// Solidity: function quoteTokenScale() pure returns(uint256)
func (_AjnaPool *AjnaPoolSession) QuoteTokenScale() (*big.Int, error) {
	return _AjnaPool.Contract.QuoteTokenScale(&_AjnaPool.CallOpts)
}

// QuoteTokenScale is a free data retrieval call binding the contract method 0x7b3f8655.
//
// Solidity: function quoteTokenScale() pure returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) QuoteTokenScale() (*big.Int, error) {
	return _AjnaPool.Contract.QuoteTokenScale(&_AjnaPool.CallOpts)
}

// ReservesInfo is a free data retrieval call binding the contract method 0x5a3b4477.
//
// Solidity: function reservesInfo() view returns(uint256, uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCaller) ReservesInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "reservesInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// ReservesInfo is a free data retrieval call binding the contract method 0x5a3b4477.
//
// Solidity: function reservesInfo() view returns(uint256, uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolSession) ReservesInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.ReservesInfo(&_AjnaPool.CallOpts)
}

// ReservesInfo is a free data retrieval call binding the contract method 0x5a3b4477.
//
// Solidity: function reservesInfo() view returns(uint256, uint256, uint256, uint256, uint256)
func (_AjnaPool *AjnaPoolCallerSession) ReservesInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AjnaPool.Contract.ReservesInfo(&_AjnaPool.CallOpts)
}

// TotalAuctionsInPool is a free data retrieval call binding the contract method 0xbcb630d7.
//
// Solidity: function totalAuctionsInPool() view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) TotalAuctionsInPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "totalAuctionsInPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAuctionsInPool is a free data retrieval call binding the contract method 0xbcb630d7.
//
// Solidity: function totalAuctionsInPool() view returns(uint256)
func (_AjnaPool *AjnaPoolSession) TotalAuctionsInPool() (*big.Int, error) {
	return _AjnaPool.Contract.TotalAuctionsInPool(&_AjnaPool.CallOpts)
}

// TotalAuctionsInPool is a free data retrieval call binding the contract method 0xbcb630d7.
//
// Solidity: function totalAuctionsInPool() view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) TotalAuctionsInPool() (*big.Int, error) {
	return _AjnaPool.Contract.TotalAuctionsInPool(&_AjnaPool.CallOpts)
}

// TotalT0Debt is a free data retrieval call binding the contract method 0x5d3637e7.
//
// Solidity: function totalT0Debt() view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) TotalT0Debt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "totalT0Debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalT0Debt is a free data retrieval call binding the contract method 0x5d3637e7.
//
// Solidity: function totalT0Debt() view returns(uint256)
func (_AjnaPool *AjnaPoolSession) TotalT0Debt() (*big.Int, error) {
	return _AjnaPool.Contract.TotalT0Debt(&_AjnaPool.CallOpts)
}

// TotalT0Debt is a free data retrieval call binding the contract method 0x5d3637e7.
//
// Solidity: function totalT0Debt() view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) TotalT0Debt() (*big.Int, error) {
	return _AjnaPool.Contract.TotalT0Debt(&_AjnaPool.CallOpts)
}

// TotalT0DebtInAuction is a free data retrieval call binding the contract method 0x870c764a.
//
// Solidity: function totalT0DebtInAuction() view returns(uint256)
func (_AjnaPool *AjnaPoolCaller) TotalT0DebtInAuction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPool.contract.Call(opts, &out, "totalT0DebtInAuction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalT0DebtInAuction is a free data retrieval call binding the contract method 0x870c764a.
//
// Solidity: function totalT0DebtInAuction() view returns(uint256)
func (_AjnaPool *AjnaPoolSession) TotalT0DebtInAuction() (*big.Int, error) {
	return _AjnaPool.Contract.TotalT0DebtInAuction(&_AjnaPool.CallOpts)
}

// TotalT0DebtInAuction is a free data retrieval call binding the contract method 0x870c764a.
//
// Solidity: function totalT0DebtInAuction() view returns(uint256)
func (_AjnaPool *AjnaPoolCallerSession) TotalT0DebtInAuction() (*big.Int, error) {
	return _AjnaPool.Contract.TotalT0DebtInAuction(&_AjnaPool.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xc861c6e6.
//
// Solidity: function addCollateral(uint256 amountToAdd_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_)
func (_AjnaPool *AjnaPoolTransactor) AddCollateral(opts *bind.TransactOpts, amountToAdd_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "addCollateral", amountToAdd_, index_, expiry_)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xc861c6e6.
//
// Solidity: function addCollateral(uint256 amountToAdd_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_)
func (_AjnaPool *AjnaPoolSession) AddCollateral(amountToAdd_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.AddCollateral(&_AjnaPool.TransactOpts, amountToAdd_, index_, expiry_)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xc861c6e6.
//
// Solidity: function addCollateral(uint256 amountToAdd_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_)
func (_AjnaPool *AjnaPoolTransactorSession) AddCollateral(amountToAdd_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.AddCollateral(&_AjnaPool.TransactOpts, amountToAdd_, index_, expiry_)
}

// AddQuoteToken is a paid mutator transaction binding the contract method 0xf78b0cce.
//
// Solidity: function addQuoteToken(uint256 amount_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_, uint256 addedAmount_)
func (_AjnaPool *AjnaPoolTransactor) AddQuoteToken(opts *bind.TransactOpts, amount_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "addQuoteToken", amount_, index_, expiry_)
}

// AddQuoteToken is a paid mutator transaction binding the contract method 0xf78b0cce.
//
// Solidity: function addQuoteToken(uint256 amount_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_, uint256 addedAmount_)
func (_AjnaPool *AjnaPoolSession) AddQuoteToken(amount_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.AddQuoteToken(&_AjnaPool.TransactOpts, amount_, index_, expiry_)
}

// AddQuoteToken is a paid mutator transaction binding the contract method 0xf78b0cce.
//
// Solidity: function addQuoteToken(uint256 amount_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_, uint256 addedAmount_)
func (_AjnaPool *AjnaPoolTransactorSession) AddQuoteToken(amount_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.AddQuoteToken(&_AjnaPool.TransactOpts, amount_, index_, expiry_)
}

// ApproveLPTransferors is a paid mutator transaction binding the contract method 0x7f8baa37.
//
// Solidity: function approveLPTransferors(address[] transferors_) returns()
func (_AjnaPool *AjnaPoolTransactor) ApproveLPTransferors(opts *bind.TransactOpts, transferors_ []common.Address) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "approveLPTransferors", transferors_)
}

// ApproveLPTransferors is a paid mutator transaction binding the contract method 0x7f8baa37.
//
// Solidity: function approveLPTransferors(address[] transferors_) returns()
func (_AjnaPool *AjnaPoolSession) ApproveLPTransferors(transferors_ []common.Address) (*types.Transaction, error) {
	return _AjnaPool.Contract.ApproveLPTransferors(&_AjnaPool.TransactOpts, transferors_)
}

// ApproveLPTransferors is a paid mutator transaction binding the contract method 0x7f8baa37.
//
// Solidity: function approveLPTransferors(address[] transferors_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) ApproveLPTransferors(transferors_ []common.Address) (*types.Transaction, error) {
	return _AjnaPool.Contract.ApproveLPTransferors(&_AjnaPool.TransactOpts, transferors_)
}

// BucketTake is a paid mutator transaction binding the contract method 0x0729f62c.
//
// Solidity: function bucketTake(address borrowerAddress_, bool depositTake_, uint256 index_) returns()
func (_AjnaPool *AjnaPoolTransactor) BucketTake(opts *bind.TransactOpts, borrowerAddress_ common.Address, depositTake_ bool, index_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "bucketTake", borrowerAddress_, depositTake_, index_)
}

// BucketTake is a paid mutator transaction binding the contract method 0x0729f62c.
//
// Solidity: function bucketTake(address borrowerAddress_, bool depositTake_, uint256 index_) returns()
func (_AjnaPool *AjnaPoolSession) BucketTake(borrowerAddress_ common.Address, depositTake_ bool, index_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.BucketTake(&_AjnaPool.TransactOpts, borrowerAddress_, depositTake_, index_)
}

// BucketTake is a paid mutator transaction binding the contract method 0x0729f62c.
//
// Solidity: function bucketTake(address borrowerAddress_, bool depositTake_, uint256 index_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) BucketTake(borrowerAddress_ common.Address, depositTake_ bool, index_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.BucketTake(&_AjnaPool.TransactOpts, borrowerAddress_, depositTake_, index_)
}

// DecreaseLPAllowance is a paid mutator transaction binding the contract method 0x987165ed.
//
// Solidity: function decreaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_AjnaPool *AjnaPoolTransactor) DecreaseLPAllowance(opts *bind.TransactOpts, spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "decreaseLPAllowance", spender_, indexes_, amounts_)
}

// DecreaseLPAllowance is a paid mutator transaction binding the contract method 0x987165ed.
//
// Solidity: function decreaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_AjnaPool *AjnaPoolSession) DecreaseLPAllowance(spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.DecreaseLPAllowance(&_AjnaPool.TransactOpts, spender_, indexes_, amounts_)
}

// DecreaseLPAllowance is a paid mutator transaction binding the contract method 0x987165ed.
//
// Solidity: function decreaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) DecreaseLPAllowance(spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.DecreaseLPAllowance(&_AjnaPool.TransactOpts, spender_, indexes_, amounts_)
}

// DrawDebt is a paid mutator transaction binding the contract method 0xcfa8ff03.
//
// Solidity: function drawDebt(address borrowerAddress_, uint256 amountToBorrow_, uint256 limitIndex_, uint256 collateralToPledge_) returns()
func (_AjnaPool *AjnaPoolTransactor) DrawDebt(opts *bind.TransactOpts, borrowerAddress_ common.Address, amountToBorrow_ *big.Int, limitIndex_ *big.Int, collateralToPledge_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "drawDebt", borrowerAddress_, amountToBorrow_, limitIndex_, collateralToPledge_)
}

// DrawDebt is a paid mutator transaction binding the contract method 0xcfa8ff03.
//
// Solidity: function drawDebt(address borrowerAddress_, uint256 amountToBorrow_, uint256 limitIndex_, uint256 collateralToPledge_) returns()
func (_AjnaPool *AjnaPoolSession) DrawDebt(borrowerAddress_ common.Address, amountToBorrow_ *big.Int, limitIndex_ *big.Int, collateralToPledge_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.DrawDebt(&_AjnaPool.TransactOpts, borrowerAddress_, amountToBorrow_, limitIndex_, collateralToPledge_)
}

// DrawDebt is a paid mutator transaction binding the contract method 0xcfa8ff03.
//
// Solidity: function drawDebt(address borrowerAddress_, uint256 amountToBorrow_, uint256 limitIndex_, uint256 collateralToPledge_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) DrawDebt(borrowerAddress_ common.Address, amountToBorrow_ *big.Int, limitIndex_ *big.Int, collateralToPledge_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.DrawDebt(&_AjnaPool.TransactOpts, borrowerAddress_, amountToBorrow_, limitIndex_, collateralToPledge_)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver_, address token_, uint256 amount_, bytes data_) returns(bool)
func (_AjnaPool *AjnaPoolTransactor) FlashLoan(opts *bind.TransactOpts, receiver_ common.Address, token_ common.Address, amount_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "flashLoan", receiver_, token_, amount_, data_)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver_, address token_, uint256 amount_, bytes data_) returns(bool)
func (_AjnaPool *AjnaPoolSession) FlashLoan(receiver_ common.Address, token_ common.Address, amount_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _AjnaPool.Contract.FlashLoan(&_AjnaPool.TransactOpts, receiver_, token_, amount_, data_)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver_, address token_, uint256 amount_, bytes data_) returns(bool)
func (_AjnaPool *AjnaPoolTransactorSession) FlashLoan(receiver_ common.Address, token_ common.Address, amount_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _AjnaPool.Contract.FlashLoan(&_AjnaPool.TransactOpts, receiver_, token_, amount_, data_)
}

// IncreaseLPAllowance is a paid mutator transaction binding the contract method 0xa918058d.
//
// Solidity: function increaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_AjnaPool *AjnaPoolTransactor) IncreaseLPAllowance(opts *bind.TransactOpts, spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "increaseLPAllowance", spender_, indexes_, amounts_)
}

// IncreaseLPAllowance is a paid mutator transaction binding the contract method 0xa918058d.
//
// Solidity: function increaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_AjnaPool *AjnaPoolSession) IncreaseLPAllowance(spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.IncreaseLPAllowance(&_AjnaPool.TransactOpts, spender_, indexes_, amounts_)
}

// IncreaseLPAllowance is a paid mutator transaction binding the contract method 0xa918058d.
//
// Solidity: function increaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) IncreaseLPAllowance(spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.IncreaseLPAllowance(&_AjnaPool.TransactOpts, spender_, indexes_, amounts_)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 rate_) returns()
func (_AjnaPool *AjnaPoolTransactor) Initialize(opts *bind.TransactOpts, rate_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "initialize", rate_)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 rate_) returns()
func (_AjnaPool *AjnaPoolSession) Initialize(rate_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.Initialize(&_AjnaPool.TransactOpts, rate_)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 rate_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) Initialize(rate_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.Initialize(&_AjnaPool.TransactOpts, rate_)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address borrower_, uint256 npLimitIndex_) returns()
func (_AjnaPool *AjnaPoolTransactor) Kick(opts *bind.TransactOpts, borrower_ common.Address, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "kick", borrower_, npLimitIndex_)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address borrower_, uint256 npLimitIndex_) returns()
func (_AjnaPool *AjnaPoolSession) Kick(borrower_ common.Address, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.Kick(&_AjnaPool.TransactOpts, borrower_, npLimitIndex_)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address borrower_, uint256 npLimitIndex_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) Kick(borrower_ common.Address, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.Kick(&_AjnaPool.TransactOpts, borrower_, npLimitIndex_)
}

// KickReserveAuction is a paid mutator transaction binding the contract method 0x5a422b92.
//
// Solidity: function kickReserveAuction() returns()
func (_AjnaPool *AjnaPoolTransactor) KickReserveAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "kickReserveAuction")
}

// KickReserveAuction is a paid mutator transaction binding the contract method 0x5a422b92.
//
// Solidity: function kickReserveAuction() returns()
func (_AjnaPool *AjnaPoolSession) KickReserveAuction() (*types.Transaction, error) {
	return _AjnaPool.Contract.KickReserveAuction(&_AjnaPool.TransactOpts)
}

// KickReserveAuction is a paid mutator transaction binding the contract method 0x5a422b92.
//
// Solidity: function kickReserveAuction() returns()
func (_AjnaPool *AjnaPoolTransactorSession) KickReserveAuction() (*types.Transaction, error) {
	return _AjnaPool.Contract.KickReserveAuction(&_AjnaPool.TransactOpts)
}

// LenderKick is a paid mutator transaction binding the contract method 0xeca48706.
//
// Solidity: function lenderKick(uint256 index_, uint256 npLimitIndex_) returns()
func (_AjnaPool *AjnaPoolTransactor) LenderKick(opts *bind.TransactOpts, index_ *big.Int, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "lenderKick", index_, npLimitIndex_)
}

// LenderKick is a paid mutator transaction binding the contract method 0xeca48706.
//
// Solidity: function lenderKick(uint256 index_, uint256 npLimitIndex_) returns()
func (_AjnaPool *AjnaPoolSession) LenderKick(index_ *big.Int, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.LenderKick(&_AjnaPool.TransactOpts, index_, npLimitIndex_)
}

// LenderKick is a paid mutator transaction binding the contract method 0xeca48706.
//
// Solidity: function lenderKick(uint256 index_, uint256 npLimitIndex_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) LenderKick(index_ *big.Int, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.LenderKick(&_AjnaPool.TransactOpts, index_, npLimitIndex_)
}

// MoveQuoteToken is a paid mutator transaction binding the contract method 0x332c0e43.
//
// Solidity: function moveQuoteToken(uint256 maxAmount_, uint256 fromIndex_, uint256 toIndex_, uint256 expiry_) returns(uint256 fromBucketLP_, uint256 toBucketLP_, uint256 movedAmount_)
func (_AjnaPool *AjnaPoolTransactor) MoveQuoteToken(opts *bind.TransactOpts, maxAmount_ *big.Int, fromIndex_ *big.Int, toIndex_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "moveQuoteToken", maxAmount_, fromIndex_, toIndex_, expiry_)
}

// MoveQuoteToken is a paid mutator transaction binding the contract method 0x332c0e43.
//
// Solidity: function moveQuoteToken(uint256 maxAmount_, uint256 fromIndex_, uint256 toIndex_, uint256 expiry_) returns(uint256 fromBucketLP_, uint256 toBucketLP_, uint256 movedAmount_)
func (_AjnaPool *AjnaPoolSession) MoveQuoteToken(maxAmount_ *big.Int, fromIndex_ *big.Int, toIndex_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.MoveQuoteToken(&_AjnaPool.TransactOpts, maxAmount_, fromIndex_, toIndex_, expiry_)
}

// MoveQuoteToken is a paid mutator transaction binding the contract method 0x332c0e43.
//
// Solidity: function moveQuoteToken(uint256 maxAmount_, uint256 fromIndex_, uint256 toIndex_, uint256 expiry_) returns(uint256 fromBucketLP_, uint256 toBucketLP_, uint256 movedAmount_)
func (_AjnaPool *AjnaPoolTransactorSession) MoveQuoteToken(maxAmount_ *big.Int, fromIndex_ *big.Int, toIndex_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.MoveQuoteToken(&_AjnaPool.TransactOpts, maxAmount_, fromIndex_, toIndex_, expiry_)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AjnaPool *AjnaPoolTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AjnaPool *AjnaPoolSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AjnaPool.Contract.Multicall(&_AjnaPool.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AjnaPool *AjnaPoolTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AjnaPool.Contract.Multicall(&_AjnaPool.TransactOpts, data)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0x6a9b1891.
//
// Solidity: function removeCollateral(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_AjnaPool *AjnaPoolTransactor) RemoveCollateral(opts *bind.TransactOpts, maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "removeCollateral", maxAmount_, index_)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0x6a9b1891.
//
// Solidity: function removeCollateral(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_AjnaPool *AjnaPoolSession) RemoveCollateral(maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.RemoveCollateral(&_AjnaPool.TransactOpts, maxAmount_, index_)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0x6a9b1891.
//
// Solidity: function removeCollateral(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_AjnaPool *AjnaPoolTransactorSession) RemoveCollateral(maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.RemoveCollateral(&_AjnaPool.TransactOpts, maxAmount_, index_)
}

// RemoveQuoteToken is a paid mutator transaction binding the contract method 0xb1f07247.
//
// Solidity: function removeQuoteToken(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_AjnaPool *AjnaPoolTransactor) RemoveQuoteToken(opts *bind.TransactOpts, maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "removeQuoteToken", maxAmount_, index_)
}

// RemoveQuoteToken is a paid mutator transaction binding the contract method 0xb1f07247.
//
// Solidity: function removeQuoteToken(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_AjnaPool *AjnaPoolSession) RemoveQuoteToken(maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.RemoveQuoteToken(&_AjnaPool.TransactOpts, maxAmount_, index_)
}

// RemoveQuoteToken is a paid mutator transaction binding the contract method 0xb1f07247.
//
// Solidity: function removeQuoteToken(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_AjnaPool *AjnaPoolTransactorSession) RemoveQuoteToken(maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.RemoveQuoteToken(&_AjnaPool.TransactOpts, maxAmount_, index_)
}

// RepayDebt is a paid mutator transaction binding the contract method 0xa9ff9f77.
//
// Solidity: function repayDebt(address borrowerAddress_, uint256 maxQuoteTokenAmountToRepay_, uint256 collateralAmountToPull_, address collateralReceiver_, uint256 limitIndex_) returns(uint256 amountRepaid_)
func (_AjnaPool *AjnaPoolTransactor) RepayDebt(opts *bind.TransactOpts, borrowerAddress_ common.Address, maxQuoteTokenAmountToRepay_ *big.Int, collateralAmountToPull_ *big.Int, collateralReceiver_ common.Address, limitIndex_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "repayDebt", borrowerAddress_, maxQuoteTokenAmountToRepay_, collateralAmountToPull_, collateralReceiver_, limitIndex_)
}

// RepayDebt is a paid mutator transaction binding the contract method 0xa9ff9f77.
//
// Solidity: function repayDebt(address borrowerAddress_, uint256 maxQuoteTokenAmountToRepay_, uint256 collateralAmountToPull_, address collateralReceiver_, uint256 limitIndex_) returns(uint256 amountRepaid_)
func (_AjnaPool *AjnaPoolSession) RepayDebt(borrowerAddress_ common.Address, maxQuoteTokenAmountToRepay_ *big.Int, collateralAmountToPull_ *big.Int, collateralReceiver_ common.Address, limitIndex_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.RepayDebt(&_AjnaPool.TransactOpts, borrowerAddress_, maxQuoteTokenAmountToRepay_, collateralAmountToPull_, collateralReceiver_, limitIndex_)
}

// RepayDebt is a paid mutator transaction binding the contract method 0xa9ff9f77.
//
// Solidity: function repayDebt(address borrowerAddress_, uint256 maxQuoteTokenAmountToRepay_, uint256 collateralAmountToPull_, address collateralReceiver_, uint256 limitIndex_) returns(uint256 amountRepaid_)
func (_AjnaPool *AjnaPoolTransactorSession) RepayDebt(borrowerAddress_ common.Address, maxQuoteTokenAmountToRepay_ *big.Int, collateralAmountToPull_ *big.Int, collateralReceiver_ common.Address, limitIndex_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.RepayDebt(&_AjnaPool.TransactOpts, borrowerAddress_, maxQuoteTokenAmountToRepay_, collateralAmountToPull_, collateralReceiver_, limitIndex_)
}

// RevokeLPAllowance is a paid mutator transaction binding the contract method 0x06e47f26.
//
// Solidity: function revokeLPAllowance(address spender_, uint256[] indexes_) returns()
func (_AjnaPool *AjnaPoolTransactor) RevokeLPAllowance(opts *bind.TransactOpts, spender_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "revokeLPAllowance", spender_, indexes_)
}

// RevokeLPAllowance is a paid mutator transaction binding the contract method 0x06e47f26.
//
// Solidity: function revokeLPAllowance(address spender_, uint256[] indexes_) returns()
func (_AjnaPool *AjnaPoolSession) RevokeLPAllowance(spender_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.RevokeLPAllowance(&_AjnaPool.TransactOpts, spender_, indexes_)
}

// RevokeLPAllowance is a paid mutator transaction binding the contract method 0x06e47f26.
//
// Solidity: function revokeLPAllowance(address spender_, uint256[] indexes_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) RevokeLPAllowance(spender_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.RevokeLPAllowance(&_AjnaPool.TransactOpts, spender_, indexes_)
}

// RevokeLPTransferors is a paid mutator transaction binding the contract method 0xd39d813f.
//
// Solidity: function revokeLPTransferors(address[] transferors_) returns()
func (_AjnaPool *AjnaPoolTransactor) RevokeLPTransferors(opts *bind.TransactOpts, transferors_ []common.Address) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "revokeLPTransferors", transferors_)
}

// RevokeLPTransferors is a paid mutator transaction binding the contract method 0xd39d813f.
//
// Solidity: function revokeLPTransferors(address[] transferors_) returns()
func (_AjnaPool *AjnaPoolSession) RevokeLPTransferors(transferors_ []common.Address) (*types.Transaction, error) {
	return _AjnaPool.Contract.RevokeLPTransferors(&_AjnaPool.TransactOpts, transferors_)
}

// RevokeLPTransferors is a paid mutator transaction binding the contract method 0xd39d813f.
//
// Solidity: function revokeLPTransferors(address[] transferors_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) RevokeLPTransferors(transferors_ []common.Address) (*types.Transaction, error) {
	return _AjnaPool.Contract.RevokeLPTransferors(&_AjnaPool.TransactOpts, transferors_)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(address borrowerAddress_, uint256 maxDepth_) returns(uint256 collateralSettled_, bool isBorrowerSettled_)
func (_AjnaPool *AjnaPoolTransactor) Settle(opts *bind.TransactOpts, borrowerAddress_ common.Address, maxDepth_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "settle", borrowerAddress_, maxDepth_)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(address borrowerAddress_, uint256 maxDepth_) returns(uint256 collateralSettled_, bool isBorrowerSettled_)
func (_AjnaPool *AjnaPoolSession) Settle(borrowerAddress_ common.Address, maxDepth_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.Settle(&_AjnaPool.TransactOpts, borrowerAddress_, maxDepth_)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(address borrowerAddress_, uint256 maxDepth_) returns(uint256 collateralSettled_, bool isBorrowerSettled_)
func (_AjnaPool *AjnaPoolTransactorSession) Settle(borrowerAddress_ common.Address, maxDepth_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.Settle(&_AjnaPool.TransactOpts, borrowerAddress_, maxDepth_)
}

// StampLoan is a paid mutator transaction binding the contract method 0xce4396d7.
//
// Solidity: function stampLoan() returns()
func (_AjnaPool *AjnaPoolTransactor) StampLoan(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "stampLoan")
}

// StampLoan is a paid mutator transaction binding the contract method 0xce4396d7.
//
// Solidity: function stampLoan() returns()
func (_AjnaPool *AjnaPoolSession) StampLoan() (*types.Transaction, error) {
	return _AjnaPool.Contract.StampLoan(&_AjnaPool.TransactOpts)
}

// StampLoan is a paid mutator transaction binding the contract method 0xce4396d7.
//
// Solidity: function stampLoan() returns()
func (_AjnaPool *AjnaPoolTransactorSession) StampLoan() (*types.Transaction, error) {
	return _AjnaPool.Contract.StampLoan(&_AjnaPool.TransactOpts)
}

// Take is a paid mutator transaction binding the contract method 0x66ae5880.
//
// Solidity: function take(address borrowerAddress_, uint256 maxAmount_, address callee_, bytes data_) returns(uint256 collateralTaken_)
func (_AjnaPool *AjnaPoolTransactor) Take(opts *bind.TransactOpts, borrowerAddress_ common.Address, maxAmount_ *big.Int, callee_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "take", borrowerAddress_, maxAmount_, callee_, data_)
}

// Take is a paid mutator transaction binding the contract method 0x66ae5880.
//
// Solidity: function take(address borrowerAddress_, uint256 maxAmount_, address callee_, bytes data_) returns(uint256 collateralTaken_)
func (_AjnaPool *AjnaPoolSession) Take(borrowerAddress_ common.Address, maxAmount_ *big.Int, callee_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _AjnaPool.Contract.Take(&_AjnaPool.TransactOpts, borrowerAddress_, maxAmount_, callee_, data_)
}

// Take is a paid mutator transaction binding the contract method 0x66ae5880.
//
// Solidity: function take(address borrowerAddress_, uint256 maxAmount_, address callee_, bytes data_) returns(uint256 collateralTaken_)
func (_AjnaPool *AjnaPoolTransactorSession) Take(borrowerAddress_ common.Address, maxAmount_ *big.Int, callee_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _AjnaPool.Contract.Take(&_AjnaPool.TransactOpts, borrowerAddress_, maxAmount_, callee_, data_)
}

// TakeReserves is a paid mutator transaction binding the contract method 0x42302a9a.
//
// Solidity: function takeReserves(uint256 maxAmount_) returns(uint256 amount_)
func (_AjnaPool *AjnaPoolTransactor) TakeReserves(opts *bind.TransactOpts, maxAmount_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "takeReserves", maxAmount_)
}

// TakeReserves is a paid mutator transaction binding the contract method 0x42302a9a.
//
// Solidity: function takeReserves(uint256 maxAmount_) returns(uint256 amount_)
func (_AjnaPool *AjnaPoolSession) TakeReserves(maxAmount_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.TakeReserves(&_AjnaPool.TransactOpts, maxAmount_)
}

// TakeReserves is a paid mutator transaction binding the contract method 0x42302a9a.
//
// Solidity: function takeReserves(uint256 maxAmount_) returns(uint256 amount_)
func (_AjnaPool *AjnaPoolTransactorSession) TakeReserves(maxAmount_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.TakeReserves(&_AjnaPool.TransactOpts, maxAmount_)
}

// TransferLP is a paid mutator transaction binding the contract method 0x4efe8af7.
//
// Solidity: function transferLP(address owner_, address newOwner_, uint256[] indexes_) returns()
func (_AjnaPool *AjnaPoolTransactor) TransferLP(opts *bind.TransactOpts, owner_ common.Address, newOwner_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "transferLP", owner_, newOwner_, indexes_)
}

// TransferLP is a paid mutator transaction binding the contract method 0x4efe8af7.
//
// Solidity: function transferLP(address owner_, address newOwner_, uint256[] indexes_) returns()
func (_AjnaPool *AjnaPoolSession) TransferLP(owner_ common.Address, newOwner_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.TransferLP(&_AjnaPool.TransactOpts, owner_, newOwner_, indexes_)
}

// TransferLP is a paid mutator transaction binding the contract method 0x4efe8af7.
//
// Solidity: function transferLP(address owner_, address newOwner_, uint256[] indexes_) returns()
func (_AjnaPool *AjnaPoolTransactorSession) TransferLP(owner_ common.Address, newOwner_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.TransferLP(&_AjnaPool.TransactOpts, owner_, newOwner_, indexes_)
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns()
func (_AjnaPool *AjnaPoolTransactor) UpdateInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "updateInterest")
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns()
func (_AjnaPool *AjnaPoolSession) UpdateInterest() (*types.Transaction, error) {
	return _AjnaPool.Contract.UpdateInterest(&_AjnaPool.TransactOpts)
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns()
func (_AjnaPool *AjnaPoolTransactorSession) UpdateInterest() (*types.Transaction, error) {
	return _AjnaPool.Contract.UpdateInterest(&_AjnaPool.TransactOpts)
}

// WithdrawBonds is a paid mutator transaction binding the contract method 0xd53e2b1b.
//
// Solidity: function withdrawBonds(address recipient_, uint256 maxAmount_) returns(uint256 withdrawnAmount_)
func (_AjnaPool *AjnaPoolTransactor) WithdrawBonds(opts *bind.TransactOpts, recipient_ common.Address, maxAmount_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.contract.Transact(opts, "withdrawBonds", recipient_, maxAmount_)
}

// WithdrawBonds is a paid mutator transaction binding the contract method 0xd53e2b1b.
//
// Solidity: function withdrawBonds(address recipient_, uint256 maxAmount_) returns(uint256 withdrawnAmount_)
func (_AjnaPool *AjnaPoolSession) WithdrawBonds(recipient_ common.Address, maxAmount_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.WithdrawBonds(&_AjnaPool.TransactOpts, recipient_, maxAmount_)
}

// WithdrawBonds is a paid mutator transaction binding the contract method 0xd53e2b1b.
//
// Solidity: function withdrawBonds(address recipient_, uint256 maxAmount_) returns(uint256 withdrawnAmount_)
func (_AjnaPool *AjnaPoolTransactorSession) WithdrawBonds(recipient_ common.Address, maxAmount_ *big.Int) (*types.Transaction, error) {
	return _AjnaPool.Contract.WithdrawBonds(&_AjnaPool.TransactOpts, recipient_, maxAmount_)
}

// AjnaPoolAddCollateralIterator is returned from FilterAddCollateral and is used to iterate over the raw logs and unpacked data for AddCollateral events raised by the AjnaPool contract.
type AjnaPoolAddCollateralIterator struct {
	Event *AjnaPoolAddCollateral // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolAddCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolAddCollateral)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolAddCollateral)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolAddCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolAddCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolAddCollateral represents a AddCollateral event raised by the AjnaPool contract.
type AjnaPoolAddCollateral struct {
	Actor     common.Address
	Index     *big.Int
	Amount    *big.Int
	LpAwarded *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddCollateral is a free log retrieval operation binding the contract event 0xa9387d09ded47dbc173eb751964c0c7b7e0a1165939b958fafc8109337597f94.
//
// Solidity: event AddCollateral(address indexed actor, uint256 indexed index, uint256 amount, uint256 lpAwarded)
func (_AjnaPool *AjnaPoolFilterer) FilterAddCollateral(opts *bind.FilterOpts, actor []common.Address, index []*big.Int) (*AjnaPoolAddCollateralIterator, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "AddCollateral", actorRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolAddCollateralIterator{contract: _AjnaPool.contract, event: "AddCollateral", logs: logs, sub: sub}, nil
}

// WatchAddCollateral is a free log subscription operation binding the contract event 0xa9387d09ded47dbc173eb751964c0c7b7e0a1165939b958fafc8109337597f94.
//
// Solidity: event AddCollateral(address indexed actor, uint256 indexed index, uint256 amount, uint256 lpAwarded)
func (_AjnaPool *AjnaPoolFilterer) WatchAddCollateral(opts *bind.WatchOpts, sink chan<- *AjnaPoolAddCollateral, actor []common.Address, index []*big.Int) (event.Subscription, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "AddCollateral", actorRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolAddCollateral)
				if err := _AjnaPool.contract.UnpackLog(event, "AddCollateral", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddCollateral is a log parse operation binding the contract event 0xa9387d09ded47dbc173eb751964c0c7b7e0a1165939b958fafc8109337597f94.
//
// Solidity: event AddCollateral(address indexed actor, uint256 indexed index, uint256 amount, uint256 lpAwarded)
func (_AjnaPool *AjnaPoolFilterer) ParseAddCollateral(log types.Log) (*AjnaPoolAddCollateral, error) {
	event := new(AjnaPoolAddCollateral)
	if err := _AjnaPool.contract.UnpackLog(event, "AddCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolAddQuoteTokenIterator is returned from FilterAddQuoteToken and is used to iterate over the raw logs and unpacked data for AddQuoteToken events raised by the AjnaPool contract.
type AjnaPoolAddQuoteTokenIterator struct {
	Event *AjnaPoolAddQuoteToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolAddQuoteTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolAddQuoteToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolAddQuoteToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolAddQuoteTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolAddQuoteTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolAddQuoteToken represents a AddQuoteToken event raised by the AjnaPool contract.
type AjnaPoolAddQuoteToken struct {
	Lender    common.Address
	Index     *big.Int
	Amount    *big.Int
	LpAwarded *big.Int
	Lup       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddQuoteToken is a free log retrieval operation binding the contract event 0x8b24a9808cf05d3d8e48ac09e4f649054994a88cfa657b3f4bf340b62137df1e.
//
// Solidity: event AddQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpAwarded, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) FilterAddQuoteToken(opts *bind.FilterOpts, lender []common.Address, index []*big.Int) (*AjnaPoolAddQuoteTokenIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "AddQuoteToken", lenderRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolAddQuoteTokenIterator{contract: _AjnaPool.contract, event: "AddQuoteToken", logs: logs, sub: sub}, nil
}

// WatchAddQuoteToken is a free log subscription operation binding the contract event 0x8b24a9808cf05d3d8e48ac09e4f649054994a88cfa657b3f4bf340b62137df1e.
//
// Solidity: event AddQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpAwarded, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) WatchAddQuoteToken(opts *bind.WatchOpts, sink chan<- *AjnaPoolAddQuoteToken, lender []common.Address, index []*big.Int) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "AddQuoteToken", lenderRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolAddQuoteToken)
				if err := _AjnaPool.contract.UnpackLog(event, "AddQuoteToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddQuoteToken is a log parse operation binding the contract event 0x8b24a9808cf05d3d8e48ac09e4f649054994a88cfa657b3f4bf340b62137df1e.
//
// Solidity: event AddQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpAwarded, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) ParseAddQuoteToken(log types.Log) (*AjnaPoolAddQuoteToken, error) {
	event := new(AjnaPoolAddQuoteToken)
	if err := _AjnaPool.contract.UnpackLog(event, "AddQuoteToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolApproveLPTransferorsIterator is returned from FilterApproveLPTransferors and is used to iterate over the raw logs and unpacked data for ApproveLPTransferors events raised by the AjnaPool contract.
type AjnaPoolApproveLPTransferorsIterator struct {
	Event *AjnaPoolApproveLPTransferors // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolApproveLPTransferorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolApproveLPTransferors)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolApproveLPTransferors)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolApproveLPTransferorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolApproveLPTransferorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolApproveLPTransferors represents a ApproveLPTransferors event raised by the AjnaPool contract.
type AjnaPoolApproveLPTransferors struct {
	Lender      common.Address
	Transferors []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterApproveLPTransferors is a free log retrieval operation binding the contract event 0x6dceda33b951e30872202c47c5e3b449112437927dbc475dfaedd3a22889aa54.
//
// Solidity: event ApproveLPTransferors(address indexed lender, address[] transferors)
func (_AjnaPool *AjnaPoolFilterer) FilterApproveLPTransferors(opts *bind.FilterOpts, lender []common.Address) (*AjnaPoolApproveLPTransferorsIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "ApproveLPTransferors", lenderRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolApproveLPTransferorsIterator{contract: _AjnaPool.contract, event: "ApproveLPTransferors", logs: logs, sub: sub}, nil
}

// WatchApproveLPTransferors is a free log subscription operation binding the contract event 0x6dceda33b951e30872202c47c5e3b449112437927dbc475dfaedd3a22889aa54.
//
// Solidity: event ApproveLPTransferors(address indexed lender, address[] transferors)
func (_AjnaPool *AjnaPoolFilterer) WatchApproveLPTransferors(opts *bind.WatchOpts, sink chan<- *AjnaPoolApproveLPTransferors, lender []common.Address) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "ApproveLPTransferors", lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolApproveLPTransferors)
				if err := _AjnaPool.contract.UnpackLog(event, "ApproveLPTransferors", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproveLPTransferors is a log parse operation binding the contract event 0x6dceda33b951e30872202c47c5e3b449112437927dbc475dfaedd3a22889aa54.
//
// Solidity: event ApproveLPTransferors(address indexed lender, address[] transferors)
func (_AjnaPool *AjnaPoolFilterer) ParseApproveLPTransferors(log types.Log) (*AjnaPoolApproveLPTransferors, error) {
	event := new(AjnaPoolApproveLPTransferors)
	if err := _AjnaPool.contract.UnpackLog(event, "ApproveLPTransferors", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolAuctionNFTSettleIterator is returned from FilterAuctionNFTSettle and is used to iterate over the raw logs and unpacked data for AuctionNFTSettle events raised by the AjnaPool contract.
type AjnaPoolAuctionNFTSettleIterator struct {
	Event *AjnaPoolAuctionNFTSettle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolAuctionNFTSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolAuctionNFTSettle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolAuctionNFTSettle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolAuctionNFTSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolAuctionNFTSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolAuctionNFTSettle represents a AuctionNFTSettle event raised by the AjnaPool contract.
type AjnaPoolAuctionNFTSettle struct {
	Borrower   common.Address
	Collateral *big.Int
	Lp         *big.Int
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionNFTSettle is a free log retrieval operation binding the contract event 0xddd6b496b84171d179d9874158b1cbbe422dd482e5523f1b09cb69ebef287841.
//
// Solidity: event AuctionNFTSettle(address indexed borrower, uint256 collateral, uint256 lp, uint256 index)
func (_AjnaPool *AjnaPoolFilterer) FilterAuctionNFTSettle(opts *bind.FilterOpts, borrower []common.Address) (*AjnaPoolAuctionNFTSettleIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "AuctionNFTSettle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolAuctionNFTSettleIterator{contract: _AjnaPool.contract, event: "AuctionNFTSettle", logs: logs, sub: sub}, nil
}

// WatchAuctionNFTSettle is a free log subscription operation binding the contract event 0xddd6b496b84171d179d9874158b1cbbe422dd482e5523f1b09cb69ebef287841.
//
// Solidity: event AuctionNFTSettle(address indexed borrower, uint256 collateral, uint256 lp, uint256 index)
func (_AjnaPool *AjnaPoolFilterer) WatchAuctionNFTSettle(opts *bind.WatchOpts, sink chan<- *AjnaPoolAuctionNFTSettle, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "AuctionNFTSettle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolAuctionNFTSettle)
				if err := _AjnaPool.contract.UnpackLog(event, "AuctionNFTSettle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionNFTSettle is a log parse operation binding the contract event 0xddd6b496b84171d179d9874158b1cbbe422dd482e5523f1b09cb69ebef287841.
//
// Solidity: event AuctionNFTSettle(address indexed borrower, uint256 collateral, uint256 lp, uint256 index)
func (_AjnaPool *AjnaPoolFilterer) ParseAuctionNFTSettle(log types.Log) (*AjnaPoolAuctionNFTSettle, error) {
	event := new(AjnaPoolAuctionNFTSettle)
	if err := _AjnaPool.contract.UnpackLog(event, "AuctionNFTSettle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolAuctionSettleIterator is returned from FilterAuctionSettle and is used to iterate over the raw logs and unpacked data for AuctionSettle events raised by the AjnaPool contract.
type AjnaPoolAuctionSettleIterator struct {
	Event *AjnaPoolAuctionSettle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolAuctionSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolAuctionSettle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolAuctionSettle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolAuctionSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolAuctionSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolAuctionSettle represents a AuctionSettle event raised by the AjnaPool contract.
type AjnaPoolAuctionSettle struct {
	Borrower   common.Address
	Collateral *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettle is a free log retrieval operation binding the contract event 0x91a9dcdd01df8b934f14307641e884e0ea6e414bf05fe8daf8c74a28f69b55ee.
//
// Solidity: event AuctionSettle(address indexed borrower, uint256 collateral)
func (_AjnaPool *AjnaPoolFilterer) FilterAuctionSettle(opts *bind.FilterOpts, borrower []common.Address) (*AjnaPoolAuctionSettleIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "AuctionSettle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolAuctionSettleIterator{contract: _AjnaPool.contract, event: "AuctionSettle", logs: logs, sub: sub}, nil
}

// WatchAuctionSettle is a free log subscription operation binding the contract event 0x91a9dcdd01df8b934f14307641e884e0ea6e414bf05fe8daf8c74a28f69b55ee.
//
// Solidity: event AuctionSettle(address indexed borrower, uint256 collateral)
func (_AjnaPool *AjnaPoolFilterer) WatchAuctionSettle(opts *bind.WatchOpts, sink chan<- *AjnaPoolAuctionSettle, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "AuctionSettle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolAuctionSettle)
				if err := _AjnaPool.contract.UnpackLog(event, "AuctionSettle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionSettle is a log parse operation binding the contract event 0x91a9dcdd01df8b934f14307641e884e0ea6e414bf05fe8daf8c74a28f69b55ee.
//
// Solidity: event AuctionSettle(address indexed borrower, uint256 collateral)
func (_AjnaPool *AjnaPoolFilterer) ParseAuctionSettle(log types.Log) (*AjnaPoolAuctionSettle, error) {
	event := new(AjnaPoolAuctionSettle)
	if err := _AjnaPool.contract.UnpackLog(event, "AuctionSettle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolBondWithdrawnIterator is returned from FilterBondWithdrawn and is used to iterate over the raw logs and unpacked data for BondWithdrawn events raised by the AjnaPool contract.
type AjnaPoolBondWithdrawnIterator struct {
	Event *AjnaPoolBondWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolBondWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolBondWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolBondWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolBondWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolBondWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolBondWithdrawn represents a BondWithdrawn event raised by the AjnaPool contract.
type AjnaPoolBondWithdrawn struct {
	Kicker   common.Address
	Reciever common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBondWithdrawn is a free log retrieval operation binding the contract event 0x1b6622b92ce16ed648b5b93fe47df1cd4c763fdcafe3281bc1dfd5ff7998a94d.
//
// Solidity: event BondWithdrawn(address indexed kicker, address indexed reciever, uint256 amount)
func (_AjnaPool *AjnaPoolFilterer) FilterBondWithdrawn(opts *bind.FilterOpts, kicker []common.Address, reciever []common.Address) (*AjnaPoolBondWithdrawnIterator, error) {

	var kickerRule []interface{}
	for _, kickerItem := range kicker {
		kickerRule = append(kickerRule, kickerItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "BondWithdrawn", kickerRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolBondWithdrawnIterator{contract: _AjnaPool.contract, event: "BondWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBondWithdrawn is a free log subscription operation binding the contract event 0x1b6622b92ce16ed648b5b93fe47df1cd4c763fdcafe3281bc1dfd5ff7998a94d.
//
// Solidity: event BondWithdrawn(address indexed kicker, address indexed reciever, uint256 amount)
func (_AjnaPool *AjnaPoolFilterer) WatchBondWithdrawn(opts *bind.WatchOpts, sink chan<- *AjnaPoolBondWithdrawn, kicker []common.Address, reciever []common.Address) (event.Subscription, error) {

	var kickerRule []interface{}
	for _, kickerItem := range kicker {
		kickerRule = append(kickerRule, kickerItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "BondWithdrawn", kickerRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolBondWithdrawn)
				if err := _AjnaPool.contract.UnpackLog(event, "BondWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBondWithdrawn is a log parse operation binding the contract event 0x1b6622b92ce16ed648b5b93fe47df1cd4c763fdcafe3281bc1dfd5ff7998a94d.
//
// Solidity: event BondWithdrawn(address indexed kicker, address indexed reciever, uint256 amount)
func (_AjnaPool *AjnaPoolFilterer) ParseBondWithdrawn(log types.Log) (*AjnaPoolBondWithdrawn, error) {
	event := new(AjnaPoolBondWithdrawn)
	if err := _AjnaPool.contract.UnpackLog(event, "BondWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolBucketBankruptcyIterator is returned from FilterBucketBankruptcy and is used to iterate over the raw logs and unpacked data for BucketBankruptcy events raised by the AjnaPool contract.
type AjnaPoolBucketBankruptcyIterator struct {
	Event *AjnaPoolBucketBankruptcy // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolBucketBankruptcyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolBucketBankruptcy)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolBucketBankruptcy)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolBucketBankruptcyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolBucketBankruptcyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolBucketBankruptcy represents a BucketBankruptcy event raised by the AjnaPool contract.
type AjnaPoolBucketBankruptcy struct {
	Index       *big.Int
	LpForfeited *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBucketBankruptcy is a free log retrieval operation binding the contract event 0x30ee43613aaa48d222b158aab1123c5a29d452f8b1a849e5f815dd355923ba85.
//
// Solidity: event BucketBankruptcy(uint256 indexed index, uint256 lpForfeited)
func (_AjnaPool *AjnaPoolFilterer) FilterBucketBankruptcy(opts *bind.FilterOpts, index []*big.Int) (*AjnaPoolBucketBankruptcyIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "BucketBankruptcy", indexRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolBucketBankruptcyIterator{contract: _AjnaPool.contract, event: "BucketBankruptcy", logs: logs, sub: sub}, nil
}

// WatchBucketBankruptcy is a free log subscription operation binding the contract event 0x30ee43613aaa48d222b158aab1123c5a29d452f8b1a849e5f815dd355923ba85.
//
// Solidity: event BucketBankruptcy(uint256 indexed index, uint256 lpForfeited)
func (_AjnaPool *AjnaPoolFilterer) WatchBucketBankruptcy(opts *bind.WatchOpts, sink chan<- *AjnaPoolBucketBankruptcy, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "BucketBankruptcy", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolBucketBankruptcy)
				if err := _AjnaPool.contract.UnpackLog(event, "BucketBankruptcy", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBucketBankruptcy is a log parse operation binding the contract event 0x30ee43613aaa48d222b158aab1123c5a29d452f8b1a849e5f815dd355923ba85.
//
// Solidity: event BucketBankruptcy(uint256 indexed index, uint256 lpForfeited)
func (_AjnaPool *AjnaPoolFilterer) ParseBucketBankruptcy(log types.Log) (*AjnaPoolBucketBankruptcy, error) {
	event := new(AjnaPoolBucketBankruptcy)
	if err := _AjnaPool.contract.UnpackLog(event, "BucketBankruptcy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolBucketTakeIterator is returned from FilterBucketTake and is used to iterate over the raw logs and unpacked data for BucketTake events raised by the AjnaPool contract.
type AjnaPoolBucketTakeIterator struct {
	Event *AjnaPoolBucketTake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolBucketTakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolBucketTake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolBucketTake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolBucketTakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolBucketTakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolBucketTake represents a BucketTake event raised by the AjnaPool contract.
type AjnaPoolBucketTake struct {
	Borrower   common.Address
	Index      *big.Int
	Amount     *big.Int
	Collateral *big.Int
	BondChange *big.Int
	IsReward   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBucketTake is a free log retrieval operation binding the contract event 0xcb6905a59200cc485bfe6d2168392e96a0f204daf9e3937dff19180cb0d796a4.
//
// Solidity: event BucketTake(address indexed borrower, uint256 index, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_AjnaPool *AjnaPoolFilterer) FilterBucketTake(opts *bind.FilterOpts, borrower []common.Address) (*AjnaPoolBucketTakeIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "BucketTake", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolBucketTakeIterator{contract: _AjnaPool.contract, event: "BucketTake", logs: logs, sub: sub}, nil
}

// WatchBucketTake is a free log subscription operation binding the contract event 0xcb6905a59200cc485bfe6d2168392e96a0f204daf9e3937dff19180cb0d796a4.
//
// Solidity: event BucketTake(address indexed borrower, uint256 index, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_AjnaPool *AjnaPoolFilterer) WatchBucketTake(opts *bind.WatchOpts, sink chan<- *AjnaPoolBucketTake, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "BucketTake", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolBucketTake)
				if err := _AjnaPool.contract.UnpackLog(event, "BucketTake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBucketTake is a log parse operation binding the contract event 0xcb6905a59200cc485bfe6d2168392e96a0f204daf9e3937dff19180cb0d796a4.
//
// Solidity: event BucketTake(address indexed borrower, uint256 index, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_AjnaPool *AjnaPoolFilterer) ParseBucketTake(log types.Log) (*AjnaPoolBucketTake, error) {
	event := new(AjnaPoolBucketTake)
	if err := _AjnaPool.contract.UnpackLog(event, "BucketTake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolBucketTakeLPAwardedIterator is returned from FilterBucketTakeLPAwarded and is used to iterate over the raw logs and unpacked data for BucketTakeLPAwarded events raised by the AjnaPool contract.
type AjnaPoolBucketTakeLPAwardedIterator struct {
	Event *AjnaPoolBucketTakeLPAwarded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolBucketTakeLPAwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolBucketTakeLPAwarded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolBucketTakeLPAwarded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolBucketTakeLPAwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolBucketTakeLPAwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolBucketTakeLPAwarded represents a BucketTakeLPAwarded event raised by the AjnaPool contract.
type AjnaPoolBucketTakeLPAwarded struct {
	Taker           common.Address
	Kicker          common.Address
	LpAwardedTaker  *big.Int
	LpAwardedKicker *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBucketTakeLPAwarded is a free log retrieval operation binding the contract event 0xd43d7a2b808648e0814e6795ea3a9b723df6eae4046a7e279a036458f38bc644.
//
// Solidity: event BucketTakeLPAwarded(address indexed taker, address indexed kicker, uint256 lpAwardedTaker, uint256 lpAwardedKicker)
func (_AjnaPool *AjnaPoolFilterer) FilterBucketTakeLPAwarded(opts *bind.FilterOpts, taker []common.Address, kicker []common.Address) (*AjnaPoolBucketTakeLPAwardedIterator, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var kickerRule []interface{}
	for _, kickerItem := range kicker {
		kickerRule = append(kickerRule, kickerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "BucketTakeLPAwarded", takerRule, kickerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolBucketTakeLPAwardedIterator{contract: _AjnaPool.contract, event: "BucketTakeLPAwarded", logs: logs, sub: sub}, nil
}

// WatchBucketTakeLPAwarded is a free log subscription operation binding the contract event 0xd43d7a2b808648e0814e6795ea3a9b723df6eae4046a7e279a036458f38bc644.
//
// Solidity: event BucketTakeLPAwarded(address indexed taker, address indexed kicker, uint256 lpAwardedTaker, uint256 lpAwardedKicker)
func (_AjnaPool *AjnaPoolFilterer) WatchBucketTakeLPAwarded(opts *bind.WatchOpts, sink chan<- *AjnaPoolBucketTakeLPAwarded, taker []common.Address, kicker []common.Address) (event.Subscription, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var kickerRule []interface{}
	for _, kickerItem := range kicker {
		kickerRule = append(kickerRule, kickerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "BucketTakeLPAwarded", takerRule, kickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolBucketTakeLPAwarded)
				if err := _AjnaPool.contract.UnpackLog(event, "BucketTakeLPAwarded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBucketTakeLPAwarded is a log parse operation binding the contract event 0xd43d7a2b808648e0814e6795ea3a9b723df6eae4046a7e279a036458f38bc644.
//
// Solidity: event BucketTakeLPAwarded(address indexed taker, address indexed kicker, uint256 lpAwardedTaker, uint256 lpAwardedKicker)
func (_AjnaPool *AjnaPoolFilterer) ParseBucketTakeLPAwarded(log types.Log) (*AjnaPoolBucketTakeLPAwarded, error) {
	event := new(AjnaPoolBucketTakeLPAwarded)
	if err := _AjnaPool.contract.UnpackLog(event, "BucketTakeLPAwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolDecreaseLPAllowanceIterator is returned from FilterDecreaseLPAllowance and is used to iterate over the raw logs and unpacked data for DecreaseLPAllowance events raised by the AjnaPool contract.
type AjnaPoolDecreaseLPAllowanceIterator struct {
	Event *AjnaPoolDecreaseLPAllowance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolDecreaseLPAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolDecreaseLPAllowance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolDecreaseLPAllowance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolDecreaseLPAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolDecreaseLPAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolDecreaseLPAllowance represents a DecreaseLPAllowance event raised by the AjnaPool contract.
type AjnaPoolDecreaseLPAllowance struct {
	Owner   common.Address
	Spender common.Address
	Indexes []*big.Int
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDecreaseLPAllowance is a free log retrieval operation binding the contract event 0x4a7a52e2fe7e10addaa7f875ecf9ec17563ec12be96c885457061cfc04e05660.
//
// Solidity: event DecreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_AjnaPool *AjnaPoolFilterer) FilterDecreaseLPAllowance(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AjnaPoolDecreaseLPAllowanceIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "DecreaseLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolDecreaseLPAllowanceIterator{contract: _AjnaPool.contract, event: "DecreaseLPAllowance", logs: logs, sub: sub}, nil
}

// WatchDecreaseLPAllowance is a free log subscription operation binding the contract event 0x4a7a52e2fe7e10addaa7f875ecf9ec17563ec12be96c885457061cfc04e05660.
//
// Solidity: event DecreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_AjnaPool *AjnaPoolFilterer) WatchDecreaseLPAllowance(opts *bind.WatchOpts, sink chan<- *AjnaPoolDecreaseLPAllowance, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "DecreaseLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolDecreaseLPAllowance)
				if err := _AjnaPool.contract.UnpackLog(event, "DecreaseLPAllowance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreaseLPAllowance is a log parse operation binding the contract event 0x4a7a52e2fe7e10addaa7f875ecf9ec17563ec12be96c885457061cfc04e05660.
//
// Solidity: event DecreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_AjnaPool *AjnaPoolFilterer) ParseDecreaseLPAllowance(log types.Log) (*AjnaPoolDecreaseLPAllowance, error) {
	event := new(AjnaPoolDecreaseLPAllowance)
	if err := _AjnaPool.contract.UnpackLog(event, "DecreaseLPAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolDrawDebtIterator is returned from FilterDrawDebt and is used to iterate over the raw logs and unpacked data for DrawDebt events raised by the AjnaPool contract.
type AjnaPoolDrawDebtIterator struct {
	Event *AjnaPoolDrawDebt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolDrawDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolDrawDebt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolDrawDebt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolDrawDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolDrawDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolDrawDebt represents a DrawDebt event raised by the AjnaPool contract.
type AjnaPoolDrawDebt struct {
	Borrower          common.Address
	AmountBorrowed    *big.Int
	CollateralPledged *big.Int
	Lup               *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDrawDebt is a free log retrieval operation binding the contract event 0x49a2aab2f4f7ca5c6ba6d413b46a0a09d91d10188fd94b8e23c3225362d12b50.
//
// Solidity: event DrawDebt(address indexed borrower, uint256 amountBorrowed, uint256 collateralPledged, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) FilterDrawDebt(opts *bind.FilterOpts, borrower []common.Address) (*AjnaPoolDrawDebtIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "DrawDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolDrawDebtIterator{contract: _AjnaPool.contract, event: "DrawDebt", logs: logs, sub: sub}, nil
}

// WatchDrawDebt is a free log subscription operation binding the contract event 0x49a2aab2f4f7ca5c6ba6d413b46a0a09d91d10188fd94b8e23c3225362d12b50.
//
// Solidity: event DrawDebt(address indexed borrower, uint256 amountBorrowed, uint256 collateralPledged, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) WatchDrawDebt(opts *bind.WatchOpts, sink chan<- *AjnaPoolDrawDebt, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "DrawDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolDrawDebt)
				if err := _AjnaPool.contract.UnpackLog(event, "DrawDebt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDrawDebt is a log parse operation binding the contract event 0x49a2aab2f4f7ca5c6ba6d413b46a0a09d91d10188fd94b8e23c3225362d12b50.
//
// Solidity: event DrawDebt(address indexed borrower, uint256 amountBorrowed, uint256 collateralPledged, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) ParseDrawDebt(log types.Log) (*AjnaPoolDrawDebt, error) {
	event := new(AjnaPoolDrawDebt)
	if err := _AjnaPool.contract.UnpackLog(event, "DrawDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolFlashloanIterator is returned from FilterFlashloan and is used to iterate over the raw logs and unpacked data for Flashloan events raised by the AjnaPool contract.
type AjnaPoolFlashloanIterator struct {
	Event *AjnaPoolFlashloan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolFlashloanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolFlashloan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolFlashloan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolFlashloanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolFlashloanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolFlashloan represents a Flashloan event raised by the AjnaPool contract.
type AjnaPoolFlashloan struct {
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFlashloan is a free log retrieval operation binding the contract event 0x6b15284fe89dbd5c436c2e0b06b1bf72e3a0a8e96d1b4a2abd61dfae2d7849a6.
//
// Solidity: event Flashloan(address indexed receiver, address indexed token, uint256 amount)
func (_AjnaPool *AjnaPoolFilterer) FilterFlashloan(opts *bind.FilterOpts, receiver []common.Address, token []common.Address) (*AjnaPoolFlashloanIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "Flashloan", receiverRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolFlashloanIterator{contract: _AjnaPool.contract, event: "Flashloan", logs: logs, sub: sub}, nil
}

// WatchFlashloan is a free log subscription operation binding the contract event 0x6b15284fe89dbd5c436c2e0b06b1bf72e3a0a8e96d1b4a2abd61dfae2d7849a6.
//
// Solidity: event Flashloan(address indexed receiver, address indexed token, uint256 amount)
func (_AjnaPool *AjnaPoolFilterer) WatchFlashloan(opts *bind.WatchOpts, sink chan<- *AjnaPoolFlashloan, receiver []common.Address, token []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "Flashloan", receiverRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolFlashloan)
				if err := _AjnaPool.contract.UnpackLog(event, "Flashloan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashloan is a log parse operation binding the contract event 0x6b15284fe89dbd5c436c2e0b06b1bf72e3a0a8e96d1b4a2abd61dfae2d7849a6.
//
// Solidity: event Flashloan(address indexed receiver, address indexed token, uint256 amount)
func (_AjnaPool *AjnaPoolFilterer) ParseFlashloan(log types.Log) (*AjnaPoolFlashloan, error) {
	event := new(AjnaPoolFlashloan)
	if err := _AjnaPool.contract.UnpackLog(event, "Flashloan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolIncreaseLPAllowanceIterator is returned from FilterIncreaseLPAllowance and is used to iterate over the raw logs and unpacked data for IncreaseLPAllowance events raised by the AjnaPool contract.
type AjnaPoolIncreaseLPAllowanceIterator struct {
	Event *AjnaPoolIncreaseLPAllowance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolIncreaseLPAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolIncreaseLPAllowance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolIncreaseLPAllowance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolIncreaseLPAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolIncreaseLPAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolIncreaseLPAllowance represents a IncreaseLPAllowance event raised by the AjnaPool contract.
type AjnaPoolIncreaseLPAllowance struct {
	Owner   common.Address
	Spender common.Address
	Indexes []*big.Int
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIncreaseLPAllowance is a free log retrieval operation binding the contract event 0xa9f6ab83637f87ef702b94c10d09430c40cd3d4642d14fc2a07408bde931545f.
//
// Solidity: event IncreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_AjnaPool *AjnaPoolFilterer) FilterIncreaseLPAllowance(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AjnaPoolIncreaseLPAllowanceIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "IncreaseLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolIncreaseLPAllowanceIterator{contract: _AjnaPool.contract, event: "IncreaseLPAllowance", logs: logs, sub: sub}, nil
}

// WatchIncreaseLPAllowance is a free log subscription operation binding the contract event 0xa9f6ab83637f87ef702b94c10d09430c40cd3d4642d14fc2a07408bde931545f.
//
// Solidity: event IncreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_AjnaPool *AjnaPoolFilterer) WatchIncreaseLPAllowance(opts *bind.WatchOpts, sink chan<- *AjnaPoolIncreaseLPAllowance, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "IncreaseLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolIncreaseLPAllowance)
				if err := _AjnaPool.contract.UnpackLog(event, "IncreaseLPAllowance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseLPAllowance is a log parse operation binding the contract event 0xa9f6ab83637f87ef702b94c10d09430c40cd3d4642d14fc2a07408bde931545f.
//
// Solidity: event IncreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_AjnaPool *AjnaPoolFilterer) ParseIncreaseLPAllowance(log types.Log) (*AjnaPoolIncreaseLPAllowance, error) {
	event := new(AjnaPoolIncreaseLPAllowance)
	if err := _AjnaPool.contract.UnpackLog(event, "IncreaseLPAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolInterestUpdateFailureIterator is returned from FilterInterestUpdateFailure and is used to iterate over the raw logs and unpacked data for InterestUpdateFailure events raised by the AjnaPool contract.
type AjnaPoolInterestUpdateFailureIterator struct {
	Event *AjnaPoolInterestUpdateFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolInterestUpdateFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolInterestUpdateFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolInterestUpdateFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolInterestUpdateFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolInterestUpdateFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolInterestUpdateFailure represents a InterestUpdateFailure event raised by the AjnaPool contract.
type AjnaPoolInterestUpdateFailure struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInterestUpdateFailure is a free log retrieval operation binding the contract event 0x84da056cd0ff5380ec35a74f131057a96626a24305fa137c235bdbe1b414a396.
//
// Solidity: event InterestUpdateFailure()
func (_AjnaPool *AjnaPoolFilterer) FilterInterestUpdateFailure(opts *bind.FilterOpts) (*AjnaPoolInterestUpdateFailureIterator, error) {

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "InterestUpdateFailure")
	if err != nil {
		return nil, err
	}
	return &AjnaPoolInterestUpdateFailureIterator{contract: _AjnaPool.contract, event: "InterestUpdateFailure", logs: logs, sub: sub}, nil
}

// WatchInterestUpdateFailure is a free log subscription operation binding the contract event 0x84da056cd0ff5380ec35a74f131057a96626a24305fa137c235bdbe1b414a396.
//
// Solidity: event InterestUpdateFailure()
func (_AjnaPool *AjnaPoolFilterer) WatchInterestUpdateFailure(opts *bind.WatchOpts, sink chan<- *AjnaPoolInterestUpdateFailure) (event.Subscription, error) {

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "InterestUpdateFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolInterestUpdateFailure)
				if err := _AjnaPool.contract.UnpackLog(event, "InterestUpdateFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterestUpdateFailure is a log parse operation binding the contract event 0x84da056cd0ff5380ec35a74f131057a96626a24305fa137c235bdbe1b414a396.
//
// Solidity: event InterestUpdateFailure()
func (_AjnaPool *AjnaPoolFilterer) ParseInterestUpdateFailure(log types.Log) (*AjnaPoolInterestUpdateFailure, error) {
	event := new(AjnaPoolInterestUpdateFailure)
	if err := _AjnaPool.contract.UnpackLog(event, "InterestUpdateFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolKickIterator is returned from FilterKick and is used to iterate over the raw logs and unpacked data for Kick events raised by the AjnaPool contract.
type AjnaPoolKickIterator struct {
	Event *AjnaPoolKick // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolKickIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolKick)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolKick)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolKickIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolKickIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolKick represents a Kick event raised by the AjnaPool contract.
type AjnaPoolKick struct {
	Borrower   common.Address
	Debt       *big.Int
	Collateral *big.Int
	Bond       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKick is a free log retrieval operation binding the contract event 0x9f9a32e7f0271518f9b1895d0b1f2f4f73ed305e48b0a3782932094f9d00d948.
//
// Solidity: event Kick(address indexed borrower, uint256 debt, uint256 collateral, uint256 bond)
func (_AjnaPool *AjnaPoolFilterer) FilterKick(opts *bind.FilterOpts, borrower []common.Address) (*AjnaPoolKickIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "Kick", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolKickIterator{contract: _AjnaPool.contract, event: "Kick", logs: logs, sub: sub}, nil
}

// WatchKick is a free log subscription operation binding the contract event 0x9f9a32e7f0271518f9b1895d0b1f2f4f73ed305e48b0a3782932094f9d00d948.
//
// Solidity: event Kick(address indexed borrower, uint256 debt, uint256 collateral, uint256 bond)
func (_AjnaPool *AjnaPoolFilterer) WatchKick(opts *bind.WatchOpts, sink chan<- *AjnaPoolKick, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "Kick", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolKick)
				if err := _AjnaPool.contract.UnpackLog(event, "Kick", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseKick is a log parse operation binding the contract event 0x9f9a32e7f0271518f9b1895d0b1f2f4f73ed305e48b0a3782932094f9d00d948.
//
// Solidity: event Kick(address indexed borrower, uint256 debt, uint256 collateral, uint256 bond)
func (_AjnaPool *AjnaPoolFilterer) ParseKick(log types.Log) (*AjnaPoolKick, error) {
	event := new(AjnaPoolKick)
	if err := _AjnaPool.contract.UnpackLog(event, "Kick", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolKickReserveAuctionIterator is returned from FilterKickReserveAuction and is used to iterate over the raw logs and unpacked data for KickReserveAuction events raised by the AjnaPool contract.
type AjnaPoolKickReserveAuctionIterator struct {
	Event *AjnaPoolKickReserveAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolKickReserveAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolKickReserveAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolKickReserveAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolKickReserveAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolKickReserveAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolKickReserveAuction represents a KickReserveAuction event raised by the AjnaPool contract.
type AjnaPoolKickReserveAuction struct {
	ClaimableReservesRemaining *big.Int
	AuctionPrice               *big.Int
	CurrentBurnEpoch           *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterKickReserveAuction is a free log retrieval operation binding the contract event 0x3dacf6b19b0a84358a76a3338466cd428d1d4e80e53ccfe91b15d9b8021df960.
//
// Solidity: event KickReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_AjnaPool *AjnaPoolFilterer) FilterKickReserveAuction(opts *bind.FilterOpts) (*AjnaPoolKickReserveAuctionIterator, error) {

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "KickReserveAuction")
	if err != nil {
		return nil, err
	}
	return &AjnaPoolKickReserveAuctionIterator{contract: _AjnaPool.contract, event: "KickReserveAuction", logs: logs, sub: sub}, nil
}

// WatchKickReserveAuction is a free log subscription operation binding the contract event 0x3dacf6b19b0a84358a76a3338466cd428d1d4e80e53ccfe91b15d9b8021df960.
//
// Solidity: event KickReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_AjnaPool *AjnaPoolFilterer) WatchKickReserveAuction(opts *bind.WatchOpts, sink chan<- *AjnaPoolKickReserveAuction) (event.Subscription, error) {

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "KickReserveAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolKickReserveAuction)
				if err := _AjnaPool.contract.UnpackLog(event, "KickReserveAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseKickReserveAuction is a log parse operation binding the contract event 0x3dacf6b19b0a84358a76a3338466cd428d1d4e80e53ccfe91b15d9b8021df960.
//
// Solidity: event KickReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_AjnaPool *AjnaPoolFilterer) ParseKickReserveAuction(log types.Log) (*AjnaPoolKickReserveAuction, error) {
	event := new(AjnaPoolKickReserveAuction)
	if err := _AjnaPool.contract.UnpackLog(event, "KickReserveAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolLoanStampedIterator is returned from FilterLoanStamped and is used to iterate over the raw logs and unpacked data for LoanStamped events raised by the AjnaPool contract.
type AjnaPoolLoanStampedIterator struct {
	Event *AjnaPoolLoanStamped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolLoanStampedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolLoanStamped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolLoanStamped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolLoanStampedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolLoanStampedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolLoanStamped represents a LoanStamped event raised by the AjnaPool contract.
type AjnaPoolLoanStamped struct {
	Borrower common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLoanStamped is a free log retrieval operation binding the contract event 0x8d6660b4a409414ebe386e9dd200a5c4e75591f0fc98e1272d7ba207d06d4c34.
//
// Solidity: event LoanStamped(address indexed borrower)
func (_AjnaPool *AjnaPoolFilterer) FilterLoanStamped(opts *bind.FilterOpts, borrower []common.Address) (*AjnaPoolLoanStampedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "LoanStamped", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolLoanStampedIterator{contract: _AjnaPool.contract, event: "LoanStamped", logs: logs, sub: sub}, nil
}

// WatchLoanStamped is a free log subscription operation binding the contract event 0x8d6660b4a409414ebe386e9dd200a5c4e75591f0fc98e1272d7ba207d06d4c34.
//
// Solidity: event LoanStamped(address indexed borrower)
func (_AjnaPool *AjnaPoolFilterer) WatchLoanStamped(opts *bind.WatchOpts, sink chan<- *AjnaPoolLoanStamped, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "LoanStamped", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolLoanStamped)
				if err := _AjnaPool.contract.UnpackLog(event, "LoanStamped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLoanStamped is a log parse operation binding the contract event 0x8d6660b4a409414ebe386e9dd200a5c4e75591f0fc98e1272d7ba207d06d4c34.
//
// Solidity: event LoanStamped(address indexed borrower)
func (_AjnaPool *AjnaPoolFilterer) ParseLoanStamped(log types.Log) (*AjnaPoolLoanStamped, error) {
	event := new(AjnaPoolLoanStamped)
	if err := _AjnaPool.contract.UnpackLog(event, "LoanStamped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolMoveQuoteTokenIterator is returned from FilterMoveQuoteToken and is used to iterate over the raw logs and unpacked data for MoveQuoteToken events raised by the AjnaPool contract.
type AjnaPoolMoveQuoteTokenIterator struct {
	Event *AjnaPoolMoveQuoteToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolMoveQuoteTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolMoveQuoteToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolMoveQuoteToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolMoveQuoteTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolMoveQuoteTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolMoveQuoteToken represents a MoveQuoteToken event raised by the AjnaPool contract.
type AjnaPoolMoveQuoteToken struct {
	Lender         common.Address
	From           *big.Int
	To             *big.Int
	Amount         *big.Int
	LpRedeemedFrom *big.Int
	LpAwardedTo    *big.Int
	Lup            *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMoveQuoteToken is a free log retrieval operation binding the contract event 0x9d7ab6bb30c003ae7d5b583911db0ada7a9e51b0b4ac7ac1bb5e6896e82e4dbe.
//
// Solidity: event MoveQuoteToken(address indexed lender, uint256 indexed from, uint256 indexed to, uint256 amount, uint256 lpRedeemedFrom, uint256 lpAwardedTo, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) FilterMoveQuoteToken(opts *bind.FilterOpts, lender []common.Address, from []*big.Int, to []*big.Int) (*AjnaPoolMoveQuoteTokenIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "MoveQuoteToken", lenderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolMoveQuoteTokenIterator{contract: _AjnaPool.contract, event: "MoveQuoteToken", logs: logs, sub: sub}, nil
}

// WatchMoveQuoteToken is a free log subscription operation binding the contract event 0x9d7ab6bb30c003ae7d5b583911db0ada7a9e51b0b4ac7ac1bb5e6896e82e4dbe.
//
// Solidity: event MoveQuoteToken(address indexed lender, uint256 indexed from, uint256 indexed to, uint256 amount, uint256 lpRedeemedFrom, uint256 lpAwardedTo, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) WatchMoveQuoteToken(opts *bind.WatchOpts, sink chan<- *AjnaPoolMoveQuoteToken, lender []common.Address, from []*big.Int, to []*big.Int) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "MoveQuoteToken", lenderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolMoveQuoteToken)
				if err := _AjnaPool.contract.UnpackLog(event, "MoveQuoteToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMoveQuoteToken is a log parse operation binding the contract event 0x9d7ab6bb30c003ae7d5b583911db0ada7a9e51b0b4ac7ac1bb5e6896e82e4dbe.
//
// Solidity: event MoveQuoteToken(address indexed lender, uint256 indexed from, uint256 indexed to, uint256 amount, uint256 lpRedeemedFrom, uint256 lpAwardedTo, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) ParseMoveQuoteToken(log types.Log) (*AjnaPoolMoveQuoteToken, error) {
	event := new(AjnaPoolMoveQuoteToken)
	if err := _AjnaPool.contract.UnpackLog(event, "MoveQuoteToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolRemoveCollateralIterator is returned from FilterRemoveCollateral and is used to iterate over the raw logs and unpacked data for RemoveCollateral events raised by the AjnaPool contract.
type AjnaPoolRemoveCollateralIterator struct {
	Event *AjnaPoolRemoveCollateral // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolRemoveCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolRemoveCollateral)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolRemoveCollateral)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolRemoveCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolRemoveCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolRemoveCollateral represents a RemoveCollateral event raised by the AjnaPool contract.
type AjnaPoolRemoveCollateral struct {
	Claimer    common.Address
	Index      *big.Int
	Amount     *big.Int
	LpRedeemed *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveCollateral is a free log retrieval operation binding the contract event 0x90895bc82397742e0cea4685e72279103862a03bee6bbe1d71265c7aeb111527.
//
// Solidity: event RemoveCollateral(address indexed claimer, uint256 indexed index, uint256 amount, uint256 lpRedeemed)
func (_AjnaPool *AjnaPoolFilterer) FilterRemoveCollateral(opts *bind.FilterOpts, claimer []common.Address, index []*big.Int) (*AjnaPoolRemoveCollateralIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "RemoveCollateral", claimerRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolRemoveCollateralIterator{contract: _AjnaPool.contract, event: "RemoveCollateral", logs: logs, sub: sub}, nil
}

// WatchRemoveCollateral is a free log subscription operation binding the contract event 0x90895bc82397742e0cea4685e72279103862a03bee6bbe1d71265c7aeb111527.
//
// Solidity: event RemoveCollateral(address indexed claimer, uint256 indexed index, uint256 amount, uint256 lpRedeemed)
func (_AjnaPool *AjnaPoolFilterer) WatchRemoveCollateral(opts *bind.WatchOpts, sink chan<- *AjnaPoolRemoveCollateral, claimer []common.Address, index []*big.Int) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "RemoveCollateral", claimerRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolRemoveCollateral)
				if err := _AjnaPool.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveCollateral is a log parse operation binding the contract event 0x90895bc82397742e0cea4685e72279103862a03bee6bbe1d71265c7aeb111527.
//
// Solidity: event RemoveCollateral(address indexed claimer, uint256 indexed index, uint256 amount, uint256 lpRedeemed)
func (_AjnaPool *AjnaPoolFilterer) ParseRemoveCollateral(log types.Log) (*AjnaPoolRemoveCollateral, error) {
	event := new(AjnaPoolRemoveCollateral)
	if err := _AjnaPool.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolRemoveQuoteTokenIterator is returned from FilterRemoveQuoteToken and is used to iterate over the raw logs and unpacked data for RemoveQuoteToken events raised by the AjnaPool contract.
type AjnaPoolRemoveQuoteTokenIterator struct {
	Event *AjnaPoolRemoveQuoteToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolRemoveQuoteTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolRemoveQuoteToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolRemoveQuoteToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolRemoveQuoteTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolRemoveQuoteTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolRemoveQuoteToken represents a RemoveQuoteToken event raised by the AjnaPool contract.
type AjnaPoolRemoveQuoteToken struct {
	Lender     common.Address
	Index      *big.Int
	Amount     *big.Int
	LpRedeemed *big.Int
	Lup        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveQuoteToken is a free log retrieval operation binding the contract event 0x0130a7b525bd6b1e72def1ee0b77f3467028a0e958e30174a0c95eb3860fc221.
//
// Solidity: event RemoveQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpRedeemed, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) FilterRemoveQuoteToken(opts *bind.FilterOpts, lender []common.Address, index []*big.Int) (*AjnaPoolRemoveQuoteTokenIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "RemoveQuoteToken", lenderRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolRemoveQuoteTokenIterator{contract: _AjnaPool.contract, event: "RemoveQuoteToken", logs: logs, sub: sub}, nil
}

// WatchRemoveQuoteToken is a free log subscription operation binding the contract event 0x0130a7b525bd6b1e72def1ee0b77f3467028a0e958e30174a0c95eb3860fc221.
//
// Solidity: event RemoveQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpRedeemed, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) WatchRemoveQuoteToken(opts *bind.WatchOpts, sink chan<- *AjnaPoolRemoveQuoteToken, lender []common.Address, index []*big.Int) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "RemoveQuoteToken", lenderRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolRemoveQuoteToken)
				if err := _AjnaPool.contract.UnpackLog(event, "RemoveQuoteToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveQuoteToken is a log parse operation binding the contract event 0x0130a7b525bd6b1e72def1ee0b77f3467028a0e958e30174a0c95eb3860fc221.
//
// Solidity: event RemoveQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpRedeemed, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) ParseRemoveQuoteToken(log types.Log) (*AjnaPoolRemoveQuoteToken, error) {
	event := new(AjnaPoolRemoveQuoteToken)
	if err := _AjnaPool.contract.UnpackLog(event, "RemoveQuoteToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolRepayDebtIterator is returned from FilterRepayDebt and is used to iterate over the raw logs and unpacked data for RepayDebt events raised by the AjnaPool contract.
type AjnaPoolRepayDebtIterator struct {
	Event *AjnaPoolRepayDebt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolRepayDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolRepayDebt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolRepayDebt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolRepayDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolRepayDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolRepayDebt represents a RepayDebt event raised by the AjnaPool contract.
type AjnaPoolRepayDebt struct {
	Borrower         common.Address
	QuoteRepaid      *big.Int
	CollateralPulled *big.Int
	Lup              *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRepayDebt is a free log retrieval operation binding the contract event 0xef9d6dc34b1e6893b8746b03ac07fd084909654a5cedab240265a8d1bd584dc2.
//
// Solidity: event RepayDebt(address indexed borrower, uint256 quoteRepaid, uint256 collateralPulled, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) FilterRepayDebt(opts *bind.FilterOpts, borrower []common.Address) (*AjnaPoolRepayDebtIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "RepayDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolRepayDebtIterator{contract: _AjnaPool.contract, event: "RepayDebt", logs: logs, sub: sub}, nil
}

// WatchRepayDebt is a free log subscription operation binding the contract event 0xef9d6dc34b1e6893b8746b03ac07fd084909654a5cedab240265a8d1bd584dc2.
//
// Solidity: event RepayDebt(address indexed borrower, uint256 quoteRepaid, uint256 collateralPulled, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) WatchRepayDebt(opts *bind.WatchOpts, sink chan<- *AjnaPoolRepayDebt, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "RepayDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolRepayDebt)
				if err := _AjnaPool.contract.UnpackLog(event, "RepayDebt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepayDebt is a log parse operation binding the contract event 0xef9d6dc34b1e6893b8746b03ac07fd084909654a5cedab240265a8d1bd584dc2.
//
// Solidity: event RepayDebt(address indexed borrower, uint256 quoteRepaid, uint256 collateralPulled, uint256 lup)
func (_AjnaPool *AjnaPoolFilterer) ParseRepayDebt(log types.Log) (*AjnaPoolRepayDebt, error) {
	event := new(AjnaPoolRepayDebt)
	if err := _AjnaPool.contract.UnpackLog(event, "RepayDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolReserveAuctionIterator is returned from FilterReserveAuction and is used to iterate over the raw logs and unpacked data for ReserveAuction events raised by the AjnaPool contract.
type AjnaPoolReserveAuctionIterator struct {
	Event *AjnaPoolReserveAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolReserveAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolReserveAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolReserveAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolReserveAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolReserveAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolReserveAuction represents a ReserveAuction event raised by the AjnaPool contract.
type AjnaPoolReserveAuction struct {
	ClaimableReservesRemaining *big.Int
	AuctionPrice               *big.Int
	CurrentBurnEpoch           *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterReserveAuction is a free log retrieval operation binding the contract event 0x2115712930a0e5047c8904a9cc557d2b1ba5a21876e04f59249843ce1a74d699.
//
// Solidity: event ReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_AjnaPool *AjnaPoolFilterer) FilterReserveAuction(opts *bind.FilterOpts) (*AjnaPoolReserveAuctionIterator, error) {

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "ReserveAuction")
	if err != nil {
		return nil, err
	}
	return &AjnaPoolReserveAuctionIterator{contract: _AjnaPool.contract, event: "ReserveAuction", logs: logs, sub: sub}, nil
}

// WatchReserveAuction is a free log subscription operation binding the contract event 0x2115712930a0e5047c8904a9cc557d2b1ba5a21876e04f59249843ce1a74d699.
//
// Solidity: event ReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_AjnaPool *AjnaPoolFilterer) WatchReserveAuction(opts *bind.WatchOpts, sink chan<- *AjnaPoolReserveAuction) (event.Subscription, error) {

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "ReserveAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolReserveAuction)
				if err := _AjnaPool.contract.UnpackLog(event, "ReserveAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveAuction is a log parse operation binding the contract event 0x2115712930a0e5047c8904a9cc557d2b1ba5a21876e04f59249843ce1a74d699.
//
// Solidity: event ReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_AjnaPool *AjnaPoolFilterer) ParseReserveAuction(log types.Log) (*AjnaPoolReserveAuction, error) {
	event := new(AjnaPoolReserveAuction)
	if err := _AjnaPool.contract.UnpackLog(event, "ReserveAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolResetInterestRateIterator is returned from FilterResetInterestRate and is used to iterate over the raw logs and unpacked data for ResetInterestRate events raised by the AjnaPool contract.
type AjnaPoolResetInterestRateIterator struct {
	Event *AjnaPoolResetInterestRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolResetInterestRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolResetInterestRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolResetInterestRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolResetInterestRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolResetInterestRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolResetInterestRate represents a ResetInterestRate event raised by the AjnaPool contract.
type AjnaPoolResetInterestRate struct {
	OldRate *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterResetInterestRate is a free log retrieval operation binding the contract event 0x20ae1d4a2e8d297f3820670c20fc79531e31643d4b201892680e7df3c4ab1599.
//
// Solidity: event ResetInterestRate(uint256 oldRate, uint256 newRate)
func (_AjnaPool *AjnaPoolFilterer) FilterResetInterestRate(opts *bind.FilterOpts) (*AjnaPoolResetInterestRateIterator, error) {

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "ResetInterestRate")
	if err != nil {
		return nil, err
	}
	return &AjnaPoolResetInterestRateIterator{contract: _AjnaPool.contract, event: "ResetInterestRate", logs: logs, sub: sub}, nil
}

// WatchResetInterestRate is a free log subscription operation binding the contract event 0x20ae1d4a2e8d297f3820670c20fc79531e31643d4b201892680e7df3c4ab1599.
//
// Solidity: event ResetInterestRate(uint256 oldRate, uint256 newRate)
func (_AjnaPool *AjnaPoolFilterer) WatchResetInterestRate(opts *bind.WatchOpts, sink chan<- *AjnaPoolResetInterestRate) (event.Subscription, error) {

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "ResetInterestRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolResetInterestRate)
				if err := _AjnaPool.contract.UnpackLog(event, "ResetInterestRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseResetInterestRate is a log parse operation binding the contract event 0x20ae1d4a2e8d297f3820670c20fc79531e31643d4b201892680e7df3c4ab1599.
//
// Solidity: event ResetInterestRate(uint256 oldRate, uint256 newRate)
func (_AjnaPool *AjnaPoolFilterer) ParseResetInterestRate(log types.Log) (*AjnaPoolResetInterestRate, error) {
	event := new(AjnaPoolResetInterestRate)
	if err := _AjnaPool.contract.UnpackLog(event, "ResetInterestRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolRevokeLPAllowanceIterator is returned from FilterRevokeLPAllowance and is used to iterate over the raw logs and unpacked data for RevokeLPAllowance events raised by the AjnaPool contract.
type AjnaPoolRevokeLPAllowanceIterator struct {
	Event *AjnaPoolRevokeLPAllowance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolRevokeLPAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolRevokeLPAllowance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolRevokeLPAllowance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolRevokeLPAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolRevokeLPAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolRevokeLPAllowance represents a RevokeLPAllowance event raised by the AjnaPool contract.
type AjnaPoolRevokeLPAllowance struct {
	Owner   common.Address
	Spender common.Address
	Indexes []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevokeLPAllowance is a free log retrieval operation binding the contract event 0xdf7f78d931b4e040d4598563bab17506dba0aed1a0515c51fd7dbc2a2382afdf.
//
// Solidity: event RevokeLPAllowance(address indexed owner, address indexed spender, uint256[] indexes)
func (_AjnaPool *AjnaPoolFilterer) FilterRevokeLPAllowance(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AjnaPoolRevokeLPAllowanceIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "RevokeLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolRevokeLPAllowanceIterator{contract: _AjnaPool.contract, event: "RevokeLPAllowance", logs: logs, sub: sub}, nil
}

// WatchRevokeLPAllowance is a free log subscription operation binding the contract event 0xdf7f78d931b4e040d4598563bab17506dba0aed1a0515c51fd7dbc2a2382afdf.
//
// Solidity: event RevokeLPAllowance(address indexed owner, address indexed spender, uint256[] indexes)
func (_AjnaPool *AjnaPoolFilterer) WatchRevokeLPAllowance(opts *bind.WatchOpts, sink chan<- *AjnaPoolRevokeLPAllowance, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "RevokeLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolRevokeLPAllowance)
				if err := _AjnaPool.contract.UnpackLog(event, "RevokeLPAllowance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokeLPAllowance is a log parse operation binding the contract event 0xdf7f78d931b4e040d4598563bab17506dba0aed1a0515c51fd7dbc2a2382afdf.
//
// Solidity: event RevokeLPAllowance(address indexed owner, address indexed spender, uint256[] indexes)
func (_AjnaPool *AjnaPoolFilterer) ParseRevokeLPAllowance(log types.Log) (*AjnaPoolRevokeLPAllowance, error) {
	event := new(AjnaPoolRevokeLPAllowance)
	if err := _AjnaPool.contract.UnpackLog(event, "RevokeLPAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolRevokeLPTransferorsIterator is returned from FilterRevokeLPTransferors and is used to iterate over the raw logs and unpacked data for RevokeLPTransferors events raised by the AjnaPool contract.
type AjnaPoolRevokeLPTransferorsIterator struct {
	Event *AjnaPoolRevokeLPTransferors // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolRevokeLPTransferorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolRevokeLPTransferors)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolRevokeLPTransferors)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolRevokeLPTransferorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolRevokeLPTransferorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolRevokeLPTransferors represents a RevokeLPTransferors event raised by the AjnaPool contract.
type AjnaPoolRevokeLPTransferors struct {
	Lender      common.Address
	Transferors []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokeLPTransferors is a free log retrieval operation binding the contract event 0xde63bcd4c57b5d710a16396f80a2797846720923ff52198e806544ccbb720b2b.
//
// Solidity: event RevokeLPTransferors(address indexed lender, address[] transferors)
func (_AjnaPool *AjnaPoolFilterer) FilterRevokeLPTransferors(opts *bind.FilterOpts, lender []common.Address) (*AjnaPoolRevokeLPTransferorsIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "RevokeLPTransferors", lenderRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolRevokeLPTransferorsIterator{contract: _AjnaPool.contract, event: "RevokeLPTransferors", logs: logs, sub: sub}, nil
}

// WatchRevokeLPTransferors is a free log subscription operation binding the contract event 0xde63bcd4c57b5d710a16396f80a2797846720923ff52198e806544ccbb720b2b.
//
// Solidity: event RevokeLPTransferors(address indexed lender, address[] transferors)
func (_AjnaPool *AjnaPoolFilterer) WatchRevokeLPTransferors(opts *bind.WatchOpts, sink chan<- *AjnaPoolRevokeLPTransferors, lender []common.Address) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "RevokeLPTransferors", lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolRevokeLPTransferors)
				if err := _AjnaPool.contract.UnpackLog(event, "RevokeLPTransferors", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokeLPTransferors is a log parse operation binding the contract event 0xde63bcd4c57b5d710a16396f80a2797846720923ff52198e806544ccbb720b2b.
//
// Solidity: event RevokeLPTransferors(address indexed lender, address[] transferors)
func (_AjnaPool *AjnaPoolFilterer) ParseRevokeLPTransferors(log types.Log) (*AjnaPoolRevokeLPTransferors, error) {
	event := new(AjnaPoolRevokeLPTransferors)
	if err := _AjnaPool.contract.UnpackLog(event, "RevokeLPTransferors", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the AjnaPool contract.
type AjnaPoolSettleIterator struct {
	Event *AjnaPoolSettle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolSettle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolSettle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolSettle represents a Settle event raised by the AjnaPool contract.
type AjnaPoolSettle struct {
	Borrower    common.Address
	SettledDebt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed borrower, uint256 settledDebt)
func (_AjnaPool *AjnaPoolFilterer) FilterSettle(opts *bind.FilterOpts, borrower []common.Address) (*AjnaPoolSettleIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "Settle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolSettleIterator{contract: _AjnaPool.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed borrower, uint256 settledDebt)
func (_AjnaPool *AjnaPoolFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *AjnaPoolSettle, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "Settle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolSettle)
				if err := _AjnaPool.contract.UnpackLog(event, "Settle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettle is a log parse operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed borrower, uint256 settledDebt)
func (_AjnaPool *AjnaPoolFilterer) ParseSettle(log types.Log) (*AjnaPoolSettle, error) {
	event := new(AjnaPoolSettle)
	if err := _AjnaPool.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolTakeIterator is returned from FilterTake and is used to iterate over the raw logs and unpacked data for Take events raised by the AjnaPool contract.
type AjnaPoolTakeIterator struct {
	Event *AjnaPoolTake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolTakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolTake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolTake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolTakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolTakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolTake represents a Take event raised by the AjnaPool contract.
type AjnaPoolTake struct {
	Borrower   common.Address
	Amount     *big.Int
	Collateral *big.Int
	BondChange *big.Int
	IsReward   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTake is a free log retrieval operation binding the contract event 0x4591b2dfbebff121b3ccd19ae2407e59a9cefd959b35e12d82752cb5bc6eb757.
//
// Solidity: event Take(address indexed borrower, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_AjnaPool *AjnaPoolFilterer) FilterTake(opts *bind.FilterOpts, borrower []common.Address) (*AjnaPoolTakeIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "Take", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolTakeIterator{contract: _AjnaPool.contract, event: "Take", logs: logs, sub: sub}, nil
}

// WatchTake is a free log subscription operation binding the contract event 0x4591b2dfbebff121b3ccd19ae2407e59a9cefd959b35e12d82752cb5bc6eb757.
//
// Solidity: event Take(address indexed borrower, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_AjnaPool *AjnaPoolFilterer) WatchTake(opts *bind.WatchOpts, sink chan<- *AjnaPoolTake, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "Take", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolTake)
				if err := _AjnaPool.contract.UnpackLog(event, "Take", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTake is a log parse operation binding the contract event 0x4591b2dfbebff121b3ccd19ae2407e59a9cefd959b35e12d82752cb5bc6eb757.
//
// Solidity: event Take(address indexed borrower, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_AjnaPool *AjnaPoolFilterer) ParseTake(log types.Log) (*AjnaPoolTake, error) {
	event := new(AjnaPoolTake)
	if err := _AjnaPool.contract.UnpackLog(event, "Take", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolTransferLPIterator is returned from FilterTransferLP and is used to iterate over the raw logs and unpacked data for TransferLP events raised by the AjnaPool contract.
type AjnaPoolTransferLPIterator struct {
	Event *AjnaPoolTransferLP // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolTransferLPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolTransferLP)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolTransferLP)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolTransferLPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolTransferLPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolTransferLP represents a TransferLP event raised by the AjnaPool contract.
type AjnaPoolTransferLP struct {
	Owner    common.Address
	NewOwner common.Address
	Indexes  []*big.Int
	Lp       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferLP is a free log retrieval operation binding the contract event 0x98b14a8359f3da2f702dde7a51271b67ea43d27b8712e860408b8bf8dd0eb682.
//
// Solidity: event TransferLP(address owner, address newOwner, uint256[] indexes, uint256 lp)
func (_AjnaPool *AjnaPoolFilterer) FilterTransferLP(opts *bind.FilterOpts) (*AjnaPoolTransferLPIterator, error) {

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "TransferLP")
	if err != nil {
		return nil, err
	}
	return &AjnaPoolTransferLPIterator{contract: _AjnaPool.contract, event: "TransferLP", logs: logs, sub: sub}, nil
}

// WatchTransferLP is a free log subscription operation binding the contract event 0x98b14a8359f3da2f702dde7a51271b67ea43d27b8712e860408b8bf8dd0eb682.
//
// Solidity: event TransferLP(address owner, address newOwner, uint256[] indexes, uint256 lp)
func (_AjnaPool *AjnaPoolFilterer) WatchTransferLP(opts *bind.WatchOpts, sink chan<- *AjnaPoolTransferLP) (event.Subscription, error) {

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "TransferLP")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolTransferLP)
				if err := _AjnaPool.contract.UnpackLog(event, "TransferLP", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferLP is a log parse operation binding the contract event 0x98b14a8359f3da2f702dde7a51271b67ea43d27b8712e860408b8bf8dd0eb682.
//
// Solidity: event TransferLP(address owner, address newOwner, uint256[] indexes, uint256 lp)
func (_AjnaPool *AjnaPoolFilterer) ParseTransferLP(log types.Log) (*AjnaPoolTransferLP, error) {
	event := new(AjnaPoolTransferLP)
	if err := _AjnaPool.contract.UnpackLog(event, "TransferLP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AjnaPoolUpdateInterestRateIterator is returned from FilterUpdateInterestRate and is used to iterate over the raw logs and unpacked data for UpdateInterestRate events raised by the AjnaPool contract.
type AjnaPoolUpdateInterestRateIterator struct {
	Event *AjnaPoolUpdateInterestRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AjnaPoolUpdateInterestRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolUpdateInterestRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AjnaPoolUpdateInterestRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AjnaPoolUpdateInterestRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolUpdateInterestRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolUpdateInterestRate represents a UpdateInterestRate event raised by the AjnaPool contract.
type AjnaPoolUpdateInterestRate struct {
	OldRate *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateInterestRate is a free log retrieval operation binding the contract event 0x2463616ef8e6f9bddf00e4964b853ad9050f87cd3c73985d2ee6b6c8a8336991.
//
// Solidity: event UpdateInterestRate(uint256 oldRate, uint256 newRate)
func (_AjnaPool *AjnaPoolFilterer) FilterUpdateInterestRate(opts *bind.FilterOpts) (*AjnaPoolUpdateInterestRateIterator, error) {

	logs, sub, err := _AjnaPool.contract.FilterLogs(opts, "UpdateInterestRate")
	if err != nil {
		return nil, err
	}
	return &AjnaPoolUpdateInterestRateIterator{contract: _AjnaPool.contract, event: "UpdateInterestRate", logs: logs, sub: sub}, nil
}

// WatchUpdateInterestRate is a free log subscription operation binding the contract event 0x2463616ef8e6f9bddf00e4964b853ad9050f87cd3c73985d2ee6b6c8a8336991.
//
// Solidity: event UpdateInterestRate(uint256 oldRate, uint256 newRate)
func (_AjnaPool *AjnaPoolFilterer) WatchUpdateInterestRate(opts *bind.WatchOpts, sink chan<- *AjnaPoolUpdateInterestRate) (event.Subscription, error) {

	logs, sub, err := _AjnaPool.contract.WatchLogs(opts, "UpdateInterestRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolUpdateInterestRate)
				if err := _AjnaPool.contract.UnpackLog(event, "UpdateInterestRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateInterestRate is a log parse operation binding the contract event 0x2463616ef8e6f9bddf00e4964b853ad9050f87cd3c73985d2ee6b6c8a8336991.
//
// Solidity: event UpdateInterestRate(uint256 oldRate, uint256 newRate)
func (_AjnaPool *AjnaPoolFilterer) ParseUpdateInterestRate(log types.Log) (*AjnaPoolUpdateInterestRate, error) {
	event := new(AjnaPoolUpdateInterestRate)
	if err := _AjnaPool.contract.UnpackLog(event, "UpdateInterestRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
