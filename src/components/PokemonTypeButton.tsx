interface Type {
    typeID: number,
    nameEn: string,
    nameJa: string
}

export const PokemonTypesButton = ({setType}: {setType: (type: string) => void}) => {
    const types: Type[] = [
        {typeID: 0, nameEn: "all", nameJa: "すべて"},
        {typeID: 1, nameEn: "normal", nameJa: "ノーマル"},
        {typeID: 2, nameEn: "fire", nameJa: "ほのお"},
        {typeID: 3, nameEn: "water", nameJa: "みず"},
        {typeID: 4, nameEn: "grass", nameJa: "くさ"},
        {typeID: 5, nameEn: "electric", nameJa: "でんき"},
        {typeID: 6, nameEn: "ice", nameJa: "こおり"},
        {typeID: 7, nameEn: "fighting", nameJa: "かくとう"},
        {typeID: 8, nameEn: "poison", nameJa: "どく"},
        {typeID: 9, nameEn: "ground", nameJa: "じめん"},
        {typeID: 10, nameEn: "flying", nameJa: "ひこう"},
        {typeID: 11, nameEn: "psychic", nameJa: "エスパー"},
        {typeID: 12, nameEn: "bug", nameJa: "むし"},
        {typeID: 13, nameEn: "rock", nameJa: "いわ"},
        {typeID: 14, nameEn: "ghost", nameJa: "ゴースト"},
        {typeID: 15, nameEn: "dragon", nameJa: "ドラゴン"},
        {typeID: 16, nameEn: "dark", nameJa: "あく"},
        {typeID: 17, nameEn: "steel", nameJa: "はがね"},
        {typeID: 18, nameEn: "fairy", nameJa: "フェアリー"},
    ]

    return (
        <div>
            {types.map((type) => {
                const className = `bg-${type.nameEn} text-white rounded px-6 py-3 m-2`
                return <button onClick={() => setType(type.nameJa)} className={className}>{type.nameJa}</button>
            })}
        </div>
    )
}