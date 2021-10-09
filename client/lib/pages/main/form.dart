// External
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

// Internal
import 'bloc/events.dart';
import 'bloc/bloc.dart';
import 'bloc/state.dart';

class MainForm extends StatelessWidget {
  const MainForm({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<MainPageBloc, MainPageState>(
      builder: (context, state) {
        if (state is InitedState) {
          return Center(
            child: ElevatedButton(
              key: const Key("btn_create_spot"),
              child: const Text("Create new spot"),
              onPressed: () =>
                  context.read<MainPageBloc>().add(CreateNewSpotEvent()),
            ),
          );
        } else if (state is CreatingNewSpotState) {
          return const Center(
            child: CircularProgressIndicator(),
          );
        } else if (state is NewSpotCreatedState) {
          return Center(
            child: Text("New spot UUID: ${state.spotUuid}"),
          );
        } else if (state is CreateSpotErrorState) {
          return Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Padding(
                  padding: const EdgeInsets.only(bottom: 20),
                  child: Text("Error when create new spot: ${state.error}"),
                ),
                ElevatedButton(
                  key: const Key("btn_create_spot"),
                  child: const Text("Create new spot"),
                  onPressed: () =>
                      context.read<MainPageBloc>().add(CreateNewSpotEvent()),
                ),
              ],
            ),
          );
        }

        return Text("Unkown state: $state");
      },
    );
  }
}
