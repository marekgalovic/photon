# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: models.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import core_pb2 as core__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='models.proto',
  package='photon',
  syntax='proto3',
  serialized_pb=_b('\n\x0cmodels.proto\x12\x06photon\x1a\ncore.proto\"\xcc\x01\n\x05Model\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0brunner_type\x18\x03 \x01(\t\x12\x10\n\x08replicas\x18\x04 \x01(\x05\x12&\n\x08\x66\x65\x61tures\x18\x05 \x03(\x0b\x32\x14.photon.ModelFeature\x12\x32\n\x14precomputed_features\x18\x06 \x03(\x0b\x32\x14.photon.ModelFeature\x12\x12\n\ncreated_at\x18\x07 \x01(\x05\x12\x12\n\nupdated_at\x18\x08 \x01(\x05\"\x88\x01\n\x0cModelVersion\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x10\n\x08model_id\x18\x02 \x01(\x03\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x11\n\tfile_name\x18\x04 \x01(\t\x12\x12\n\nis_primary\x18\x05 \x01(\x08\x12\x11\n\tis_shadow\x18\x06 \x01(\x08\x12\x12\n\ncreated_at\x18\x07 \x01(\x05\".\n\x0cModelFeature\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x10\n\x08required\x18\x02 \x01(\x08\"@\n\x16PrecomputedFeaturesSet\x12&\n\x08\x66\x65\x61tures\x18\x01 \x03(\x0b\x32\x14.photon.ModelFeature\"\x1e\n\x10\x46indModelRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"\xa0\x02\n\x12\x43reateModelRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0brunner_type\x18\x02 \x01(\t\x12\x10\n\x08replicas\x18\x03 \x01(\x05\x12&\n\x08\x66\x65\x61tures\x18\x04 \x03(\x0b\x32\x14.photon.ModelFeature\x12Q\n\x14precomputed_features\x18\x05 \x03(\x0b\x32\x33.photon.CreateModelRequest.PrecomputedFeaturesEntry\x1aZ\n\x18PrecomputedFeaturesEntry\x12\x0b\n\x03key\x18\x01 \x01(\x03\x12-\n\x05value\x18\x02 \x01(\x0b\x32\x1e.photon.PrecomputedFeaturesSet:\x02\x38\x01\"!\n\x13\x43reateModelResponse\x12\n\n\x02id\x18\x01 \x01(\x03\" \n\x12\x44\x65leteModelRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"\'\n\x13ListVersionsRequest\x12\x10\n\x08model_id\x18\x01 \x01(\x03\" \n\x12\x46indVersionRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"8\n\x18SetPrimaryVersionRequest\x12\x10\n\x08model_id\x18\x01 \x01(\x03\x12\n\n\x02id\x18\x02 \x01(\x03\"X\n\x14\x43reateVersionRequest\x12\'\n\x07version\x18\x01 \x01(\x0b\x32\x14.photon.ModelVersionH\x00\x12\x0e\n\x04\x64\x61ta\x18\x02 \x01(\x0cH\x00\x42\x07\n\x05value\"#\n\x15\x43reateVersionResponse\x12\n\n\x02id\x18\x01 \x01(\x03\"\"\n\x14\x44\x65leteVersionRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xdb\x04\n\rModelsService\x12-\n\x04List\x12\x14.photon.EmptyRequest\x1a\r.photon.Model0\x01\x12/\n\x04\x46ind\x12\x18.photon.FindModelRequest\x1a\r.photon.Model\x12\x41\n\x06\x43reate\x12\x1a.photon.CreateModelRequest\x1a\x1b.photon.CreateModelResponse\x12;\n\x06\x44\x65lete\x12\x1a.photon.DeleteModelRequest\x1a\x15.photon.EmptyResponse\x12\x43\n\x0cListVersions\x12\x1b.photon.ListVersionsRequest\x1a\x14.photon.ModelVersion0\x01\x12?\n\x0b\x46indVersion\x12\x1a.photon.FindVersionRequest\x1a\x14.photon.ModelVersion\x12L\n\x11SetPrimaryVersion\x12 .photon.SetPrimaryVersionRequest\x1a\x15.photon.EmptyResponse\x12P\n\rCreateVersion\x12\x1c.photon.CreateVersionRequest\x1a\x1d.photon.CreateVersionResponse(\x01\x30\x01\x12\x44\n\rDeleteVersion\x12\x1c.photon.DeleteVersionRequest\x1a\x15.photon.EmptyResponseb\x06proto3')
  ,
  dependencies=[core__pb2.DESCRIPTOR,])




