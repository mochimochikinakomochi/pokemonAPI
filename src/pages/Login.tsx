import { Link } from "react-router-dom";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useUser } from "../hooks/useUser";


export const Login = () => {
    const navigate = useNavigate()
    const {user, setUser, insertUserPassword, getUserPassword, token, getToken} = useUser()
    const [name, setName] = useState("")
    const [password, setPassword] = useState("")
    const [isAbleToLogin, setIsAbleToLogin] = useState(false)

    const handleSignUpButton = () => {
        setUser({Name: name, Password: password})
        insertUserPassword({Name: name, Password: password})
        setName("")
        setPassword("")
    }

    const handleLoginButton = () => {
        setUser({Name: name, Password: password})
        getToken({Name: name, Password: password})
        getUserPassword({Name: name, Password: password}).then(data => {
            if (password === data) {
                setIsAbleToLogin(true)
            } else {
                alert("正しいユーザー名を入力してください   ")
            }
        })
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
            <Link to={`/pokemonInfo/?userName=${user?.Name}`}>
                <button onClick={handleSignUpButton} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 m-2 rounded">
                    新規登録
                </button>
            </Link>
            <Link to={`/pokemonInfo/?userName=${user?.Name}`}>
                <button disabled={!isAbleToLogin} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 m-2 rounded">
                    ログイン
                </button>
            </Link>
            <button onClick={() => navigate('/pokemonList')} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 m-2 rounded">
                セット
            </button>
        </div>
    )
}