import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { BellSvgActive, BellSvgInactive } from '../components/BellSvg';
import HeardSvg from '../components/HeardSvg';
import SearchSvg from '../components/SearchSvg';
import FavoritosScreen from '../screens/FavoritosScreen';
import HomeScreen from '../screens/HomeScreen';
import ProfileScreen from '../screens/ProfileScreen';
import ProyetosScreen from '../screens/ProyectosScreen';
import Svg, { ClipPath, Defs, G, Path } from "react-native-svg"
const Tab = createBottomTabNavigator();
const Bell = (props) => (
    <Svg
        xmlns="http://www.w3.org/2000/svg"
        width={16.766}
        height={18.427}
        {...props}
    >
        <G data-name="Group 28" fill="#e5e1f6">
            <Path
                data-name="Path 13"
                d="M15.218 8.582c-.193-.923-.414-1.97-.534-3.006a6.349 6.349 0 0 0-12.6 0c-.118 1.021-.33 2.031-.531 2.994l-1.55 6.776h16.763ZM1.766 13.937l1.159-5.066c.2-.961.429-2.049.554-3.133a4.942 4.942 0 0 1 9.807 0c.126 1.1.353 2.18.557 3.145l1.155 5.054Z"
            />
            <Path
                data-name="Path 14"
                d="M11.563 17.019H5.2a.704.704 0 0 0 0 1.408h6.363a.704.704 0 0 0 0-1.408"
            />
        </G>
    </Svg>
)
const Profile = (props) => (
    <Svg
        xmlns="http://www.w3.org/2000/svg"
        width={21.292}
        height={21.292}
        {...props}
    >
        <Defs>
            <ClipPath id="a">
                <Path data-name="Rectangle 77" fill="none" d="M0 0h21.292v21.292H0z" />
            </ClipPath>
        </Defs>
        <G data-name="Group 56">
            <G data-name="Group 55" clipPath="url(#a)">
                <Path
                    data-name="Path 43"
                    d="M21.292 10.644a10.646 10.646 0 0 0-21.292 0 10.542 10.542 0 0 0 2.672 7.034l-.018.025.45.452a10.643 10.643 0 0 0 15.091 0l.451-.454-.019-.025a10.565 10.565 0 0 0 2.665-7.033M10.649 1.501a9.149 9.149 0 0 1 7.039 14.99 9.493 9.493 0 0 0-4.182-2.7 4.481 4.481 0 1 0-5.716 0 9.458 9.458 0 0 0-4.182 2.7A9.059 9.059 0 0 1 1.5 10.644a9.16 9.16 0 0 1 9.149-9.143m-2.986 8.834a2.985 2.985 0 1 1 2.985 2.985 2.989 2.989 0 0 1-2.985-2.985m2.114 9.419c-.095-.009-.187-.026-.281-.038a8.569 8.569 0 0 1-.575-.086c-.11-.02-.217-.049-.326-.073a8.064 8.064 0 0 1-.509-.129c-.112-.033-.223-.071-.334-.108a8.673 8.673 0 0 1-.476-.17 10.276 10.276 0 0 1-.781-.355q-.161-.083-.318-.17-.219-.125-.431-.261a7.44 7.44 0 0 1-.3-.2c-.14-.1-.277-.2-.412-.306-.093-.073-.187-.145-.277-.222-.034-.029-.067-.061-.1-.091a7.968 7.968 0 0 1 5.993-2.731 7.866 7.866 0 0 1 5.995 2.731c-.034.03-.066.063-.1.092-.089.077-.184.147-.276.221a9.137 9.137 0 0 1-.712.506 8.148 8.148 0 0 1-.749.429c-.149.076-.3.146-.45.214-.11.048-.219.1-.331.141a8.914 8.914 0 0 1-.474.17c-.112.038-.223.076-.337.108-.168.049-.338.089-.508.129-.108.024-.217.053-.326.073-.19.037-.382.062-.575.086-.095.011-.187.029-.281.038a9.011 9.011 0 0 1-1.745 0"
                    fill="#e5e1f6"
                />
            </G>
        </G>
    </Svg>
)
const BottomTabs = () => {
    return (
        <Tab.Navigator
            screenOptions={{
                headerShown: false,
                tabBarStyle: { backgroundColor: '#23232B', borderTopColor: '#E5E1F6' }
            }} initialRouteName='Z' >
            <Tab.Screen name="Search" component={HomeScreen} options={{
                tabBarIcon: ({ focused }) => (
                    <SearchSvg />
                ),
            }} />
            <Tab.Screen name="Favorites" component={ProfileScreen} options={{
                tabBarIcon: ({ focused }) => (
                    <HeardSvg />
                ),
            }} />
            <Tab.Screen name="Z" component={ProyetosScreen} />

            <Tab.Screen name="Notification" component={FavoritosScreen} options={{
                tabBarIcon: ({ focused }) => (
                    <Bell />
                ),
            }} />
            <Tab.Screen name="Profile" component={ProfileScreen} options={{
                tabBarIcon: ({ focused }) => (
                    <Profile />
                ),
            }} />
        </Tab.Navigator>
    );
}

export default BottomTabs;
