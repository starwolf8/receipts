import React, { useEffect, useState } from 'react';
import { StatusBar } from 'expo-status-bar';
import { FlatList, Text, Image, StyleSheet, ActivityIndicator, TouchableWithoutFeedback, View } from 'react-native';

import ApiContainer from '../../src/app/screen/ApiContainer';
import DemoTable from './receipt/receipt-table';

// add export statement so that this function is available to other components of the app.
export default function HomePage({ navigation }) {

    const [dataLoading, finishLoading] = useState(true);
    const [receiptData, setData] = useState([]);


    useEffect(() => {
        fetch("http://127.0.0.1:5001/api/receipts")
            .then((response) => response.json())
            .then((json) => {
                setTimeout(() => {
                    console.log("getting data from fetch", json);
                    setData(json);
                }, 500)
                // setData(json);
            })
            .catch((error) => console.error(error))
            .finally(() => finishLoading(false))
    }, []);




    const renderItem = (data) => {
        return (
            <TouchableWithoutFeedback key={data.item.receiptID}
            onPress={() =>
                navigation.navigate('ReceiptDetail', { url: item.url })}
                // navigation.navigate('ReceiptDetail') }
                >
                <View style={styles.listings}>
                    <Text style={styles.title}>Receipt ID: {data.item.receiptID}</Text>
                    <Text style={styles.blurb}>Store Name: {data.item.storeName}</Text>
                    <Text style={styles.blurb}>Date of Purchase: {data.item.dateOfPurchase}</Text>
                    <Text style={styles.blurb}>Total Cost: {data.item.totalCost}</Text>
                    {/* <Image style={styles.thumbnail} source={{uri: item.urlToImage}} /> */}
                    {/* <Text style={styles.blurb}>{item.item.receiptID}</Text> */}
                </View>
            </TouchableWithoutFeedback>
        );
    };
    return (
        <View style={styles.container}>
            {/* <ApiContainer />
            <StatusBar style="auto" /> */}

            <DemoTable/>

            {/* {dataLoading ? <ActivityIndicator /> : (
                <FlatList
                    data={receiptData}
                    renderItem={item => renderItem(item)}
                    keyExtractor={(item) => item.receiptID.toString()}
                />
            )} */}
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
        width: '100%',
        padding: 20
    },
    thumbnail: {
        height: 100,
        width: '98%'
    },
    listings: {
        paddingTop: 15,
        paddingBottom: 25,
        borderBottomColor: 'black',
        borderBottomWidth: 1
    },
    title: {
        paddingBottom: 10,
        fontFamily: 'OpenSans',
        fontWeight: 'bold'
    },
    blurb: {
        fontFamily: 'OpenSans',
        fontStyle: 'italic'
    }

})