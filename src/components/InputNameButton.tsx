import { useState } from "react"
import { Pokemon } from "../domains/pokemon/entity"
import { Move } from "../domains/move/entity"
import { usePokemons } from "../hooks/usePokemons"

export const InputNameButton = ({handleClick}: {handleClick: (name: string) => void}) => {
    const [variable, setVariable] = useState("")
    const [isFocus, setIsFocus] = useState(false)
    const {pokemons, getPokemponsByKeyWord} = usePokemons()

    return (
        <div>
            <input  
                type="text"
                value={variable} 
                onChange={(event) => setVariable(event.target.value)} 
                onFocus={() => setIsFocus(true)}
                className="w-64 px-4 py-2 rounded-lg border-2 border-gray-300 focus:outline-none focus:border-blue-500 transition-colors duration-300"
                placeholder="名前を入力!"
            />
            <button 
                onClick={() => {
                    handleClick(variable)
                    setVariable("")
                }}
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg"
            >
                検索
            </button>
        </div>
    )
}