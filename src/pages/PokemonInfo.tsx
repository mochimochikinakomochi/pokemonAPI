import React, { useRef, useEffect } from "react";
import { usePokemons } from "../hooks/usePokemons";
import { PreAndNextButton } from "../components/PreAndNextIDButton";
import { PokemonTypesButton } from "../components/PokemonTypeButton";


export const PokemonInfo = () => {
    const {pokemons, getPokemons} = usePokemons()
    const pokemonID = useRef(1)
    const pokemonType = useRef("")

    useEffect(() => {
        const ID = String(pokemonID.current)
        getPokemons({pokemonID: ID})
    }, [])

    const handleClick = (isNext: boolean) => {
        if (isNext === true) {
            pokemonID.current += 1
        } else {
            pokemonID.current -= 1
        }
        const ID = String(pokemonID.current)
        getPokemons({pokemonID: ID})
    }

    const setType = (type: string) => {
        pokemonType.current = type
        console.log("type", pokemonType.current)
        getPokemons({pokemonType: pokemonType.current})
    }

    console.log(pokemons)

    return (
        <div>
            <p>GET POKEMON</p>
            <PreAndNextButton handleClick={handleClick as (isNext: boolean) => void} />
            <PokemonTypesButton setType={setType as (type: string) => void} />
            {pokemons ? (
                pokemons.map((pokemon) => {
                return (
                    <React.Fragment>
                        <p>{pokemon.ID}</p>
                        <p>{pokemon.Name}</p>
                        <p>{pokemon.Type}</p>
                        <p>{pokemon.Stats}</p>
                    </React.Fragment>
                )
            })) : (
                <p>Loading Pokemons...</p>
            )}
        </div>
    )
}