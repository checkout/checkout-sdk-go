package issuing

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Serialization tests for the control responses returned by create, get, list and update
// control. Each case unmarshals the swagger example of one control_type subtype, checks every
// property, then marshals and unmarshals again to prove the roundtrip keeps every value.

const velocityControlJson = `{
  "id": "ctr_gp7vkmxayztufjz6top5bjcdra",
  "description": "Maximum spend of 500€ per week for restaurants",
  "control_type": "velocity_limit",
  "target_id": "crd_fa6psq42dcdd6fdn5gifcq1491",
  "is_editable": true,
  "created_date": "2021-09-09T19:41:39Z",
  "last_modified_date": "2021-09-09T19:41:39Z",
  "velocity_limit": {
    "amount_remaining": 45000,
    "amount_limit": 50000,
    "velocity_window": { "type": "weekly" },
    "mcc_list": ["4121", "4582"],
    "mid_list": ["1234567890"]
  }
}`

const mccControlJson = `{
  "id": "ctr_gp7vkmxayztufjz6top5bjcdra",
  "description": "Allow the card to be used only in restaurants and supermarkets",
  "control_type": "mcc_limit",
  "target_id": "crd_fa6psq42dcdd6fdn5gifcq1491",
  "is_editable": true,
  "created_date": "2021-09-09T19:41:39Z",
  "last_modified_date": "2021-09-09T19:41:39Z",
  "mcc_limit": { "type": "allow", "mcc_list": ["5932", "5411"] }
}`

const midControlJson = `{
  "id": "ctr_gp7vkmxayztufjz6top5bjcdra",
  "description": "Allow the card to be used only in AZ Pizza",
  "control_type": "mid_limit",
  "target_id": "crd_fa6psq42dcdd6fdn5gifcq1491",
  "is_editable": false,
  "created_date": "2021-09-09T19:41:39Z",
  "last_modified_date": "2021-09-09T19:41:39Z",
  "mid_limit": { "type": "block", "mid_list": ["593278", "541114"] }
}`

// marshalControl rebuilds the wire shape (base fields plus the limit under its spec key),
// because CardControlResponse only has a custom unmarshaller.
func marshalControl(t *testing.T, r CardControlResponse) []byte {
	m := map[string]interface{}{}
	base, err := json.Marshal(r.CardControlData)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(base, &m))
	m[string(r.Limit.GetType())] = r.Limit
	out, err := json.Marshal(m)
	require.NoError(t, err)
	return out
}

func assertBase(t *testing.T, r CardControlResponse, controlType ControlType, description string, editable bool) {
	expected := time.Date(2021, 9, 9, 19, 41, 39, 0, time.UTC)
	assert.Equal(t, "ctr_gp7vkmxayztufjz6top5bjcdra", r.Id)
	assert.Equal(t, description, r.Description)
	assert.Equal(t, controlType, r.ControlType)
	assert.Equal(t, "crd_fa6psq42dcdd6fdn5gifcq1491", r.TargetId)
	assert.Equal(t, editable, r.IsEditable)
	require.NotNil(t, r.CreatedDate)
	require.NotNil(t, r.LastModifiedDate)
	assert.True(t, expected.Equal(*r.CreatedDate))
	assert.True(t, expected.Equal(*r.LastModifiedDate))
	assert.Equal(t, controlType, r.Limit.GetType())
}

func TestCardControlResponse_Subtypes_SwaggerExampleAndRoundtrip(t *testing.T) {
	cases := []struct {
		name    string
		payload string
		check   func(t *testing.T, r CardControlResponse)
	}{
		{
			name:    "velocity_limit",
			payload: velocityControlJson,
			check: func(t *testing.T, r CardControlResponse) {
				assertBase(t, r, VelocityLimitType, "Maximum spend of 500€ per week for restaurants", true)
				limit, ok := r.Limit.(VelocityLimit)
				require.True(t, ok)
				require.NotNil(t, limit.AmountRemaining)
				assert.Equal(t, int64(45000), *limit.AmountRemaining)
				assert.Equal(t, int64(50000), limit.AmountLimit)
				assert.Equal(t, Weekly, limit.VelocityWindow.Type)
				assert.Equal(t, []string{"4121", "4582"}, limit.MccList)
				assert.Equal(t, []string{"1234567890"}, limit.MidList)
			},
		},
		{
			name:    "mcc_limit",
			payload: mccControlJson,
			check: func(t *testing.T, r CardControlResponse) {
				assertBase(t, r, MccLimitType, "Allow the card to be used only in restaurants and supermarkets", true)
				limit, ok := r.Limit.(MccLimit)
				require.True(t, ok)
				assert.Equal(t, Allow, limit.Type)
				assert.Equal(t, []string{"5932", "5411"}, limit.MccList)
			},
		},
		{
			name:    "mid_limit",
			payload: midControlJson,
			check: func(t *testing.T, r CardControlResponse) {
				assertBase(t, r, MidLimitType, "Allow the card to be used only in AZ Pizza", false)
				limit, ok := r.Limit.(MidLimit)
				require.True(t, ok)
				assert.Equal(t, Block, limit.Type)
				assert.Equal(t, []string{"593278", "541114"}, limit.MidList)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var first CardControlResponse
			require.NoError(t, json.Unmarshal([]byte(tc.payload), &first))
			tc.check(t, first)

			var second CardControlResponse
			require.NoError(t, json.Unmarshal(marshalControl(t, first), &second))
			tc.check(t, second)
		})
	}
}

func TestCardControlsQueryResponse_ListOfSubtypes(t *testing.T) {
	payload := `{"controls":[` + velocityControlJson + `,` + mccControlJson + `,` + midControlJson + `]}`
	var r CardControlsQueryResponse
	require.NoError(t, json.Unmarshal([]byte(payload), &r))
	require.Len(t, r.Controls, 3)
	assert.Equal(t, int64(45000), *r.Controls[0].Limit.(VelocityLimit).AmountRemaining)
	assert.Equal(t, MccLimitType, r.Controls[1].Limit.GetType())
	assert.Equal(t, MidLimitType, r.Controls[2].Limit.GetType())
}

func TestCardControlResponse_UnknownControlTypeFails(t *testing.T) {
	var r CardControlResponse
	assert.Error(t, json.Unmarshal([]byte(`{"control_type":"other_limit"}`), &r))
}

func TestVelocityLimit_AmountRemainingIsNeverSentOnRequests(t *testing.T) {
	request := NewVelocityCardControlRequest()
	request.VelocityLimit = VelocityLimit{AmountLimit: 500, VelocityWindow: VelocityWindow{Type: Weekly}}
	body, err := json.Marshal(request)
	require.NoError(t, err)
	assert.NotContains(t, string(body), "amount_remaining")

	update := UpdateCardControlRequest{VelocityLimit: &VelocityLimit{AmountLimit: 500}}
	body, err = json.Marshal(update)
	require.NoError(t, err)
	assert.NotContains(t, string(body), "amount_remaining")
}
