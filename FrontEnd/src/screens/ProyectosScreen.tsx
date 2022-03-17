

import { useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { useState } from 'react';
import { Button, FlatList, Image, SafeAreaView, ScrollView, Text, TouchableOpacity, View } from 'react-native';
import { StyleSheet } from 'react-native'
import styles from 'react-native-switch-toggles/src/switch/styles';
import CustomSmallBotton from '../components/CustomSmallBotton';
import PlusButton from '../components/PlusSvg';
import ProjectBotton from '../components/ProjectBotton';
import { useProjects } from '../hooks/useProjects';
import { RootStackParamList } from '../navigator/StackNavigator';
import { styleBackgrounds } from '../themes/Backgrounds';
import { styleTitles } from '../themes/Titles';
import { styleWrappers } from '../themes/Wrappers';





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
            <View style={{ ...styleWrappers.wrapperHorizontalGap }}>
                <View style={{ ...styleWrappers.wrapperTitles, flexDirection: 'row' }}>
                    <Text style={styleTitles.titleTextLight}>RROJECTS</Text>
                    <TouchableOpacity style={styleProyectosScreen.mas} onPress={() => navigation.navigate('CreateProyectScreen')}>
                        <PlusButton />
                    </TouchableOpacity>
                </View>
                <View style={{
                    marginTop: '10%', flexDirection: 'row',
                }}>


                    <CustomSmallBotton borderColor={borderColor} color={colorProject} text='My Project' onPress={projectsBotton} />
                    <View style={{ marginHorizontal: '2%', marginBottom: "12%" }} />
                    <CustomSmallBotton borderColor={borderColor1} color={colorMyProject} text='Projects' onPress={myprojectsBotton} />
                </View>

            </View>
            <FlatList contentContainerStyle={{ flexGrow: 1, top: '3%', ...styleWrappers.wrapperHorizontalGap }} data={data} renderItem={({ item }: any) => <ProjectBotton title={item.name} description={item.description} color={item.color} />} />

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
