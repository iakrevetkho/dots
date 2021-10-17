///
//  Generated code. Do not modify.
//  source: spot/v1/spot_v1.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use createSpotRequestDescriptor instead')
const CreateSpotRequest$json = const {
  '1': 'CreateSpotRequest',
  '2': const [
    const {'1': 'position', '3': 1, '4': 1, '5': 11, '6': '.spot.v1.Position', '10': 'position'},
    const {'1': 'radius', '3': 2, '4': 1, '5': 5, '10': 'radius'},
    const {'1': 'scan_period_in_seconds', '3': 3, '4': 1, '5': 5, '10': 'scanPeriodInSeconds'},
    const {'1': 'zone_period_in_seconds', '3': 4, '4': 1, '5': 5, '10': 'zonePeriodInSeconds'},
  ],
};

/// Descriptor for `CreateSpotRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createSpotRequestDescriptor = $convert.base64Decode('ChFDcmVhdGVTcG90UmVxdWVzdBItCghwb3NpdGlvbhgBIAEoCzIRLnNwb3QudjEuUG9zaXRpb25SCHBvc2l0aW9uEhYKBnJhZGl1cxgCIAEoBVIGcmFkaXVzEjMKFnNjYW5fcGVyaW9kX2luX3NlY29uZHMYAyABKAVSE3NjYW5QZXJpb2RJblNlY29uZHMSMwoWem9uZV9wZXJpb2RfaW5fc2Vjb25kcxgEIAEoBVITem9uZVBlcmlvZEluU2Vjb25kcw==');
@$core.Deprecated('Use createSpotResponseDescriptor instead')
const CreateSpotResponse$json = const {
  '1': 'CreateSpotResponse',
  '2': const [
    const {'1': 'spot_uuid', '3': 1, '4': 1, '5': 9, '10': 'spotUuid'},
  ],
};

/// Descriptor for `CreateSpotResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createSpotResponseDescriptor = $convert.base64Decode('ChJDcmVhdGVTcG90UmVzcG9uc2USGwoJc3BvdF91dWlkGAEgASgJUghzcG90VXVpZA==');
@$core.Deprecated('Use getSpotRequestDescriptor instead')
const GetSpotRequest$json = const {
  '1': 'GetSpotRequest',
  '2': const [
    const {'1': 'spot_uuid', '3': 1, '4': 1, '5': 9, '10': 'spotUuid'},
  ],
};

/// Descriptor for `GetSpotRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getSpotRequestDescriptor = $convert.base64Decode('Cg5HZXRTcG90UmVxdWVzdBIbCglzcG90X3V1aWQYASABKAlSCHNwb3RVdWlk');
@$core.Deprecated('Use getSpotResponseDescriptor instead')
const GetSpotResponse$json = const {
  '1': 'GetSpotResponse',
  '2': const [
    const {'1': 'position', '3': 1, '4': 1, '5': 11, '6': '.spot.v1.Position', '10': 'position'},
    const {'1': 'radius', '3': 3, '4': 1, '5': 5, '10': 'radius'},
    const {'1': 'scan_period_in_seconds', '3': 4, '4': 1, '5': 5, '10': 'scanPeriodInSeconds'},
    const {'1': 'zone_period_in_seconds', '3': 5, '4': 1, '5': 5, '10': 'zonePeriodInSeconds'},
  ],
};

/// Descriptor for `GetSpotResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getSpotResponseDescriptor = $convert.base64Decode('Cg9HZXRTcG90UmVzcG9uc2USLQoIcG9zaXRpb24YASABKAsyES5zcG90LnYxLlBvc2l0aW9uUghwb3NpdGlvbhIWCgZyYWRpdXMYAyABKAVSBnJhZGl1cxIzChZzY2FuX3BlcmlvZF9pbl9zZWNvbmRzGAQgASgFUhNzY2FuUGVyaW9kSW5TZWNvbmRzEjMKFnpvbmVfcGVyaW9kX2luX3NlY29uZHMYBSABKAVSE3pvbmVQZXJpb2RJblNlY29uZHM=');
@$core.Deprecated('Use sendPlayerPositionRequestDescriptor instead')
const SendPlayerPositionRequest$json = const {
  '1': 'SendPlayerPositionRequest',
  '2': const [
    const {'1': 'spot_uuid', '3': 1, '4': 1, '5': 9, '10': 'spotUuid'},
    const {'1': 'player_uuid', '3': 2, '4': 1, '5': 9, '10': 'playerUuid'},
    const {'1': 'position', '3': 3, '4': 1, '5': 11, '6': '.spot.v1.Position', '10': 'position'},
  ],
};

