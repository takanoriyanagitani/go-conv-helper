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
