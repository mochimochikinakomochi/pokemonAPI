import React, { useRef, useEffect } from "react";
import { useMoves } from "../hooks/useMoves";

import { MoveArrayTable } from "../components/MoveArrayTable";


export const MoveList = () => {
    const {moves, getAllMoves} = useMoves({ID: 1})

    useEffect(() => {
        console.log("200")
        getAllMoves()
    }, [])


    return (
        <div>  
            <div className="flex justify-center p-6">
                <p className="text-2xl font-bold">技一覧</p> 
            </div>         
            {moves ? (
                <div className="flex justify-center items-center pt-4">
                    <MoveArrayTable moves={moves} />
                </div>
            ) : (
                <p className="font-xl text-center">Loading Moves...</p>
            )}
        </div>
    )
}