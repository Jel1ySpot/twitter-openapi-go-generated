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
	"bytes"
	"fmt"
)

// checks if the BirdwatchPivot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BirdwatchPivot{}

// BirdwatchPivot struct for BirdwatchPivot
type BirdwatchPivot struct {
	CallToAction *BirdwatchPivotCallToAction `json:"callToAction,omitempty"`
	DestinationUrl string `json:"destinationUrl"`
	Footer BirdwatchPivotFooter `json:"footer"`
	IconType string `json:"iconType"`
	Note BirdwatchPivotNote `json:"note"`
	Shorttitle string `json:"shorttitle"`
	Subtitle BirdwatchPivotSubtitle `json:"subtitle"`
	Title string `json:"title"`
	VisualStyle *string `json:"visualStyle,omitempty"`
}

type _BirdwatchPivot BirdwatchPivot

// NewBirdwatchPivot instantiates a new BirdwatchPivot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBirdwatchPivot(destinationUrl string, footer BirdwatchPivotFooter, iconType string, note BirdwatchPivotNote, shorttitle string, subtitle BirdwatchPivotSubtitle, title string) *BirdwatchPivot {
	this := BirdwatchPivot{}
	this.DestinationUrl = destinationUrl
	this.Footer = footer
	this.IconType = iconType
	this.Note = note
	this.Shorttitle = shorttitle
	this.Subtitle = subtitle
	this.Title = title
	return &this
}

// NewBirdwatchPivotWithDefaults instantiates a new BirdwatchPivot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBirdwatchPivotWithDefaults() *BirdwatchPivot {
	this := BirdwatchPivot{}
	return &this
}

// GetCallToAction returns the CallToAction field value if set, zero value otherwise.
func (o *BirdwatchPivot) GetCallToAction() BirdwatchPivotCallToAction {
	if o == nil || IsNil(o.CallToAction) {
		var ret BirdwatchPivotCallToAction
		return ret
	}
	return *o.CallToAction
}

// GetCallToActionOk returns a tuple with the CallToAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BirdwatchPivot) GetCallToActionOk() (*BirdwatchPivotCallToAction, bool) {
	if o == nil || IsNil(o.CallToAction) {
		return nil, false
	}
	return o.CallToAction, true
}

// HasCallToAction returns a boolean if a field has been set.
func (o *BirdwatchPivot) HasCallToAction() bool {
	if o != nil && !IsNil(o.CallToAction) {
		return true
	}

	return false
}

// SetCallToAction gets a reference to the given BirdwatchPivotCallToAction and assigns it to the CallToAction field.
func (o *BirdwatchPivot) SetCallToAction(v BirdwatchPivotCallToAction) {
	o.CallToAction = &v
}

// GetDestinationUrl returns the DestinationUrl field value
func (o *BirdwatchPivot) GetDestinationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationUrl
}

// GetDestinationUrlOk returns a tuple with the DestinationUrl field value
// and a boolean to check if the value has been set.
func (o *BirdwatchPivot) GetDestinationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationUrl, true
}

// SetDestinationUrl sets field value
func (o *BirdwatchPivot) SetDestinationUrl(v string) {
	o.DestinationUrl = v
}

// GetFooter returns the Footer field value
func (o *BirdwatchPivot) GetFooter() BirdwatchPivotFooter {
	if o == nil {
		var ret BirdwatchPivotFooter
		return ret
	}

	return o.Footer
}

// GetFooterOk returns a tuple with the Footer field value
// and a boolean to check if the value has been set.
func (o *BirdwatchPivot) GetFooterOk() (*BirdwatchPivotFooter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Footer, true
}

// SetFooter sets field value
func (o *BirdwatchPivot) SetFooter(v BirdwatchPivotFooter) {
	o.Footer = v
}

// GetIconType returns the IconType field value
func (o *BirdwatchPivot) GetIconType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IconType
}

// GetIconTypeOk returns a tuple with the IconType field value
// and a boolean to check if the value has been set.
func (o *BirdwatchPivot) GetIconTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconType, true
}

// SetIconType sets field value
func (o *BirdwatchPivot) SetIconType(v string) {
	o.IconType = v
}

