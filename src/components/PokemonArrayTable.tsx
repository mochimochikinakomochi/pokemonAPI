import { Link } from 'react-router-dom'
import { Pokemon } from '../domains/pokemon/entity'

export const PokemonArrayTable = ({pokemons}: {pokemons: Pokemon[]}) => {

    return (
        <div className='grid grid-cols-4 gap-3'>
        {pokemons.map(
            (pokemon) => {
                return (
                    <Link to={`../pokemonInfo/?pokemonID=${pokemon.ID}`}>
                        <div className="grid-item border pl-2 bg-gray-200 hover:bg-gray-300 rounded-[16px] shadow-md">
                            <p className="text-xl text-red-800 pt-5 font-semibold">{pokemon.ID}</p>
                            <p className="text-2xl font-bold border-b-4 border-red-500 w-fit">{pokemon.Name}</p>
                            <div className="flex pt-2">
                                <div className="flex-graw text-xl font-bold w-fit">タイプ :&nbsp;</div>
                                {pokemon.Types?.map((type) => (
                                    <div className={`flex-graw bg-${type} text-base text-white rounded w-fit m-1 px-1`}>{type}</div>
                                ))}
                            </div>
                            {/* <p className='text-xl'>H: {pokemon.Status.H}</p>
                            <p className='text-xl'>A: {pokemon.Status.A}</p>
                            <p className='text-xl'>B: {pokemon.Status.B}</p>
                            <p className='text-xl'>C: {pokemon.Status.C}</p>
                            <p className='text-xl'>D: {pokemon.Status.D}</p>
                            <p className='text-xl'>S: {pokemon.Status.S}</p> */}
                            <img src={pokemon?.ImageURL} className="w-64 h-64"/>
                            <b></b>
                        </div>
                    </Link>
                )
            }
        )}
        </div>
    )
}