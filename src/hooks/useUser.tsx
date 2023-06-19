import { useState } from "react";
import axios from "axios";

interface User {
    Name: string;
    Password: string;
}

export const useUser = () => {
    const [user, setUser] = useState<User>()
    const baseURL = 'http://localhost:8080/insert/'

    const insertUserInfo = async (newUser: User) => {
        try {
            const params: User = {
                Name: newUser?.Name || "hiroki",
                Password: newUser?.Password || "1111",
            }
            const response = await axios.post(baseURL, null, {params})
            console.log(response.data)
        } catch (error) {
            console.error(error)
        }
    }

    return {user, setUser, insertUserInfo}
}