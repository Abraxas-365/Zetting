import React from 'react';
import { StyleSheet, View } from 'react-native';

const MembersPhotos = () => {
    return (
        <View style={{ flex: 1, flexDirection: 'row' }}>
            <View style={[styles.container]}></View>
            <View style={[styles.container]}></View>
            <View style={[styles.container]}></View>

        </View>
    );
};

export default MembersPhotos;

const styles = StyleSheet.create({
    container: {
        marginHorizontal: '-7%',
        backgroundColor: 'red',
        borderRadius: 100,
        height: '100%',
        width: '33.3%'
    }

})
