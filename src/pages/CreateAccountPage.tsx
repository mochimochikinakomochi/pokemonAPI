import { useState, useRef } from "react";
import { useUser } from "../hooks/useUser";

interface User {
    Name: string;
    PassWord: string;
}

export const CreateAccountPage = () => {
    const {user, setUser, insertUserInfo} = useUser()
    const [name, setName] = useState("")
    const [password, setPassword] = useState("")

    const handleClick = () => {
        console.log("handleClick")
        console.log("name:", name, "password", password)
        setUser({Name: name, Password: password})
        console.log(user?.Name, user?.Password)
        insertUserInfo({Name: name, Password: password})
        setName("")
        setPassword("")
    }
    
    return (
        <div>
            <div>
                <input value={name} onChange={(event) => setName(event.target.value)} type="text" className="px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-500" placeholder="ユーザー名"></input>
            </div>
            <div>
                <input value={password} onChange={(event) => setPassword(event.target.value)} type="text" className="px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-500" placeholder="パスワード"></input>
            </div>
            <button onClick={handleClick}>セット</button>
        </div>
    )
}