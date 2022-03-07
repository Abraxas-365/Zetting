import { useEffect, useState } from "react";
import { log } from "react-native-reanimated";
import { apiCalls } from "../api/apiCalls";

//agregar isPlay para que aparezca iswin con la plantilla
export const tryGame = (email: string) => {

    const [exists, setExists] = useState(false);
    const url = 'http://localhost:3000/api/users';
    const getgame = async () => {
        const res = await apiCalls.get(url + "/" + email);
        setExists(res.data.exist)
        console.log("ddd", res.data)

    }
    return { exists, getgame, setExists };

}
