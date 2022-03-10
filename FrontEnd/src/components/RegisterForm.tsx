
import React, { useContext, useState } from 'react';
import { Pressable, ScrollView, StyleSheet, Text, View } from 'react-native';
import StakeHolders from './StakeHolders';
import MyCheckbox from './CheckBox';
import { useForm } from '../hooks/useForm';
import { AuthContext } from '../context/AuthContext';
import { User } from '../interfaces/appInterfaces';
import CustomCenterBotton from './CustomCenterBotton';

interface FormProps {
    email: string
}

const RegisterForm = ({ email }: FormProps) => {
    console.log('---RegisterForm---', email);

    //aqui acumular los botones de profesiones
    var professions: string[] = ["hola"]
    const { first_name, last_name, country, phone, identifierDocument, password, onChange } = useForm({

        first_name: '',
        last_name: '',
        phone: '',
        country: '',
        identifierDocument: '',
        password: '',
    })
    const { signUp } = useContext(AuthContext);
    const onSingUpBotton = async () => {
        console.log('register')
        var user: User = {
            first_name: first_name,
            last_name: last_name,
            email: email,
            phone: phone,
            professions: professions,
            country: country,
            identifierDocument: '',
            password: password,
            verified: false

        }
        signUp(user)

    }
    return (

        <ScrollView style={{ top: '8%', flex: 1 }} contentContainerStyle={{ paddingBottom: 90 }} >

            <StakeHolders texto='First Name' top='0%' stakeHold={onChange} stakeHoldText={first_name} valueText='first_name' color='#23232B' />
            <StakeHolders texto='Last Name' stakeHold={onChange} stakeHoldText={last_name} valueText='last_name' color='#23232B' />
            <StakeHolders texto='Phone Numbe' stakeHold={onChange} stakeHoldText={phone} valueText='phone' color='#23232B' />
            <StakeHolders texto='Country' stakeHold={onChange} stakeHoldText={country} valueText='country' color='#23232B' />
            <StakeHolders texto='Password' stakeHold={onChange} stakeHoldText={password} valueText='password' color='#23232B' />
            <View style={styles.section}>
                <MyCheckbox />
                <View style={{ width: '60%', flexDirection: 'row' }}>
                    <Text style={styles.paragraph}>I have read and agree to the
                        <Text style={styles.paragraphOrange}> Privacy Policy</Text>
                        <Text style={styles.paragraph}>{'\n'}Our</Text><Text style={styles.paragraphOrange}> Terms</Text>
                        <Text style={styles.paragraph}> and </Text><Text style={styles.paragraphOrange}> Conditions</Text>
                    </Text>
                </View>
            </View>
            <CustomCenterBotton onPress={() => onSingUpBotton()} text='Register' top='8%' />
        </ScrollView>
    );
};

export default RegisterForm
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
