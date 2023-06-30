import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

interface User {
    Name: string;
    Password: string;
}

export const useUser = () => {
    const [user, setUser] = useState<User>()
    const [token, setToken] = useState("")
    const baseURL = 'http://localhost:8080/user/'
    const navigate = useNavigate()

    const insertUserPassword = async (newUser: User) => {
        const URL = baseURL + 'insert/'
        try {
            const params: User = {
                Name: newUser.Name,
                Password: newUser.Password,
            }
            const response = await axios.post(URL, null, {params})
            console.log(response.data)
        } catch (error) {
            console.error(error)
        }
    }

    const getUserPassword = async (newUser: User) => {
        const URL = baseURL + 'get/password/'
        let correctPassword
        try {
            const response = await axios.get(URL, {params: {userName: newUser.Name}})
            correctPassword = response.data
        } catch (error) {
            console.error(error)
        }
        console.log("before ")
        return correctPassword
    }

    // LoginAPIをたたいてtokenを取得
    const getToken = async (newUser: User) => {
        const URL = baseURL + `get/token/?userName=${newUser.Name}&password=${newUser.Password}`
        try {
            const response = await axios.get(URL)
            setToken(response.data)
            document.cookie = `jwtToken=${response.data}; path=/;`
        } catch (error) {
            console.error(error)
        }
    }

    return {user, setUser, insertUserPassword, getUserPassword, token, getToken}
}