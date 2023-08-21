from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Rgb(_message.Message):
    __slots__ = ["red", "green", "blue"]
    RED_FIELD_NUMBER: _ClassVar[int]
    GREEN_FIELD_NUMBER: _ClassVar[int]
    BLUE_FIELD_NUMBER: _ClassVar[int]
    red: float
    green: float
    blue: float
    def __init__(self, red: _Optional[float] = ..., green: _Optional[float] = ..., blue: _Optional[float] = ...) -> None: ...

class Hsv(_message.Message):
    __slots__ = ["hue", "saturation", "value"]
    HUE_FIELD_NUMBER: _ClassVar[int]
    SATURATION_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    hue: float
    saturation: float
    value: float
    def __init__(self, hue: _Optional[float] = ..., saturation: _Optional[float] = ..., value: _Optional[float] = ...) -> None: ...

class FromHsvRequest(_message.Message):
    __slots__ = ["hsv"]
    HSV_FIELD_NUMBER: _ClassVar[int]
    hsv: Hsv
    def __init__(self, hsv: _Optional[_Union[Hsv, _Mapping]] = ...) -> None: ...

class FromHsvResponse(_message.Message):
    __slots__ = ["rgb"]
    RGB_FIELD_NUMBER: _ClassVar[int]
    rgb: Rgb
    def __init__(self, rgb: _Optional[_Union[Rgb, _Mapping]] = ...) -> None: ...

class Timestamp(_message.Message):
    __slots__ = ["unixtime_micros"]
    UNIXTIME_MICROS_FIELD_NUMBER: _ClassVar[int]
    unixtime_micros: int
    def __init__(self, unixtime_micros: _Optional[int] = ...) -> None: ...

class HsvEvt(_message.Message):
    __slots__ = []
    class ConvertedRequest(_message.Message):
        __slots__ = ["meta", "converted"]
        META_FIELD_NUMBER: _ClassVar[int]
        CONVERTED_FIELD_NUMBER: _ClassVar[int]
        meta: Meta
        converted: Rgb
        def __init__(self, meta: _Optional[_Union[Meta, _Mapping]] = ..., converted: _Optional[_Union[Rgb, _Mapping]] = ...) -> None: ...
    class ConvertedResponse(_message.Message):
        __slots__ = ["sent"]
        SENT_FIELD_NUMBER: _ClassVar[int]
        sent: Timestamp
        def __init__(self, sent: _Optional[_Union[Timestamp, _Mapping]] = ...) -> None: ...
    def __init__(self) -> None: ...

class Uuid(_message.Message):
    __slots__ = ["hi", "lo"]
    HI_FIELD_NUMBER: _ClassVar[int]
    LO_FIELD_NUMBER: _ClassVar[int]
    hi: int
    lo: int
    def __init__(self, hi: _Optional[int] = ..., lo: _Optional[int] = ...) -> None: ...

class Meta(_message.Message):
    __slots__ = ["request_id", "reply_id"]
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    REPLY_ID_FIELD_NUMBER: _ClassVar[int]
    request_id: Uuid
    reply_id: Uuid
    def __init__(self, request_id: _Optional[_Union[Uuid, _Mapping]] = ..., reply_id: _Optional[_Union[Uuid, _Mapping]] = ...) -> None: ...

class HsvCmd(_message.Message):
    __slots__ = []
    class ConvertRequest(_message.Message):
        __slots__ = ["meta", "hsv"]
        META_FIELD_NUMBER: _ClassVar[int]
        HSV_FIELD_NUMBER: _ClassVar[int]
        meta: Meta
        hsv: Hsv
        def __init__(self, meta: _Optional[_Union[Meta, _Mapping]] = ..., hsv: _Optional[_Union[Hsv, _Mapping]] = ...) -> None: ...
    class ConvertResponse(_message.Message):
        __slots__ = ["rgb"]
        RGB_FIELD_NUMBER: _ClassVar[int]
        rgb: Rgb
        def __init__(self, rgb: _Optional[_Union[Rgb, _Mapping]] = ...) -> None: ...
    def __init__(self) -> None: ...

class HsvOrder(_message.Message):
    __slots__ = []
    class GetCmdRequest(_message.Message):
        __slots__ = ["request_id"]
        REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
        request_id: Uuid
        def __init__(self, request_id: _Optional[_Union[Uuid, _Mapping]] = ...) -> None: ...
    class GetCmdResponse(_message.Message):
        __slots__ = ["commands"]
        COMMANDS_FIELD_NUMBER: _ClassVar[int]
        commands: HsvOrder.Commands
        def __init__(self, commands: _Optional[_Union[HsvOrder.Commands, _Mapping]] = ...) -> None: ...
    class InvalidCmd(_message.Message):
        __slots__ = ["meta", "error"]
        META_FIELD_NUMBER: _ClassVar[int]
        ERROR_FIELD_NUMBER: _ClassVar[int]
        meta: Meta
        error: str
        def __init__(self, meta: _Optional[_Union[Meta, _Mapping]] = ..., error: _Optional[str] = ...) -> None: ...
    class Commands(_message.Message):
        __slots__ = ["invalid", "fatal", "convert"]
        INVALID_FIELD_NUMBER: _ClassVar[int]
        FATAL_FIELD_NUMBER: _ClassVar[int]
        CONVERT_FIELD_NUMBER: _ClassVar[int]
        invalid: HsvOrder.InvalidCmd
        fatal: str
        convert: HsvCmd.ConvertRequest
        def __init__(self, invalid: _Optional[_Union[HsvOrder.InvalidCmd, _Mapping]] = ..., fatal: _Optional[str] = ..., convert: _Optional[_Union[HsvCmd.ConvertRequest, _Mapping]] = ...) -> None: ...
    def __init__(self) -> None: ...
