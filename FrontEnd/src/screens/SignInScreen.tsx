
import { useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { useContext, useEffect } from 'react';
import { Alert, SafeAreaView, StyleSheet, Text, View } from "react-native";
import { TextInput } from 'react-native-gesture-handler';
import BottonsPlatforms from '../components/BottonPlatform';
import BottonSignIn from '../components/BottonSignIn';
import StakeHolders from '../components/StakeHolders';
import { AuthContext } from '../context/AuthContext';
import { useForm } from '../hooks/useForm';
import { onLoginBotton, UseLogin } from '../hooks/useLogin';
import { RootStackParamList } from '../navigator/StackNavigator';

const SignInScreen = () => {

    const { email, password, onChange } = useForm({
        email: '',
        password: '',
    })

    const { signIn, errorMessage, removeError } = useContext(AuthContext);
    useEffect(() => {
        if (errorMessage == null) return;
        Alert.alert('Login incorrecto', errorMessage, [{ text: 'Ok', onPress: removeError }])
    }, [errorMessage])

    const onLoginBotton = async (email: string, password: string) => {
        console.log('register')

        signIn(email, password)

    }

    return (
        <SafeAreaView style={styleSignInScreen.fondo}>
            <View style={{ alignSelf: 'center', justifyContent: 'center', top: '4%' }}>
                <Text style={styleSignInScreen.title}>LOG IN</Text>
            </View>

            <StakeHolders top="15%" texto="Your Email" keyboardType='email-address' stakeHold={onChange} stakeHoldText={email} valueText='email' color='#23232B' />
            <StakeHolders top='15%' texto="Password" stakeHold={onChange} stakeHoldText={password} valueText='password' secureTextEntry={true} color='#23232B' />

            <BottonSignIn onPress={() => onLoginBotton(email, password)} />

            <BottonsPlatforms bottom='23%' />
        </SafeAreaView>
    )
};

export default SignInScreen;

export const styleSignInScreen = StyleSheet.create({
    fondo: {
        flex: 1,
        backgroundColor: '#E2E0F3',
    },
    title: {
        fontSize: 30,
        color: '#23232B',
        fontWeight: 'bold'
    },

})
