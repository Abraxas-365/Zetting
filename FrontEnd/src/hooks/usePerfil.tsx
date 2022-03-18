
import { useContext, useEffect, useState } from "react"
import { apiCalls } from "../api/apiCalls"
import { AuthContext } from "../context/AuthContext"
import { User } from "../interfaces/appInterfaces"



export const usePerfil = () => {
    const { token } = useContext(AuthContext)
    const [name, setName] = useState("")
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };
    const getUser = async () => {
        console.log("opteniendo ususario");
        const { data } = await apiCalls.get<User>('/api/users/', config)
        setName(data.first_name!)
    }
    useEffect(() => {
        getUser()
    }, [])
    return (name)

}
