import { StackScreenProps } from '@react-navigation/stack';
import React, { Component } from 'react';
import { Button, Image, SafeAreaView, Text, View } from 'react-native';
import { StyleSheet } from 'react-native'
import { StretchyScrollView } from 'react-native-stretchy';
import BlackInfoBottons from '../../components/perfilComponets/BlackInfoBottons';
import FeaturesBox from '../../components/perfilComponets/FeaturesBox';
import { RootStackParamList } from '../../navigator/StackNavigator';
import { styleBackgrounds } from '../../themes/Backgrounds';

class SystretchyScrollView extends Component {
    render() {
        return (
            <StretchyScrollView
                image={{ uri: 'https://dvyvvujm9h0uq.cloudfront.net/com/articles/1525891879-886386-sam-burriss-457746-unsplashjpg.jpg' }}
            >

                <View style={styleBackgrounds.fondoDark}>
                    <View style={styleViews.wrapper}>
                        <View style={styleViews.nameAndPerfil}>
                            <View style={styleViews.nameAndProfession}>
                                <Text style={styleText.name}>LUIS FERNANDO MIRANDA</Text>
                                <Text style={styleText.professions}>ACTOR</Text>
                            </View>
                            <View style={styleViews.perfilPercents}></View>
                        </View>
                        <View style={styleViews.infoBoxes}>
                            <BlackInfoBottons width={86} marginRight={'4%'} numbers={1} text={'Projects'} />
                            <BlackInfoBottons width={103} marginRight={'4%'} numbers={5} text={'Comments'} />
                            <BlackInfoBottons width={109} numbers={'_'} text={'Price'} />
                        </View>
                        <FeaturesBox age={22} gender={"alien"} tamano={180} />
                        <View style={styleViews.aboutBox}>
                            <Text style={styleText.smalTitles}>About</Text>
                            <Text style={styleText.normal}>osdfjhfskhfgousdhgdnfsakjfnisjdfhosguih</Text>
                        </View>
                        <View style={styleViews.SkillsBox}>
                            <Text style={styleText.smalTitles}>About</Text>
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
}

interface Props extends StackScreenProps<RootStackParamList, 'WorkerActorScreen'> { }
const WorkerActorScreen = () => {
    return (
        <View style={styleBackgrounds.fondoDark}>
            <SystretchyScrollView />
        </View>
    );
};

export default WorkerActorScreen;



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
