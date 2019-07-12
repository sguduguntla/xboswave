# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: energise.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import c37_pb2 as c37__pb2
from . import nullabletypes_pb2 as nullabletypes__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='energise.proto',
  package='xbospb',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0e\x65nergise.proto\x12\x06xbospb\x1a\tc37.proto\x1a\x13nullabletypes.proto\"\x7f\n\x0f\x45nergiseMessage\x12\x1a\n\x04SPBC\x18\x01 \x01(\x0b\x32\x0c.xbospb.SPBC\x12&\n\nLPBCStatus\x18\x02 \x01(\x0b\x32\x12.xbospb.LPBCStatus\x12(\n\x0bLPBCCommand\x18\x03 \x01(\x0b\x32\x13.xbospb.LPBCCommand\"\x1c\n\rEnergiseError\x12\x0b\n\x03msg\x18\x01 \x01(\t\"}\n\x14\x45nergisePhasorTarget\x12\x0e\n\x06nodeID\x18\x01 \x01(\t\x12\x13\n\x0b\x63hannelName\x18\x02 \x01(\t\x12\r\n\x05\x61ngle\x18\x03 \x01(\x01\x12\x11\n\tmagnitude\x18\x04 \x01(\x01\x12\x1e\n\x06kvbase\x18\x05 \x01(\x0b\x32\x0e.xbospb.Double\"p\n\x04SPBC\x12\x0c\n\x04time\x18\x01 \x01(\x03\x12\x34\n\x0ephasor_targets\x18\x02 \x03(\x0b\x32\x1c.xbospb.EnergisePhasorTarget\x12$\n\x05\x65rror\x18\x03 \x01(\x0b\x32\x15.xbospb.EnergiseError\"\xe3\x01\n\nLPBCStatus\x12\x0c\n\x04time\x18\x01 \x01(\x03\x12$\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x15.xbospb.EnergiseError\x12%\n\rphasor_errors\x18\x03 \x01(\x0b\x32\x0e.xbospb.Phasor\x12\x13\n\x0bp_saturated\x18\x04 \x01(\x08\x12\x13\n\x0bq_saturated\x18\x05 \x01(\x08\x12\x12\n\ndo_control\x18\x06 \x01(\x08\x12\x1d\n\x05p_max\x18\x07 \x01(\x0b\x32\x0e.xbospb.Double\x12\x1d\n\x05q_max\x18\x08 \x01(\x0b\x32\x0e.xbospb.Double\"V\n\x0bLPBCCommand\x12\x0c\n\x04time\x18\x01 \x01(\x03\x12%\n\rphasor_target\x18\x02 \x01(\x0b\x32\x0e.xbospb.Phasor\x12\x12\n\ndo_control\x18\x03 \x01(\x08\x62\x06proto3')
  ,
  dependencies=[c37__pb2.DESCRIPTOR,nullabletypes__pb2.DESCRIPTOR,])




_ENERGISEMESSAGE = _descriptor.Descriptor(
  name='EnergiseMessage',
  full_name='xbospb.EnergiseMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='SPBC', full_name='xbospb.EnergiseMessage.SPBC', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='LPBCStatus', full_name='xbospb.EnergiseMessage.LPBCStatus', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='LPBCCommand', full_name='xbospb.EnergiseMessage.LPBCCommand', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=58,
  serialized_end=185,
)


_ENERGISEERROR = _descriptor.Descriptor(
  name='EnergiseError',
  full_name='xbospb.EnergiseError',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='msg', full_name='xbospb.EnergiseError.msg', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=187,
  serialized_end=215,
)


_ENERGISEPHASORTARGET = _descriptor.Descriptor(
  name='EnergisePhasorTarget',
  full_name='xbospb.EnergisePhasorTarget',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='nodeID', full_name='xbospb.EnergisePhasorTarget.nodeID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='channelName', full_name='xbospb.EnergisePhasorTarget.channelName', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='angle', full_name='xbospb.EnergisePhasorTarget.angle', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='magnitude', full_name='xbospb.EnergisePhasorTarget.magnitude', index=3,
      number=4, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='kvbase', full_name='xbospb.EnergisePhasorTarget.kvbase', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=217,
  serialized_end=342,
)


_SPBC = _descriptor.Descriptor(
  name='SPBC',
  full_name='xbospb.SPBC',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='xbospb.SPBC.time', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phasor_targets', full_name='xbospb.SPBC.phasor_targets', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='xbospb.SPBC.error', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=344,
  serialized_end=456,
)


