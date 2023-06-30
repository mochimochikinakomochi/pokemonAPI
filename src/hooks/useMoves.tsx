import { useState, useEffect } from "react"
import axios from "axios"
import { Move, newMove } from "../domains/move/entity"


export const useMoves = ({ID}: {ID: number | string | undefined | null}) => {

    // moveIDの型をnumberに変換
    const pokemonID = typeof(ID) === "number" ? ID : typeof(ID) === "string" ? parseInt(ID, 10) : 1

    const [moves, setMoves] = useState<Move[]>()

    /* useEffect(() => {
        console.log("pokemonID:", pokemonID)
        getMovesByPokemonID({pokemonID: pokemonID})
    }, []) */

    const getMovesByPokemonID = async ({pokemonID}: {pokemonID: number}) => {
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

                const URL = "http://localhost:8080/move/byPokemonID/"
                const response = await axios.get(URL, config)
                const data = response.data
                setMoves(data)
            } catch (error) {
                console.error(error)
            }
        }
    }

    const getAllMoves = async () => {
        const cookies = document.cookie.split(";").map(cookie => cookie.trim())
        const jwtCookie = cookies.find(cookie => cookie.startsWith("jwtToken"))
        if (jwtCookie) {
            const token = jwtCookie.split("=")[1]
            try {
                const config = {
                    params: {
                    },
                    headers: {}
                }

                if (token) {
                    config.headers = {
                        'Authorization': token
                    }
                }

                const URL = "http://localhost:8080/move/all/"
                const response = await axios.get(URL, config)
                const data = response.data
                setMoves(data)
            } catch (error) {
                console.error(error)
            }
        }
    }

    return {moves, getMovesByPokemonID, getAllMoves}
}