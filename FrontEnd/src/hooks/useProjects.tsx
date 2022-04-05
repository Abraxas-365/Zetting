import { useContext, useEffect, useState } from "react"
import { apiCalls } from "../api/apiCalls"
import { AuthContext } from "../context/AuthContext"




export const useProjects = () => {
    const [projects, setState] = useState([])
    const [page, setPageProject] = useState(1)
    const [isLoadingProjects, setIsLoading] = useState(true)
    const { token } = useContext(AuthContext)
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    const getProjects = async () => {

        const { data } = await apiCalls.get('/api/projects/projects/page=' + page, config)
        console.log(data)
        setState(data)

    }
    useEffect(() => {
        getProjects();

    }, [page])
    return { projects, isLoadingProjects }

}

export const useMyProjects = () => {
    const [myProjects, setmyProjects] = useState([])
    const [page, setPageMyProject] = useState(1)
    const [isLoadingMyProjects, setIsLoading] = useState(true)
    const { token } = useContext(AuthContext)
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };
    const getMyProjects = async () => {

        const { data } = await apiCalls.get('/api/projects/myprojects/page=' + page, config)
        console.log(data);
        setmyProjects(data)
    }

    useEffect(() => {
        getMyProjects();

    }, [page])
    return { myProjects, isLoadingMyProjects }

}
