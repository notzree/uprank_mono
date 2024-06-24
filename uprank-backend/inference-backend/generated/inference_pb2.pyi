from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class EmbedVectorRequest(_message.Message):
    __slots__ = ("text",)
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...

class EmbedVectorResponse(_message.Message):
    __slots__ = ("vector",)
    VECTOR_FIELD_NUMBER: _ClassVar[int]
    vector: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, vector: _Optional[_Iterable[float]] = ...) -> None: ...

class UpsertVectorRequest(_message.Message):
    __slots__ = ("id", "vector", "metadata")
    class MetadataEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    VECTOR_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    id: str
    vector: _containers.RepeatedScalarFieldContainer[float]
    metadata: _containers.ScalarMap[str, str]
    def __init__(self, id: _Optional[str] = ..., vector: _Optional[_Iterable[float]] = ..., metadata: _Optional[_Mapping[str, str]] = ...) -> None: ...

class UpsertVectorResponse(_message.Message):
    __slots__ = ("id", "ok", "error")
    ID_FIELD_NUMBER: _ClassVar[int]
    OK_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    id: str
    ok: bool
    error: str
    def __init__(self, id: _Optional[str] = ..., ok: bool = ..., error: _Optional[str] = ...) -> None: ...

class QueryVectorRequest(_message.Message):
    __slots__ = ("namespace", "vector", "top_k", "filter")
    class FilterEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    VECTOR_FIELD_NUMBER: _ClassVar[int]
    TOP_K_FIELD_NUMBER: _ClassVar[int]
    FILTER_FIELD_NUMBER: _ClassVar[int]
    namespace: str
    vector: _containers.RepeatedScalarFieldContainer[float]
    top_k: int
    filter: _containers.ScalarMap[str, str]
    def __init__(self, namespace: _Optional[str] = ..., vector: _Optional[_Iterable[float]] = ..., top_k: _Optional[int] = ..., filter: _Optional[_Mapping[str, str]] = ...) -> None: ...

class QueryVectorResponse(_message.Message):
    __slots__ = ("matches", "namespace", "usage")
    MATCHES_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    USAGE_FIELD_NUMBER: _ClassVar[int]
    matches: _containers.RepeatedCompositeFieldContainer[Match]
    namespace: str
    usage: Usage
    def __init__(self, matches: _Optional[_Iterable[_Union[Match, _Mapping]]] = ..., namespace: _Optional[str] = ..., usage: _Optional[_Union[Usage, _Mapping]] = ...) -> None: ...

class Match(_message.Message):
    __slots__ = ("id", "score", "values")
    ID_FIELD_NUMBER: _ClassVar[int]
    SCORE_FIELD_NUMBER: _ClassVar[int]
    VALUES_FIELD_NUMBER: _ClassVar[int]
    id: str
    score: float
    values: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, id: _Optional[str] = ..., score: _Optional[float] = ..., values: _Optional[_Iterable[float]] = ...) -> None: ...

class Usage(_message.Message):
    __slots__ = ("read_units",)
    READ_UNITS_FIELD_NUMBER: _ClassVar[int]
    read_units: int
    def __init__(self, read_units: _Optional[int] = ...) -> None: ...