_MODEL = _descriptor.Descriptor(
  name='Model',
  full_name='photon.Model',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='photon.Model.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='photon.Model.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='runner_type', full_name='photon.Model.runner_type', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='replicas', full_name='photon.Model.replicas', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='features', full_name='photon.Model.features', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='precomputed_features', full_name='photon.Model.precomputed_features', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='photon.Model.created_at', index=6,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='photon.Model.updated_at', index=7,
      number=8, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=37,
  serialized_end=241,
)


_MODELVERSION = _descriptor.Descriptor(
  name='ModelVersion',
  full_name='photon.ModelVersion',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='photon.ModelVersion.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='model_id', full_name='photon.ModelVersion.model_id', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='photon.ModelVersion.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='file_name', full_name='photon.ModelVersion.file_name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='is_primary', full_name='photon.ModelVersion.is_primary', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='is_shadow', full_name='photon.ModelVersion.is_shadow', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='photon.ModelVersion.created_at', index=6,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=244,
  serialized_end=380,
)


_MODELFEATURE = _descriptor.Descriptor(
  name='ModelFeature',
  full_name='photon.ModelFeature',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='photon.ModelFeature.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='required', full_name='photon.ModelFeature.required', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=382,
  serialized_end=428,
)


_PRECOMPUTEDFEATURESSET = _descriptor.Descriptor(
  name='PrecomputedFeaturesSet',
  full_name='photon.PrecomputedFeaturesSet',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='features', full_name='photon.PrecomputedFeaturesSet.features', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=430,
  serialized_end=494,
)


_FINDMODELREQUEST = _descriptor.Descriptor(
  name='FindModelRequest',
  full_name='photon.FindModelRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='photon.FindModelRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=496,
  serialized_end=526,
)


_CREATEMODELREQUEST_PRECOMPUTEDFEATURESENTRY = _descriptor.Descriptor(
  name='PrecomputedFeaturesEntry',
  full_name='photon.CreateModelRequest.PrecomputedFeaturesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='photon.CreateModelRequest.PrecomputedFeaturesEntry.key', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='photon.CreateModelRequest.PrecomputedFeaturesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=727,
  serialized_end=817,
)

_CREATEMODELREQUEST = _descriptor.Descriptor(
  name='CreateModelRequest',
  full_name='photon.CreateModelRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='photon.CreateModelRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='runner_type', full_name='photon.CreateModelRequest.runner_type', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='replicas', full_name='photon.CreateModelRequest.replicas', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='features', full_name='photon.CreateModelRequest.features', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='precomputed_features', full_name='photon.CreateModelRequest.precomputed_features', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_CREATEMODELREQUEST_PRECOMPUTEDFEATURESENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=529,
  serialized_end=817,
)


_CREATEMODELRESPONSE = _descriptor.Descriptor(
  name='CreateModelResponse',
  full_name='photon.CreateModelResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='photon.CreateModelResponse.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=819,
  serialized_end=852,
)


_DELETEMODELREQUEST = _descriptor.Descriptor(
  name='DeleteModelRequest',
  full_name='photon.DeleteModelRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='photon.DeleteModelRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=854,
  serialized_end=886,
)


_LISTVERSIONSREQUEST = _descriptor.Descriptor(
  name='ListVersionsRequest',
  full_name='photon.ListVersionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='model_id', full_name='photon.ListVersionsRequest.model_id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=888,
  serialized_end=927,
)


_FINDVERSIONREQUEST = _descriptor.Descriptor(
  name='FindVersionRequest',
  full_name='photon.FindVersionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='photon.FindVersionRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=929,
  serialized_end=961,
)


_SETPRIMARYVERSIONREQUEST = _descriptor.Descriptor(
  name='SetPrimaryVersionRequest',
  full_name='photon.SetPrimaryVersionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='model_id', full_name='photon.SetPrimaryVersionRequest.model_id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='photon.SetPrimaryVersionRequest.id', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=963,
  serialized_end=1019,
)


