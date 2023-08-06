# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/bytes2bz/v1/b2b.proto](#proto_bytes2bz_v1_b2b-proto)
    - [CompressRequest](#bytes2bz-v1-CompressRequest)
    - [CompressResponse](#bytes2bz-v1-CompressResponse)
    - [Compression](#bytes2bz-v1-Compression)
  
    - [Compression.Level](#bytes2bz-v1-Compression-Level)
  
    - [CompressService](#bytes2bz-v1-CompressService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_bytes2bz_v1_b2b-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/bytes2bz/v1/b2b.proto



<a name="bytes2bz-v1-CompressRequest"></a>

### CompressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| level | [Compression.Level](#bytes2bz-v1-Compression-Level) |  |  |
| input | [bytes](#bytes) |  |  |






<a name="bytes2bz-v1-CompressResponse"></a>

### CompressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| compressed | [bytes](#bytes) |  |  |






<a name="bytes2bz-v1-Compression"></a>

### Compression






 


<a name="bytes2bz-v1-Compression-Level"></a>

### Compression.Level


| Name | Number | Description |
| ---- | ------ | ----------- |
| LEVEL_NUM_UNSPECIFIED | 0 |  |
| LEVEL_NUM_FAST | 1 |  |
| LEVEL_NUM_BEST | 2 |  |
| LEVEL_NUM_0 | 3 |  |
| LEVEL_NUM_1 | 4 |  |
| LEVEL_NUM_2 | 5 |  |
| LEVEL_NUM_3 | 6 |  |
| LEVEL_NUM_4 | 7 |  |
| LEVEL_NUM_5 | 8 |  |
| LEVEL_NUM_6 | 9 |  |
| LEVEL_NUM_7 | 10 |  |
| LEVEL_NUM_8 | 11 |  |
| LEVEL_NUM_9 | 12 |  |


 

 


<a name="bytes2bz-v1-CompressService"></a>

### CompressService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Compress | [CompressRequest](#bytes2bz-v1-CompressRequest) | [CompressResponse](#bytes2bz-v1-CompressResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

