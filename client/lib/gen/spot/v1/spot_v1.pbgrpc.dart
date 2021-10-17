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
  static final _$getSpot =
      $grpc.ClientMethod<$0.GetSpotRequest, $0.GetSpotResponse>(
          '/spot.v1.SpotService/GetSpot',
          ($0.GetSpotRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.GetSpotResponse.fromBuffer(value));
  static final _$sendPlayerPosition = $grpc.ClientMethod<
          $0.SendPlayerPositionRequest, $0.SendPlayerPositionResponse>(
      '/spot.v1.SpotService/SendPlayerPosition',
      ($0.SendPlayerPositionRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.SendPlayerPositionResponse.fromBuffer(value));
  static final _$getPlayersStates = $grpc.ClientMethod<
          $0.GetPlayersStatesRequest, $0.GetPlayersStatesResponse>(
      '/spot.v1.SpotService/GetPlayersStates',
      ($0.GetPlayersStatesRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.GetPlayersStatesResponse.fromBuffer(value));

  SpotServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.CreateSpotResponse> createSpot(
      $0.CreateSpotRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createSpot, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetSpotResponse> getSpot($0.GetSpotRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getSpot, request, options: options);
  }

  $grpc.ResponseFuture<$0.SendPlayerPositionResponse> sendPlayerPosition(
      $async.Stream<$0.SendPlayerPositionRequest> request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$sendPlayerPosition, request, options: options)
        .single;
  }

  $grpc.ResponseStream<$0.GetPlayersStatesResponse> getPlayersStates(
      $0.GetPlayersStatesRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$getPlayersStates, $async.Stream.fromIterable([request]),
        options: options);
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
    $addMethod($grpc.ServiceMethod<$0.GetSpotRequest, $0.GetSpotResponse>(
        'GetSpot',
        getSpot_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetSpotRequest.fromBuffer(value),
        ($0.GetSpotResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.SendPlayerPositionRequest,
            $0.SendPlayerPositionResponse>(
        'SendPlayerPosition',
        sendPlayerPosition,
        true,
        false,
        ($core.List<$core.int> value) =>
            $0.SendPlayerPositionRequest.fromBuffer(value),
        ($0.SendPlayerPositionResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetPlayersStatesRequest,
            $0.GetPlayersStatesResponse>(
        'GetPlayersStates',
        getPlayersStates_Pre,
        false,
        true,
        ($core.List<$core.int> value) =>
            $0.GetPlayersStatesRequest.fromBuffer(value),
        ($0.GetPlayersStatesResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.CreateSpotResponse> createSpot_Pre($grpc.ServiceCall call,
      $async.Future<$0.CreateSpotRequest> request) async {
    return createSpot(call, await request);
  }

  $async.Future<$0.GetSpotResponse> getSpot_Pre(
      $grpc.ServiceCall call, $async.Future<$0.GetSpotRequest> request) async {
    return getSpot(call, await request);
  }

  $async.Stream<$0.GetPlayersStatesResponse> getPlayersStates_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.GetPlayersStatesRequest> request) async* {
    yield* getPlayersStates(call, await request);
  }

  $async.Future<$0.CreateSpotResponse> createSpot(
      $grpc.ServiceCall call, $0.CreateSpotRequest request);
  $async.Future<$0.GetSpotResponse> getSpot(
      $grpc.ServiceCall call, $0.GetSpotRequest request);
  $async.Future<$0.SendPlayerPositionResponse> sendPlayerPosition(
      $grpc.ServiceCall call,
      $async.Stream<$0.SendPlayerPositionRequest> request);
  $async.Stream<$0.GetPlayersStatesResponse> getPlayersStates(
      $grpc.ServiceCall call, $0.GetPlayersStatesRequest request);
}
