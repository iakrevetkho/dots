// External
import 'package:geolocator/geolocator.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:latlong2/latlong.dart';
import 'package:logging/logging.dart';

// Internal
import 'events.dart';
import 'state.dart';

class SpotPageBloc extends Bloc<SpotPageEvent, SpotPageState> {
  final _logger = Logger("SpotPageBloc");

  final LatLng spotPosition;

  SpotPageBloc({
    required this.spotPosition,
  }) : super(InitingState(spotPosition: spotPosition)) {
    on<InitEvent>((event, emit) async {
      // Couldn't check permission because we checked it before
      _logger.fine("Get last known position");
      final position = await Geolocator.getLastKnownPosition();
      if (position != null) {
        emit(InitedState(
          spotPosition: spotPosition,
          position: LatLng(position.latitude, position.longitude),
        ));
      } else {
        throw Exception("Position is null");
      }

      _logger.fine("Subscribe on location");
      Geolocator.getPositionStream(
        desiredAccuracy: LocationAccuracy.high,
      ).listen((position) => add(NewGeoPositionEvent(
              position: LatLng(
            position.latitude,
            position.longitude,
          ))));
    });
    on<NewGeoPositionEvent>((event, emit) async {
      if (state is InitedState) {
        emit(InitedState(
          spotPosition: spotPosition,
          position: event.position,
        ));
      }
    });

    add(InitEvent());
  }
}