// GetNote returns the Note field value
func (o *BirdwatchPivot) GetNote() BirdwatchPivotNote {
	if o == nil {
		var ret BirdwatchPivotNote
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *BirdwatchPivot) GetNoteOk() (*BirdwatchPivotNote, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *BirdwatchPivot) SetNote(v BirdwatchPivotNote) {
	o.Note = v
}

// GetShorttitle returns the Shorttitle field value
func (o *BirdwatchPivot) GetShorttitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Shorttitle
}

// GetShorttitleOk returns a tuple with the Shorttitle field value
// and a boolean to check if the value has been set.
func (o *BirdwatchPivot) GetShorttitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shorttitle, true
}

// SetShorttitle sets field value
func (o *BirdwatchPivot) SetShorttitle(v string) {
	o.Shorttitle = v
}

// GetSubtitle returns the Subtitle field value
func (o *BirdwatchPivot) GetSubtitle() BirdwatchPivotSubtitle {
	if o == nil {
		var ret BirdwatchPivotSubtitle
		return ret
	}

	return o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value
// and a boolean to check if the value has been set.
func (o *BirdwatchPivot) GetSubtitleOk() (*BirdwatchPivotSubtitle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtitle, true
}

// SetSubtitle sets field value
func (o *BirdwatchPivot) SetSubtitle(v BirdwatchPivotSubtitle) {
	o.Subtitle = v
}

// GetTitle returns the Title field value
func (o *BirdwatchPivot) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BirdwatchPivot) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BirdwatchPivot) SetTitle(v string) {
	o.Title = v
}

// GetVisualStyle returns the VisualStyle field value if set, zero value otherwise.
func (o *BirdwatchPivot) GetVisualStyle() string {
	if o == nil || IsNil(o.VisualStyle) {
		var ret string
		return ret
	}
	return *o.VisualStyle
}

// GetVisualStyleOk returns a tuple with the VisualStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BirdwatchPivot) GetVisualStyleOk() (*string, bool) {
	if o == nil || IsNil(o.VisualStyle) {
		return nil, false
	}
	return o.VisualStyle, true
}

// HasVisualStyle returns a boolean if a field has been set.
func (o *BirdwatchPivot) HasVisualStyle() bool {
	if o != nil && !IsNil(o.VisualStyle) {
		return true
	}

	return false
}

// SetVisualStyle gets a reference to the given string and assigns it to the VisualStyle field.
func (o *BirdwatchPivot) SetVisualStyle(v string) {
	o.VisualStyle = &v
}

func (o BirdwatchPivot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BirdwatchPivot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallToAction) {
		toSerialize["callToAction"] = o.CallToAction
	}
	toSerialize["destinationUrl"] = o.DestinationUrl
	toSerialize["footer"] = o.Footer
	toSerialize["iconType"] = o.IconType
	toSerialize["note"] = o.Note
	toSerialize["shorttitle"] = o.Shorttitle
	toSerialize["subtitle"] = o.Subtitle
	toSerialize["title"] = o.Title
	if !IsNil(o.VisualStyle) {
		toSerialize["visualStyle"] = o.VisualStyle
	}
	return toSerialize, nil
}

func (o *BirdwatchPivot) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destinationUrl",
		"footer",
		"iconType",
		"note",
		"shorttitle",
		"subtitle",
		"title",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBirdwatchPivot := _BirdwatchPivot{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBirdwatchPivot)

	if err != nil {
		return err
	}

	*o = BirdwatchPivot(varBirdwatchPivot)

	return err
}

type NullableBirdwatchPivot struct {
	value *BirdwatchPivot
	isSet bool
}

func (v NullableBirdwatchPivot) Get() *BirdwatchPivot {
	return v.value
}

func (v *NullableBirdwatchPivot) Set(val *BirdwatchPivot) {
	v.value = val
	v.isSet = true
}

func (v NullableBirdwatchPivot) IsSet() bool {
	return v.isSet
}

func (v *NullableBirdwatchPivot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBirdwatchPivot(val *BirdwatchPivot) *NullableBirdwatchPivot {
	return &NullableBirdwatchPivot{value: val, isSet: true}
}

func (v NullableBirdwatchPivot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBirdwatchPivot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

