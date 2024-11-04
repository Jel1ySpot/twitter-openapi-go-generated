# Timeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instructions** | [**[]InstructionUnion**](InstructionUnion.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ResponseObjects** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTimeline

`func NewTimeline(instructions []InstructionUnion, ) *Timeline`

NewTimeline instantiates a new Timeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineWithDefaults

`func NewTimelineWithDefaults() *Timeline`

NewTimelineWithDefaults instantiates a new Timeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *Timeline) GetInstructions() []InstructionUnion`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Timeline) GetInstructionsOk() (*[]InstructionUnion, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Timeline) SetInstructions(v []InstructionUnion)`

SetInstructions sets Instructions field to given value.


### GetMetadata

`func (o *Timeline) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Timeline) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Timeline) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Timeline) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResponseObjects

`func (o *Timeline) GetResponseObjects() map[string]interface{}`

GetResponseObjects returns the ResponseObjects field if non-nil, zero value otherwise.

### GetResponseObjectsOk

`func (o *Timeline) GetResponseObjectsOk() (*map[string]interface{}, bool)`

GetResponseObjectsOk returns a tuple with the ResponseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseObjects

`func (o *Timeline) SetResponseObjects(v map[string]interface{})`

SetResponseObjects sets ResponseObjects field to given value.

### HasResponseObjects

`func (o *Timeline) HasResponseObjects() bool`

HasResponseObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