_LPBCSTATUS = _descriptor.Descriptor(
  name='LPBCStatus',
  full_name='xbospb.LPBCStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='xbospb.LPBCStatus.time', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='xbospb.LPBCStatus.error', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phasor_errors', full_name='xbospb.LPBCStatus.phasor_errors', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='p_saturated', full_name='xbospb.LPBCStatus.p_saturated', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='q_saturated', full_name='xbospb.LPBCStatus.q_saturated', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='do_control', full_name='xbospb.LPBCStatus.do_control', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='p_max', full_name='xbospb.LPBCStatus.p_max', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='q_max', full_name='xbospb.LPBCStatus.q_max', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=459,
  serialized_end=686,
)


_LPBCCOMMAND = _descriptor.Descriptor(
  name='LPBCCommand',
  full_name='xbospb.LPBCCommand',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='xbospb.LPBCCommand.time', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phasor_target', full_name='xbospb.LPBCCommand.phasor_target', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='do_control', full_name='xbospb.LPBCCommand.do_control', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=688,
  serialized_end=774,
)

_ENERGISEMESSAGE.fields_by_name['SPBC'].message_type = _SPBC
_ENERGISEMESSAGE.fields_by_name['LPBCStatus'].message_type = _LPBCSTATUS
_ENERGISEMESSAGE.fields_by_name['LPBCCommand'].message_type = _LPBCCOMMAND
_ENERGISEPHASORTARGET.fields_by_name['kvbase'].message_type = nullabletypes__pb2._DOUBLE
_SPBC.fields_by_name['phasor_targets'].message_type = _ENERGISEPHASORTARGET
_SPBC.fields_by_name['error'].message_type = _ENERGISEERROR
_LPBCSTATUS.fields_by_name['error'].message_type = _ENERGISEERROR
_LPBCSTATUS.fields_by_name['phasor_errors'].message_type = c37__pb2._PHASOR
_LPBCSTATUS.fields_by_name['p_max'].message_type = nullabletypes__pb2._DOUBLE
_LPBCSTATUS.fields_by_name['q_max'].message_type = nullabletypes__pb2._DOUBLE
_LPBCCOMMAND.fields_by_name['phasor_target'].message_type = c37__pb2._PHASOR
DESCRIPTOR.message_types_by_name['EnergiseMessage'] = _ENERGISEMESSAGE
DESCRIPTOR.message_types_by_name['EnergiseError'] = _ENERGISEERROR
DESCRIPTOR.message_types_by_name['EnergisePhasorTarget'] = _ENERGISEPHASORTARGET
DESCRIPTOR.message_types_by_name['SPBC'] = _SPBC
DESCRIPTOR.message_types_by_name['LPBCStatus'] = _LPBCSTATUS
DESCRIPTOR.message_types_by_name['LPBCCommand'] = _LPBCCOMMAND
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

EnergiseMessage = _reflection.GeneratedProtocolMessageType('EnergiseMessage', (_message.Message,), dict(
  DESCRIPTOR = _ENERGISEMESSAGE,
  __module__ = 'energise_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.EnergiseMessage)
  ))
_sym_db.RegisterMessage(EnergiseMessage)

EnergiseError = _reflection.GeneratedProtocolMessageType('EnergiseError', (_message.Message,), dict(
  DESCRIPTOR = _ENERGISEERROR,
  __module__ = 'energise_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.EnergiseError)
  ))
_sym_db.RegisterMessage(EnergiseError)

EnergisePhasorTarget = _reflection.GeneratedProtocolMessageType('EnergisePhasorTarget', (_message.Message,), dict(
  DESCRIPTOR = _ENERGISEPHASORTARGET,
  __module__ = 'energise_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.EnergisePhasorTarget)
  ))
_sym_db.RegisterMessage(EnergisePhasorTarget)

SPBC = _reflection.GeneratedProtocolMessageType('SPBC', (_message.Message,), dict(
  DESCRIPTOR = _SPBC,
  __module__ = 'energise_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.SPBC)
  ))
_sym_db.RegisterMessage(SPBC)

LPBCStatus = _reflection.GeneratedProtocolMessageType('LPBCStatus', (_message.Message,), dict(
  DESCRIPTOR = _LPBCSTATUS,
  __module__ = 'energise_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.LPBCStatus)
  ))
_sym_db.RegisterMessage(LPBCStatus)

LPBCCommand = _reflection.GeneratedProtocolMessageType('LPBCCommand', (_message.Message,), dict(
  DESCRIPTOR = _LPBCCOMMAND,
  __module__ = 'energise_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.LPBCCommand)
  ))
_sym_db.RegisterMessage(LPBCCommand)


# @@protoc_insertion_point(module_scope)
