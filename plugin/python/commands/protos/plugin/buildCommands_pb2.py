# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/output/global/buildCommands.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import importlib
protos_dot_output_dot_global_dot_command__pb2 = importlib.import_module('protos.output.global.command_pb2')


DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/output/global/buildCommands.proto',
  package='protos.output.global',
  syntax='proto3',
  serialized_options=b'Z%code-analyser/protos/pb/output/global',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n(protos/output/global/buildCommands.proto\x12\x14protos.output.global\x1a\"protos/output/global/command.proto\"Y\n\x13\x42uildCommandsOutput\x12\x0c\n\x04used\x18\x01 \x01(\x08\x12\x34\n\rbuildCommands\x18\x02 \x03(\x0b\x32\x1d.protos.output.global.CommandB\'Z%code-analyser/protos/pb/output/globalb\x06proto3'
  ,
  dependencies=[protos_dot_output_dot_global_dot_command__pb2.DESCRIPTOR,])




_BUILDCOMMANDSOUTPUT = _descriptor.Descriptor(
  name='BuildCommandsOutput',
  full_name='protos.output.global.BuildCommandsOutput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='used', full_name='protos.output.global.BuildCommandsOutput.used', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='buildCommands', full_name='protos.output.global.BuildCommandsOutput.buildCommands', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=102,
  serialized_end=191,
)

_BUILDCOMMANDSOUTPUT.fields_by_name['buildCommands'].message_type = protos_dot_output_dot_global_dot_command__pb2._COMMAND
DESCRIPTOR.message_types_by_name['BuildCommandsOutput'] = _BUILDCOMMANDSOUTPUT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

BuildCommandsOutput = _reflection.GeneratedProtocolMessageType('BuildCommandsOutput', (_message.Message,), {
  'DESCRIPTOR' : _BUILDCOMMANDSOUTPUT,
  '__module__' : 'protos.output.global.buildCommands_pb2'
  # @@protoc_insertion_point(class_scope:protos.output.global.BuildCommandsOutput)
  })
_sym_db.RegisterMessage(BuildCommandsOutput)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
