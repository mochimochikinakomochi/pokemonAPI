import { useState, useEffect } from "react";
import axios from "axios";
import { Pokemon, newPokemon } from "../domains/pokemon/entity";
import { useMoves } from "./useMoves";


export const usePokemon = ({ID, Name}: {ID?: number | string | undefined | null, Name?: string | undefined | null}) => {

    // pokemonIDの型をnumberに変換
    const pokemonID = typeof(ID) === "number" ? ID : typeof(ID) === "string" ? parseInt(ID, 10) : 1
    const pokemonName = typeof(Name) === "string" ? Name : "フシギダネ"

    const [pokemon, setPokemon] = useState<Pokemon>(newPokemon())
    const {moves, getMovesByPokemonID} = useMoves({ID: pokemonID})
    console.log("moves", moves)

    useEffect(() => {
        getMovesByPokemonID({pokemonID: pokemon.ID})
    }, [pokemon])

    const getPokemonByName = async ({pokemonName}: {pokemonName: string}) => { 
        const cookies = document.cookie.split(";").map(cookie => cookie.trim())
        const jwtCookie = cookies.find(cookie => cookie.startsWith("jwtToken"))
        if (jwtCookie) {
            const token = jwtCookie.split("=")[1]
            try {
                const config = {
                    params: {
                        pokemonName: pokemonName
                    },
                    headers: {}
                }
    
                if (token) {
                    config.headers = {
                        'Authorization': token
                    }
                }

                const URL = "http://localhost:8080/pokemon/byName/"
                const response = await axios.get(URL, config)
                const data = response.data
                setPokemon(data)
            } catch (error) {
                console.error(error)
            }
        }

    }

    const getPokemonByID = async ({pokemonID}: {pokemonID: number}) => { 
        const cookies = document.cookie.split(";").map(cookie => cookie.trim())
        const jwtCookie = cookies.find(cookie => cookie.startsWith("jwtToken"))
        if (jwtCookie) {
            const token = jwtCookie.split("=")[1]
            try {
                const config = {
                    params: {
                        pokemonID: pokemonID
                    },
                    headers: {}
                }
    
                if (token) {
                    config.headers = {
                        'Authorization': token
                    }
                }

                const URL = "http://localhost:8080/pokemon/byID/"
                const response = await axios.get(URL, config)
                const data = response.data
                setPokemon(data)
            } catch (error) {
                console.error(error)
            }
        }

    }


    return {pokemon, moves, getPokemonByName, getPokemonByID, getMovesByPokemonID}
}