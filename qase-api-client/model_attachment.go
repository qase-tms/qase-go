/*
Qase.io TestOps API v1

Qase TestOps API v1 Specification.

API version: 1.0.0
Contact: support@qase.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v1_client

import (
	"encoding/json"
)

// checks if the Attachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attachment{}

// Attachment struct for Attachment
type Attachment struct {
	Size     *int32  `json:"size,omitempty"`
	Mime     *string `json:"mime,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Url      *string `json:"url,omitempty"`
}

// NewAttachment instantiates a new Attachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachment() *Attachment {
	this := Attachment{}
	return &this
}

// NewAttachmentWithDefaults instantiates a new Attachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentWithDefaults() *Attachment {
	this := Attachment{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Attachment) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Attachment) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *Attachment) SetSize(v int32) {
	o.Size = &v
}

// GetMime returns the Mime field value if set, zero value otherwise.
func (o *Attachment) GetMime() string {
	if o == nil || IsNil(o.Mime) {
		var ret string
		return ret
	}
	return *o.Mime
}

// GetMimeOk returns a tuple with the Mime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetMimeOk() (*string, bool) {
	if o == nil || IsNil(o.Mime) {
		return nil, false
	}
	return o.Mime, true
}

// HasMime returns a boolean if a field has been set.
func (o *Attachment) HasMime() bool {
	if o != nil && !IsNil(o.Mime) {
		return true
	}

	return false
}

// SetMime gets a reference to the given string and assigns it to the Mime field.
func (o *Attachment) SetMime(v string) {
	o.Mime = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *Attachment) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *Attachment) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *Attachment) SetFilename(v string) {
	o.Filename = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Attachment) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Attachment) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Attachment) SetUrl(v string) {
	o.Url = &v
}

func (o Attachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Mime) {
		toSerialize["mime"] = o.Mime
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAttachment struct {
	value *Attachment
	isSet bool
}

func (v NullableAttachment) Get() *Attachment {
	return v.value
}

func (v *NullableAttachment) Set(val *Attachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachment(val *Attachment) *NullableAttachment {
	return &NullableAttachment{value: val, isSet: true}
}

func (v NullableAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
