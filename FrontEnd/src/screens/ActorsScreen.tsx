import React, { useState } from 'react';
import { SafeAreaView, Text, View } from 'react-native';
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';

const ActorsScreen = () => {



    return (
        <SafeAreaView style={styleBackgrounds.fondoDark}>
            <View style={styleTitles.titleCenterView}>
                <Text style={styleTitles.titleTextLight}>ACTORS</Text>
            </View>
        </SafeAreaView>
    );
};

export default ActorsScreen;
