
import { useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { useContext } from 'react';
import { SafeAreaView, StyleSheet, Text, View } from "react-native";
import { apiCalls } from '../api/apiCalls';
import BottonsPlatforms from '../components/BottonPlatform';
import BottonRegister from '../components/BottonRegister';
import StakeHolders from '../components/StakeHolders';
import { AuthContext } from '../context/AuthContext';
import { useForm } from '../hooks/useForm';
import { RootStackParamList } from '../navigator/StackNavigator';

const RegisterScreen = () => {

    const { email, onChange } = useForm({
        email: ''
    })
    const { registerIn } = useContext(AuthContext);
    const onRegisterBotton = async () => {
        console.log('register')
        registerIn(email)

    }
    return (
        <SafeAreaView style={styleRegisterScreen.fondo}>
            <View style={{ alignSelf: 'center', justifyContent: 'center', top: '4%' }}>
                <Text style={styleRegisterScreen.title}>Register</Text>
            </View>

            <StakeHolders top="15%" texto="Your Email" keyboardType='email-address' stakeHold={onChange} stakeHoldText={email} valueText='email' color='#23232B' />

            <BottonRegister onPress={() => onRegisterBotton()} />

            <BottonsPlatforms bottom='39%' />
        </SafeAreaView>
    )
};

export default RegisterScreen;

export const styleRegisterScreen = StyleSheet.create({
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
