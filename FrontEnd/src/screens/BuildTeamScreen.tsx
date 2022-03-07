
import React from 'react';
import { SafeAreaView, ScrollView, StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import { styleBackgrounds } from '../themes/Backgrounds';

const OmitBottom = () => {
    return (
        <TouchableOpacity
            style={styles.boton}
            onPress={() => { }}>
            <Text style={styles.botonText}>Omitir</Text>
        </TouchableOpacity>


    )
}

const BuildTeamScreen = () => {
    return (
        <View style={styleBackgrounds.fondoOrange}>
            <View style={styles.topContainer}>
                <Text style={styles.title}>BUILD YOUR TEAM</Text>
                <View style={{ flexDirection: 'row', width: '100%' }}>
                    <Text style={styles.slogan}>The Sound of Music</Text>
                    <OmitBottom />
                </View>
            </View>
            <View style={{ ...styleBackgrounds.fondoDark, top: '13%' }}>
                <ScrollView style={{ ...styleBackgrounds.fondoDark, marginHorizontal: '10%', backgroundColor: 'white' }}>

                </ScrollView>
            </View>
        </View>
    );
};

export default BuildTeamScreen;

const styles = StyleSheet.create({
    title: {
        fontSize: 30,
        fontWeight: 'bold',
        color: '#E9E5F7',
    },
    topContainer: {
        top: '8%',
        marginHorizontal: '10%'
    },
    slogan: {
        fontSize: 20,
        marginTop: '5%',
        color: '#E9E5F7',
    },

    boton: {
        width: '25%',
        backgroundColor: '#E9E5F7',
        borderRadius: 40,
        justifyContent: 'center',
        position: 'absolute',
        right: 0,
        bottom: 0
    },
    botonText: {
        textAlign: 'center',
        padding: 5,
        fontSize: 12,
        color: '#FE893C',
        fontWeight: 'bold',
    }

})