/// Descriptor for `SendPlayerPositionRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sendPlayerPositionRequestDescriptor = $convert.base64Decode('ChlTZW5kUGxheWVyUG9zaXRpb25SZXF1ZXN0EhsKCXNwb3RfdXVpZBgBIAEoCVIIc3BvdFV1aWQSHwoLcGxheWVyX3V1aWQYAiABKAlSCnBsYXllclV1aWQSLQoIcG9zaXRpb24YAyABKAsyES5zcG90LnYxLlBvc2l0aW9uUghwb3NpdGlvbg==');
@$core.Deprecated('Use sendPlayerPositionResponseDescriptor instead')
const SendPlayerPositionResponse$json = const {
  '1': 'SendPlayerPositionResponse',
};

/// Descriptor for `SendPlayerPositionResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sendPlayerPositionResponseDescriptor = $convert.base64Decode('ChpTZW5kUGxheWVyUG9zaXRpb25SZXNwb25zZQ==');
@$core.Deprecated('Use getPlayersStatesRequestDescriptor instead')
const GetPlayersStatesRequest$json = const {
  '1': 'GetPlayersStatesRequest',
  '2': const [
    const {'1': 'spot_uuid', '3': 1, '4': 1, '5': 9, '10': 'spotUuid'},
    const {'1': 'player_uuid', '3': 2, '4': 1, '5': 9, '10': 'playerUuid'},
  ],
};

/// Descriptor for `GetPlayersStatesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getPlayersStatesRequestDescriptor = $convert.base64Decode('ChdHZXRQbGF5ZXJzU3RhdGVzUmVxdWVzdBIbCglzcG90X3V1aWQYASABKAlSCHNwb3RVdWlkEh8KC3BsYXllcl91dWlkGAIgASgJUgpwbGF5ZXJVdWlk');
@$core.Deprecated('Use getPlayersStatesResponseDescriptor instead')
const GetPlayersStatesResponse$json = const {
  '1': 'GetPlayersStatesResponse',
  '2': const [
    const {'1': 'player_state', '3': 1, '4': 1, '5': 11, '6': '.spot.v1.PlayerState', '10': 'playerState'},
    const {'1': 'other_players_states', '3': 2, '4': 3, '5': 11, '6': '.spot.v1.PlayerState', '10': 'otherPlayersStates'},
  ],
};

/// Descriptor for `GetPlayersStatesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getPlayersStatesResponseDescriptor = $convert.base64Decode('ChhHZXRQbGF5ZXJzU3RhdGVzUmVzcG9uc2USNwoMcGxheWVyX3N0YXRlGAEgASgLMhQuc3BvdC52MS5QbGF5ZXJTdGF0ZVILcGxheWVyU3RhdGUSRgoUb3RoZXJfcGxheWVyc19zdGF0ZXMYAiADKAsyFC5zcG90LnYxLlBsYXllclN0YXRlUhJvdGhlclBsYXllcnNTdGF0ZXM=');
@$core.Deprecated('Use playerStateDescriptor instead')
const PlayerState$json = const {
  '1': 'PlayerState',
  '2': const [
    const {'1': 'player_uuid', '3': 1, '4': 1, '5': 9, '10': 'playerUuid'},
    const {'1': 'position', '3': 2, '4': 1, '5': 11, '6': '.spot.v1.Position', '10': 'position'},
    const {'1': 'health', '3': 3, '4': 1, '5': 5, '10': 'health'},
  ],
};

/// Descriptor for `PlayerState`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List playerStateDescriptor = $convert.base64Decode('CgtQbGF5ZXJTdGF0ZRIfCgtwbGF5ZXJfdXVpZBgBIAEoCVIKcGxheWVyVXVpZBItCghwb3NpdGlvbhgCIAEoCzIRLnNwb3QudjEuUG9zaXRpb25SCHBvc2l0aW9uEhYKBmhlYWx0aBgDIAEoBVIGaGVhbHRo');
@$core.Deprecated('Use positionDescriptor instead')
const Position$json = const {
  '1': 'Position',
  '2': const [
    const {'1': 'longitude', '3': 1, '4': 1, '5': 1, '10': 'longitude'},
    const {'1': 'latitude', '3': 2, '4': 1, '5': 1, '10': 'latitude'},
  ],
};

/// Descriptor for `Position`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List positionDescriptor = $convert.base64Decode('CghQb3NpdGlvbhIcCglsb25naXR1ZGUYASABKAFSCWxvbmdpdHVkZRIaCghsYXRpdHVkZRgCIAEoAVIIbGF0aXR1ZGU=');
