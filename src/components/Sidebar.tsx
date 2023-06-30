import { Link } from "react-router-dom"

export const Sidebar = () => {
    return (
        <div>
            <nav className="w-full ">
                <div className={`flex justify-between max-w-5xl mx-auto`}>
                    <div className="text-xl font-bold  bg-ほのお bg-みず bg-くさ bg-でんき bg-こおり bg-かくとう bg-どく bg-じめん bg-ひこう bg-エスパー bg-むし bg-いわ bg-ゴースト bg-ドラゴン bg-あく bg-はがね bg-フェアリー bg-ノーマル">
                        ポケモン図鑑
                    </div>
                    <div>
                        <div className="flex lg:hidden">
                            <button 
                                type="button"
                                className="text-gray-500 dark:text-gray-200 hover:text-gray-600 dark:hover:text-gray-300 focus:outline-none focus:text-gray-600 dark:focus:text-gray-400"
                                aria-label="toggle menu"
                            >   
                            </button>
                        </div>

                        <div className="hidden -mx-4 lg:flex lg:items-center">
                            <Link
                                to="/"
                                className="block mx-4 mt-2 text-sm text-gray-700 capitalize lg:mt-0 dark:text-gray-200 hover:text-blue-600 dark:hover:text-indigo-400"
                            >ホーム</Link>
                            <Link
                                to="/PokemonInfo/"
                                className="block mx-4 mt-2 text-sm text-gray-700 capitalize lg:mt-0 dark:text-gray-200 hover:text-blue-600 dark:hover:text-indigo-400"
                            >ポケモン詳細検索</Link>
                            <Link
                                to="/PokemonList"
                                className="block mx-4 mt-2 text-sm text-gray-700 capitalize lg:mt-0 dark:text-gray-200 hover:text-blue-600 dark:hover:text-indigo-400"
                            >ポケモンリスト</Link>
                            <Link
                                to="/MoveInfo"
                                className="block mx-4 mt-2 text-sm text-gray-700 capitalize lg:mt-0 dark:text-gray-200 hover:text-blue-600 dark:hover:text-indigo-400"
                            >技詳細検索</Link>
                            <Link
                                to="/MoveList"
                                className="block mx-4 mt-2 text-sm text-gray-700 capitalize lg:mt-0 dark:text-gray-200 hover:text-blue-600 dark:hover:text-indigo-400"
                            >技リスト</Link>
                        </div>
                    </div>
                </div>
            </nav>
        </div>
    )
}