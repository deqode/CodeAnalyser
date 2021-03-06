# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/plugin/dependencies.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from protos.plugin import common_pb2 as protos_dot_plugin_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/plugin/dependencies.proto',
  package='proto',
  syntax='proto3',
  serialized_options=b'Z\036code-analyser/protos/pb/plugin',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n protos/plugin/dependencies.proto\x12\x05proto\x1a\x1aprotos/plugin/common.proto2\\\n\x13\x44\x65pendenciesService\x12\x45\n\x0fGetDependencies\x12\x13.proto.ServiceInput\x1a\x1d.proto.ServiceOutputStringMapB Z\x1e\x63ode-analyser/protos/pb/pluginb\x06proto3'
  ,
  dependencies=[protos_dot_plugin_dot_common__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_DEPENDENCIESSERVICE = _descriptor.ServiceDescriptor(
  name='DependenciesService',
  full_name='proto.DependenciesService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=71,
  serialized_end=163,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetDependencies',
    full_name='proto.DependenciesService.GetDependencies',
    index=0,
    containing_service=None,
    input_type=protos_dot_plugin_dot_common__pb2._SERVICEINPUT,
    output_type=protos_dot_plugin_dot_common__pb2._SERVICEOUTPUTSTRINGMAP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_DEPENDENCIESSERVICE)

DESCRIPTOR.services_by_name['DependenciesService'] = _DEPENDENCIESSERVICE

# @@protoc_insertion_point(module_scope)
