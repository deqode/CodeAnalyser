# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/plugin/detectRuntime.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from protos.plugin import common_pb2 as protos_dot_plugin_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/plugin/detectRuntime.proto',
  package='proto',
  syntax='proto3',
  serialized_options=b'Z\036code-analyser/protos/pb/plugin',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n!protos/plugin/detectRuntime.proto\x12\x05proto\x1a\x1aprotos/plugin/common.proto2^\n\x14\x44\x65tectRuntimeService\x12\x46\n\rDetectRuntime\x12\x19.proto.ServiceInputString\x1a\x1a.proto.ServiceOutputStringB Z\x1e\x63ode-analyser/protos/pb/pluginb\x06proto3'
  ,
  dependencies=[protos_dot_plugin_dot_common__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_DETECTRUNTIMESERVICE = _descriptor.ServiceDescriptor(
  name='DetectRuntimeService',
  full_name='proto.DetectRuntimeService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=72,
  serialized_end=166,
  methods=[
  _descriptor.MethodDescriptor(
    name='DetectRuntime',
    full_name='proto.DetectRuntimeService.DetectRuntime',
    index=0,
    containing_service=None,
    input_type=protos_dot_plugin_dot_common__pb2._SERVICEINPUTSTRING,
    output_type=protos_dot_plugin_dot_common__pb2._SERVICEOUTPUTSTRING,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_DETECTRUNTIMESERVICE)

DESCRIPTOR.services_by_name['DetectRuntimeService'] = _DETECTRUNTIMESERVICE

# @@protoc_insertion_point(module_scope)