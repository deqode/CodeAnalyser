# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/plugin/commands.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from protos.plugin import common_pb2 as protos_dot_plugin_dot_common__pb2
import importlib
protos_dot_output_dot_global_dot_buildCommands__pb2 = importlib.import_module('protos.output.global.buildCommands_pb2')
import importlib
protos_dot_output_dot_global_dot_seedCommands__pb2 = importlib.import_module('protos.output.global.seedCommands_pb2')
import importlib
protos_dot_output_dot_global_dot_migrationCommands__pb2 = importlib.import_module('protos.output.global.migrationCommands_pb2')
import importlib
protos_dot_output_dot_global_dot_adHocScripts__pb2 = importlib.import_module('protos.output.global.adHocScripts_pb2')
import importlib
protos_dot_output_dot_global_dot_startUpCommands__pb2 = importlib.import_module('protos.output.global.startUpCommands_pb2')


DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/plugin/commands.proto',
  package='proto',
  syntax='proto3',
  serialized_options=b'Z\036code-analyser/protos/pb/plugin',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1cprotos/plugin/commands.proto\x12\x05proto\x1a\x1aprotos/plugin/common.proto\x1a(protos/output/global/buildCommands.proto\x1a\'protos/output/global/seedCommands.proto\x1a,protos/output/global/migrationCommands.proto\x1a\'protos/output/global/adHocScripts.proto\x1a*protos/output/global/startUpCommands.proto\"6\n\x14ServiceCommandsInput\x12\x0c\n\x04root\x18\x01 \x01(\t\x12\x10\n\x08language\x18\x02 \x01(\t\"\x88\x01\n ServiceOutputDetectBuildCommands\x12\"\n\x05\x65rror\x18\x01 \x01(\x0b\x32\x13.proto.ServiceError\x12@\n\rbuildCommands\x18\x02 \x01(\x0b\x32).protos.output.global.BuildCommandsOutput\"\x88\x01\n\x1cServiceOutputStartUpCommands\x12\"\n\x05\x65rror\x18\x01 \x01(\x0b\x32\x13.proto.ServiceError\x12\x44\n\x0fstartUpCommands\x18\x02 \x01(\x0b\x32+.protos.output.global.StartUpCommandsOutput\"\x7f\n\x19ServiceOutputSeedCommands\x12\"\n\x05\x65rror\x18\x01 \x01(\x0b\x32\x13.proto.ServiceError\x12>\n\x0cseedCommands\x18\x02 \x01(\x0b\x32(.protos.output.global.SeedCommandsOutput\"\x88\x01\n\x18ServiceMigrationCommands\x12\"\n\x05\x65rror\x18\x01 \x01(\x0b\x32\x13.proto.ServiceError\x12H\n\x11migrationCommands\x18\x02 \x01(\x0b\x32-.protos.output.global.MigrationCommandsOutput\"y\n\x13ServiceAdHocScripts\x12\"\n\x05\x65rror\x18\x01 \x01(\x0b\x32\x13.proto.ServiceError\x12>\n\x0c\x61\x64HocScripts\x18\x02 \x01(\x0b\x32(.protos.output.global.AdHocScriptsOutput2\xc6\x03\n\x0f\x43ommandsService\x12[\n\x13\x44\x65tectBuildCommands\x12\x1b.proto.ServiceCommandsInput\x1a\'.proto.ServiceOutputDetectBuildCommands\x12Y\n\x15\x44\x65tectStartUpCommands\x12\x1b.proto.ServiceCommandsInput\x1a#.proto.ServiceOutputStartUpCommands\x12S\n\x12\x44\x65tectSeedCommands\x12\x1b.proto.ServiceCommandsInput\x1a .proto.ServiceOutputSeedCommands\x12W\n\x17\x44\x65tectMigrationCommands\x12\x1b.proto.ServiceCommandsInput\x1a\x1f.proto.ServiceMigrationCommands\x12M\n\x12\x44\x65tectAdHocScripts\x12\x1b.proto.ServiceCommandsInput\x1a\x1a.proto.ServiceAdHocScriptsB Z\x1e\x63ode-analyser/protos/pb/pluginb\x06proto3'
  ,
  dependencies=[protos_dot_plugin_dot_common__pb2.DESCRIPTOR,protos_dot_output_dot_global_dot_buildCommands__pb2.DESCRIPTOR,protos_dot_output_dot_global_dot_seedCommands__pb2.DESCRIPTOR,protos_dot_output_dot_global_dot_migrationCommands__pb2.DESCRIPTOR,protos_dot_output_dot_global_dot_adHocScripts__pb2.DESCRIPTOR,protos_dot_output_dot_global_dot_startUpCommands__pb2.DESCRIPTOR,])




