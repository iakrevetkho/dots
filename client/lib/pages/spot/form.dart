// External
import 'package:dots_client/pages/spot/resources/player_position.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_map/plugin_api.dart';
import 'package:latlong2/latlong.dart';

// Internal
import 'bloc/bloc.dart';
import 'bloc/state.dart';

class SpotForm extends StatelessWidget {
  const SpotForm({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<SpotPageBloc, SpotPageState>(
      builder: (context, state) {
        if (state is InitingState) {
          return const Center(
            child: CircularProgressIndicator(),
          );
        } else if (state is InitedState) {
          return _InitedStateView(
            playerHealth: state.playerHealth,
            playerPosition: state.playerPosition,
            otherPlayersPositions: state.otherPlayersPositions,
            spotPosition: state.spotPosition,
            zoneRadius: state.zoneRadius,
          );
        }

        return Text("Unkown state: $state");
      },
    );
  }
}

class _InitedStateView extends StatelessWidget {
  final LatLng playerPosition;
  final int playerHealth;
  final List<PlayerPosition> otherPlayersPositions;
  final LatLng spotPosition;
  // Spot radius in meters
  final int zoneRadius;

  const _InitedStateView({
    required this.playerPosition,
    required this.playerHealth,
    required this.otherPlayersPositions,
    required this.spotPosition,
    required this.zoneRadius,
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final markers = <Marker>[];
    // Spot position pointer
    markers.add(Marker(
      point: spotPosition,
      builder: (ctx) => const Icon(
        Icons.location_pin,
        size: 40,
        color: Colors.red,
      ),
    ));
    // Player position pointer
    markers.add(Marker(
      point: playerPosition,
      builder: (ctx) => const Icon(
        Icons.circle,
      ),
    ));

    // Other players pointers
    for (var otherPlayerPosition in otherPlayersPositions) {
      markers.add(Marker(
        point: otherPlayerPosition.position,
        builder: (ctx) => const Icon(
          Icons.circle,
          color: Colors.amber,
        ),
      ));
    }

    return Stack(
      children: [
        FlutterMap(
          options: MapOptions(
            center: playerPosition,
            zoom: 17.0,
          ),
          layers: [
            TileLayerOptions(
              urlTemplate: "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
              subdomains: ['a', 'b', 'c'],
            ),
            MarkerLayerOptions(markers: markers),
            CircleLayerOptions(circles: [
              CircleMarker(
                point: spotPosition,
                useRadiusInMeter: true,
                color: const Color.fromRGBO(0, 0, 0, 0),
                borderColor: Colors.red,
                borderStrokeWidth: 2,
                radius: zoneRadius.toDouble(),
              ),
            ])
          ],
        ),
        Positioned(
          right: 0,
          top: 0,
          child: Padding(
            padding: const EdgeInsets.all(10),
            child: Row(
              children: [
                Padding(
                  padding: const EdgeInsets.only(right: 10),
                  child: Text(
                    "Health: $playerHealth",
                    style: Theme.of(context)
                        .textTheme
                        .headline4!
                        .copyWith(color: _healthColor(playerHealth)),
                  ),
                ),
                CircularProgressIndicator(
                  value: playerHealth / 100,
                  strokeWidth: 10,
                  color: _healthColor(playerHealth),
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }

  Color _healthColor(int health) {
    if (health > 60) {
      return Colors.green;
    } else if (health > 30) {
      return Colors.orange;
    } else if (health > 0) {
      return Colors.red;
    } else {
      return Colors.black;
    }
  }
}
