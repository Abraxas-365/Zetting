
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
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';

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
        <SafeAreaView style={styleBackgrounds.fondoLight}>
            <View style={{ alignSelf: 'center', justifyContent: 'center', top: '4%' }}>
                <Text style={styleTitles.titleTextDark}>REGISTER</Text>
            </View>

            <StakeHolders top="15%" texto="Your Email" keyboardType='email-address' stakeHold={onChange} stakeHoldText={email} valueText='email' color='#23232B' />

            <BottonRegister onPress={() => onRegisterBotton()} />

            <BottonsPlatforms bottom='39%' />
        </SafeAreaView>
    )
};

export default RegisterScreen;


