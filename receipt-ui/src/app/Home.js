import React from 'react';
import { StatusBar } from 'expo-status-bar';
import { StyleSheet, View } from 'react-native';

import ApiContainer from '../../src/app/screen/ApiContainer';

// add export statement so that this function is available to other components of the app. 
export default function HomePage() {
    return (
        <View style={styles.container}>
            <ApiContainer />
          <StatusBar style="auto" />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center'
    }
})