
import React from 'react';
import { StyleSheet, Text, TouchableOpacity, View } from 'react-native'



type BottonProps = {
    onPress: any
    top?: string
    color?: string
    text?: string
}
const CustomCenterBotton = ({ onPress = () => { }, color = '#FF7F39', top = '13%', text = "texto" }: BottonProps) => {
    return (
        <TouchableOpacity
            style={{ ...style.boton, backgroundColor: color, marginTop: top }}
            onPress={onPress}>
            <Text style={style.botonText}>{text}</Text>
        </TouchableOpacity>
    );
};

const style = StyleSheet.create({
    boton: {
        alignSelf: 'center',
        width: '50%',
        backgroundColor: '#FF7F39',
        borderRadius: 40,
        justifyContent: 'center',
        marginTop: '13%',
    },
    botonText: {
        textAlign: 'center',
        padding: 10,
        fontSize: 16,
        color: '#E5E1F6',
        fontWeight: 'bold',
    }

})
export default CustomCenterBotton;
