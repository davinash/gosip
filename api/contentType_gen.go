// Package api :: This is auto generated file, do not edit manually
package api

import "encoding/json"

// Conf receives custom request config definition, e.g. custom headers, custom OData mod
func (contentType *ContentType) Conf(config *RequestConfig) *ContentType {
	contentType.config = config
	return contentType
}

// Select adds $select OData modifier
func (contentType *ContentType) Select(oDataSelect string) *ContentType {
	contentType.modifiers.AddSelect(oDataSelect)
	return contentType
}

// Expand adds $expand OData modifier
func (contentType *ContentType) Expand(oDataExpand string) *ContentType {
	contentType.modifiers.AddExpand(oDataExpand)
	return contentType
}

/* Response helpers */

// Data response helper
func (contentTypeResp *ContentTypeResp) Data() *ContentTypeInfo {
	data := NormalizeODataItem(*contentTypeResp)
	res := &ContentTypeInfo{}
	json.Unmarshal(data, &res)
	return res
}

// Normalized returns normalized body
func (contentTypeResp *ContentTypeResp) Normalized() []byte {
	return NormalizeODataItem(*contentTypeResp)
}