_CREATEVERSIONREQUEST = _descriptor.Descriptor(
  name='CreateVersionRequest',
  full_name='photon.CreateVersionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='version', full_name='photon.CreateVersionRequest.version', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='data', full_name='photon.CreateVersionRequest.data', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='value', full_name='photon.CreateVersionRequest.value',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=1021,
  serialized_end=1109,
)


_CREATEVERSIONRESPONSE = _descriptor.Descriptor(
  name='CreateVersionResponse',
  full_name='photon.CreateVersionResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='photon.CreateVersionResponse.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1111,
  serialized_end=1146,
)


_DELETEVERSIONREQUEST = _descriptor.Descriptor(
  name='DeleteVersionRequest',
  full_name='photon.DeleteVersionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='photon.DeleteVersionRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1148,
  serialized_end=1182,
)

_MODEL.fields_by_name['features'].message_type = _MODELFEATURE
_MODEL.fields_by_name['precomputed_features'].message_type = _MODELFEATURE
_PRECOMPUTEDFEATURESSET.fields_by_name['features'].message_type = _MODELFEATURE
_CREATEMODELREQUEST_PRECOMPUTEDFEATURESENTRY.fields_by_name['value'].message_type = _PRECOMPUTEDFEATURESSET
_CREATEMODELREQUEST_PRECOMPUTEDFEATURESENTRY.containing_type = _CREATEMODELREQUEST
_CREATEMODELREQUEST.fields_by_name['features'].message_type = _MODELFEATURE
_CREATEMODELREQUEST.fields_by_name['precomputed_features'].message_type = _CREATEMODELREQUEST_PRECOMPUTEDFEATURESENTRY
_CREATEVERSIONREQUEST.fields_by_name['version'].message_type = _MODELVERSION
_CREATEVERSIONREQUEST.oneofs_by_name['value'].fields.append(
  _CREATEVERSIONREQUEST.fields_by_name['version'])
_CREATEVERSIONREQUEST.fields_by_name['version'].containing_oneof = _CREATEVERSIONREQUEST.oneofs_by_name['value']
_CREATEVERSIONREQUEST.oneofs_by_name['value'].fields.append(
  _CREATEVERSIONREQUEST.fields_by_name['data'])
_CREATEVERSIONREQUEST.fields_by_name['data'].containing_oneof = _CREATEVERSIONREQUEST.oneofs_by_name['value']
DESCRIPTOR.message_types_by_name['Model'] = _MODEL
DESCRIPTOR.message_types_by_name['ModelVersion'] = _MODELVERSION
DESCRIPTOR.message_types_by_name['ModelFeature'] = _MODELFEATURE
DESCRIPTOR.message_types_by_name['PrecomputedFeaturesSet'] = _PRECOMPUTEDFEATURESSET
DESCRIPTOR.message_types_by_name['FindModelRequest'] = _FINDMODELREQUEST
DESCRIPTOR.message_types_by_name['CreateModelRequest'] = _CREATEMODELREQUEST
DESCRIPTOR.message_types_by_name['CreateModelResponse'] = _CREATEMODELRESPONSE
DESCRIPTOR.message_types_by_name['DeleteModelRequest'] = _DELETEMODELREQUEST
DESCRIPTOR.message_types_by_name['ListVersionsRequest'] = _LISTVERSIONSREQUEST
DESCRIPTOR.message_types_by_name['FindVersionRequest'] = _FINDVERSIONREQUEST
DESCRIPTOR.message_types_by_name['SetPrimaryVersionRequest'] = _SETPRIMARYVERSIONREQUEST
DESCRIPTOR.message_types_by_name['CreateVersionRequest'] = _CREATEVERSIONREQUEST
DESCRIPTOR.message_types_by_name['CreateVersionResponse'] = _CREATEVERSIONRESPONSE
DESCRIPTOR.message_types_by_name['DeleteVersionRequest'] = _DELETEVERSIONREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Model = _reflection.GeneratedProtocolMessageType('Model', (_message.Message,), dict(
  DESCRIPTOR = _MODEL,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.Model)
  ))
_sym_db.RegisterMessage(Model)

ModelVersion = _reflection.GeneratedProtocolMessageType('ModelVersion', (_message.Message,), dict(
  DESCRIPTOR = _MODELVERSION,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.ModelVersion)
  ))
