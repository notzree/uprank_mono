# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: inference.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0finference.proto\x12\tinference\"\x9a\x01\n\x10\x45mbedTextRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04text\x18\x02 \x01(\t\x12;\n\x08metadata\x18\x03 \x03(\x0b\x32).inference.EmbedTextRequest.MetadataEntry\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"E\n\x11\x45mbedTextResponse\x12!\n\x06vector\x18\x01 \x01(\x0b\x32\x11.inference.Vector\x12\r\n\x05\x65rror\x18\x02 \x01(\t\"L\n\x13UpsertVectorRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\"\n\x07vectors\x18\x02 \x03(\x0b\x32\x11.inference.Vector\"\x88\x01\n\x06Vector\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0e\n\x06vector\x18\x02 \x03(\x02\x12\x31\n\x08metadata\x18\x03 \x03(\x0b\x32\x1f.inference.Vector.MetadataEntry\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"%\n\x14UpsertVectorResponse\x12\r\n\x05\x65rror\x18\x01 \x01(\t\"\xb0\x01\n\x12QueryVectorRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0e\n\x06vector\x18\x02 \x03(\x02\x12\r\n\x05top_k\x18\x03 \x01(\x05\x12\x39\n\x06\x66ilter\x18\x04 \x03(\x0b\x32).inference.QueryVectorRequest.FilterEntry\x1a-\n\x0b\x46ilterEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"l\n\x13QueryVectorResponse\x12!\n\x07matches\x18\x01 \x03(\x0b\x32\x10.inference.Match\x12\x11\n\tnamespace\x18\x02 \x01(\t\x12\x1f\n\x05usage\x18\x03 \x01(\x0b\x32\x10.inference.Usage\"\x85\x01\n\x05Match\x12\r\n\x05score\x18\x01 \x01(\x02\x12\n\n\x02id\x18\x02 \x01(\t\x12\x30\n\x08metadata\x18\x03 \x03(\x0b\x32\x1e.inference.Match.MetadataEntry\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x1b\n\x05Usage\x12\x12\n\nread_units\x18\x01 \x01(\x05\"5\n\x13\x44\x65leteVectorRequest\x12\x0b\n\x03ids\x18\x01 \x03(\t\x12\x11\n\tnamespace\x18\x02 \x01(\t\">\n\x14\x44\x65leteVectorResponse\x12\x0b\n\x03ids\x18\x01 \x03(\t\x12\n\n\x02ok\x18\x02 \x01(\x08\x12\r\n\x05\x65rror\x18\x03 \x01(\t2\xc3\x02\n\tInference\x12\x46\n\tEmbedText\x12\x1b.inference.EmbedTextRequest\x1a\x1c.inference.EmbedTextResponse\x12O\n\x0cUpsertVector\x12\x1e.inference.UpsertVectorRequest\x1a\x1f.inference.UpsertVectorResponse\x12L\n\x0bQueryVector\x12\x1d.inference.QueryVectorRequest\x1a\x1e.inference.QueryVectorResponse\x12O\n\x0c\x44\x65leteVector\x12\x1e.inference.DeleteVectorRequest\x1a\x1f.inference.DeleteVectorResponseB8Z6github.com/notzree/uprank.ai/backend/inference-backendb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'inference_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z6github.com/notzree/uprank.ai/backend/inference-backend'
  _EMBEDTEXTREQUEST_METADATAENTRY._options = None
  _EMBEDTEXTREQUEST_METADATAENTRY._serialized_options = b'8\001'
  _VECTOR_METADATAENTRY._options = None
  _VECTOR_METADATAENTRY._serialized_options = b'8\001'
  _QUERYVECTORREQUEST_FILTERENTRY._options = None
  _QUERYVECTORREQUEST_FILTERENTRY._serialized_options = b'8\001'
  _MATCH_METADATAENTRY._options = None
  _MATCH_METADATAENTRY._serialized_options = b'8\001'
  _EMBEDTEXTREQUEST._serialized_start=31
  _EMBEDTEXTREQUEST._serialized_end=185
  _EMBEDTEXTREQUEST_METADATAENTRY._serialized_start=138
  _EMBEDTEXTREQUEST_METADATAENTRY._serialized_end=185
  _EMBEDTEXTRESPONSE._serialized_start=187
  _EMBEDTEXTRESPONSE._serialized_end=256
  _UPSERTVECTORREQUEST._serialized_start=258
  _UPSERTVECTORREQUEST._serialized_end=334
  _VECTOR._serialized_start=337
  _VECTOR._serialized_end=473
  _VECTOR_METADATAENTRY._serialized_start=138
  _VECTOR_METADATAENTRY._serialized_end=185
  _UPSERTVECTORRESPONSE._serialized_start=475
  _UPSERTVECTORRESPONSE._serialized_end=512
  _QUERYVECTORREQUEST._serialized_start=515
  _QUERYVECTORREQUEST._serialized_end=691
  _QUERYVECTORREQUEST_FILTERENTRY._serialized_start=646
  _QUERYVECTORREQUEST_FILTERENTRY._serialized_end=691
  _QUERYVECTORRESPONSE._serialized_start=693
  _QUERYVECTORRESPONSE._serialized_end=801
  _MATCH._serialized_start=804
  _MATCH._serialized_end=937
  _MATCH_METADATAENTRY._serialized_start=138
  _MATCH_METADATAENTRY._serialized_end=185
  _USAGE._serialized_start=939
  _USAGE._serialized_end=966
  _DELETEVECTORREQUEST._serialized_start=968
  _DELETEVECTORREQUEST._serialized_end=1021
  _DELETEVECTORRESPONSE._serialized_start=1023
  _DELETEVECTORRESPONSE._serialized_end=1085
  _INFERENCE._serialized_start=1088
  _INFERENCE._serialized_end=1411
# @@protoc_insertion_point(module_scope)
