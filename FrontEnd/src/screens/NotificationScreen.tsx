

import React from 'react';
import { Button, Image, SafeAreaView, Text, View } from 'react-native';
import { StyleSheet } from 'react-native'
import BottonsLoginScreen from '../components/BottonsLoginScreen';
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';
import { styleWrappers } from '../themes/Wrappers';


const NotificationScreen = () => {
    return (
        <SafeAreaView style={styleBackgrounds.fondoDark}>

            <View style={styleWrappers.wrapperHorizontalGap}>
                <View style={{ ...styleWrappers.wrapperTitles }}>
                    <Text style={{ ...styleTitles.titleTextLight }}>NOTIFICATION</Text>
                </View>
                <Text style={{ ...styleText.notificationsNum }}>You have n notifications today</Text>
            </View>
        </SafeAreaView>
    );
};

export default NotificationScreen;



export const styleText = StyleSheet.create({
    notificationsNum: {
        color: '#e5e1f6',
        fontSize: 15,
        marginTop: 15
    }

})
