export const PokemonTypesButton = ({setType}: {setType: (type: string) => void}) => {
    const types = ["normal", "fire", "water", "grass", "electric", "ice", "fighting", "poison", "ground", "flying", "psychic", "bug", "rock", "ghost", "dragon", "steel", "fairy"]

    return (
        <div>
            {types.map((type) => {
                const className = `bg-${type} text-black rounded px-6 py-3 m-2`
                return <button onClick={() => setType(type)} className={className}>{type}</button>
            })}
        </div>
    )
}