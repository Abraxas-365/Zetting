

import { useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React from 'react';
import { Button, FlatList, Image, SafeAreaView, ScrollView, Text, TouchableOpacity, View } from 'react-native';
import { StyleSheet } from 'react-native'
import BottonsLoginScreen from '../components/BottonsLoginScreen';
import ProjectBotton from '../components/ProjectBotton';
import { useProjects } from '../hooks/useProjects';
import { RootStackParamList } from '../navigator/StackNavigator';


const ProyetosScreen = () => {

    const navigation = useNavigation<StackNavigationProp<RootStackParamList>>();
    const { projects } = useProjects();
    console.log(projects)
    return (
        <SafeAreaView style={styleProyectosScreen.fondo}>
            <View style={{ top: '4%', flexDirection: 'row', marginHorizontal: '10%' }}>
                <Text style={styleProyectosScreen.title}>RROJECTS</Text>
                <TouchableOpacity style={styleProyectosScreen.mas} onPress={() => navigation.navigate('CreateProyectScreen')}>
                    <Image source={require('../../assets/icons/2x/mas.png')} />
                </TouchableOpacity>
            </View>
            <View style={{ marginTop: "10%" }}>
                <FlatList
                    contentContainerStyle={{ flexGrow: 1 }}
                    data={projects}
                    renderItem={({ item }: any) => <ProjectBotton title={item.name} color={item.color} />}

                />
            </View>

        </SafeAreaView>
    );
};

export default ProyetosScreen;



export const styleProyectosScreen = StyleSheet.create({
    fondo: {
        flex: 1,
        backgroundColor: '#23232B',
    },
    slogan: {
        fontSize: 40,
        color: '#E5E1F6',
        fontWeight: 'bold'
    },
    title: {
        fontSize: 30,
        color: '#E5E1F6',
        fontWeight: 'bold',
    },
    mas: {
        position: 'absolute',
        right: 0,

    }


})
