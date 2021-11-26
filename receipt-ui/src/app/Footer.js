import React from 'react';
import {StyleSheet, Text, View, TouchableOpacity } from 'react-native';

import * as RootNavigation from './RootNavigation';

export default function Footer() {
    return (
        <View style={styles.footer}>
            <TouchableOpacity
            style={styles.button}
            onPress={() => RootNavigation.navigate('HomePage')}
            >
            <Text>Home</Text>
            </TouchableOpacity>

            <TouchableOpacity 
            style={styles.button}
            onPress={() => RootNavigation.navigate('ReceiptPage')}>
                <Text>Receipts</Text>
            </TouchableOpacity>
            <TouchableOpacity style={styles.button}>
                <Text>Products</Text>
            </TouchableOpacity>
        </View>
    )
}

const styles = StyleSheet.create({
    footer: {
        width: '',
        height: 80,
        flexDirection: 'row',
        alignItems: 'flex-start',
        justifyContent: 'center'
    },
    button: {
        padding: 20
    }
})