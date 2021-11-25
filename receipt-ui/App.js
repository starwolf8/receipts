import 'react-native-gesture-handler';
import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import HomePage from './src/app/Home';

import { Platform } from 'react-native';
import { useFonts } from 'expo-font';
import AppLoading from 'expo-app-loading';
import Header from './src/app/Header';
import Footer from './src/app/Footer';
import { navigationRef } from './src/app/RootNavigation';

import { StatusBar } from 'expo-status-bar';

const Stack = createStackNavigator();

export default function App() {

  let [fontsLoaded] = useFonts({
    'OpenSans': require('./assets/fonts/OpenSans-Regular.ttf'),
  });


  if (!fontsLoaded) {
    return (
      <AppLoading />
    );
  } else {

    return (

      <NavigationContainer
        style={{ paddingTop: Platform.OS === 'android' ? StatusBar.currentHeight : 0 }}
        ref={navigationRef}>
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
        <Footer />
      </NavigationContainer>

    );
  }
}
