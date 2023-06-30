import { useState, useEffect } from "react"
import axios from "axios"
import { Move, newMove } from "../domains/move/entity"


export const useMove = ({ID}: {ID: number | string | undefined | null}) => {

    // moveIDの型をnumberに変換
    const moveID = typeof(ID) === "number" ? ID : typeof(ID) === "string" ? parseInt(ID, 10) : 1

    const [move, setMove] = useState<Move>(newMove)

    useEffect(() => {
        console.log("moveID:", moveID)
        getMoveByID({moveID: moveID})
    }, [moveID])

    const getMoveByID = async ({moveID}: {moveID: number}) => {
        const cookies = document.cookie.split(";").map(cookie => cookie.trim())
        const jwtCookie = cookies.find(cookie => cookie.startsWith("jwtToken"))
        if (jwtCookie) {
            const token = jwtCookie.split("=")[1]
            try {
                const config = {
                    params: {
                        moveID: moveID
                    },
                    headers: {}
                }

                if (token) {
                    config.headers = {
                        'Authorization': token
                    }
                }

                const URL = "http://localhost:8080/move/byID/"
                const response = await axios.get(URL, config)
                const data = response.data
                setMove(data)
            } catch (error) {
                console.error(error)
            }
        }
    }

    const getMoveByName = async ({moveName}: {moveName: string}) => {
        const cookies = document.cookie.split(";").map(cookie => cookie.trim())
        const jwtCookie = cookies.find(cookie => cookie.startsWith("jwtToken"))
        if (jwtCookie) {
            const token = jwtCookie.split("=")[1]
            try {
                const config = {
                    params: {
                        moveName: moveName
                    },
                    headers: {}
                }

                if (token) {
                    config.headers = {
                        'Authorization': token
                    }
                }

                const URL = "http://localhost:8080/move/byName/"
                const response = await axios.get(URL, config)
                const data = response.data
                setMove(data)
            } catch (error) {
                console.error(error)
            }
        }
    }

    return {move, getMoveByID, getMoveByName}
}