import React, { useRef, useEffect } from "react";
import { useLocation } from 'react-router-dom';
import { usePokemons } from "../hooks/usePokemons";

import { PokemonTypesButton } from "../components/PokemonTypeButton";
import { PokemonStatusButton } from "../components/PokemonstatusButton";
import { PokemonArrayTable } from "../components/PokemonArrayTable";


export const PokemonList = () => {
    const location = useLocation()
    const queryParams = new URLSearchParams(location.search)
    const userName = queryParams.get("userName")
    const token = queryParams.get("token") || ""
    console.log("token:", token)

    const {pokemons, getPokemons} = usePokemons()
    const pokemonType = useRef("すべて")
    const pokemonStatus = useRef("N")

    useEffect(() => {
        getPokemons({pokemonType: pokemonType.current, pokemonStatus: pokemonStatus.current})
    }, [])


    const setType = (type: string) => {
        console.log("Type", type)
        pokemonType.current = type
        getPokemons({pokemonType: pokemonType.current, pokemonStatus: pokemonStatus.current})
    }

    const setStatus = (stat: string) => {
        pokemonStatus.current = stat
        getPokemons({pokemonType: pokemonType.current, pokemonStatus: pokemonStatus.current})
    }

    console.log("pokemons:",pokemons)   

    return (
        <div>
            <p className="text-2xl text-center p-6">ポケモン図鑑</p>
            <p className="text-xl text-center pb-4">トレーナー名: {userName}</p>
            <PokemonTypesButton setType={setType as (type: string) => void} />
            <PokemonStatusButton setStatus={setStatus as (stat: string) => void}/>
            <p className="text-2xl text-center pt-5">SelectBy: {pokemonType.current}, SortBy: {pokemonStatus.current}</p>
            
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