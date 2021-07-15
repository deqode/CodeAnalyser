# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/plugin/common.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/plugin/common.proto',
  package='proto',
  syntax='proto3',
  serialized_options=b'Z\036code-analyser/protos/pb/plugin',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1aprotos/plugin/common.proto\x12\x05proto\"\x0e\n\x0cServiceEmpty\"H\n\x13ServiceOutputString\x12\r\n\x05value\x18\x01 \x01(\t\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x13.proto.ServiceError\"#\n\x12ServiceInputString\x12\r\n\x05value\x18\x01 \x01(\t\"E\n\x10ServiceOutputInt\x12\r\n\x05value\x18\x01 \x01(\x03\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x13.proto.ServiceError\"8\n\x12ServiceEmptyOutput\x12\"\n\x05\x65rror\x18\x01 \x01(\x0b\x32\x13.proto.ServiceError\"G\n\x12ServiceOutputFloat\x12\r\n\x05value\x18\x01 \x01(\x02\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x13.proto.ServiceError\"F\n\x11ServiceOutputBool\x12\r\n\x05value\x18\x01 \x01(\x08\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x13.proto.ServiceError\"4\n\x0cServiceInput\x12\x16\n\x0eruntimeVersion\x18\x01 \x01(\t\x12\x0c\n\x04root\x18\x02 \x01(\t\"\xa3\x01\n\x16ServiceOutputStringMap\x12\x37\n\x05value\x18\x01 \x03(\x0b\x32(.proto.ServiceOutputStringMap.ValueEntry\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x13.proto.ServiceError\x1a,\n\nValueEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"x\n\x0cServiceError\x12\x0f\n\x07message\x18\x01 \x01(\t\x12\x30\n\terrorType\x18\x02 \x01(\x0e\x32\x1d.proto.ServiceError.ErrorCode\x12\r\n\x05\x63\x61use\x18\x03 \x01(\t\"\x16\n\tErrorCode\x12\t\n\x05\x45rror\x10\x00\x42 Z\x1e\x63ode-analyser/protos/pb/pluginb\x06proto3'
)



_SERVICEERROR_ERRORCODE = _descriptor.EnumDescriptor(
  name='ErrorCode',
  full_name='proto.ServiceError.ErrorCode',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Error', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=756,
  serialized_end=778,
)
_sym_db.RegisterEnumDescriptor(_SERVICEERROR_ERRORCODE)


_SERVICEEMPTY = _descriptor.Descriptor(
  name='ServiceEmpty',
  full_name='proto.ServiceEmpty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=37,
  serialized_end=51,
)


_SERVICEOUTPUTSTRING = _descriptor.Descriptor(
  name='ServiceOutputString',
  full_name='proto.ServiceOutputString',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='proto.ServiceOutputString.value', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceOutputString.error', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=53,
  serialized_end=125,
)


_SERVICEINPUTSTRING = _descriptor.Descriptor(
  name='ServiceInputString',
  full_name='proto.ServiceInputString',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='proto.ServiceInputString.value', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=127,
  serialized_end=162,
)


_SERVICEOUTPUTINT = _descriptor.Descriptor(
  name='ServiceOutputInt',
  full_name='proto.ServiceOutputInt',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='proto.ServiceOutputInt.value', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceOutputInt.error', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=164,
  serialized_end=233,
)


_SERVICEEMPTYOUTPUT = _descriptor.Descriptor(
  name='ServiceEmptyOutput',
  full_name='proto.ServiceEmptyOutput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceEmptyOutput.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=235,
  serialized_end=291,
)


_SERVICEOUTPUTFLOAT = _descriptor.Descriptor(
  name='ServiceOutputFloat',
  full_name='proto.ServiceOutputFloat',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='proto.ServiceOutputFloat.value', index=0,
      number=1, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceOutputFloat.error', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=293,
  serialized_end=364,
)


_SERVICEOUTPUTBOOL = _descriptor.Descriptor(
  name='ServiceOutputBool',
  full_name='proto.ServiceOutputBool',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='proto.ServiceOutputBool.value', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceOutputBool.error', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=366,
  serialized_end=436,
)


_SERVICEINPUT = _descriptor.Descriptor(
  name='ServiceInput',
  full_name='proto.ServiceInput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='runtimeVersion', full_name='proto.ServiceInput.runtimeVersion', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='root', full_name='proto.ServiceInput.root', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=438,
  serialized_end=490,
)


_SERVICEOUTPUTSTRINGMAP_VALUEENTRY = _descriptor.Descriptor(
  name='ValueEntry',
  full_name='proto.ServiceOutputStringMap.ValueEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='proto.ServiceOutputStringMap.ValueEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='proto.ServiceOutputStringMap.ValueEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=612,
  serialized_end=656,
)

_SERVICEOUTPUTSTRINGMAP = _descriptor.Descriptor(
  name='ServiceOutputStringMap',
  full_name='proto.ServiceOutputStringMap',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='proto.ServiceOutputStringMap.value', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceOutputStringMap.error', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_SERVICEOUTPUTSTRINGMAP_VALUEENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=493,
  serialized_end=656,
)


_SERVICEERROR = _descriptor.Descriptor(
  name='ServiceError',
  full_name='proto.ServiceError',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='proto.ServiceError.message', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='errorType', full_name='proto.ServiceError.errorType', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cause', full_name='proto.ServiceError.cause', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _SERVICEERROR_ERRORCODE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=658,
  serialized_end=778,
)

_SERVICEOUTPUTSTRING.fields_by_name['error'].message_type = _SERVICEERROR
_SERVICEOUTPUTINT.fields_by_name['error'].message_type = _SERVICEERROR
_SERVICEEMPTYOUTPUT.fields_by_name['error'].message_type = _SERVICEERROR
_SERVICEOUTPUTFLOAT.fields_by_name['error'].message_type = _SERVICEERROR
_SERVICEOUTPUTBOOL.fields_by_name['error'].message_type = _SERVICEERROR
_SERVICEOUTPUTSTRINGMAP_VALUEENTRY.containing_type = _SERVICEOUTPUTSTRINGMAP
_SERVICEOUTPUTSTRINGMAP.fields_by_name['value'].message_type = _SERVICEOUTPUTSTRINGMAP_VALUEENTRY
_SERVICEOUTPUTSTRINGMAP.fields_by_name['error'].message_type = _SERVICEERROR
_SERVICEERROR.fields_by_name['errorType'].enum_type = _SERVICEERROR_ERRORCODE
_SERVICEERROR_ERRORCODE.containing_type = _SERVICEERROR
DESCRIPTOR.message_types_by_name['ServiceEmpty'] = _SERVICEEMPTY
DESCRIPTOR.message_types_by_name['ServiceOutputString'] = _SERVICEOUTPUTSTRING
DESCRIPTOR.message_types_by_name['ServiceInputString'] = _SERVICEINPUTSTRING
DESCRIPTOR.message_types_by_name['ServiceOutputInt'] = _SERVICEOUTPUTINT
DESCRIPTOR.message_types_by_name['ServiceEmptyOutput'] = _SERVICEEMPTYOUTPUT
DESCRIPTOR.message_types_by_name['ServiceOutputFloat'] = _SERVICEOUTPUTFLOAT
DESCRIPTOR.message_types_by_name['ServiceOutputBool'] = _SERVICEOUTPUTBOOL
DESCRIPTOR.message_types_by_name['ServiceInput'] = _SERVICEINPUT
DESCRIPTOR.message_types_by_name['ServiceOutputStringMap'] = _SERVICEOUTPUTSTRINGMAP
DESCRIPTOR.message_types_by_name['ServiceError'] = _SERVICEERROR
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ServiceEmpty = _reflection.GeneratedProtocolMessageType('ServiceEmpty', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEEMPTY,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceEmpty)
  })
