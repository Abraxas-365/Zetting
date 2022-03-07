import { CommonActions } from '@react-navigation/native';
import React, { useEffect } from 'react';
import { View } from 'react-native';

export const UseLogin = () => {
    const [email, setEmail] = React.useState("");
    const [password, setPassword] = React.useState("");
    //aqui hacer la logica para el login
    return { setEmail, setPassword, email, password }

};


export function onLoginBotton(navigation: any) {

    console.log('onSigInpBotton')
    navigation.dispatch(
        CommonActions.reset({
            index: 0,
            routes: [{ name: 'BottomTabs' }
            ]

        }))
}
