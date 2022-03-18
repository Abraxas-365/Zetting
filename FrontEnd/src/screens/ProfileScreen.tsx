

import React, { Component, useContext, useEffect } from 'react';
import { Button, Image, SafeAreaView, Text, View } from 'react-native';
import { StyleSheet } from 'react-native'
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';
import { StretchyScrollView } from 'react-native-stretchy';
import BlackInfoBottons from '../components/perfilComponets/BlackInfoBottons';
import FeaturesBox from '../components/perfilComponets/FeaturesBox';
import { AuthContext } from '../context/AuthContext';
import { User } from '../interfaces/appInterfaces';
import { serveDefaultImages } from '../api/apiCalls';


type PropsScroll = {
    user?: User | null,

}
const SystretchyScrollView = ({ user = {} }: PropsScroll) => {

    let price = user!.profession!.price == 0 ? "_" : "S/" + user!.profession!.price
    let tamano = user!.features!.height == 0 ? "-" : user!.features!.height
    let age = user!.features!.age == 0 ? "-" : user!.features!.age
    let gender = user!.features!.gender == "" ? "alien" : user!.features!.gender
    let foto = user!.perfil_image == "" ? serveDefaultImages + "noPerfil.png" : serveDefaultImages + user!.perfil_image
    let description = user!.profession!.description == "" ? "si no sabes donde ir no puedes estar perdido" : user!.profession!.description
    return (
        <StretchyScrollView
            image={{ uri: foto }}
        >

            <View style={styleBackgrounds.fondoDark}>
                <View style={styleViews.wrapper}>
                    <View style={styleViews.nameAndPerfil}>
                        <View style={styleViews.nameAndProfession}>
                            <Text style={styleText.name}>{user!.first_name} {user!.last_name}</Text>
                            <Text style={styleText.professions}>Actor</Text>
                        </View>
                        <View style={styleViews.perfilPercents}></View>
                    </View>
                    <View style={styleViews.infoBoxes}>
                        <BlackInfoBottons width={86} marginRight={'4%'} numbers={1} text={'Projects'} />
                        <BlackInfoBottons width={103} marginRight={'4%'} numbers={5} text={'Comments'} />
                        <BlackInfoBottons width={109} numbers={price} text={'Price'} />
                    </View>
                    <FeaturesBox age={age} gender={gender} tamano={tamano} />
                    <View style={styleViews.aboutBox}>
                        <Text style={styleText.smalTitles}>About</Text>
                        <Text style={styleText.normal}>{description}</Text>
                    </View>
                    <View style={styleViews.SkillsBox}>
                        <Text style={styleText.smalTitles}>Skills/Qualities</Text>
                    </View>
                    <View style={styleViews.ReelBox}>
                        <Text style={styleText.smalTitles}>Acting Reel</Text>
                    </View>
                    <View style={styleViews.workExpBox}>
                        <Text style={styleText.smalTitles}>Work Experience</Text>
                    </View>

                </View>
            </View>
        </StretchyScrollView>
    );
}

const ProfileScreen = () => {
    const { user } = useContext(AuthContext)
    return (
        <View style={styleBackgrounds.fondoDark}>
            <SystretchyScrollView user={user} />
        </View>
    );
};

export default ProfileScreen;



const styleViews = StyleSheet.create({
    wrapper: {
        marginTop: 20,
        marginHorizontal: '5%',
    },
    nameAndPerfil: {
        flexDirection: 'row',

    },
    nameAndProfession: {
        width: '80%',
        flexDirection: 'column',
    },

    perfilPercents: {

    },
    infoBoxes: {
        marginTop: 35,
        flexDirection: 'row'

    },
    aboutBox: {
        width: '70%',
        marginTop: 35

    },

    SkillsBox: {
        width: '70%',
        marginTop: 35

    },
    ReelBox: {
        width: '70%',
        marginTop: 35

    },
    workExpBox: {
        width: '70%',
        marginTop: 35

    },

})

const styleText = StyleSheet.create({
    normal: {

        color: '#e5e1f6',
        fontSize: 12,
        marginTop: 10,
        fontWeight: '500',
    },
    name: {
        fontSize: 18,
        fontWeight: 'bold',
        color: '#e5e1f6',
    },
    professions: {
        color: '#e5e1f6',
        fontSize: 12,
        marginTop: 10,
        fontWeight: '600',
    },

    smalTitles: {
        fontSize: 20,
        color: '#e5e1f6',
    },
})
