

import { useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { useState } from 'react';
import { Button, FlatList, Image, SafeAreaView, ScrollView, Text, TouchableOpacity, View } from 'react-native';
import { StyleSheet } from 'react-native'
import BottonsLoginScreen from '../components/BottonsLoginScreen';
import CustomSmallBotton from '../components/CustomSmallBotton';
import ProjectBotton from '../components/ProjectBotton';
import { useProjects } from '../hooks/useProjects';
import { RootStackParamList } from '../navigator/StackNavigator';
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';


const ProyetosScreen = () => {

    const navigation = useNavigation<StackNavigationProp<RootStackParamList>>();
    //naranja
    const [colorProject, setColorProject] = useState('#FF7F39');
    //gis
    const [colorMyProject, setColorMyProject] = useState('#23232B');
    //blaanco
    const [borderColor, setBorderColor] = useState('#FF7F39');
    const [borderColor1, setBorderColor1] = useState('#E5E1F6');
    const [tipo, setTipo] = useState(false);

    const { projects, myProjects } = useProjects();
    var data = myProjects
    if (tipo == false) {
        data = myProjects
    } else {
        data = projects
    }
    const projectsBotton = () => {

        setColorProject('#FF7F39')
        setBorderColor('#FF7F39')
        setColorMyProject('#23232B')
        setBorderColor1('#E5E1F6')
        setTipo(false)

    }
    const myprojectsBotton = () => {
        setColorProject('#23232B')
        setBorderColor('#E5E1F6')
        setColorMyProject('#FF7F39')
        setBorderColor1('#FF7F39')
        setTipo(true)
    }

    return (
        <SafeAreaView style={styleBackgrounds.fondoDark}>
            <View style={{ top: '4%', flexDirection: 'row', marginHorizontal: '10%' }}>
                <Text style={styleTitles.titleTextLight}>RROJECTS</Text>
                <TouchableOpacity style={styleProyectosScreen.mas} onPress={() => navigation.navigate('CreateProyectScreen')}>
                    <Image source={require('../../assets/icons/2x/mas.png')} />
                </TouchableOpacity>
            </View>
            <View style={{ marginTop: '13%', flexDirection: 'row', marginLeft: '10%' }}>

                <CustomSmallBotton borderColor={borderColor} color={colorProject} text='My Project' onPress={projectsBotton} />
                <View style={{ marginHorizontal: '2%' }} />
                <CustomSmallBotton borderColor={borderColor1} color={colorMyProject} text='Projects' onPress={myprojectsBotton} />
            </View>
            <View style={{ marginTop: "10%" }}>
                <FlatList contentContainerStyle={{ flexGrow: 1 }} data={data} renderItem={({ item }: any) => <ProjectBotton title={item.name} color={item.color} />} />
            </View>

        </SafeAreaView >
    );
};

export default ProyetosScreen;



export const styleProyectosScreen = StyleSheet.create({
    mas: {
        position: 'absolute',
        right: 0,

    }


})
