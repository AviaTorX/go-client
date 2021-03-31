/*
 * NFT Storage API
 *
 * # API Reference 
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nftstorage

import (
	"encoding/json"
)

// LinksFile struct for LinksFile
type LinksFile struct {
	Ipfs *string `json:"ipfs,omitempty"`
	Http *string `json:"http,omitempty"`
}

// NewLinksFile instantiates a new LinksFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksFile() *LinksFile {
	this := LinksFile{}
	return &this
}

// NewLinksFileWithDefaults instantiates a new LinksFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksFileWithDefaults() *LinksFile {
	this := LinksFile{}
	return &this
}

// GetIpfs returns the Ipfs field value if set, zero value otherwise.
func (o *LinksFile) GetIpfs() string {
	if o == nil || o.Ipfs == nil {
		var ret string
		return ret
	}
	return *o.Ipfs
}

// GetIpfsOk returns a tuple with the Ipfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksFile) GetIpfsOk() (*string, bool) {
	if o == nil || o.Ipfs == nil {
		return nil, false
	}
	return o.Ipfs, true
}

// HasIpfs returns a boolean if a field has been set.
func (o *LinksFile) HasIpfs() bool {
	if o != nil && o.Ipfs != nil {
		return true
	}

	return false
}

// SetIpfs gets a reference to the given string and assigns it to the Ipfs field.
func (o *LinksFile) SetIpfs(v string) {
	o.Ipfs = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *LinksFile) GetHttp() string {
	if o == nil || o.Http == nil {
		var ret string
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksFile) GetHttpOk() (*string, bool) {
	if o == nil || o.Http == nil {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *LinksFile) HasHttp() bool {
	if o != nil && o.Http != nil {
		return true
	}

	return false
}

// SetHttp gets a reference to the given string and assigns it to the Http field.
func (o *LinksFile) SetHttp(v string) {
	o.Http = &v
}

func (o LinksFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ipfs != nil {
		toSerialize["ipfs"] = o.Ipfs
	}
	if o.Http != nil {
		toSerialize["http"] = o.Http
	}
	return json.Marshal(toSerialize)
}

type NullableLinksFile struct {
	value *LinksFile
	isSet bool
}

func (v NullableLinksFile) Get() *LinksFile {
	return v.value
}

func (v *NullableLinksFile) Set(val *LinksFile) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksFile) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksFile(val *LinksFile) *NullableLinksFile {
	return &NullableLinksFile{value: val, isSet: true}
}

func (v NullableLinksFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


