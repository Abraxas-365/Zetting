import React from 'react';
import { StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import Switch from 'react-native-switch-toggles';
import { FontAwesome5 } from '@expo/vector-icons';
import { useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '../navigator/StackNavigator';
import MembersPhotos from './MembersPhotos';
import { Feather } from '@expo/vector-icons';

type Props = {
    title?: string
    color?: string
    members?: Array<string>
    notifications?: boolean
}

const ProjectBotton = ({ title = 'Titulo', color = '#FF7F39' }: Props) => {
    const navigation = useNavigation<StackNavigationProp<RootStackParamList>>();
    const [isEnabled, setIsEnabled] = React.useState(true);
    return (
        <TouchableOpacity style={{ ...styles.boton, backgroundColor: color }} onPress={() => navigation.navigate('ProyectoScreen')}>
            <View style={styles.wrapper}>
                <View style={styles.viewLeft}>
                    <Text style={styles.title}>{title}</Text>
                    <View style={styles.members}>
                        <MembersPhotos />
                    </View>
                </View>
                <View style={styles.viewRight}>
                    <View style={styles.switchMoon}>
                        <Switch
                            size={33}
                            value={isEnabled}
                            onChange={(value) => setIsEnabled(value)}
                            activeTrackColor={'#6ab04c'}
                            renderInactiveThumbIcon={() => (
                                <Feather name="bell-off" size={20} color="black" />)}
                            renderActiveThumbIcon={() => (
                                <Feather name="bell" size={20} color="black" />)}
                        />
                    </View>
                    <TouchableOpacity style={styles.chat} onPress={() => navigation.navigate('ChatScreen')}>
                        <Text style={styles.chatText}>Chat</Text>
                    </TouchableOpacity>
                </View>

            </View>
        </TouchableOpacity>
    );
};

export default ProjectBotton;

const styles = StyleSheet.create({
    boton: {
        alignSelf: 'center',
        width: '83%',
        height: 138,
        marginBottom: 20,
        // backgroundColor: '#FF7F39',
        borderRadius: 9,

    },
    title: {
        fontSize: 20,
        color: '#E5E1F6',
        fontWeight: 'bold',
    },
    members: {
        position: 'absolute',
        bottom: 0,
        width: '100%',
        height: '50%',
        marginLeft: '7%'

    },
    chat: {
        borderColor: '#E5E1F6',
        borderWidth: 1.5,
        borderRadius: 40,
        position: 'absolute',
        width: 64,
        height: 26,
        right: 0,
        bottom: 0,

    },

    chatText: {
        textAlign: 'center',
        fontSize: 15,
        color: 'white',
        fontWeight: 'bold',
        padding: '3%'
    },
    switchMoon: {
        position: 'absolute',
        right: 0,
    },
    viewLeft: {
        width: '50%',
        flexDirection: 'column',

    },
    viewRight: {
        width: '50%',
        flexDirection: 'column',

    },
    wrapper: {
        flex: 1,
        flexDirection: 'row',
        marginVertical: '7%',
        marginHorizontal: '7%',
    }

})
