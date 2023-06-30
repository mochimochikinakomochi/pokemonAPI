export type PokemonStatus = {
    H: number;
    A: number;
    B: number;
    C: number;
    D: number;
    S: number;
}

export type Pokemon = {
    ID: number;
    Name: string;
    Types: string[];
    Status: PokemonStatus;
    Category: string;
    Height: number;
    Weight: number;
    ImageURL: string;
}

const newPokemonStats = (): PokemonStatus => ({
    H: 1,
    A: 1,
    B: 1,
    C: 1,
    D: 1,
    S: 1,
})

export const newPokemon = (): Pokemon => ({
    ID: 0, 
    Name: "?", 
    Types: ["?"], 
    Status: newPokemonStats(),
    Category: "?",
    Height: 0,
    Weight: 0,
    ImageURL: "",
})