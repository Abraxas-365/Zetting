import React, { useState } from 'react';
import { SafeAreaView, Text, View } from 'react-native';
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';

const ActorsScreen = () => {



    return (
        <SafeAreaView style={styleBackgrounds.fondoDark}>
            <View style={{ top: '4%' }}>
                <View style={{ ...styleTitles.titleCenterView, top: 0 }}>
                    <Text style={styleTitles.titleTextLight}>ACTORS</Text>
                </View>
                <View>
                    <Text style={{ ...styleTitles.subTitle, marginTop: '3%', alignSelf: 'center' }} >The sound of Music</Text>
                </View>
            </View>
        </SafeAreaView>
    );
};

export default ActorsScreen;
