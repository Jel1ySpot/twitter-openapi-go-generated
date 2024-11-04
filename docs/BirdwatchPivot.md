# BirdwatchPivot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallToAction** | Pointer to [**BirdwatchPivotCallToAction**](BirdwatchPivotCallToAction.md) |  | [optional] 
**DestinationUrl** | **string** |  | 
**Footer** | [**BirdwatchPivotFooter**](BirdwatchPivotFooter.md) |  | 
**IconType** | **string** |  | 
**Note** | [**BirdwatchPivotNote**](BirdwatchPivotNote.md) |  | 
**Shorttitle** | **string** |  | 
**Subtitle** | [**BirdwatchPivotSubtitle**](BirdwatchPivotSubtitle.md) |  | 
**Title** | **string** |  | 
**VisualStyle** | Pointer to **string** |  | [optional] 

## Methods

### NewBirdwatchPivot

`func NewBirdwatchPivot(destinationUrl string, footer BirdwatchPivotFooter, iconType string, note BirdwatchPivotNote, shorttitle string, subtitle BirdwatchPivotSubtitle, title string, ) *BirdwatchPivot`

NewBirdwatchPivot instantiates a new BirdwatchPivot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBirdwatchPivotWithDefaults

`func NewBirdwatchPivotWithDefaults() *BirdwatchPivot`

NewBirdwatchPivotWithDefaults instantiates a new BirdwatchPivot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallToAction

`func (o *BirdwatchPivot) GetCallToAction() BirdwatchPivotCallToAction`

GetCallToAction returns the CallToAction field if non-nil, zero value otherwise.

### GetCallToActionOk

`func (o *BirdwatchPivot) GetCallToActionOk() (*BirdwatchPivotCallToAction, bool)`

GetCallToActionOk returns a tuple with the CallToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToAction

`func (o *BirdwatchPivot) SetCallToAction(v BirdwatchPivotCallToAction)`

SetCallToAction sets CallToAction field to given value.

### HasCallToAction

`func (o *BirdwatchPivot) HasCallToAction() bool`

HasCallToAction returns a boolean if a field has been set.

### GetDestinationUrl

`func (o *BirdwatchPivot) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *BirdwatchPivot) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *BirdwatchPivot) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.


### GetFooter

`func (o *BirdwatchPivot) GetFooter() BirdwatchPivotFooter`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *BirdwatchPivot) GetFooterOk() (*BirdwatchPivotFooter, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *BirdwatchPivot) SetFooter(v BirdwatchPivotFooter)`

SetFooter sets Footer field to given value.


### GetIconType

`func (o *BirdwatchPivot) GetIconType() string`

GetIconType returns the IconType field if non-nil, zero value otherwise.

### GetIconTypeOk

`func (o *BirdwatchPivot) GetIconTypeOk() (*string, bool)`

GetIconTypeOk returns a tuple with the IconType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconType

`func (o *BirdwatchPivot) SetIconType(v string)`

SetIconType sets IconType field to given value.


### GetNote

`func (o *BirdwatchPivot) GetNote() BirdwatchPivotNote`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BirdwatchPivot) GetNoteOk() (*BirdwatchPivotNote, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BirdwatchPivot) SetNote(v BirdwatchPivotNote)`

SetNote sets Note field to given value.


### GetShorttitle

`func (o *BirdwatchPivot) GetShorttitle() string`

GetShorttitle returns the Shorttitle field if non-nil, zero value otherwise.

### GetShorttitleOk

`func (o *BirdwatchPivot) GetShorttitleOk() (*string, bool)`

GetShorttitleOk returns a tuple with the Shorttitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShorttitle

`func (o *BirdwatchPivot) SetShorttitle(v string)`

SetShorttitle sets Shorttitle field to given value.


### GetSubtitle

`func (o *BirdwatchPivot) GetSubtitle() BirdwatchPivotSubtitle`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *BirdwatchPivot) GetSubtitleOk() (*BirdwatchPivotSubtitle, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *BirdwatchPivot) SetSubtitle(v BirdwatchPivotSubtitle)`

SetSubtitle sets Subtitle field to given value.


### GetTitle

`func (o *BirdwatchPivot) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BirdwatchPivot) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BirdwatchPivot) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVisualStyle

`func (o *BirdwatchPivot) GetVisualStyle() string`

GetVisualStyle returns the VisualStyle field if non-nil, zero value otherwise.

### GetVisualStyleOk

`func (o *BirdwatchPivot) GetVisualStyleOk() (*string, bool)`

GetVisualStyleOk returns a tuple with the VisualStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualStyle

`func (o *BirdwatchPivot) SetVisualStyle(v string)`

SetVisualStyle sets VisualStyle field to given value.

### HasVisualStyle

`func (o *BirdwatchPivot) HasVisualStyle() bool`

HasVisualStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


