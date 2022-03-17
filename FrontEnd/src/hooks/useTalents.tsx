
import { useContext, useEffect, useState } from "react"
import { apiCalls } from "../api/apiCalls"
import { AuthContext } from "../context/AuthContext"




export const useTalents = (talent: string) => {
    const [talents, setTalents] = useState([])
    const { token } = useContext(AuthContext)
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };
    const getTalets = async () => {

        const { data } = await apiCalls.get('/api/profession/' + talent, config)
        console.log(data);
        setTalents(data)
    }

    useEffect(() => {
        getTalets();

    }, [])
    return { talents }

}
