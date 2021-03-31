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

// Links struct for Links
type Links struct {
	Ipfs *string `json:"ipfs,omitempty"`
	Http *string `json:"http,omitempty"`
	File *[]LinksFile `json:"file,omitempty"`
}

// NewLinks instantiates a new Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinks() *Links {
	this := Links{}
	return &this
}

// NewLinksWithDefaults instantiates a new Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetIpfs returns the Ipfs field value if set, zero value otherwise.
func (o *Links) GetIpfs() string {
	if o == nil || o.Ipfs == nil {
		var ret string
		return ret
	}
	return *o.Ipfs
}

// GetIpfsOk returns a tuple with the Ipfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetIpfsOk() (*string, bool) {
	if o == nil || o.Ipfs == nil {
		return nil, false
	}
	return o.Ipfs, true
}

// HasIpfs returns a boolean if a field has been set.
func (o *Links) HasIpfs() bool {
	if o != nil && o.Ipfs != nil {
		return true
	}

	return false
}

// SetIpfs gets a reference to the given string and assigns it to the Ipfs field.
func (o *Links) SetIpfs(v string) {
	o.Ipfs = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *Links) GetHttp() string {
	if o == nil || o.Http == nil {
		var ret string
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetHttpOk() (*string, bool) {
	if o == nil || o.Http == nil {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *Links) HasHttp() bool {
	if o != nil && o.Http != nil {
		return true
	}

	return false
}

// SetHttp gets a reference to the given string and assigns it to the Http field.
func (o *Links) SetHttp(v string) {
	o.Http = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *Links) GetFile() []LinksFile {
	if o == nil || o.File == nil {
		var ret []LinksFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetFileOk() (*[]LinksFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *Links) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given []LinksFile and assigns it to the File field.
func (o *Links) SetFile(v []LinksFile) {
	o.File = &v
}

func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ipfs != nil {
		toSerialize["ipfs"] = o.Ipfs
	}
	if o.Http != nil {
		toSerialize["http"] = o.Http
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableLinks struct {
	value *Links
	isSet bool
}

func (v NullableLinks) Get() *Links {
	return v.value
}

func (v *NullableLinks) Set(val *Links) {
	v.value = val
	v.isSet = true
}

func (v NullableLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinks(val *Links) *NullableLinks {
	return &NullableLinks{value: val, isSet: true}
}

func (v NullableLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


