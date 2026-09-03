package apm

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// PaymentRequestAchSource is the only schema declaring savings / checking / cash.
// requestAchSource previously used common.AccountType, which declares "current" instead of
// "checking", so a valid account type could not be sent and an invalid one was offered.
func TestAchSourceAccountTypeValues(t *testing.T) {
	assert.Equal(t, "savings", string(AchSourceSavings))
	assert.Equal(t, "checking", string(AchSourceChecking))
	assert.Equal(t, "cash", string(AchSourceCash))
}

func TestAchSourceAccountTypeDiffersFromSharedAccountType(t *testing.T) {
	// The shared enum offers Current, which this position rejects, and cannot express checking.
	// If these are ever unified, this fails.
	assert.Equal(t, "current", string(common.Current))
	assert.NotEqual(t, string(common.Current), string(AchSourceChecking))
	assert.NotEqual(t, string(common.Savings), string(AchSourceChecking))
}

func TestAchSourceSerializesCheckingAccountType(t *testing.T) {
	source := NewRequestAchSource()
	source.AccountType = AchSourceChecking
	source.Country = common.US
	source.AccountNumber = "136549956"
	source.BankCode = "021000021"

	body, err := json.Marshal(source)
	assert.Nil(t, err)

	var got map[string]interface{}
	assert.Nil(t, json.Unmarshal(body, &got))

	assert.Equal(t, "ach", got["type"])
	assert.Equal(t, "checking", got["account_type"])
	assert.Equal(t, "136549956", got["account_number"])
	assert.Equal(t, "021000021", got["bank_code"])
}

func TestAchSourceSerializesEveryDeclaredAccountType(t *testing.T) {
	for _, accountType := range []AchSourceAccountType{AchSourceSavings, AchSourceChecking, AchSourceCash} {
		source := NewRequestAchSource()
		source.AccountType = accountType

		body, err := json.Marshal(source)
		assert.Nil(t, err)

		var got map[string]interface{}
		assert.Nil(t, json.Unmarshal(body, &got))
		assert.Equal(t, string(accountType), got["account_type"])
	}
}

func TestAchSourceAccountHolderIsNarrowerThanTheSharedType(t *testing.T) {
	source := NewRequestAchSource()
	source.AccountType = AchSourceChecking
	source.AccountHolder = &AchSourceAccountHolder{
		Type:      common.Individual,
		FirstName: "John",
		LastName:  "Smith",
	}

	body, err := json.Marshal(source)
	assert.Nil(t, err)

	var got map[string]interface{}
	assert.Nil(t, json.Unmarshal(body, &got))

	holder := got["account_holder"].(map[string]interface{})
	// AccountHolderAch declares seven properties. The shared common.AccountHolder carries 18,
	// including a phone and a tax ID this position does not accept.
	assert.Equal(t, "individual", holder["type"])
	assert.Equal(t, "John", holder["first_name"])
	assert.Len(t, holder, 3)
}
