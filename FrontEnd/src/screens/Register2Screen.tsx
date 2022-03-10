

import { StackScreenProps } from '@react-navigation/stack';
import React from 'react';
import { SafeAreaView, ScrollView, StyleSheet, Text, View } from "react-native";
import RegisterForm from '../components/RegisterForm';
import { RootStackParamList } from '../navigator/StackNavigator';
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';

interface Props extends StackScreenProps<RootStackParamList, 'ProyectoScreen'> { }
const Register2Screen = ({ route }: Props) => {
    console.log('Register2Screen')
    console.log("email", route.params.email);

    return (
        <SafeAreaView style={styleBackgrounds.fondoLight}>
            <View style={{ alignSelf: 'center', justifyContent: 'center', top: '4%' }}>
                <Text style={styleTitles.titleTextDark}>Register</Text>
            </View>

            <RegisterForm email={route.params.email} />


        </SafeAreaView>
    )
};

export default Register2Screen;

