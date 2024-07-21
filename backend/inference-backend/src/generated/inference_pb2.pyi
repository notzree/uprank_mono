from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DeleteVectorRequest(_message.Message):
    __slots__ = ["ids", "namespace"]
    IDS_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    ids: _containers.RepeatedScalarFieldContainer[str]
    namespace: str
    def __init__(self, ids: _Optional[_Iterable[str]] = ..., namespace: _Optional[str] = ...) -> None: ...

class DeleteVectorResponse(_message.Message):
    __slots__ = ["error", "ids", "ok"]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    IDS_FIELD_NUMBER: _ClassVar[int]
    OK_FIELD_NUMBER: _ClassVar[int]
    error: str
    ids: _containers.RepeatedScalarFieldContainer[str]
    ok: bool
    def __init__(self, ids: _Optional[_Iterable[str]] = ..., ok: bool = ..., error: _Optional[str] = ...) -> None: ...

class EmbedTextRequest(_message.Message):
    __slots__ = ["id", "metadata", "text"]
    class MetadataEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    id: str
    metadata: _containers.ScalarMap[str, str]
    text: str
    def __init__(self, id: _Optional[str] = ..., text: _Optional[str] = ..., metadata: _Optional[_Mapping[str, str]] = ...) -> None: ...

class EmbedTextResponse(_message.Message):
    __slots__ = ["error", "vector"]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    VECTOR_FIELD_NUMBER: _ClassVar[int]
    error: str
    vector: Vector
    def __init__(self, vector: _Optional[_Union[Vector, _Mapping]] = ..., error: _Optional[str] = ...) -> None: ...

class Match(_message.Message):
    __slots__ = ["id", "metadata", "score"]
    class MetadataEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    SCORE_FIELD_NUMBER: _ClassVar[int]
    id: str
    metadata: _containers.ScalarMap[str, str]
    score: float
    def __init__(self, score: _Optional[float] = ..., id: _Optional[str] = ..., metadata: _Optional[_Mapping[str, str]] = ...) -> None: ...

class QueryVectorRequest(_message.Message):
    __slots__ = ["filter", "namespace", "top_k", "vector"]
    class FilterEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    FILTER_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    TOP_K_FIELD_NUMBER: _ClassVar[int]
    VECTOR_FIELD_NUMBER: _ClassVar[int]
    filter: _containers.ScalarMap[str, str]
    namespace: str
    top_k: int
    vector: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, namespace: _Optional[str] = ..., vector: _Optional[_Iterable[float]] = ..., top_k: _Optional[int] = ..., filter: _Optional[_Mapping[str, str]] = ...) -> None: ...

class QueryVectorResponse(_message.Message):
    __slots__ = ["matches", "namespace", "usage"]
    MATCHES_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    USAGE_FIELD_NUMBER: _ClassVar[int]
    matches: _containers.RepeatedCompositeFieldContainer[Match]
    namespace: str
    usage: Usage
    def __init__(self, matches: _Optional[_Iterable[_Union[Match, _Mapping]]] = ..., namespace: _Optional[str] = ..., usage: _Optional[_Union[Usage, _Mapping]] = ...) -> None: ...

class UpsertVectorRequest(_message.Message):
    __slots__ = ["namespace", "vectors"]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    VECTORS_FIELD_NUMBER: _ClassVar[int]
    namespace: str
    vectors: _containers.RepeatedCompositeFieldContainer[Vector]
    def __init__(self, namespace: _Optional[str] = ..., vectors: _Optional[_Iterable[_Union[Vector, _Mapping]]] = ...) -> None: ...

class UpsertVectorResponse(_message.Message):
    __slots__ = ["error"]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    error: str
    def __init__(self, error: _Optional[str] = ...) -> None: ...

class Usage(_message.Message):
    __slots__ = ["read_units"]
    READ_UNITS_FIELD_NUMBER: _ClassVar[int]
    read_units: int
    def __init__(self, read_units: _Optional[int] = ...) -> None: ...

class Vector(_message.Message):
    __slots__ = ["id", "metadata", "vector"]
    class MetadataEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    VECTOR_FIELD_NUMBER: _ClassVar[int]
    id: str
    metadata: _containers.ScalarMap[str, str]
    vector: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, id: _Optional[str] = ..., vector: _Optional[_Iterable[float]] = ..., metadata: _Optional[_Mapping[str, str]] = ...) -> None: ...
