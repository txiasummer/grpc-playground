# proto buffers

## Scalar Types
|  Type   |      Keyword      |             Value              |       Default        |                   Note                    |
|:-------:|:-----------------:|:------------------------------:|:--------------------:|:-----------------------------------------:|
| Numeric | See details below |         numeric value          |          0           |                                           |
| Boolean |      `bool`       |          true / false          |       `false`        |                                           |
| String  |     `string`      |     arbitrary length text      |     empty string     |                                           |
| String  |     `string`      |     arbitrary length text      |     empty string     | Only accepts UTF-8 encoded or 7-bit ASCII |
|  Bytes  |      `bytes`      | arbitrary length byte sequence | empty array of bytes |      Up to you to interpret in code       |

### Numeric
Numeric data types can be of the following:
* signed integers: `int32, int64, sint32, sint64`
* unsigned integers: `uint32, uint64`
* fixed: `fixed32, fixed64, sfixed32, sfixed64`
* decimals: `float, double`

## Tags
Variable names are _NOT_ used to serialize protobuf messages, tags are! The smallest tag is `1` and the 
largest tag is `536,870,911`. Google reserves tags from `19000` to `19999`.

It is best practice to reserve the smallest tags for the most frequently populated fields:
* tags 1 ~ 15 takes up 1 byte of memory.
* tags 16 ~ 2047 takes up 2 bytes of memory.

## Repeated Fields
* Syntax: `repeated <type> <name> = <tag>;`
* Value: Any number of elements (0 or more) of the specified type.
* Default: an empty list of the specified type

## Enums
* Keyword: enum
* Default: first value
* Import note: 
  * First tag is `0` instead of `1`.
