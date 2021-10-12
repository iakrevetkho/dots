// External
import 'package:dots_client/widgets/map.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_map/plugin_api.dart';
import 'package:latlong2/latlong.dart';
import 'package:numberpicker/numberpicker.dart';

// Internal
import 'bloc/events.dart';
import 'bloc/bloc.dart';
import 'bloc/state.dart';

class SpotSettingsForm extends StatelessWidget {
  final MapController mapController = MapController();

  SpotSettingsForm({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<SpotSettingsPageBloc, SpotSettingsPageState>(
      builder: (context, state) {
        final curState = state;
        if (curState is InitedState) {
          return Center(
            child: SingleChildScrollView(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Text("Location"),
                  SizedBox(
                    width: 300,
                    height: 300,
                    child: MapWidget(
                      mapController: mapController,
                      position: curState.position,
                      zoom: 17.0,
                    ),
                  ),
                  const Text("Radius in meters"),
                  NumberPicker(
                    axis: Axis.horizontal,
                    step: 10,
                    minValue: 0,
                    maxValue: 1000,
                    value: curState.radius,
                    onChanged: (value) => context
                        .read<SpotSettingsPageBloc>()
                        .add(NewRadiusEvent(value: value)),
                  ),
                  const Text("Scan time in seconds"),
                  NumberPicker(
                    axis: Axis.horizontal,
                    step: 10,
                    minValue: 0,
                    maxValue: 1000,
                    value: curState.scanPeriod.inSeconds,
                    onChanged: (value) => context
                        .read<SpotSettingsPageBloc>()
                        .add(NewScanDurationEvent(
                            value: Duration(seconds: value))),
                  ),
                  const Text("Zone time in seconds"),
                  NumberPicker(
                    axis: Axis.horizontal,
                    step: 10,
                    minValue: 0,
                    maxValue: 1000,
                    value: curState.zonePeriod.inSeconds,
                    onChanged: (value) => context
                        .read<SpotSettingsPageBloc>()
                        .add(NewZoneDurationEvent(
                            value: Duration(seconds: value))),
                  ),
                  _CreateNewSpotBtn(
                    position: curState.position,
                    radius: curState.radius,
                    scanPeriod: curState.scanPeriod,
                    zonePeriod: curState.zonePeriod,
                  ),
                ],
              ),
            ),
          );
        }

        return Text("Unkown state: $state");
      },
    );
  }
}

class _CreateNewSpotBtn extends StatelessWidget {
  final LatLng position;
  // Spot radius in meters
  final int radius;
  final Duration scanPeriod;
  final Duration zonePeriod;

  const _CreateNewSpotBtn({
    required this.position,
    required this.radius,
    required this.scanPeriod,
    required this.zonePeriod,
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        mainAxisSize: MainAxisSize.max,
        mainAxisAlignment: MainAxisAlignment.end,
        children: <Widget>[
          Padding(
            padding: const EdgeInsets.only(bottom: 40),
            child: ElevatedButton(
              key: const Key("btn_create_spot"),
              child: const Text("Create"),
              onPressed: () =>
                  context.read<SpotSettingsPageBloc>().add(CreateNewSpotEvent(
                        position: position,
                        radius: radius,
                        scanPeriod: scanPeriod,
                        zonePeriod: zonePeriod,
                      )),
            ),
          ),
        ],
      ),
    );
  }
}
