// Mocks generated by Mockito 5.0.15 from annotations
// in dots_client/test/widget_test.dart.
// Do not manually edit this file.

import 'dart:async' as _i3;

import 'package:dots_client/settings/controller.dart' as _i2;
import 'package:dots_client/settings/settings.dart' as _i4;
import 'package:mockito/mockito.dart' as _i1;

// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis

/// A class which mocks [AppSettingsController].
///
/// See the documentation for Mockito's code generation for more information.
class MockAppSettingsController extends _i1.Mock
    implements _i2.AppSettingsController {
  MockAppSettingsController() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i3.Future<void> init() => (super.noSuchMethod(Invocation.method(#init, []),
      returnValue: Future<void>.value(),
      returnValueForMissingStub: Future<void>.value()) as _i3.Future<void>);
  @override
  _i3.Future<bool> save(_i4.AppSettings? newSettings) =>
      (super.noSuchMethod(Invocation.method(#save, [newSettings]),
          returnValue: Future<bool>.value(false)) as _i3.Future<bool>);
  @override
  String toString() => super.toString();
}
