# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/plugin/predetectcommand.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from protos.plugin import common_pb2 as protos_dot_plugin_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/plugin/predetectcommand.proto',
  package='proto',
  syntax='proto3',
  serialized_options=b'Z\036code-analyser/protos/pb/plugin',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n$protos/plugin/predetectcommand.proto\x12\x05proto\x1a\x1aprotos/plugin/common.proto2R\n\x10PreDetectCommand\x12>\n\x0cRunPreDetect\x12\x13.proto.ServiceInput\x1a\x19.proto.ServiceEmptyOutputB Z\x1e\x63ode-analyser/protos/pb/pluginb\x06proto3'
  ,
  dependencies=[protos_dot_plugin_dot_common__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_PREDETECTCOMMAND = _descriptor.ServiceDescriptor(
  name='PreDetectCommand',
  full_name='proto.PreDetectCommand',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=75,
  serialized_end=157,
  methods=[
  _descriptor.MethodDescriptor(
    name='RunPreDetect',
    full_name='proto.PreDetectCommand.RunPreDetect',
    index=0,
    containing_service=None,
    input_type=protos_dot_plugin_dot_common__pb2._SERVICEINPUT,
    output_type=protos_dot_plugin_dot_common__pb2._SERVICEEMPTYOUTPUT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_PREDETECTCOMMAND)

DESCRIPTOR.services_by_name['PreDetectCommand'] = _PREDETECTCOMMAND

# @@protoc_insertion_point(module_scope)
