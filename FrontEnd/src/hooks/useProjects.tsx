import { useContext, useEffect, useState } from "react"
import { apiCalls } from "../api/apiCalls"
import { AuthContext } from "../context/AuthContext"




export const useProjects = () => {
    const [projects, setState] = useState([])
    const [myProjects, setmyProjects] = useState([])
    const { token } = useContext(AuthContext)
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };
    const getMyProjects = async () => {

        const { data } = await apiCalls.get('/api/projects/myprojects', config)
        console.log(data);
        setmyProjects(data)
    }
    // useEffect(() => {
    //     getProjects();

    // }, [])

    const getProjects = async () => {

        const { data } = await apiCalls.get('/api/projects/projects', config)
        console.log(data)
        setState(data)

    }
    useEffect(() => {
        getMyProjects();
        getProjects();

    }, [])
    return { projects, myProjects }

}
