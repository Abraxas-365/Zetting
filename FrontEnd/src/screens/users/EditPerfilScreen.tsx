import React from 'react';
import { SafeAreaView, ScrollView, Text, View } from 'react-native';
import { styleBackgrounds } from '../../themes/Backgrounds';
import { styleTitles } from '../../themes/Titles';
import { styleWrappers } from '../../themes/Wrappers';

const EditPerfilScreen = () => {
    return (
        <SafeAreaView style={styleBackgrounds.fondoDark}>
            <View style={{ ...styleWrappers.wrapperTitles }}>
                <Text style={{ ...styleTitles.titleTextLight, alignSelf: 'center' }}>REGISTER</Text>
            </View>
        </SafeAreaView>
    );
};

export default EditPerfilScreen;
