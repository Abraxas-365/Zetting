

import React from 'react';
import { Button, Image, SafeAreaView, Text, View } from 'react-native';
import { StyleSheet } from 'react-native'
import BottonsLoginScreen from '../components/BottonsLoginScreen';


const FavoritosScreen = () => {
    return (
        <SafeAreaView style={styleFavoritosScreen.fondo}>

            <View style={{ alignSelf: 'center', justifyContent: 'center', top: '30%' }}>
                <Text style={styleFavoritosScreen.slogan}>HOME</Text>
            </View>
        </SafeAreaView>
    );
};

export default FavoritosScreen;



export const styleFavoritosScreen = StyleSheet.create({
    fondo: {
        flex: 1,
        backgroundColor: '#23232B',
    },
    slogan: {
        fontSize: 40,
        color: '#E5E1F6',
        fontWeight: 'bold'
    }

})
