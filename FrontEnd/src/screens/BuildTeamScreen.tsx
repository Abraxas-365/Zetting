
import React from 'react';
import { Image, SafeAreaView, ScrollView, StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';

const OmitBottom = () => {
    return (
        <TouchableOpacity
            style={styles.boton}
            onPress={() => { }}>
            <Text style={styles.botonText}>Omitir</Text>
        </TouchableOpacity>


    )
}

type BottonProps = {
    onPress?: any
    text?: string
    image?: string
}
const CategoriesBotton = ({ onPress = () => { }, text = 'test', image = 'https://media.istockphoto.com/photos/director-chairmovie-clapper-in-studioconcept-for-film-industry3d-picture-id1305968604?s=612x612' }: BottonProps) => {
    return (
        <TouchableOpacity
            style={{ width: 140, height: 92, borderRadius: 9 }}
            onPress={onPress}>
            <Image style={{ flex: 1, borderRadius: 9 }} resizeMode='cover' source={{ uri: image }} />
            <Text style={{ position: 'absolute', fontSize: 13, color: '#E5E1F6', fontWeight: 'bold', left: 10, top: 10 }}>{text}</Text>
        </TouchableOpacity>
    )

}

const CategoriesBottons = () => {
    return (
        <View>
            <View style={{ flexDirection: 'row', marginTop: 20, position: 'relative' }}>
                <View style={{ position: 'relative', left: 0, }}>
                    <CategoriesBotton />
                </View>
                <View style={{ position: 'absolute', right: 0 }}>
                    <CategoriesBotton />
                </View>
            </View>

        </View>
    )
}

const BuildTeamScreen = () => {
    return (
        <View style={styleBackgrounds.fondoOrange}>
            <View style={styles.topContainer}>
                <Text style={styleTitles.titleTextLight}>BUILD YOUR TEAM</Text>
                <View style={{ flexDirection: 'row', width: '100%' }}>
                    <Text style={styles.slogan}>The Sound of Music</Text>
                    <OmitBottom />
                </View>
            </View>
            <View style={{ ...styleBackgrounds.fondoDark, top: '13%' }}>
                <ScrollView style={{ ...styleBackgrounds.fondoDark, marginHorizontal: '10%' }}>
                    <CategoriesBottons />
                    <CategoriesBottons />
                    <CategoriesBottons />
                    <CategoriesBottons />
                    <CategoriesBottons />
                </ScrollView>
            </View>
        </View>
    );
};

export default BuildTeamScreen;

const styles = StyleSheet.create({
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
