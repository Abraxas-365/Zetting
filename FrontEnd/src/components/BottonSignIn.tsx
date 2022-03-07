
import React from 'react';
import { StyleSheet, Text, TouchableOpacity, View } from 'react-native'



type BottonProps = {
    onPress: any
}
const BottonSignIn = ({ onPress }: BottonProps) => {
    return (
        <TouchableOpacity
            style={style.boton}
            onPress={onPress}>
            <Text style={style.botonText}>Sign in</Text>
        </TouchableOpacity>
    );
};

const style = StyleSheet.create({
    boton: {
        alignSelf: 'center',
        width: '50%',
        position: 'absolute',
        backgroundColor: '#FF7F39',
        borderRadius: 40,
        justifyContent: 'center',
        bottom: '57%',
    },
    botonText: {
        textAlign: 'center',
        padding: 10,
        fontSize: 16,
        color: 'white',
        fontWeight: 'bold',
    }

})
export default BottonSignIn;
