import { useContext, useEffect, useState } from "react"
import { apiCalls } from "../api/apiCalls"
import { AuthContext } from "../context/AuthContext"




export const useProjects = () => {
    const [projects, setState] = useState([])
    const { token } = useContext(AuthContext)
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };
    const getProjects = async () => {

        const { data } = await apiCalls.get('http://192.168.1.56:3000/api/projects/myprojects', config)
        setState(data)
    }
    useEffect(() => {
        getProjects();

    }, [])
    return { projects }

}