_SERVICECOMMANDSINPUT = _descriptor.Descriptor(
  name='ServiceCommandsInput',
  full_name='proto.ServiceCommandsInput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='root', full_name='proto.ServiceCommandsInput.root', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='language', full_name='proto.ServiceCommandsInput.language', index=1,
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
  serialized_start=281,
  serialized_end=335,
)


_SERVICEOUTPUTDETECTBUILDCOMMANDS = _descriptor.Descriptor(
  name='ServiceOutputDetectBuildCommands',
  full_name='proto.ServiceOutputDetectBuildCommands',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceOutputDetectBuildCommands.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='buildCommands', full_name='proto.ServiceOutputDetectBuildCommands.buildCommands', index=1,
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
  serialized_start=338,
  serialized_end=474,
)


_SERVICEOUTPUTSTARTUPCOMMANDS = _descriptor.Descriptor(
  name='ServiceOutputStartUpCommands',
  full_name='proto.ServiceOutputStartUpCommands',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceOutputStartUpCommands.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='startUpCommands', full_name='proto.ServiceOutputStartUpCommands.startUpCommands', index=1,
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
  serialized_start=477,
  serialized_end=613,
)


_SERVICEOUTPUTSEEDCOMMANDS = _descriptor.Descriptor(
  name='ServiceOutputSeedCommands',
  full_name='proto.ServiceOutputSeedCommands',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceOutputSeedCommands.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='seedCommands', full_name='proto.ServiceOutputSeedCommands.seedCommands', index=1,
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
  serialized_start=615,
  serialized_end=742,
)


_SERVICEMIGRATIONCOMMANDS = _descriptor.Descriptor(
  name='ServiceMigrationCommands',
  full_name='proto.ServiceMigrationCommands',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceMigrationCommands.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='migrationCommands', full_name='proto.ServiceMigrationCommands.migrationCommands', index=1,
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
  serialized_start=745,
  serialized_end=881,
)


_SERVICEADHOCSCRIPTS = _descriptor.Descriptor(
  name='ServiceAdHocScripts',
  full_name='proto.ServiceAdHocScripts',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceAdHocScripts.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='adHocScripts', full_name='proto.ServiceAdHocScripts.adHocScripts', index=1,
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
  serialized_start=883,
  serialized_end=1004,
)

_SERVICEOUTPUTDETECTBUILDCOMMANDS.fields_by_name['error'].message_type = protos_dot_plugin_dot_common__pb2._SERVICEERROR
_SERVICEOUTPUTDETECTBUILDCOMMANDS.fields_by_name['buildCommands'].message_type = protos_dot_output_dot_global_dot_buildCommands__pb2._BUILDCOMMANDSOUTPUT
_SERVICEOUTPUTSTARTUPCOMMANDS.fields_by_name['error'].message_type = protos_dot_plugin_dot_common__pb2._SERVICEERROR
_SERVICEOUTPUTSTARTUPCOMMANDS.fields_by_name['startUpCommands'].message_type = protos_dot_output_dot_global_dot_startUpCommands__pb2._STARTUPCOMMANDSOUTPUT
_SERVICEOUTPUTSEEDCOMMANDS.fields_by_name['error'].message_type = protos_dot_plugin_dot_common__pb2._SERVICEERROR
_SERVICEOUTPUTSEEDCOMMANDS.fields_by_name['seedCommands'].message_type = protos_dot_output_dot_global_dot_seedCommands__pb2._SEEDCOMMANDSOUTPUT
_SERVICEMIGRATIONCOMMANDS.fields_by_name['error'].message_type = protos_dot_plugin_dot_common__pb2._SERVICEERROR
_SERVICEMIGRATIONCOMMANDS.fields_by_name['migrationCommands'].message_type = protos_dot_output_dot_global_dot_migrationCommands__pb2._MIGRATIONCOMMANDSOUTPUT
_SERVICEADHOCSCRIPTS.fields_by_name['error'].message_type = protos_dot_plugin_dot_common__pb2._SERVICEERROR
_SERVICEADHOCSCRIPTS.fields_by_name['adHocScripts'].message_type = protos_dot_output_dot_global_dot_adHocScripts__pb2._ADHOCSCRIPTSOUTPUT
DESCRIPTOR.message_types_by_name['ServiceCommandsInput'] = _SERVICECOMMANDSINPUT
DESCRIPTOR.message_types_by_name['ServiceOutputDetectBuildCommands'] = _SERVICEOUTPUTDETECTBUILDCOMMANDS
DESCRIPTOR.message_types_by_name['ServiceOutputStartUpCommands'] = _SERVICEOUTPUTSTARTUPCOMMANDS
DESCRIPTOR.message_types_by_name['ServiceOutputSeedCommands'] = _SERVICEOUTPUTSEEDCOMMANDS
DESCRIPTOR.message_types_by_name['ServiceMigrationCommands'] = _SERVICEMIGRATIONCOMMANDS
DESCRIPTOR.message_types_by_name['ServiceAdHocScripts'] = _SERVICEADHOCSCRIPTS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ServiceCommandsInput = _reflection.GeneratedProtocolMessageType('ServiceCommandsInput', (_message.Message,), {
  'DESCRIPTOR' : _SERVICECOMMANDSINPUT,
  '__module__' : 'protos.plugin.commands_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceCommandsInput)
  })
