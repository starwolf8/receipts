import React, { Component } from 'react';
import ApiView from './ApiView';
import axios from 'axios';
import styles from './ApiStyles';
import {
    View,
    Text,
    TouchableOpacity
} from "react-native";
class ApiContainer extends Component {
    constructor(props) {
        super(props);
        this.state = {
            loading: false,
            fromFetch: false,
            fromAxios: false,
            dataSource: [],
            axiosData: null
        };
    }
    goForFetch = () => {
        this.setState({
            fromFetch: true,
            loading: true,

        })
        // fetch("https://jsonplaceholder.typicode.com/users")
        fetch("http://localhost:5001/api/receipts")
            .then(response => response.json())
            .then((responseJson) => {
                console.log('getting data from fetch', responseJson)
                setTimeout(() => {
                    this.setState({
                        loading: false,
                        dataSource: responseJson
                    })
                }, 2000)

            })
            .catch(error => console.log(error))
    }
    goForAxios = () => {
        this.setState({
            fromFetch: false,
            loading: true,

        })
        // axios.get("https://jsonplaceholder.typicode.com/users")
        axios.get("http://localhost:5001/api/receipts")
            .then(response => {
                console.log('getting data from axios', response.data);
                setTimeout(() => {
                    this.setState({
                        loading: false,
                        axiosData: response.data
                    })
                }, 2000)
            })
            .catch(error => {
                console.log(error);
            });
    }
    FlatListSeparator = () => {
        return (
            <View style={{
                height: .5,
                width: "100%",
                backgroundColor: "rgba(0,0,0,0.5)",
            }}
            />
        );
    }
    renderItem = (data) => {
        return (
            <TouchableOpacity style={styles.list}>
                <Text style={styles.lightText}>ReceiptID: {data.item.receiptID}</Text>
                <Text style={styles.lightText}>StoreName: {data.item.storeName}</Text>
                <Text style={styles.lightText}>Total Cost: {data.item.totalCost}</Text>
            </TouchableOpacity>
        )

    }


    render() {
        const { dataSource, fromFetch, fromAxios, loading, axiosData } = this.state
        return (
            <ApiView
                goForFetch={this.goForFetch}
                goForAxios={this.goForAxios}
                dataSource={dataSource}
                loading={loading}
                fromFetch={fromFetch}
                fromAxios={fromAxios}
                axiosData={axiosData}
                FlatListSeparator={this.FlatListSeparator}
                renderItem={this.renderItem}
            />
        );
    }
}

export default ApiContainer;