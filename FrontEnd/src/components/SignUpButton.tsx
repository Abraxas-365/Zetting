

import React from 'react';
import { StyleSheet, Text, TouchableOpacity, View } from 'react-native'


type BottonProps = {
    onPress: any
}

const BottonSignUp = ({ onPress }: BottonProps) => {
    return (
        <TouchableOpacity
            style={style.boton}
            onPress={onPress}>
            <Text style={style.botonText}>Sign Up</Text>
        </TouchableOpacity>
    );
};

const style = StyleSheet.create({
    boton: {
        alignSelf: 'center',
        width: '50%',
        position: 'relative',
        backgroundColor: '#FF7F39',
        borderRadius: 40,
        justifyContent: 'center',
        marginTop: '10%',
    },
    botonText: {
        textAlign: 'center',
        padding: 10,
        fontSize: 16,
        color: 'white',
        fontWeight: 'bold',
    }

})
export default BottonSignUp;
