

import { useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { useContext } from 'react';
import { ActivityIndicator, Text, TouchableOpacity, View } from 'react-native';
import { StyleSheet } from 'react-native'
import { StretchyScrollView } from 'react-native-stretchy';
import { serveDefaultImages } from '../../api/apiCalls';
import BlackInfoBottons from '../../components/perfilComponets/BlackInfoBottons';
import FeaturesBox from '../../components/perfilComponets/FeaturesBox';
import { AuthContext } from '../../context/AuthContext';
import { useUser } from '../../hooks/useUser';
import { User } from '../../interfaces/appInterfaces';
import { RootStackParamList } from '../../navigator/StackNavigator';
import { styleBackgrounds } from '../../themes/Backgrounds';


type PropsScroll = {
    user?: User,

}
const SystretchyScrollView = ({ user }: PropsScroll) => {
    const navigation = useNavigation<StackNavigationProp<RootStackParamList>>();
    let price = user!.profession!.price == null ? "_" : "S/" + user!.profession!.price
    let tamano = user!.features!.height == null ? "-" : user!.features!.height
    let age = user!.features!.age == null ? "-" : user!.features!.age
    let gender = user!.features!.gender == null ? "alien" : user!.features!.gender
    let foto = user!.perfil_image == null ? serveDefaultImages + "noPerfil.png" : serveDefaultImages + user!.perfil_image
    let description = user!.profession!.description == null ? "si no sabes donde ir no puedes estar perdido" : user!.profession!.description

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
                        <TouchableOpacity style={styleViews.perfilPercents} onPress={() => { navigation.navigate('EditPerfilScreen') }}>

                        </TouchableOpacity>
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
    const { userId } = useContext(AuthContext)
    const { user, isLoading } = useUser(userId)
    return (

        isLoading
            ?
            <View>
                <ActivityIndicator size={50} />
            </View>
            : <View style={styleBackgrounds.fondoDark}>
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
        borderRadius: 100,
        width: 45,
        height: 45,
        borderColor: '#E5E1F6',
        borderWidth: 1.5,
        alignSelf: 'center'
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