_sym_db.RegisterMessage(ModelVersion)

ModelFeature = _reflection.GeneratedProtocolMessageType('ModelFeature', (_message.Message,), dict(
  DESCRIPTOR = _MODELFEATURE,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.ModelFeature)
  ))
_sym_db.RegisterMessage(ModelFeature)

PrecomputedFeaturesSet = _reflection.GeneratedProtocolMessageType('PrecomputedFeaturesSet', (_message.Message,), dict(
  DESCRIPTOR = _PRECOMPUTEDFEATURESSET,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.PrecomputedFeaturesSet)
  ))
_sym_db.RegisterMessage(PrecomputedFeaturesSet)

FindModelRequest = _reflection.GeneratedProtocolMessageType('FindModelRequest', (_message.Message,), dict(
  DESCRIPTOR = _FINDMODELREQUEST,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.FindModelRequest)
  ))
_sym_db.RegisterMessage(FindModelRequest)

CreateModelRequest = _reflection.GeneratedProtocolMessageType('CreateModelRequest', (_message.Message,), dict(

  PrecomputedFeaturesEntry = _reflection.GeneratedProtocolMessageType('PrecomputedFeaturesEntry', (_message.Message,), dict(
    DESCRIPTOR = _CREATEMODELREQUEST_PRECOMPUTEDFEATURESENTRY,
    __module__ = 'models_pb2'
    # @@protoc_insertion_point(class_scope:photon.CreateModelRequest.PrecomputedFeaturesEntry)
    ))
  ,
  DESCRIPTOR = _CREATEMODELREQUEST,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.CreateModelRequest)
  ))
_sym_db.RegisterMessage(CreateModelRequest)
_sym_db.RegisterMessage(CreateModelRequest.PrecomputedFeaturesEntry)

CreateModelResponse = _reflection.GeneratedProtocolMessageType('CreateModelResponse', (_message.Message,), dict(
  DESCRIPTOR = _CREATEMODELRESPONSE,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.CreateModelResponse)
  ))
_sym_db.RegisterMessage(CreateModelResponse)

DeleteModelRequest = _reflection.GeneratedProtocolMessageType('DeleteModelRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETEMODELREQUEST,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.DeleteModelRequest)
  ))
_sym_db.RegisterMessage(DeleteModelRequest)

ListVersionsRequest = _reflection.GeneratedProtocolMessageType('ListVersionsRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTVERSIONSREQUEST,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.ListVersionsRequest)
  ))
_sym_db.RegisterMessage(ListVersionsRequest)

FindVersionRequest = _reflection.GeneratedProtocolMessageType('FindVersionRequest', (_message.Message,), dict(
  DESCRIPTOR = _FINDVERSIONREQUEST,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.FindVersionRequest)
  ))
_sym_db.RegisterMessage(FindVersionRequest)

SetPrimaryVersionRequest = _reflection.GeneratedProtocolMessageType('SetPrimaryVersionRequest', (_message.Message,), dict(
  DESCRIPTOR = _SETPRIMARYVERSIONREQUEST,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.SetPrimaryVersionRequest)
  ))
_sym_db.RegisterMessage(SetPrimaryVersionRequest)

CreateVersionRequest = _reflection.GeneratedProtocolMessageType('CreateVersionRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATEVERSIONREQUEST,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.CreateVersionRequest)
  ))
_sym_db.RegisterMessage(CreateVersionRequest)

CreateVersionResponse = _reflection.GeneratedProtocolMessageType('CreateVersionResponse', (_message.Message,), dict(
  DESCRIPTOR = _CREATEVERSIONRESPONSE,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.CreateVersionResponse)
  ))
_sym_db.RegisterMessage(CreateVersionResponse)

DeleteVersionRequest = _reflection.GeneratedProtocolMessageType('DeleteVersionRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETEVERSIONREQUEST,
  __module__ = 'models_pb2'
  # @@protoc_insertion_point(class_scope:photon.DeleteVersionRequest)
  ))
_sym_db.RegisterMessage(DeleteVersionRequest)


_CREATEMODELREQUEST_PRECOMPUTEDFEATURESENTRY.has_options = True
_CREATEMODELREQUEST_PRECOMPUTEDFEATURESENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
# @@protoc_insertion_point(module_scope)