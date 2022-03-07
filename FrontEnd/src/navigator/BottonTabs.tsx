import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import FavoritosScreen from '../screens/FavoritosScreen';
import HomeScreen from '../screens/HomeScreen';
import ProfileScreen from '../screens/ProfileScreen';
import ProyetosScreen from '../screens/ProyectosScreen';

const Tab = createBottomTabNavigator();

const BottomTabs = () => {
    return (
        <Tab.Navigator
            screenOptions={{
                headerShown: false,
                tabBarStyle: { backgroundColor: '#23232B', borderTopColor: '#E5E1F6' }
            }}  >
            <Tab.Screen name="Z" component={ProyetosScreen} />
            <Tab.Screen name="H" component={HomeScreen} />
            <Tab.Screen name="P" component={ProfileScreen} />
            <Tab.Screen name="F" component={FavoritosScreen} />
        </Tab.Navigator>
    );
}

export default BottomTabs;
