# ClientEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** |  | [optional] 
**Element** | Pointer to **string** |  | [optional] 

## Methods

### NewClientEventInfo

`func NewClientEventInfo() *ClientEventInfo`

NewClientEventInfo instantiates a new ClientEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientEventInfoWithDefaults

`func NewClientEventInfoWithDefaults() *ClientEventInfo`

NewClientEventInfoWithDefaults instantiates a new ClientEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ClientEventInfo) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ClientEventInfo) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ClientEventInfo) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ClientEventInfo) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDetails

`func (o *ClientEventInfo) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ClientEventInfo) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ClientEventInfo) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ClientEventInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetElement

`func (o *ClientEventInfo) GetElement() string`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *ClientEventInfo) GetElementOk() (*string, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *ClientEventInfo) SetElement(v string)`

SetElement sets Element field to given value.

### HasElement

`func (o *ClientEventInfo) HasElement() bool`

HasElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


