package creditcardcheck

import (
	//"strconv"
	"testing"
)

func TestCheckCreditCard_1(t *testing.T) {
	cNum := "4111111111111111"
	expected := &CreditCardStatus{
		CardType:   "VISA",
		CardNumber: "4111111111111111",
		CardStatus: "valid",
		Error:      nil,
	}
	r := CheckCreditCard(cNum)
	if r.CardType != expected.CardType || r.CardStatus != expected.CardStatus || r.CardNumber != expected.CardNumber || r.Error != expected.Error {
		t.Error(
			"For", cNum,
			"expected", expected,
			"got", r,
		)
	}
}

func TestCheckCreditCard_2(t *testing.T) {
	cNum := "4111111111111"
	expected := &CreditCardStatus{
		CardType:   "VISA",
		CardNumber: "4111111111111",
		CardStatus: "invalid",
		Error:      nil,
	}
	r := CheckCreditCard(cNum)
	if r.CardType != expected.CardType || r.CardStatus != expected.CardStatus || r.CardNumber != expected.CardNumber || r.Error != expected.Error {
		t.Error(
			"For", cNum,
			"expected", expected,
			"got", r,
		)
	}
}

func TestCheckCreditCard_3(t *testing.T) {
	cNum := "4012888888881881"
	expected := &CreditCardStatus{
		CardType:   "VISA",
		CardNumber: "4012888888881881",
		CardStatus: "valid",
		Error:      nil,
	}
	r := CheckCreditCard(cNum)
	if r.CardType != expected.CardType || r.CardStatus != expected.CardStatus || r.CardNumber != expected.CardNumber || r.Error != expected.Error {
		t.Error(
			"For", cNum,
			"expected", expected,
			"got", r,
		)
	}
}

func TestCheckCreditCard_4(t *testing.T) {
	cNum := "378282246310005"
	expected := &CreditCardStatus{
		CardType:   "AMEX",
		CardNumber: "378282246310005",
		CardStatus: "valid",
		Error:      nil,
	}
	r := CheckCreditCard(cNum)
	if r.CardType != expected.CardType || r.CardStatus != expected.CardStatus || r.CardNumber != expected.CardNumber || r.Error != expected.Error {
		t.Error(
			"For", cNum,
			"expected", expected,
			"got", r,
		)
	}
}

func TestCheckCreditCard_5(t *testing.T) {
	cNum := "6011111111111117"
	expected := &CreditCardStatus{
		CardType:   "Discover",
		CardNumber: "6011111111111117",
		CardStatus: "valid",
		Error:      nil,
	}
	r := CheckCreditCard(cNum)
	if r.CardType != expected.CardType || r.CardStatus != expected.CardStatus || r.CardNumber != expected.CardNumber || r.Error != expected.Error {
		t.Error(
			"For", cNum,
			"expected", expected,
			"got", r,
		)
	}
}

func TestCheckCreditCard_6(t *testing.T) {
	cNum := "5105105105105100"
	expected := &CreditCardStatus{
		CardType:   "MasterCard",
		CardNumber: "5105105105105100",
		CardStatus: "valid",
		Error:      nil,
	}
	r := CheckCreditCard(cNum)
	if r.CardType != expected.CardType || r.CardStatus != expected.CardStatus || r.CardNumber != expected.CardNumber || r.Error != expected.Error {
		t.Error(
			"For", cNum,
			"expected", expected,
			"got", r,
		)
	}
}

func TestCheckCreditCard_7(t *testing.T) {
	cNum := "5105 1051 0510 5106"
	expected := &CreditCardStatus{
		CardType:   "MasterCard",
		CardNumber: "5105105105105106",
		CardStatus: "invalid",
		Error:      nil,
	}
	r := CheckCreditCard(cNum)
	if r.CardType != expected.CardType || r.CardStatus != expected.CardStatus || r.CardNumber != expected.CardNumber || r.Error != expected.Error {
		t.Error(
			"For", cNum,
			"expected", expected,
			"got", r,
		)
	}
}

func TestCheckCreditCard_8(t *testing.T) {
	cNum := "9111111111111111"
	expected := &CreditCardStatus{
		CardType:   "Unknown",
		CardNumber: "9111111111111111",
		CardStatus: "invalid",
		Error:      nil,
	}
	r := CheckCreditCard(cNum)
	if r.CardType != expected.CardType || r.CardStatus != expected.CardStatus || r.CardNumber != expected.CardNumber || r.Error != expected.Error {
		t.Error(
			"For", cNum,
			"expected", expected,
			"got", r,
		)
	}
}
