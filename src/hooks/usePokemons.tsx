import { useState } from "react";
import axios from "axios";

interface Pokemon {
    ID: number;
    Name: string;
    Type: string;
    Stats: string;
}

export const usePokemons = () => {
    const [pokemons, setPokemons] = useState<Pokemon[]>()
    const baseURL = "http://localhost:8080/"

    const getPokemons = async ({pokemonID, pokemonType}: {pokemonID?: string, pokemonType?: string}) => { 
          try {
            const URL = baseURL/*  + "?ID=" + pokemonID */ + "?Type=" + pokemonType
            const response = await axios.get(URL)
            const data = response.data
            setPokemons(data.map((pokemon: Pokemon) => pokemon))
        } catch (error) {
            console.error(error)
        }
    }

    return {pokemons, getPokemons}
}