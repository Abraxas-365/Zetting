
import React from 'react';
import { Image, StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import Switch from 'react-native-switch-toggles';
import { useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '../navigator/StackNavigator';
import MembersPhotos from './MembersPhotos';
import { Feather } from '@expo/vector-icons';

type Props = {
    name?: string
    lastName?: string
    description?: string
    color?: string
    members?: Array<string>
    notifications?: boolean
}

const TalentsButtonCard = ({ name = 'name', lastName = 'apellido', description = 'description', color = '#FF7F39' }: Props) => {
    const navigation = useNavigation<StackNavigationProp<RootStackParamList>>();
    return (
        <TouchableOpacity style={{ ...styles.boton, backgroundColor: color }} onPress={() => navigation.navigate('ProyectoScreen', { name: name, description: description })}>
            <View style={styles.wrapper}>
                <View style={{ ...styles.viewLeft, overflow: 'hidden' }}>
                    <Image style={{ flex: 1, overflow: 'hidden', borderTopLeftRadius: 9, borderBottomLeftRadius: 9 }} source={{ uri: 'https://i.pinimg.com/736x/8a/52/d3/8a52d3863a272e2e0556861ba98dceb1.jpg' }} />
                </View>
                <View style={styles.viewRight}>
                    <Text style={{ ...styles.title, marginTop: '8%', marginHorizontal: '8%' }}>{name} {lastName}</Text>
                    <Text style={{ ...styles.description, marginTop: '8%', marginHorizontal: '8%' }}>{description}</Text>
                    <View style={{ position: 'absolute', bottom: 10, marginHorizontal: '8%', flexDirection: 'row' }}>
                        <Text style={{ ...styles.preciou }}>S/120</Text>

                        <Text style={{ fontSize: 12, bottom: -8, marginHorizontal: '8%', color: '#E5E1F6' }}>xH</Text>
                    </View>
                </View>

            </View>
        </TouchableOpacity>
    );
};

export default TalentsButtonCard;

const styles = StyleSheet.create({
    boton: {
        height: 138,
        marginBottom: 20,
        borderRadius: 9,

    },
    title: {
        fontSize: 18,
        color: '#E5E1F6',
        fontWeight: "500",
    },
    description: {
        fontSize: 12,
        color: '#E5E1F6'
    },
    viewLeft: {
        width: '38%',
        flexDirection: 'column',

    },
    viewRight: {
        width: '50%',
        flexDirection: 'column',

    },
    wrapper: {
        flex: 1,
        flexDirection: 'row',
    },
    preciou: {
        fontSize: 20,
        color: '#E5E1F6',
    }

})