_sym_db.RegisterMessage(ServiceCommandsInput)

ServiceOutputDetectBuildCommands = _reflection.GeneratedProtocolMessageType('ServiceOutputDetectBuildCommands', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEOUTPUTDETECTBUILDCOMMANDS,
  '__module__' : 'protos.plugin.commands_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceOutputDetectBuildCommands)
  })
_sym_db.RegisterMessage(ServiceOutputDetectBuildCommands)

ServiceOutputStartUpCommands = _reflection.GeneratedProtocolMessageType('ServiceOutputStartUpCommands', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEOUTPUTSTARTUPCOMMANDS,
  '__module__' : 'protos.plugin.commands_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceOutputStartUpCommands)
  })
_sym_db.RegisterMessage(ServiceOutputStartUpCommands)

ServiceOutputSeedCommands = _reflection.GeneratedProtocolMessageType('ServiceOutputSeedCommands', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEOUTPUTSEEDCOMMANDS,
  '__module__' : 'protos.plugin.commands_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceOutputSeedCommands)
  })
_sym_db.RegisterMessage(ServiceOutputSeedCommands)

ServiceMigrationCommands = _reflection.GeneratedProtocolMessageType('ServiceMigrationCommands', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEMIGRATIONCOMMANDS,
  '__module__' : 'protos.plugin.commands_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceMigrationCommands)
  })
_sym_db.RegisterMessage(ServiceMigrationCommands)

ServiceAdHocScripts = _reflection.GeneratedProtocolMessageType('ServiceAdHocScripts', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEADHOCSCRIPTS,
  '__module__' : 'protos.plugin.commands_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceAdHocScripts)
  })
_sym_db.RegisterMessage(ServiceAdHocScripts)


DESCRIPTOR._options = None

_COMMANDSSERVICE = _descriptor.ServiceDescriptor(
  name='CommandsService',
  full_name='proto.CommandsService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1007,
  serialized_end=1461,
  methods=[
  _descriptor.MethodDescriptor(
    name='DetectBuildCommands',
    full_name='proto.CommandsService.DetectBuildCommands',
    index=0,
    containing_service=None,
    input_type=_SERVICECOMMANDSINPUT,
    output_type=_SERVICEOUTPUTDETECTBUILDCOMMANDS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DetectStartUpCommands',
    full_name='proto.CommandsService.DetectStartUpCommands',
    index=1,
    containing_service=None,
    input_type=_SERVICECOMMANDSINPUT,
    output_type=_SERVICEOUTPUTSTARTUPCOMMANDS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DetectSeedCommands',
    full_name='proto.CommandsService.DetectSeedCommands',
    index=2,
    containing_service=None,
    input_type=_SERVICECOMMANDSINPUT,
    output_type=_SERVICEOUTPUTSEEDCOMMANDS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DetectMigrationCommands',
    full_name='proto.CommandsService.DetectMigrationCommands',
    index=3,
    containing_service=None,
    input_type=_SERVICECOMMANDSINPUT,
    output_type=_SERVICEMIGRATIONCOMMANDS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DetectAdHocScripts',
    full_name='proto.CommandsService.DetectAdHocScripts',
    index=4,
    containing_service=None,
    input_type=_SERVICECOMMANDSINPUT,
    output_type=_SERVICEADHOCSCRIPTS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_COMMANDSSERVICE)

DESCRIPTOR.services_by_name['CommandsService'] = _COMMANDSSERVICE

# @@protoc_insertion_point(module_scope)
