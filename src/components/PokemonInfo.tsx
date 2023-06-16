import React, { useRef, useEffect } from "react";
import { usePokemons } from "../hooks/usePokemons";


export const PokemonInfo = () => {
    const {pokemons, getPokemons} = usePokemons()
    const pokemonID = useRef(1)

    useEffect(() => {
        const ID = String(pokemonID.current)
        getPokemons({pokemonID: ID})
    }, [])

    const handleClick = () => {
       pokemonID.current += 1
       const ID = String(pokemonID.current)
       getPokemons({pokemonID: ID})
    }

    console.log(pokemons)

    return (
        <div>
            <p>GET POKEMON</p>
            <button onClick={() => handleClick()}>Next</button>
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