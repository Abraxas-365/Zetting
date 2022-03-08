import React, { useContext, useState } from 'react';
import { Pressable, ScrollView, StyleSheet, Text, TouchableOpacity, View } from 'react-native';
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
import { Entypo } from '@expo/vector-icons';
const creteProject = async (project: Project, navigation: any, token: any) => {
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };
    console.log(config);

    try {
        console.log("el color es", project.color);
        const { data } = await apiCalls.post('/api/projects/new', project, config);

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


    const [checked, onChake] = useState(false);
    const MyColorbox = () => {

        function onCheckmarkPress() {
            onChake(!checked);
        }

        return (
            <Pressable
                style={[styles.checkboxBase, checked && styles.checkboxChecked]}
                onPress={onCheckmarkPress}>
                {checked && <Entypo name="check" size={15} color='#23232B' />}
            </Pressable>
        );
    }

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
            <View style={{
                borderBottomColor: '#E5E1F6',
                borderBottomWidth: 1,
                marginTop: '15%',
                height: '15%',
                marginHorizontal: '10%',
            }}>
                <Text style={{ color: '#E5E1F6' }}>Color</Text>

                <View style={{ flexDirection: 'row', width: '100%', height: '40%', position: 'absolute', bottom: 5 }}>
                    <MyColorbox />
                    <MyColorbox />
                </View>
            </View>
            <View style={styles.section}>
            </View>
            <View>
                <CustomBotton onPress={() => onCreatePtojectBotton()} text="Create Project" />
            </View>
        </ScrollView >
    );
};


export default CreateProjectForm
const styles = StyleSheet.create({
    checkboxBase: {
        width: 26,
        height: 26,
        justifyContent: 'center',
        alignItems: 'center',
        borderRadius: 100,
        borderWidth: 2,
        marginRight: '1%',
        borderColor: '#23232B',
        backgroundColor: 'red',
    },

    checkboxChecked: {
        backgroundColor: 'red',
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
