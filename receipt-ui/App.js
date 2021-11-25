import 'react-native-gesture-handler';
import React from 'react';
import { useState } from 'react';
import { Asset } from 'expo-asset';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import HomePage from './Home';

import { Platform } from 'react-native';
import { useFonts } from 'expo-font';
import AppLoading from 'expo-app-loading';
import Header from './Header';

import { StatusBar } from 'expo-status-bar';

const Stack = createStackNavigator();

export default function App() {

  let [fontsLoaded] = useFonts({
    'OpenSans': require('./assets/fonts/OpenSans-Regular.ttf'),
  });


  if ( !fontsLoaded ) {
    return (
      <AppLoading />
    );
  } else {

    return (

      <NavigationContainer
        style={{ paddingTop: Platform.OS === 'android' ? StatusBar.currentHeight : 0 }}>
        <Stack.Navigator
          initiateRouteName="Receipts"
          headerMode="screen"
        >
          <Stack.Screen
            name="Receipts"
            component={HomePage}
            options={{
              header: () => <Header headerDisplay="Receipt" />
            }}
          />
        </Stack.Navigator>
      </NavigationContainer>

    );
  }
}
