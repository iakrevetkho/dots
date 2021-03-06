import 'package:dots_client/pages/settings/widget/environment.dart';
import 'package:dots_client/utils/nav.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_settings_ui/flutter_settings_ui.dart';
import 'bloc/events.dart';
import 'bloc/bloc.dart';
import 'bloc/state.dart';

class SettingsForm extends StatelessWidget {
  const SettingsForm({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<SettingsPageBloc, SettingsPageState>(
      builder: (context, state) {
        if (state is InitedState) {
          return SettingsList(
            sections: [
              SettingsSection(
                title: 'Theme',
                tiles: [
                  SettingsTile.switchTile(
                    title: 'Use OS theme',
                    subtitle: 'Use OS ligth/dark theme setting.',
                    leading: const Icon(Icons.wb_twilight_rounded),
                    onToggle: (value) => context
                        .read<SettingsPageBloc>()
                        .add(ChangeUseOsThemeEvent(
                          context: context,
                          value: value,
                        )),
                    switchValue: state.settings.useOsThemeSettings,
                  ),
                  SettingsTile.switchTile(
                    title: 'Ligth theme',
                    subtitle: 'Use ligth or dark theme.',
                    leading: const Icon(Icons.wb_sunny_rounded),
                    onToggle: (value) => context
                        .read<SettingsPageBloc>()
                        .add(ChangeLightThemeEvent(
                          context: context,
                          value: value,
                        )),
                    switchValue: state.settings.ligthTheme,
                    enabled: !state.settings.useOsThemeSettings,
                  ),
                ],
              ),
              SettingsSection(
                title: 'Common',
                tiles: [
                  SettingsTile(
                    title: 'Environment',
                    subtitle: state.settings.environment.name,
                    leading: const Icon(Icons.cloud_queue),
                    trailing: const Icon(Icons.arrow_forward_ios),
                    onPressed: (context) async {
                      final index =
                          await navPush(context, const EnvironmentPage());
                      context
                          .read<SettingsPageBloc>()
                          .add(ChangeEnvironmentEvent(index: index));
                    },
                  ),
                ],
              ),
              SettingsSection(
                title: 'Misc',
                tiles: [
                  SettingsTile(
                    title: 'Terms of Service',
                    leading: const Icon(Icons.description),
                    trailing: const SizedBox(),
                  ),
                  SettingsTile(
                    title: 'Changelog',
                    leading: const Icon(Icons.list),
                    trailing: const SizedBox(),
                  ),
                  SettingsTile(
                    title: 'Open source licenses',
                    leading: const Icon(Icons.collections_bookmark),
                    trailing: const SizedBox(),
                  ),
                ],
              ),
              CustomSection(
                child: Column(
                  children: const [
                    Padding(
                      padding: EdgeInsets.only(top: 22, bottom: 8),
                      child: Icon(Icons.settings),
                    ),
                    Text(
                      'Version: 2.4.0 (287)',
                      style: TextStyle(color: Color(0xFF777777)),
                    ),
                  ],
                ),
              ),
            ],
          );
        }

        return Text("Unkown state: $state");
      },
    );
  }
}
