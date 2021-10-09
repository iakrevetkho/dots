///
//  Generated code. Do not modify.
//  source: spot/v1/spot_v1.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'spot_v1.pb.dart' as $0;
export 'spot_v1.pb.dart';

class SpotServiceClient extends $grpc.Client {
  static final _$createSpot =
      $grpc.ClientMethod<$0.CreateSpotRequest, $0.CreateSpotResponse>(
          '/spot.v1.SpotService/CreateSpot',
          ($0.CreateSpotRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.CreateSpotResponse.fromBuffer(value));

  SpotServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.CreateSpotResponse> createSpot(
      $0.CreateSpotRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createSpot, request, options: options);
  }
}

abstract class SpotServiceBase extends $grpc.Service {
  $core.String get $name => 'spot.v1.SpotService';

  SpotServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.CreateSpotRequest, $0.CreateSpotResponse>(
        'CreateSpot',
        createSpot_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateSpotRequest.fromBuffer(value),
        ($0.CreateSpotResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.CreateSpotResponse> createSpot_Pre($grpc.ServiceCall call,
      $async.Future<$0.CreateSpotRequest> request) async {
    return createSpot(call, await request);
  }

  $async.Future<$0.CreateSpotResponse> createSpot(
      $grpc.ServiceCall call, $0.CreateSpotRequest request);
}