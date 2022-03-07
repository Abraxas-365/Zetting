import React, { useContext, useState } from 'react';
import { Pressable, ScrollView, StyleSheet, Text, View } from 'react-native';
import StakeHolders from './StakeHolders';
import { Ionicons } from '@expo/vector-icons';
import MyCheckbox from './CheckBox';
import BottonSignUp from './SignUpButton';
import { useForm } from '../hooks/useForm';
import { AuthContext } from '../context/AuthContext';
import { Project, User } from '../interfaces/appInterfaces';
import CustomBotton from './CustomBotton';
import { apiCalls } from '../api/apiCalls';
import { CommonActions, useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '../navigator/StackNavigator';

const creteProject = async (project: Project, navigation: any, token: any) => {
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };
    console.log(config);

    try {
        console.log("el color es", project.color);
        const { data } = await apiCalls.post('http://192.168.1.56:3000/api/projects/new', project, config);

        console.log(data.data);
        navigation.dispatch(
            CommonActions.reset({
                index: 0,
                routes: [{ name: 'BottomTabs' }
                ]

            }))

    } catch (err) {
        console.error(err);
    }


}


const CreateProjectForm = () => {



    const { token } = useContext(AuthContext)
    const navigation = useNavigation<StackNavigationProp<RootStackParamList>>();
    const { name, description, calendar, collaboration, color, onChange } = useForm({

        name: '',
        description: '',
        calendar: '',
        collaboration: '',
        color: '',
    })
    console.log(name)
    const onCreatePtojectBotton = async () => {
        console.log('creando proyecto', name)
        var project: Project = {
            name: name,
            description: description,
            // collaboration: collaboration,
            calendar: calendar,
            color: color,

        }

        creteProject(project, navigation, token)
    }
    return (

        <ScrollView style={{ top: '8%', flex: 1 }} contentContainerStyle={{ paddingBottom: 90 }} >

            <StakeHolders texto='Project Name' stakeHold={onChange} stakeHoldText={name} valueText='name' color='#E5E1F6' />
            <StakeHolders texto='Project description' stakeHold={onChange} stakeHoldText={description} valueText='description' color='#E5E1F6' />
            <StakeHolders texto='Colaboration' stakeHold={onChange} stakeHoldText={collaboration} valueText='collaboration' color='#E5E1F6' />
            <StakeHolders texto='Color' stakeHold={onChange} stakeHoldText={color} valueText='color' color='#E5E1F6' />
            <View style={styles.section}>
                <MyCheckbox />
            </View>
            <View>
                <CustomBotton onPress={() => onCreatePtojectBotton()} text="Create Project" />
            </View>
        </ScrollView>
    );
};

export default CreateProjectForm
const styles = StyleSheet.create({
    checkboxBase: {
        width: 24,
        height: 24,
        justifyContent: 'center',
        alignItems: 'center',
        borderRadius: 4,
        borderWidth: 2,
        borderColor: 'coral',
        backgroundColor: 'transparent',
    },

    checkboxChecked: {
        backgroundColor: 'coral',
    },
    section: {
        width: '100%',
        alignSelf: 'center',
        justifyContent: 'center',
        marginTop: '15%',
        flexDirection: 'row',
        alignItems: 'center',
    },
    paragraph: {
        color: '#23232B',
        fontSize: 10,
    },

    paragraphOrange: {
        color: '#FF7F39',
        fontSize: 10,
    },
    checkbox: {
        borderRadius: 5,
        margin: 8,
        color: 'red'
    },
});
