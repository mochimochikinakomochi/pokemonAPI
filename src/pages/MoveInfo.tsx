import { useState, useEffect, useRef } from "react";
import { useSearchParams } from "react-router-dom";
import { useNavigate } from "react-router-dom";

import { useMove } from "../hooks/useMove";
import { usePokemons } from "../hooks/usePokemons";
import { PokemonArrayTable } from "../components/PokemonArrayTable";
import { InputIDButton } from "../components/InputIDButton";
import { InputNameButton } from "../components/InputNameButton";

export const MoveInfo = () => {
    const [searchParams] = useSearchParams()
    const ID = searchParams.get("moveID") || "1"
    const {move, getMoveByID, getMoveByName} = useMove({ID: ID})
    const {pokemons, getPokemponsByMoveID} = usePokemons()
    const moveID = useRef(1)
    const moveName = useRef("")

    /* useEffect(() => {
        getPokemponsByMoveID({moveID: parseInt(ID, 10)})
    }, []) */

    useEffect(() => {
        moveID.current = move.ID
        console.log("222", moveID.current)
        getPokemponsByMoveID({moveID: moveID.current})
    }, [move])

    const handleClickID = (id: string) => {
        moveID.current = parseInt(id, 10)
        getMoveByID({moveID: moveID.current})
        getPokemponsByMoveID({moveID: moveID.current})
    } 

    const handleClickName = (name: string) => {
        moveName.current = name
        getMoveByName({moveName: moveName.current})
        getPokemponsByMoveID({moveID: move.ID})
    }

    return (
        <div>
            <div className="flex justify-end mt-6">
                <InputIDButton handleClick={handleClickID} />
            </div>
            <div className="flex justify-end mt-2">
                <InputNameButton handleClick={handleClickName} />
            </div>
            <div className="flex items-center justify-center pt-6">
                <div className="flex-graw flex items-center grid-item border p-2 bg-gray-200 rounded-[16px] shadow-md w-2/3">
                    <div className="flex-graw items-center">
                        <div className="flex p-4">
                        <div className="flex-graw text-xl font-bold w-fit">タイプ :&nbsp;</div>
                            <div className={`flex-graw bg-${move.Type} text-xl text-white rounded w-fit m-1 px-1`}>{move.Type}</div>

                        </div>
                        <p className='text-xl font-bold pl-2'>分類 : {move?.Class}</p>
                        <p className='text-xl font-bold pl-2'>命中率 :&nbsp;{move?.Accuracy === 0 ? "必中" : move?.Accuracy}</p>
                        <p className='text-xl font-bold pl-2'>PP :&nbsp;{move?.PP}</p>
                        <p className='text-xl font-bold pl-2'>威力 :&nbsp;{move?.Class === "変化" ? " ー" : move?.Power}</p>
                        <p className='text-xl font-bold pl-2'>優先度 :&nbsp;{move?.Priority}</p>
                    </div>
                    <div className="flex-graw justify-center items-center bg-white w-1/3 h-full rounded-[16px] shadow-md ml-32">
                        <p className="text-xl text-red-800 p-4 font-semibold">No.{move?.ID}</p>
                        <p className="text-4xl font-bold w-fit p-4">{move?.Name}</p>
                    </div>
                </div>
            </div>
            <div className="flex justify-center">
                <div className="grid-item border bg-gray-200 rounded-[16px] shadow-md w-2/3 p-3 m-3">
                    <p className="text-2xl">技の説明文</p>
                    <p className="text-xl font-bold pl-2">{move.Description}</p>
                </div>
            </div>
            <div className="flex justify-center p-4">
                <p className="text-2xl">{move.Name}を覚えるポケモン</p>
            </div>
            {pokemons ? (
                <div className="flex justify-center items-center pt-4">
                    <PokemonArrayTable pokemons={pokemons} />
                </div>
            ) : (
                <p className="font-xl text-center font-bold p-6">Loading Pokemons...</p>
            )}
        </div>
    )
}