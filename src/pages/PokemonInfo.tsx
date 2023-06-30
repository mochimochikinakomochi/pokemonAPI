import { useRef, useEffect } from "react";
import { useParams, useNavigate, useSearchParams } from "react-router-dom";

import { usePokemon } from "../hooks/usePokemon"
import { InputIDButton } from "../components/InputIDButton";
import { InputNameButton } from "../components/InputNameButton";
import { MoveArrayTable } from "../components/MoveArrayTable";
import { NextIDButton } from "../components/NextIDButton";
import { PreIDButton } from "../components/PreIDButton";


export const PokemonInfo = () => {
    const [searchParams] = useSearchParams()
    const ID = searchParams.get("pokemonID") || "1"
    const name = searchParams.get("Name")
    const {pokemon, moves, getPokemonByName, getPokemonByID, getMovesByPokemonID} = usePokemon({ID: ID, Name: name})
    const pokemonID = useRef(1)

    useEffect(() => {
        getPokemonByID({pokemonID: parseInt(ID)})
    }, [])

    useEffect(() => {
        pokemonID.current = pokemon.ID
    }, [pokemon])

    const handleClickID = (id: string) => {
        pokemonID.current = parseInt(id, 10)
        getPokemonByID({pokemonID: pokemonID.current})
        getMovesByPokemonID({pokemonID: pokemonID.current})
    }

    const handleClickName = (name: string) => {
        getPokemonByName({pokemonName: name})
    }

    const handleClick = (isNext: boolean) => {
        if (isNext && pokemonID.current < 1010) {
            pokemonID.current += 1
            getPokemonByID({pokemonID: pokemonID.current})
            getMovesByPokemonID({pokemonID: pokemonID.current})
        } else if (!isNext && pokemonID.current > 1) {
            pokemonID.current -= 1
            getPokemonByID({pokemonID: pokemonID.current})
            getMovesByPokemonID({pokemonID: pokemonID.current})
        }
    }

    return (
        <div>
            <div className="flex justify-end mt-6">
                    <InputIDButton handleClick={handleClickID} />
            </div>
            <div className="flex justify-end mt-2">
                    <InputNameButton handleClick={handleClickName} />
            </div>
            <div className="flex justify-center pt-6">
                <div className="flex-graw flex items-center grid-item border pl-2 bg-gray-200 rounded-[16px] shadow-md w-full h-fit">
                    <div className="flex-start">
                        <PreIDButton handleClick={handleClick} />
                    </div>
                    <div className="flex-graw items-center pl-32">
                        <img src={pokemon?.ImageURL} className="w-80 h-80"/>
                    </div>
                    <div className="flex-graw bg-white w-1/3 h-full rounded-[16px] shadow-md">
                        <p className="text-xl text-red-800 p-5 font-semibold">No.{pokemon.ID}</p>
                        <p className="text-4xl font-bold p-5 w-fit">{pokemon.Name}</p>
                    </div>
                    <div className="ml-auto">
                        <NextIDButton handleClick={handleClick} />
                    </div>
                </div>
            </div>
            <div className='flex items-start mt-4'>
                <div className="flex-graw border pl-4 pr-4 ml-4 bg-gray-200 rounded-[16px] shadow-md w-fit">
                    <p className="text-2xl font-bold w-fit p-4">分類 :&nbsp;{pokemon.Category}</p>
                    <div className="flex p-4">
                        <div className="flex-graw text-2xl font-bold w-fit">タイプ :&nbsp;</div>
                        {pokemon.Types?.map((type) => (
                            <div className={`flex-graw bg-${type} text-xl text-white rounded w-fit m-1 px-1`}>{type}</div>
                        ))}
                    </div>
                    <p className="text-2xl font-bold w-fit p-4">高さ :&nbsp;{pokemon.Height / 10}m</p>
                    <p className="text-2xl font-bold w-fit p-4">重さ :&nbsp;{pokemon.Weight / 10}kg</p>
                </div>
                <div className="flex-graw border pl-4 ml-4 bg-gray-200 rounded-[16px] shadow-md w-fit">
                    <table className="p-2">
                        <thead>
                            <tr>
                                <th className='text-xl p-2'>種族値</th>
                                <th className='text-xl p-2'>数値</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td className='text-xl p-2'>HP</td>
                                <td>
                                    <div className={`bg-blue-500 h-6 rounded-lg shadow-2xl`} style={{width: `${pokemon.Status.H}px`}}></div>
                                </td>
                                <td className='text-xl p-2'>{pokemon.Status.H}</td>
                            </tr>
                            <tr>
                                <td className='text-xl p-2'>攻撃</td>
                                <td>
                                    <div className={`bg-blue-500 h-6 rounded-lg`} style={{width: `${pokemon.Status.A}px`}}></div>
                                </td>
                                <td className='text-xl p-2'>{pokemon.Status.A}</td>
                            </tr>
                            <tr>
                                <td className='text-xl p-2'>防御</td>
                                <td>
                                    <div className={`bg-blue-500 h-6 rounded-lg`} style={{width: `${pokemon.Status.B}px`}}></div>
                                </td>
                                <td className='text-xl p-2'>{pokemon.Status.B}</td>
                            </tr>
                            <tr>
                                <td className='text-xl p-2'>特攻</td>
                                <td>
                                    <div className={`bg-blue-500 h-6 rounded-lg`} style={{width: `${pokemon.Status.C}px`}}></div>
                                </td>
                                <td className='text-xl p-2'>{pokemon.Status.C}</td>
                            </tr>
                            <tr>
                                <td className='text-xl p-2'>特防</td>
                                <td>
                                    <div className={`bg-blue-500 h-6 rounded-lg`} style={{width: `${pokemon.Status.D}px`}}></div>
                                </td>
                                <td className='text-xl p-2'>{pokemon.Status.D}</td>
                            </tr>
                            <tr>
                                <td className='text-xl p-2'>素早さ</td>
                                <td>
                                    <div className={`bg-blue-500 h-6 rounded-lg`} style={{width: `${pokemon.Status.S}px`}}></div>
                                </td>
                                <td className='text-xl p-2'>{pokemon.Status.S}</td>
                            </tr>
                            <tr>
                                <td className='text-xl p-2'>平均/合計</td>
                                <td>
                                    <div className={`bg-blue-500 h-6 rounded-lg`} style={{width: `${(pokemon.Status.H + pokemon.Status.A + pokemon.Status.B + pokemon.Status.C + pokemon.Status.D + pokemon.Status.S) / 6}px`}}></div>
                                </td>
                                <td className='text-xl p-2'>{((pokemon.Status.H + pokemon.Status.A + pokemon.Status.B + pokemon.Status.C + pokemon.Status.D + pokemon.Status.S) / 6).toFixed(1)} / {(pokemon.Status.H + pokemon.Status.A + pokemon.Status.B + pokemon.Status.C + pokemon.Status.D + pokemon.Status.S)}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                {moves ? (
                    <div className="flex-graw border pl-4 ml-4 bg-gray-200 rounded-[16px] shadow-md w-fit">
                        <p className="text-2xl">覚える技</p>
                        <MoveArrayTable moves={moves} />
                    </div>
                ) : (
                    <p className="font-xl text-center font-bold p-6">Loading Moves...</p>
                )}
            </div>
        </div>
    )
}