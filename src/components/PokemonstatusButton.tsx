export const PokemonStatusButton = ({setStatus}: {setStatus: (stat: string) => void}) => {
    const status = ["N", "H", "A", "B", "C", "D", "S", "SUM"]

    return (
        <div>
            {status.map((stat) => {
                return <button onClick={() => setStatus(stat)} className="bg-lime-500 text-black rounded px-6 py-3 m-2">{stat}</button>
            })}
        </div>
    )
}