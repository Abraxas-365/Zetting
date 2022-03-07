

import React from 'react';
import { SafeAreaView, ScrollView, StyleSheet, Text, View } from "react-native";
import RegisterForm from '../components/RegisterForm';

const Register2Screen = () => {

    return (
        <SafeAreaView style={styleRegisterScreen.fondo}>
            <View style={{ alignSelf: 'center', justifyContent: 'center', top: '4%' }}>
                <Text style={styleRegisterScreen.title}>Register</Text>

            </View>

            <RegisterForm />


        </SafeAreaView>
    )
};

export default Register2Screen;

const styleRegisterScreen = StyleSheet.create({
    fondo: {
        flex: 1,
        backgroundColor: '#E2E0F3',
    },
    title: {
        fontSize: 30,
        color: '#E5E1F6',
        fontWeight: 'bold'
    },

})
