package types

import (
	"bytes"
	"errors"
	math "math"
	"math/big"
	"net/url"
	"reflect"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// AccAddress a wrapper around bytes meant to represent an account address.
// When marshaled to a string or JSON, it uses Bech32.
type AccAddress []byte

//nolint:staticcheck
func AssertInt32(val int32) error {
	if val >= 0 && val <= math.MaxInt32 {
		return nil
	}
	return ErrInt32OverFlow
}

//nolint:staticcheck
func AssertInt64(val uint64) error {
	if val >= 0 && val <= math.MaxInt64 {
		return nil
	}
	return ErrInt64OverFlow
}

func AssertProofType(val string) error {
	if val == "specimen" || val == "result" {
		return nil
	}
	return ErrInvalidProofType
}

func AssertHashLength(val string) error {
	countBytes := bytes.Count([]byte(val), []byte(""))
	if countBytes >= 32 && countBytes <= 67 {
		return nil
	}
	return ErrInvalidBlockHash
}

func AssertUrl(val string) error {
	_, err := url.ParseRequestURI(val)
	if err != nil {
		return ErrStringNotUrl
	}
	return nil
}

func AssertSessionMembers(val string) error {
	operators := strings.Split(val, ",")
	for _, address := range operators {
		_, err := CovenetAccAddressFromBech32(address)
		if err != nil {
			return sdkerrors.ErrInvalidAddress
		}
	}
	return nil
}

func AssertTimeSecond(val string) error {
	if len(val) > 0 {
		if val[len(val)-1] != 's' {
			return ErrInvalidSecondTimeSuffix
		}

		if len(val) > 1 && (val[len(val)-2] == 'n' || val[len(val)-2] == 'p' || val[len(val)-2] == 'u' || val[len(val)-2] == 'm') {
			return ErrInvalidSecondTimeSuffix
		}

		_, err := time.ParseDuration(val)
		if err != nil {
			return err
		}
	} else {
		return ErrInvalidSecondTimeSuffix
	}

	return nil
}

func AssertBigInt(val string) error {
	if len(val) == 0 {
		return ErrBigIntString
	}

	// Handle the sign if present
	if val[0] == '-' {
		return ErrBigIntString
	}

	_, success := new(big.Int).SetString(val, 10)
	if !success {
		return ErrBigIntString
	}
	return nil
}

func AssertBool(val interface{}) error {
	if val == nil {
		return ErrInvalidBoolean
	}

	valType := reflect.TypeOf(val)
	if valType.Kind() == reflect.Ptr {
		valType = valType.Elem() // Get the underlying type if val is a pointer
	}

	if valType.Kind() != reflect.Bool {
		return ErrInvalidBoolean
	}

	return nil
}

func AssertStructFieldsNotEmpty(obj interface{}) error {
	val := reflect.ValueOf(obj)

	// Check if obj is a pointer and dereference it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// Check if the field is a zero value for its type
		if field.Interface() == reflect.Zero(field.Type()).Interface() {
			return ErrStructFieldIsEmpty
		}
	}
	return nil
}

func JoinErrors(errs ...error) error {
	var errStrings []string
	for _, err := range errs {
		if err != nil {
			errStrings = append(errStrings, err.Error())
		}
	}
	if len(errStrings) == 0 {
		return nil
	}
	return errors.New(strings.Join(errStrings, "\n"))
}

// CovenetAccAddressFromBech32 creates an AccAddress with "cxtmos" prefix from a Bech32 string.
func CovenetAccAddressFromBech32(address string) (addr sdk.AccAddress, err error) {

	if address == "" {
		return nil, errors.New("empty address string is not allowed")
	}

	if len(strings.TrimSpace(address)) == 0 {
		return sdk.AccAddress{}, errors.New("empty address string is not allowed")
	}

	bech32PrefixAccAddr := "cxtmos"

	bz, err := sdk.GetFromBech32(address, bech32PrefixAccAddr)
	if err != nil {
		return nil, err
	}

	err = sdk.VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return sdk.AccAddress(bz), nil
}
