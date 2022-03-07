import React from 'react';
import { SafeAreaView, Text, View } from 'react-native';
import CreateProjectForm from '../components/CreateProyectForm';
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';

const CreateProyectScreen = () => {
    return (
        <SafeAreaView style={styleBackgrounds.fondoDark}>

            <View style={styleTitles.titleCenterView}>
                <Text style={styleTitles.titleTextLight}>NEW PROYECT</Text>
            </View>
            <CreateProjectForm />
        </SafeAreaView>
    );
};

export default CreateProyectScreen;