_sym_db.RegisterMessage(ServiceEmpty)

ServiceOutputString = _reflection.GeneratedProtocolMessageType('ServiceOutputString', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEOUTPUTSTRING,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceOutputString)
  })
_sym_db.RegisterMessage(ServiceOutputString)

ServiceInputString = _reflection.GeneratedProtocolMessageType('ServiceInputString', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEINPUTSTRING,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceInputString)
  })
_sym_db.RegisterMessage(ServiceInputString)

ServiceOutputInt = _reflection.GeneratedProtocolMessageType('ServiceOutputInt', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEOUTPUTINT,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceOutputInt)
  })
_sym_db.RegisterMessage(ServiceOutputInt)

ServiceEmptyOutput = _reflection.GeneratedProtocolMessageType('ServiceEmptyOutput', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEEMPTYOUTPUT,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceEmptyOutput)
  })
_sym_db.RegisterMessage(ServiceEmptyOutput)

ServiceOutputFloat = _reflection.GeneratedProtocolMessageType('ServiceOutputFloat', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEOUTPUTFLOAT,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceOutputFloat)
  })
_sym_db.RegisterMessage(ServiceOutputFloat)

ServiceOutputBool = _reflection.GeneratedProtocolMessageType('ServiceOutputBool', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEOUTPUTBOOL,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceOutputBool)
  })
_sym_db.RegisterMessage(ServiceOutputBool)

ServiceInput = _reflection.GeneratedProtocolMessageType('ServiceInput', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEINPUT,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceInput)
  })
_sym_db.RegisterMessage(ServiceInput)

ServiceOutputStringMap = _reflection.GeneratedProtocolMessageType('ServiceOutputStringMap', (_message.Message,), {

  'ValueEntry' : _reflection.GeneratedProtocolMessageType('ValueEntry', (_message.Message,), {
    'DESCRIPTOR' : _SERVICEOUTPUTSTRINGMAP_VALUEENTRY,
    '__module__' : 'protos.plugin.common_pb2'
    # @@protoc_insertion_point(class_scope:proto.ServiceOutputStringMap.ValueEntry)
    })
  ,
  'DESCRIPTOR' : _SERVICEOUTPUTSTRINGMAP,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceOutputStringMap)
  })
_sym_db.RegisterMessage(ServiceOutputStringMap)
_sym_db.RegisterMessage(ServiceOutputStringMap.ValueEntry)

ServiceError = _reflection.GeneratedProtocolMessageType('ServiceError', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEERROR,
  '__module__' : 'protos.plugin.common_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceError)
  })
_sym_db.RegisterMessage(ServiceError)


DESCRIPTOR._options = None
_SERVICEOUTPUTSTRINGMAP_VALUEENTRY._options = None
# @@protoc_insertion_point(module_scope)