import { useState, useEffect } from "react";
import axios from "axios";
import { Pokemon, newPokemon } from "../domains/pokemon/entity";


export const usePokemons = () => {
    const [pokemons, setPokemons] = useState<Pokemon[]>()
    const baseURL = "http://localhost:8080/pokemon/"

    /* useEffect(() => {
        
    }, [pokemons]) */


    const getPokemons = async ({pokemonType, pokemonStatus}: {pokemonType?: string, pokemonStatus?: string}) => { 
        const cookies = document.cookie.split(";").map(cookie => cookie.trim())
        const jwtCookie = cookies.find(cookie => cookie.startsWith("jwtToken"))
        try {
            if (jwtCookie) {
                const token = jwtCookie.split("=")[1]
                if (pokemonType === "すべて") {
                    const config = {
                        params: {
                            Status: pokemonStatus
                        },
                        headers: {
                            'Authorization': token
                        }
                    }
    
                    const URL = baseURL + "all/"
                    const response = await axios.get(URL, config)
                    const data = response.data
                    setPokemons(data.map((pokemon: Pokemon) => pokemon))
                } else {
                    const config = {
                        params: {
                            Type: pokemonType,
                            Status: pokemonStatus
                        },
                        headers: {
                            'Authorization': token
                        }
                    }
                    
                    const URL = baseURL + "list/type/"
                    const response = await axios.get(URL, config)
                    const data = response.data
                    setPokemons(data.map((pokemon: Pokemon) => pokemon))
                }
                
            }
        } catch (error) {
            console.error(error)
        }  
    }

    const getPokemponsByMoveID = async ({moveID}: {moveID: number}) => {
        const cookies = document.cookie.split(";").map(cookie => cookie.trim())
        const jwtCookie = cookies.find(cookie => cookie.startsWith("jwtToken"))
        try {
            if (jwtCookie && moveID > 0) {
                const token = jwtCookie.split("=")[1]
                const config = {
                    params: {
                        moveID: moveID
                    },
                    headers: {
                        'Authorization': token
                    }
                }

                const URL = baseURL + "byMoveID/"
                const response = await axios.get(URL, config)
                const data = response.data
                console.log(data)
                data ? setPokemons(data.map((pokemon: Pokemon) => pokemon)) : setPokemons([newPokemon()])
            }
        } catch (error) {
            console.error(error)
        }
    }

    const getPokemponsByKeyWord = async ({keyWord}: {keyWord: string}) => {
        const cookies = document.cookie.split(";").map(cookie => cookie.trim())
        const jwtCookie = cookies.find(cookie => cookie.startsWith("jwtToken"))
        try {
            if (jwtCookie) {
                const token = jwtCookie.split("=")[1]
                const config = {
                    params: {
                        KeyWord: keyWord
                    },
                    headers: {
                        'Authorization': token
                    }
                }

                const URL = baseURL + "byKeyWord/"
                const response = await axios.get(URL, config)
                const data = response.data
                console.log(data)
                data ? setPokemons(data.map((pokemon: Pokemon) => pokemon)) : setPokemons([newPokemon()])
            }
        } catch (error) {
            console.error(error)
        }
    }

    return {pokemons, getPokemons, getPokemponsByMoveID, getPokemponsByKeyWord}
}