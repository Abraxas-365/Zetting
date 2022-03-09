import React from 'react';
import { createStackNavigator } from '@react-navigation/stack';
import LoginScreen from '../screens/LoginScreen';
import SignInScreen from '../screens/SignInScreen';
import RegisterScreen from '../screens/RegisterScreen';
import Register2Screen from '../screens/Register2Screen';
import HomeScreen from '../screens/HomeScreen';
import BottomTabs from './BottonTabs';
import ChatScreen from '../screens/ChatScreen';
import ProyectoScreen from '../screens/ProyectScreen';
import CreateProyectScreen from '../screens/CreateProyectScreen';
import BuildTeamScreen from '../screens/BuildTeamScreen';
import ActorsScreen from '../screens/ActorsScreen';

export type RootStackParamList = {
    LoginScreen: undefined;
    SignInScreen: undefined;
    Register2Screen: undefined;
    RegisterScreen: undefined;
    HomeScreen: undefined;
    BottomTabs: undefined;
    ChatScreen: undefined;
    ProyectoScreen: { name: string, description: string, };
    CreateProyectScreen: undefined;
    BuildTeamScreen: undefined;
    ActorsScreen: undefined;
};


const Stack = createStackNavigator();

export const StackNavigator = () => {
    return (
        <Stack.Navigator
            // initialRouteName="Pagina2Screen"
            screenOptions={{
                headerShown: false,
                headerStyle: {
                    elevation: 0,
                    shadowColor: 'transparent'
                },
                cardStyle: {
                    backgroundColor: 'white'
                }
            }}
        >
            <Stack.Screen name='LoginScreen' options={{ title: "LoginScreen" }} component={LoginScreen} />
            <Stack.Screen name="SignInScreen" options={{ title: "SignInScreen" }} component={SignInScreen} />
            <Stack.Screen name="RegisterScreen" options={{ title: "RegisterScreen" }} component={RegisterScreen} />
            <Stack.Screen name="Register2Screen" options={{ title: "Register2Screen" }} component={Register2Screen} />
            <Stack.Screen name="HomeScreen" options={{ title: "HomeScreen" }} component={HomeScreen} />
            <Stack.Screen name="BottomTabs" options={{ title: "BottomTabs" }} component={BottomTabs} />
            <Stack.Screen name="ChatScreen" options={{ title: "ChatScreen" }} component={ChatScreen} />
            <Stack.Screen name="ProyectoScreen" options={{ title: "ProyectoScreen" }} component={ProyectoScreen} />
            <Stack.Screen name="CreateProyectScreen" options={{ title: "CreateProyectScreen" }} component={CreateProyectScreen} />
            <Stack.Screen name="BuildTeamScreen" options={{ title: "BuildTeamScreen" }} component={BuildTeamScreen} />
            <Stack.Screen name="ActorsScreen" options={{ title: "ActorsScreen" }} component={ActorsScreen} />
        </Stack.Navigator>
    );
}
