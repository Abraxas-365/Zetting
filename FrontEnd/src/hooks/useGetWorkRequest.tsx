import { useContext, useEffect, useState } from "react"
import { apiCalls } from "../api/apiCalls"
import { AuthContext } from "../context/AuthContext"


export const useGetWorkRequest = () => {
    const [worRequest, setWorkRequest] = useState([])
    const [page, setPageProject] = useState(1)
    const [isLoadingProjects, setIsLoading] = useState(true)
    const { token, userId } = useContext(AuthContext)
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    const getWorkRequest = async () => {

        const { data } = await apiCalls.get('/api/work-request/worker_id' + userId + '/page=' + page, config)
        console.log(data)
        setWorkRequest(data)
        setIsLoading(false)

    }
    useEffect(() => {
        getWorkRequest();

    }, [page])
    return { worRequest, isLoadingProjects }

}
