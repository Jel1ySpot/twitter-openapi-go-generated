/*
Twitter OpenAPI

Twitter OpenAPI(Swagger) specification

API version: 0.0.1
Contact: yuki@yuki0311.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserTipJarSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTipJarSettings{}

// UserTipJarSettings struct for UserTipJarSettings
type UserTipJarSettings struct {
	BandcampHandle *string `json:"bandcamp_handle,omitempty"`
	BitcoinHandle *string `json:"bitcoin_handle,omitempty"`
	CashAppHandle *string `json:"cash_app_handle,omitempty"`
	EthereumHandle *string `json:"ethereum_handle,omitempty"`
	GofundmeHandle *string `json:"gofundme_handle,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	PatreonHandle *string `json:"patreon_handle,omitempty"`
	VenmoHandle *string `json:"venmo_handle,omitempty"`
}

// NewUserTipJarSettings instantiates a new UserTipJarSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTipJarSettings() *UserTipJarSettings {
	this := UserTipJarSettings{}
	return &this
}

// NewUserTipJarSettingsWithDefaults instantiates a new UserTipJarSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTipJarSettingsWithDefaults() *UserTipJarSettings {
	this := UserTipJarSettings{}
	return &this
}

// GetBandcampHandle returns the BandcampHandle field value if set, zero value otherwise.
func (o *UserTipJarSettings) GetBandcampHandle() string {
	if o == nil || IsNil(o.BandcampHandle) {
		var ret string
		return ret
	}
	return *o.BandcampHandle
}

// GetBandcampHandleOk returns a tuple with the BandcampHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTipJarSettings) GetBandcampHandleOk() (*string, bool) {
	if o == nil || IsNil(o.BandcampHandle) {
		return nil, false
	}
	return o.BandcampHandle, true
}

// HasBandcampHandle returns a boolean if a field has been set.
func (o *UserTipJarSettings) HasBandcampHandle() bool {
	if o != nil && !IsNil(o.BandcampHandle) {
		return true
	}

	return false
}

// SetBandcampHandle gets a reference to the given string and assigns it to the BandcampHandle field.
func (o *UserTipJarSettings) SetBandcampHandle(v string) {
	o.BandcampHandle = &v
}

// GetBitcoinHandle returns the BitcoinHandle field value if set, zero value otherwise.
func (o *UserTipJarSettings) GetBitcoinHandle() string {
	if o == nil || IsNil(o.BitcoinHandle) {
		var ret string
		return ret
	}
	return *o.BitcoinHandle
}

// GetBitcoinHandleOk returns a tuple with the BitcoinHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTipJarSettings) GetBitcoinHandleOk() (*string, bool) {
	if o == nil || IsNil(o.BitcoinHandle) {
		return nil, false
	}
	return o.BitcoinHandle, true
}

// HasBitcoinHandle returns a boolean if a field has been set.
func (o *UserTipJarSettings) HasBitcoinHandle() bool {
	if o != nil && !IsNil(o.BitcoinHandle) {
		return true
	}

	return false
}

// SetBitcoinHandle gets a reference to the given string and assigns it to the BitcoinHandle field.
func (o *UserTipJarSettings) SetBitcoinHandle(v string) {
	o.BitcoinHandle = &v
}

// GetCashAppHandle returns the CashAppHandle field value if set, zero value otherwise.
func (o *UserTipJarSettings) GetCashAppHandle() string {
	if o == nil || IsNil(o.CashAppHandle) {
		var ret string
		return ret
	}
	return *o.CashAppHandle
}

// GetCashAppHandleOk returns a tuple with the CashAppHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTipJarSettings) GetCashAppHandleOk() (*string, bool) {
	if o == nil || IsNil(o.CashAppHandle) {
		return nil, false
	}
	return o.CashAppHandle, true
}

// HasCashAppHandle returns a boolean if a field has been set.
func (o *UserTipJarSettings) HasCashAppHandle() bool {
	if o != nil && !IsNil(o.CashAppHandle) {
		return true
	}

	return false
}

// SetCashAppHandle gets a reference to the given string and assigns it to the CashAppHandle field.
func (o *UserTipJarSettings) SetCashAppHandle(v string) {
	o.CashAppHandle = &v
}

// GetEthereumHandle returns the EthereumHandle field value if set, zero value otherwise.
func (o *UserTipJarSettings) GetEthereumHandle() string {
	if o == nil || IsNil(o.EthereumHandle) {
		var ret string
		return ret
	}
	return *o.EthereumHandle
}

// GetEthereumHandleOk returns a tuple with the EthereumHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTipJarSettings) GetEthereumHandleOk() (*string, bool) {
	if o == nil || IsNil(o.EthereumHandle) {
		return nil, false
	}
	return o.EthereumHandle, true
}

// HasEthereumHandle returns a boolean if a field has been set.
func (o *UserTipJarSettings) HasEthereumHandle() bool {
	if o != nil && !IsNil(o.EthereumHandle) {
		return true
	}

	return false
}

// SetEthereumHandle gets a reference to the given string and assigns it to the EthereumHandle field.
func (o *UserTipJarSettings) SetEthereumHandle(v string) {
	o.EthereumHandle = &v
}

// GetGofundmeHandle returns the GofundmeHandle field value if set, zero value otherwise.
func (o *UserTipJarSettings) GetGofundmeHandle() string {
	if o == nil || IsNil(o.GofundmeHandle) {
		var ret string
		return ret
	}
	return *o.GofundmeHandle
}

// GetGofundmeHandleOk returns a tuple with the GofundmeHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTipJarSettings) GetGofundmeHandleOk() (*string, bool) {
	if o == nil || IsNil(o.GofundmeHandle) {
		return nil, false
	}
	return o.GofundmeHandle, true
}

// HasGofundmeHandle returns a boolean if a field has been set.
func (o *UserTipJarSettings) HasGofundmeHandle() bool {
	if o != nil && !IsNil(o.GofundmeHandle) {
		return true
	}

	return false
}

// SetGofundmeHandle gets a reference to the given string and assigns it to the GofundmeHandle field.
func (o *UserTipJarSettings) SetGofundmeHandle(v string) {
	o.GofundmeHandle = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *UserTipJarSettings) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTipJarSettings) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *UserTipJarSettings) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *UserTipJarSettings) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetPatreonHandle returns the PatreonHandle field value if set, zero value otherwise.
func (o *UserTipJarSettings) GetPatreonHandle() string {
	if o == nil || IsNil(o.PatreonHandle) {
		var ret string
		return ret
	}
	return *o.PatreonHandle
}

// GetPatreonHandleOk returns a tuple with the PatreonHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTipJarSettings) GetPatreonHandleOk() (*string, bool) {
	if o == nil || IsNil(o.PatreonHandle) {
		return nil, false
	}
	return o.PatreonHandle, true
}

// HasPatreonHandle returns a boolean if a field has been set.
func (o *UserTipJarSettings) HasPatreonHandle() bool {
	if o != nil && !IsNil(o.PatreonHandle) {
		return true
	}

	return false
}

// SetPatreonHandle gets a reference to the given string and assigns it to the PatreonHandle field.
func (o *UserTipJarSettings) SetPatreonHandle(v string) {
	o.PatreonHandle = &v
}

// GetVenmoHandle returns the VenmoHandle field value if set, zero value otherwise.
func (o *UserTipJarSettings) GetVenmoHandle() string {
	if o == nil || IsNil(o.VenmoHandle) {
		var ret string
		return ret
	}
	return *o.VenmoHandle
}

// GetVenmoHandleOk returns a tuple with the VenmoHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTipJarSettings) GetVenmoHandleOk() (*string, bool) {
	if o == nil || IsNil(o.VenmoHandle) {
		return nil, false
	}
	return o.VenmoHandle, true
}

// HasVenmoHandle returns a boolean if a field has been set.
func (o *UserTipJarSettings) HasVenmoHandle() bool {
	if o != nil && !IsNil(o.VenmoHandle) {
		return true
	}

	return false
}

// SetVenmoHandle gets a reference to the given string and assigns it to the VenmoHandle field.
func (o *UserTipJarSettings) SetVenmoHandle(v string) {
	o.VenmoHandle = &v
}

func (o UserTipJarSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTipJarSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BandcampHandle) {
		toSerialize["bandcamp_handle"] = o.BandcampHandle
	}
	if !IsNil(o.BitcoinHandle) {
		toSerialize["bitcoin_handle"] = o.BitcoinHandle
	}
	if !IsNil(o.CashAppHandle) {
		toSerialize["cash_app_handle"] = o.CashAppHandle
	}
	if !IsNil(o.EthereumHandle) {
		toSerialize["ethereum_handle"] = o.EthereumHandle
	}
	if !IsNil(o.GofundmeHandle) {
		toSerialize["gofundme_handle"] = o.GofundmeHandle
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.PatreonHandle) {
		toSerialize["patreon_handle"] = o.PatreonHandle
	}
	if !IsNil(o.VenmoHandle) {
		toSerialize["venmo_handle"] = o.VenmoHandle
	}
	return toSerialize, nil
}

type NullableUserTipJarSettings struct {
	value *UserTipJarSettings
	isSet bool
}

func (v NullableUserTipJarSettings) Get() *UserTipJarSettings {
	return v.value
}

func (v *NullableUserTipJarSettings) Set(val *UserTipJarSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTipJarSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTipJarSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTipJarSettings(val *UserTipJarSettings) *NullableUserTipJarSettings {
	return &NullableUserTipJarSettings{value: val, isSet: true}
}

func (v NullableUserTipJarSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTipJarSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


